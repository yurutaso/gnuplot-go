package goplot

import (
	"fmt"
)

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

func NewAxisOption(name string, values map[string]interface{}) (*AxisOption, error) {
	axis := &AxisOption{
		name:        name,
		min:         0,
		max:         10,
		tics:        10,
		mtics:       2,
		label:       "label",
		labelOffset: 0,
		log:         false,
	}
	if values != nil {
		for key, value := range values {
			if err := axis.Set(key, value); err != nil {
				return nil, err
			}
		}
	}
	return axis, nil
}

func (axis *AxisOption) Set(key string, value interface{}) error {
	switch key {
	case `name`:
		axis.name = value.(string)
	case `min`:
		axis.min = value.(float64)
	case `max`:
		axis.max = value.(float64)
	case `tics`:
		axis.tics = value.(float64)
	case `mtics`:
		axis.mtics = value.(float64)
	case `label`:
		axis.label = value.(string)
	case `labelOffset`, `labeloffset`:
		axis.labelOffset = value.(float64)
	case `log`:
		axis.log = value.(bool)
	default:
		return fmt.Errorf(`Unknow key %v`, key)
	}
	return nil
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

func NewPanelOption(values map[string]interface{}) (*PanelOption, error) {
	xaxis, err := NewAxisOption(`x`, nil)
	if err != nil {
		return nil, err
	}
	yaxis, err := NewAxisOption(`y`, nil)
	if err != nil {
		return nil, err
	}
	opt := &PanelOption{
		xaxis: xaxis,
		yaxis: yaxis,
		//zaxis:  NewAxisOption(),
		sample: 1000,
		grid:   "",
		key:    "",
	}
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
		opt.xaxis = value.(*AxisOption)
	case `yaxis`:
		opt.yaxis = value.(*AxisOption)
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
`, opt.xaxis.String(), opt.yaxis.String(), opt.sample, opt.grid, opt.key)
	if len(opt.grid) == 0 {
		s += "\nunset grid\n"
	}
	return s
}
