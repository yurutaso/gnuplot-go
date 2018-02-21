package goplot

import (
	"strings"
)

/* type Panel */
type Panel struct {
	Data       []*PanelData
	Opt        *PanelOption
	Annotation []*PanelAnnotation
}

func NewPanel(opt *PanelOption) (*Panel, error) {
	var err error
	if opt == nil {
		opt, err = NewPanelOption(nil)
		if err != nil {
			return nil, err
		}
	}
	return &Panel{
		Data: make([]*PanelData, 0, 0),
		Opt:  opt,
	}, nil
}

func (panel *Panel) String() string {
	s := panel.Opt.String()
	strs := make([]string, len(panel.Data), len(panel.Data))
	for i, data := range panel.Data {
		strs[i] = data.String()
	}
	s += `plot `
	s += strings.Join(strs, `,`)
	return s
}

func (panel *Panel) SetOption(opt *PanelOption) {
	panel.Opt = opt
}

func (panel *Panel) AddData(data *PanelData) {
	panel.Data = append(panel.Data, data)
}
