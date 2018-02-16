package goplot

import (
	"fmt"
)

/* type PanelData */
type PanelData struct {
	name   string
	option *DataOption
}

func NewPanelData(name string, opt *DataOption) (*PanelData, error) {
	var err error
	if opt == nil {
		opt, err = NewDataOption(nil)
		if err != nil {
			return nil, err
		}
	}
	return &PanelData{
		name:   name,
		option: opt,
	}, nil
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
