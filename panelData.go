package goplot

import (
	"fmt"
	"io/ioutil"
	"os"
)

/* type PanelData */
type PanelData struct {
	Name   string
	Opt    *PanelDataOption
	atexit func()
}

func NewPanelData(name string, opt *PanelDataOption) *PanelData {
	if opt == nil {
		opt = NewPanelDataOption()
	}
	return &PanelData{
		Name:   name,
		Opt:    opt,
		atexit: func() { return },
	}
}

func NewPanelDataFromArray(xdata, ydata, zdata []float64, opt *PanelDataOption) (*PanelData, error) {
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
	if data.Opt.IsFunc {
		return fmt.Sprintf(`%s %s`, data.Name, data.Opt)
	}
	return fmt.Sprintf(`"%s" %s`, data.Name, data.Opt)
}

func (data *PanelData) SetData(name string) {
	data.Name = name
}

func (data *PanelData) SetOption(opt *PanelDataOption) {
	data.Opt = opt
}
