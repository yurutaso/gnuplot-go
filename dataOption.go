package goplot

import (
	"fmt"
)

/* type DataOption */
type DataOption struct {
	isFunc    bool
	using     string
	index     int
	with      string
	lineStyle *LineStyle
	title     string
}

func NewDataOption() *DataOption {
	return &DataOption{
		isFunc:    false,
		using:     "1:2",
		index:     0,
		with:      "line",
		lineStyle: NewLineStyle(),
		title:     "",
	}
}

func (opt *DataOption) String() string {
	return fmt.Sprintf(`using %s index %d with %s title "%s" %s`, opt.using, opt.index, opt.with, opt.title, opt.lineStyle)
}

func (opt *DataOption) SetIsFunc(isFunc bool) {
	opt.isFunc = isFunc
}

func (opt *DataOption) SetUsing(using string) {
	opt.using = using
}

func (opt *DataOption) SetIndex(index int) {
	opt.index = index
}

func (opt *DataOption) SetWith(with string) {
	opt.with = with
}
func (opt *DataOption) SetLineStyle(lineStyle *LineStyle) {
	opt.lineStyle = lineStyle
}
func (opt *DataOption) SetTitle(title string) {
	opt.title = title
}
