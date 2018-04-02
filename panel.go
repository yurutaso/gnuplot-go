package goplot

import (
	"strings"
	"fmt"
)

/* type Panel */
type Plot struct {
	Data []*PanelData `xml:"Data"`
	Func []*PanelFunction `xml:"Function"`
}

func NewPlot() *Plot {
	return &Plot{
		Data: make([]*PanelData, 0, 0),
		Func: make([]*PanelFunction, 0, 0),
	}
}

type Panel struct {
	Plot       *Plot
	Opt        *PanelOption
	Xaxis      *Axis
	Yaxis      *Axis
	//Zaxis      *Axis
	Annotation struct {
		Labels []*AnnotationLabel `xml:"label"`
		Arrows []*AnnotationArrow `xml:"arrow"`
	}
}

func NewPanel(opt *PanelOption) *Panel {
	if opt == nil {
		opt = NewPanelOption()
	}
	return &Panel{
		Xaxis: NewAxis(`x`),
		Yaxis: NewAxis(`y`),
		Plot: NewPlot(),
		Opt:  opt,
	}
}

func (panel *Panel) String() string {
	s := fmt.Sprintf("%s\n%s\n%s\n", panel.Xaxis, panel.Yaxis, panel.Opt)
	numPlot := len(panel.Plot.Data) + len(panel.Plot.Func)
	strPlot := make([]string, numPlot, numPlot)
	for i, ann := range panel.Annotation.Labels {
		ann.SetID(i + 1) // ID must be larger than zero
		s += ann.String()
	}
	for i, ann := range panel.Annotation.Arrows {
		ann.SetID(i + 1) // ID must be larger than zero
		s += ann.String()
	}
	for i, data := range panel.Plot.Data {
		strPlot[i] = data.String()
	}
	num := len(panel.Plot.Data)
	for i, data := range panel.Plot.Func {
		strPlot[num+i] = data.String()
	}
	s += `plot `
	s += strings.Join(strPlot, `,`)
	for _, ann := range panel.Annotation.Labels {
		s += ann.Clear()
	}
	for _, ann := range panel.Annotation.Arrows {
		s += ann.Clear()
	}
	return s
}

func (panel *Panel) SetOption(opt *PanelOption) {
	panel.Opt = opt
}

func (panel *Panel) SetXAxis(axis *Axis) {
	panel.Xaxis = axis
}

func (panel *Panel) SetYAxis(axis *Axis) {
	panel.Yaxis = axis
}

func (panel *Panel) AddData(data *PanelData) {
	panel.Plot.Data = append(panel.Plot.Data, data)
}

func (panel *Panel) AddFunc(data *PanelFunction) {
	panel.Plot.Func = append(panel.Plot.Func, data)
}

func (panel *Panel) AddAnnotationLabel(ann *AnnotationLabel) {
	panel.Annotation.Labels = append(panel.Annotation.Labels, ann)
}

func (panel *Panel) AddAnnotationArrow(ann *AnnotationArrow) {
	panel.Annotation.Arrows = append(panel.Annotation.Arrows, ann)
}
