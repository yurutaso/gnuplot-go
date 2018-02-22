package goplot

import (
	"fmt"
)

/* type PanelDataOption */
type PanelDataOption struct {
	isFunc    bool
	using     string
	index     int
	with      string
	lineStyle *LineStyle
	title     string
}

func NewPanelDataOption() *PanelDataOption {
	ls := NewLineStyle()
	opt := &PanelDataOption{
		isFunc:    false,
		using:     "1:2",
		index:     0,
		with:      "line",
		lineStyle: ls,
		title:     "",
	}
	return opt
}

func NewPanelDataOptionFromMap(values map[string]interface{}) (*PanelDataOption, error) {
	opt := NewPanelDataOption()
	if values != nil {
		for key, value := range values {
			if err := opt.Set(key, value); err != nil {
				return nil, err
			}
		}
	}
	return opt, nil
}

func (opt *PanelDataOption) Set(key string, value interface{}) error {
	switch key {
	case `isfunc`, `isFunc`:
		opt.isFunc = value.(bool)
	case `using`, `u`:
		opt.using = value.(string)
	case `index`, `ind`:
		opt.index = value.(int)
	case `with`, `w`:
		opt.with = value.(string)
	case `LineStyle`, `lineStyle`, `linestyle`, `ls`:
		opt.lineStyle = value.(*LineStyle)
	case `title`:
		opt.title = value.(string)
	default:
		return fmt.Errorf(`Unknown key %v`, key)
	}
	return nil
}

func (opt *PanelDataOption) String() string {
	return fmt.Sprintf("using %s index %d with %s title \"%s\" %s\n", opt.using, opt.index, opt.with, opt.title, opt.lineStyle)
}

func (opt *PanelDataOption) Copy() *PanelDataOption {
	opt2 := &PanelDataOption{}
	*opt2 = *opt
	return opt2
}
