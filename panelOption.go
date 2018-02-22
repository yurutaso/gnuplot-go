package goplot

import (
	"fmt"
)

/* type PanelOption */
type PanelOption struct {
	Xaxis     *Axis
	Yaxis     *Axis
	showXaxis bool
	showYaxis bool
	//zaxis  *Axis
	sample int
	grid   string
	key    string
}

func NewPanelOption() *PanelOption {
	return &PanelOption{
		Xaxis: NewAxis(`x`),
		Yaxis: NewAxis(`y`),
		//Zaxis:  NewAxis(`z`),
		sample: 1000,
		grid:   "",
		key:    "",
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
		opt.grid = value.(string)
	case `sample`:
		opt.sample = value.(int)
	case `key`:
		opt.key = value.(string)
	case `xaxis`:
		opt.Xaxis = value.(*Axis)
	case `yaxis`:
		opt.Yaxis = value.(*Axis)
	default:
		return fmt.Errorf(`Unknown key %v`, key)
	}
	return nil
}

func (opt *PanelOption) String() string {
	s := fmt.Sprintf(`
%s
%s
set sample %d
set grid %s
set key %s
`, opt.Xaxis, opt.Yaxis, opt.sample, opt.grid, opt.key)
	if len(opt.grid) == 0 {
		s += "unset grid\n"
	}
	return s
}

func (opt *PanelOption) Copy() *PanelOption {
	opt2 := &PanelOption{}
	*opt2 = *opt
	opt2.Set(`xaxis`, opt.Xaxis.Copy())
	opt2.Set(`yaxis`, opt.Yaxis.Copy())
	return opt2
}
