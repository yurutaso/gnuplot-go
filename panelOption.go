package goplot

import (
	"fmt"
)

/* type PanelOption */
type PanelOption struct {
	Name   string `xml:"name,attr"`
	Sample *int
	Grid   *string
	Key    *string
}

func NewPanelOption() *PanelOption {
	sample := 1000
	grid := ""
	key := ""
	return &PanelOption{
		Sample: &sample,
		Grid:   &grid,
		Key:    &key,
	}
}

func (opt *PanelOption) Apply(new *PanelOption) {
	// Apply new panelOption if the old is not nil
	if opt.Sample == nil {
		opt.Sample = new.Sample
	}
	if opt.Grid == nil {
		opt.Grid = new.Grid
	}
	if opt.Key == nil {
		opt.Key = new.Key
	}
}

func (opt *PanelOption) Set(key string, value interface{}) error {
	switch key {
	case `grid`:
		opt.Grid = value.(*string)
	case `sample`:
		opt.Sample = value.(*int)
	case `key`:
		opt.Key = value.(*string)
	default:
		return fmt.Errorf(`Unknown key %v`, key)
	}
	return nil
}

func (opt *PanelOption) String() string {
	s := fmt.Sprintf("set sample %d\nset grid %s\nset key %s\n", *opt.Sample, *opt.Grid, *opt.Key)
	if len(*opt.Grid) == 0 {
		s += "unset grid\n"
	}
	return s
}

func (opt *PanelOption) Copy() *PanelOption {
	opt2 := &PanelOption{}
	*opt2 = *opt
	return opt2
}
