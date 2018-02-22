package goplot

import (
	"strings"
)

/* type Panel */
type Panel struct {
	Data       []*PanelData
	Opt        *PanelOption
	Annotation []PanelAnnotation
}

func NewPanel(opt *PanelOption) *Panel {
	if opt == nil {
		opt = NewPanelOption()
	}
	return &Panel{
		Data: make([]*PanelData, 0, 0),
		Opt:  opt,
	}
}

func (panel *Panel) String() string {
	s := panel.Opt.String()
	strs := make([]string, len(panel.Data), len(panel.Data))
	for i, ann := range panel.Annotation {
		ann.SetID(i + 1) // ID must be larger than zero
		s += ann.String()
	}
	for i, data := range panel.Data {
		strs[i] = data.String()
	}
	s += `plot `
	s += strings.Join(strs, `,`)
	for _, ann := range panel.Annotation {
		s += ann.Clear()
	}
	return s
}

func (panel *Panel) SetOption(opt *PanelOption) {
	panel.Opt = opt
}

func (panel *Panel) AddData(data *PanelData) {
	panel.Data = append(panel.Data, data)
}

func (panel *Panel) AddAnnotation(ann PanelAnnotation) {
	panel.Annotation = append(panel.Annotation, ann)
}
