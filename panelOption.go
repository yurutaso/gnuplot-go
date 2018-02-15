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

func NewAxisOption(name string) *AxisOption {
	return &AxisOption{
		name:        name,
		min:         0,
		max:         10,
		tics:        10,
		mtics:       2,
		label:       "label",
		labelOffset: 0,
		log:         false,
	}
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

func NewPanelOption() *PanelOption {
	return &PanelOption{
		xaxis: NewAxisOption(`x`),
		yaxis: NewAxisOption(`y`),
		//zaxis:  NewAxisOption(),
		sample: 1000,
		grid:   "",
		key:    "",
	}
}

func (option *PanelOption) String() string {
	return fmt.Sprintf(`
%s
%s
set sample %d
set grid %s
set key %s
`,
		option.xaxis.String(), option.yaxis.String(), option.sample, option.grid, option.key)
}
