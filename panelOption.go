package goplot

import (
	"fmt"
)

/* type PanelOption */
type PanelOption struct {
	Name string `xml:"name,attr"`
	Sample int
	Grid   string
	Key    string
}

func NewPanelOption() *PanelOption {
	return &PanelOption{
		Sample: 1000,
		Grid:   "",
		Key:    "",
	}
}

func NewPanelOptionFromMap(values map[string]interface{}) (*PanelOption, error) {
	opt := NewPanelOption()
	if values != nil {
		for key, value := range values {
			if err := opt.Set(key, value); err != nil {
				return nil, err
			}
		}
	}
	return opt, nil
}

func (opt *PanelOption) Set(key string, value interface{}) error {
	switch key {
	case `grid`:
		opt.Grid = value.(string)
	case `sample`:
		opt.Sample = value.(int)
	case `key`:
		opt.Key = value.(string)
	default:
		return fmt.Errorf(`Unknown key %v`, key)
	}
	return nil
}

func (opt *PanelOption) String() string {
	s := fmt.Sprintf("set sample %d\nset grid %s\nset key %s", opt.Sample, opt.Grid, opt.Key)
	if len(opt.Grid) == 0 {
		s += "unset grid\n"
	}
	return s
}

func (opt *PanelOption) Copy() *PanelOption {
	opt2 := &PanelOption{}
	*opt2 = *opt
	return opt2
}
