package goplot

import (
	"fmt"
)

/* type Axis */
type Axis struct {
	name        string
	min         float64
	max         float64
	tics        float64
	mtics       float64
	format      string
	label       string
	labelOffset *Location
	log         bool
	show        bool
}

func NewAxis(name string) *Axis {
	return &Axis{
		name:        name,
		min:         0,
		max:         10,
		tics:        10,
		mtics:       2,
		format:      `% g`,
		label:       "label",
		labelOffset: NewLocation(`character`, 0., 0.),
		log:         false,
		show:        true,
	}
}

func NewAxisFromMap(name string, values map[string]interface{}) (*Axis, error) {
	axis := NewAxis(name)
	if values != nil {
		for key, value := range values {
			if err := axis.Set(key, value); err != nil {
				return nil, err
			}
		}
	}
	return axis, nil
}

func (axis *Axis) Show() {
	axis.show = true
}

func (axis *Axis) Hide() {
	axis.show = false
}

func (axis *Axis) Set(key string, value interface{}) error {
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
		axis.labelOffset = value.(*Location)
	case `log`:
		axis.log = value.(bool)
	default:
		return fmt.Errorf(`Unknow key %v`, key)
	}
	return nil
}

func (axis *Axis) String() string {
	s := fmt.Sprintf(`
set %srange [%f:%f]
set %stics %f
set m%stics %f
`,
		axis.name, axis.min, axis.max,
		axis.name, axis.tics,
		axis.name, axis.mtics,
	)
	if axis.show {
		s += fmt.Sprintf(`set format %s "%s"
set %slabel "%s" offset %s
`,
			axis.name, axis.format,
			axis.name, axis.label, axis.labelOffset,
		)
	} else {
		s += fmt.Sprintf("set format %s \"\"\nset %slabel \"\"\n", axis.name, axis.name)
	}
	if axis.log {
		s += fmt.Sprintf("set log %s\n", axis.name)
	}
	return s
}

func (axis *Axis) Copy() *Axis {
	axis2 := &Axis{}
	*axis2 = *axis
	return axis2
}
