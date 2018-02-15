package goplot

import (
	"fmt"
)

/* type PanelData */
type PanelData struct {
	name   string
	option *DataOption
}

func NewPanelData(name string, option *DataOption) *PanelData {
	if option == nil {
		option = NewDataOption()
	}
	return &PanelData{
		name:   name,
		option: option,
	}
}

func (data *PanelData) String() string {
	if data.option.isFunc {
		return fmt.Sprintf(`%s %s`, data.name, data.option)
	}
	return fmt.Sprintf(`"%s" %s`, data.name, data.option)
}

func (data *PanelData) SetData(name string) {
	data.name = name
}

func (data *PanelData) SetOption(opt *DataOption) {
	data.option = opt
}
