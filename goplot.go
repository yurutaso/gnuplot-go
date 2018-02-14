package goplot

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

/* type LineStyle */
type LineStyle struct {
	lineWidth int
	lineType  int
	lineColor string
	pointType int
	pointSize int
	dashType  string
}

func NewLineStyle() *LineStyle {
	return &LineStyle{
		lineWidth: 1,
		lineType:  1,
		lineColor: "#000000",
		pointType: 1,
		pointSize: 1,
		dashType:  "",
	}
}

func (lineStyle *LineStyle) String() string {
	return fmt.Sprintf(`lw %d lt %d lc rgb "%s" pt %d ps %d dt "%s"`,
		lineStyle.lineWidth, lineStyle.lineType, lineStyle.lineColor, lineStyle.pointType, lineStyle.pointSize, lineStyle.dashType)
}

/* type PlotOption */
type PlotOption struct {
	isFunc    bool
	using     string
	index     int
	with      string
	lineStyle *LineStyle
	title     string
}

func NewPlotOption() *PlotOption {
	return &PlotOption{
		isFunc:    false,
		using:     "1:2",
		index:     0,
		with:      "line",
		lineStyle: NewLineStyle(),
		title:     "",
	}
}

/* type PlotData */
type PlotData struct {
	filename string
	option   *PlotOption
}

func NewPlotData(filename string, option *PlotOption) *PlotData {
	if option == nil {
		option = NewPlotOption()
	}
	return &PlotData{
		filename: filename,
		option:   option,
	}
}

func (data *PlotData) String() string {
	opt := data.option
	s := fmt.Sprintf(`plot "%s" using %s index %d with %s title "%s" %s`,
		data.filename, opt.using, opt.index, opt.with, opt.title, opt.lineStyle)
	return s
}

/* type AxisOption */
type AxisOption struct {
	name        string
	min         float64
	max         float64
	tics        float64
	mtics       float64
	label       string
	labelOffset float64
	log         bool
}

func NewAxisOption(name string) *AxisOption {
	return &AxisOption{
		name:        name,
		min:         0,
		max:         10,
		tics:        10,
		mtics:       2,
		label:       "label",
		labelOffset: 0,
		log:         false,
	}
}

func (axis *AxisOption) String() string {
	s := fmt.Sprintf(`
set %srange [%f:%f]
set %stics %f
set m%stics %f
set %slabel "%s" offset %f`,
		axis.name, axis.min, axis.max,
		axis.name, axis.tics,
		axis.name, axis.mtics,
		axis.name, axis.label, axis.labelOffset,
	)
	if axis.log {
		s += `set log` + axis.name
	}
	return s
}

/* type PanelOption */
type PanelOption struct {
	xaxis *AxisOption
	yaxis *AxisOption
	//zaxis  *AxisOption
	sample int
	grid   string
	key    string
}

func NewPanelOption() *PanelOption {
	return &PanelOption{
		xaxis: NewAxisOption(`x`),
		yaxis: NewAxisOption(`y`),
		//zaxis:  NewAxisOption(),
		sample: 1000,
		grid:   "",
		key:    "",
	}
}

func (option *PanelOption) String() string {
	return fmt.Sprintf(`
%s
%s
set sample %d
set grid %s
set key %s
`,
		option.xaxis.String(), option.yaxis.String(), option.sample, option.grid, option.key)
}

/* type Panel */
type Panel struct {
	data   []*PlotData
	option *PanelOption
}

func NewPanel() *Panel {
	return &Panel{
		data:   make([]*PlotData, 0, 0),
		option: NewPanelOption(),
	}
}

func (panel *Panel) String() string {
	s := panel.option.String()
	for _, data := range panel.data {
		s += data.String()
	}
	return s
}

func (panel *Panel) SetOption(opt *PanelOption) {
	panel.option = opt
}

func (panel *Panel) AddData(data *PlotData) {
	panel.data = append(panel.data, data)
}

/* type plotter*/
type Plotter struct {
	panels  []*Panel
	font    *FontConfig
	figname string
}

func NewPlotter() *Plotter {
	return &Plotter{
		panels:  make([]*Panel, 0, 0),
		font:    NewFontConfig(),
		figname: `output.eps`,
	}
}

func (plotter *Plotter) String() string {
	s := fmt.Sprintf(`#!/usr/bin/env/gnuplot
%s
set terminal postscript eps enhanced color
set output "%s"`,
		plotter.font,
		plotter.figname,
	)
	for _, panel := range plotter.panels {
		s += panel.String()
	}
	return s
}

func (plotter *Plotter) AddPanel(panel *Panel) {
	plotter.panels = append(plotter.panels, panel)
}

func (plotter *Plotter) SetOutput(figname string) {
	plotter.figname = figname
}

func (plotter *Plotter) Plot() error {
	tmp, err := ioutil.TempFile("", "goplot_exec_temp")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	tmp.WriteString(plotter.String())
	err = tmp.Close()
	if err != nil {
		return err
	}
	cmd := exec.Command("gnuplot", tmp.Name())
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
