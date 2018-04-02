package goplot

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

/* type Function */
type Function struct {
	Name string `xml:"name,attr"`
	Text string `xml:",chardata"`
	Para []float64 `xml:"para,attr"`
}

func (f *Function) String() string {
	form := f.Text
	for i, p := range f.Para {
		form = strings.Replace(form, "$"+strconv.Itoa(i), fmt.Sprintf("%f", p), -1)
	}
	return fmt.Sprintf("%s", form)
}

/* type PanelFunction */
type PanelFunction struct {
	Function   *Function `xml:"Func"`
	Opt    *PlotOption
	atexit func()
}

func NewPanelFunction(f *Function, opt *PlotOption) *PanelFunction {
	if opt == nil {
		opt = NewPlotOption()
	}
	return &PanelFunction{
		Function:   f,
		Opt:    opt,
		atexit: func() { return },
	}
}

func (data *PanelFunction) String() string {
	return fmt.Sprintf("%s with %s title \"%s\" %s\n", data.Function, data.Opt.With, data.Opt.Title, data.Opt.LineStyle)
}

func (data *PanelFunction) SetFunction(f *Function) {
	data.Function = f
}

func (data *PanelFunction) SetOption(opt *PlotOption) {
	data.Opt = opt
}

/* type PanelData */
type PanelData struct {
	FileName   string `xml:""`
	Opt    *PlotOption
	atexit func()
}

func NewPanelData(name string, opt *PlotOption) *PanelData {
	if opt == nil {
		opt = NewPlotOption()
	}
	return &PanelData{
		FileName:   name,
		Opt:    opt,
		atexit: func() { return },
	}
}

func NewPanelDataFromArray(xdata, ydata, zdata []float64, opt *PlotOption) (*PanelData, error) {
	/* Generate tempfile to load data from gnuplot */
	tmpfile, err := ioutil.TempFile("", "goplot")
	if err != nil {
		return nil, err
	}
	data := &PanelData{}
	// register os.Remove(tmpfile) to atexit(), which is called when closing plotter.
	// NOTE: Don't remove this tmpfile in this function. Remove it after plotter.Plot().
	data.atexit = func() { os.Remove(tmpfile.Name()) }
	if zdata == nil {
		if len(xdata) != len(ydata) {
			return nil, fmt.Errorf(`xdata and ydata have different size %d %d`, len(xdata), len(ydata))
		}
	} else {
		if len(xdata) != len(ydata) || len(xdata) != len(zdata) {
			return nil, fmt.Errorf(`xdata, ydata and zdata have different size %d %d %d`, len(xdata), len(ydata), len(zdata))
		}
	}
	// Write to tempfile
	x := 0.
	y := 0.
	z := 0.
	for i := 0; i < len(xdata); i++ {
		x = xdata[i]
		y = ydata[i]
		if zdata != nil {
			z = zdata[i]
		}
		if _, err := tmpfile.WriteString(fmt.Sprintf("%f %f %f\n", x, y, z)); err != nil {
			return nil, err
		}
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}

	return NewPanelData(tmpfile.Name(), opt), nil
}

func (data *PanelData) String() string {
	return fmt.Sprintf("\"%s\" using %s index %d with %s title \"%s\" %s", data.FileName, data.Opt.Using, data.Opt.Index, data.Opt.With, data.Opt.Title, data.Opt.LineStyle)
}

func (data *PanelData) SetData(name string) {
	data.FileName = name
}

func (data *PanelData) SetOption(opt *PlotOption) {
	data.Opt = opt
}

/* type PlotOption */
type PlotOption struct {
	Name      string `xml:"name,attr"`
	Using     string `xml:"using"`
	Index     int `xml:"index"`
	With      string `xml:"with"`
	LineStyle *LineStyle
	Title     string `xml:"title"`
}

func NewPlotOption() *PlotOption {
	ls := NewLineStyle()
	opt := &PlotOption{
		Using:     "1:2",
		Index:     0,
		With:      "line",
		LineStyle: ls,
		Title:     "",
	}
	return opt
}

func NewPlotOptionFromMap(values map[string]interface{}) (*PlotOption, error) {
	opt := NewPlotOption()
	if values != nil {
		for key, value := range values {
			if err := opt.Set(key, value); err != nil {
				return nil, err
			}
		}
	}
	return opt, nil
}

func (opt *PlotOption) Set(key string, value interface{}) error {
	switch key {
	case `using`, `u`:
		opt.Using = value.(string)
	case `index`, `ind`:
		opt.Index = value.(int)
	case `with`, `w`:
		opt.With = value.(string)
	case `LineStyle`, `lineStyle`, `linestyle`, `ls`:
		opt.LineStyle = value.(*LineStyle)
	case `title`:
		opt.Title = value.(string)
	default:
		return fmt.Errorf(`Unknown key %v`, key)
	}
	return nil
}

func (opt *PlotOption) Copy() *PlotOption {
	opt2 := &PlotOption{}
	*opt2 = *opt
	return opt2
}