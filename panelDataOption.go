package goplot

import (
	"fmt"
)

/* type PanelDataOption */
type PanelDataOption struct {
	IsFunc    bool `xml:"isFunc"`
	Using     string `xml:"using"`
	Index     int `xml:"index"`
	With      string `xml:"with"`
	LineStyle *LineStyle `xml:"lineStyle"`
	Title     string `xml:"title"`
}

func NewPanelDataOption() *PanelDataOption {
	ls := NewLineStyle()
	opt := &PanelDataOption{
		IsFunc:    false,
		Using:     "1:2",
		Index:     0,
		With:      "line",
		LineStyle: ls,
		Title:     "",
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
		opt.IsFunc = value.(bool)
	case `using`, `u`:
		opt.Using = value.(string)
	case `index`, `ind`:
		opt.Index = value.(int)
	case `with`, `w`:
		opt.With = value.(string)
	case `LineStyle`, `lineStyle`, `linestyle`, `ls`:
		opt.LineStyle = value.(*LineStyle)
	case `title`:
		opt.Title = value.(string)
	default:
		return fmt.Errorf(`Unknown key %v`, key)
	}
	return nil
}

func (opt *PanelDataOption) String() string {
	return fmt.Sprintf("using %s index %d with %s title \"%s\" %s\n", opt.Using, opt.Index, opt.With, opt.Title, opt.LineStyle)
}

func (opt *PanelDataOption) Copy() *PanelDataOption {
	opt2 := &PanelDataOption{}
	*opt2 = *opt
	return opt2
}
