package goplot

import (
	"strings"
)

/* type Panel */
type Panel struct {
	data   []*PanelData
	option *PanelOption
}

func NewPanel() *Panel {
	return &Panel{
		data:   make([]*PanelData, 0, 0),
		option: NewPanelOption(),
	}
}

func (panel *Panel) String() string {
	s := panel.option.String()
	strs := make([]string, len(panel.data), len(panel.data))
	for i, data := range panel.data {
		strs[i] = data.String()
	}
	s += `plot `
	s += strings.Join(strs, `,`)
	return s
}

func (panel *Panel) SetOption(opt *PanelOption) {
	panel.option = opt
}

func (panel *Panel) AddData(data *PanelData) {
	panel.data = append(panel.data, data)
}
