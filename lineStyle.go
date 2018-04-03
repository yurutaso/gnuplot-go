package goplot

import (
	"fmt"
	"strings"
)

/* type LineStyle */
type LineStyle struct {
	Name      string  `xml:"name,attr"`
	LineWidth *int    `xml:"lw,attr"`
	LineType  *int    `xml:"lt,attr"`
	LineColor *string `xml:"lc,attr"`
	PointType *int    `xml:"pt,attr"`
	PointSize *int    `xml:"ps,attr"`
	DashType  *string `xml:"dt,attr"`
}

func NewLineStyle() *LineStyle {
	lw := 1
	lt := 1
	lc := "#000000"
	pt := 1
	ps := 1
	dt := ""
	return &LineStyle{
		LineWidth: &lw,
		LineType:  &lt,
		LineColor: &lc,
		PointType: &pt,
		PointSize: &ps,
		DashType:  &dt,
	}
}

func (ls *LineStyle) Apply(new *LineStyle) {
	if ls.LineWidth == nil {
		ls.LineWidth = new.LineWidth
	}
	if ls.LineType == nil {
		ls.LineType = new.LineType
	}
	if ls.LineColor == nil {
		ls.LineColor = new.LineColor
	}
	if ls.PointType == nil {
		ls.PointType = new.PointType
	}
	if ls.PointSize == nil {
		ls.PointSize = new.PointSize
	}
	if ls.DashType == nil {
		ls.DashType = new.DashType
	}
}

func (ls *LineStyle) String() string {
	return fmt.Sprintf(`lw %d lt %d lc rgb "%s" pt %d ps %d dt "%s"`,
		*ls.LineWidth, *ls.LineType, *ls.LineColor, *ls.PointType, *ls.PointSize, *ls.DashType)
}

func (ls *LineStyle) Set(key string, value interface{}) error {
	switch strings.ToLower(key) {
	case `linewidth`, `lw`:
		ls.LineWidth = value.(*int)
	case `linetype`, `lt`:
		ls.LineType = value.(*int)
	case `linecolor`, `lc`:
		ls.LineColor = value.(*string)
	case `pointtype`, `pt`:
		ls.PointType = value.(*int)
	case `pointsize`, `ps`:
		ls.PointSize = value.(*int)
	case `dashtype`, `dt`:
		ls.DashType = value.(*string)
	default:
		return fmt.Errorf("Unknown key %v", value)
	}
	return nil
}
