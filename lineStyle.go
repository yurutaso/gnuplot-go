package goplot

import (
	"fmt"
)

/* type LineStyle */
type LineStyle struct {
	LineWidth int `xml:"lw,attr"`
	LineType  int `xml:"lt,attr"`
	LineColor string `xml:"lc,attr"`
	PointType int `xml:"pt,attr"`
	PointSize int `xml:"ps,attr"`
	DashType  string `xml:"dt,attr"`
}

func NewLineStyle() *LineStyle {
	return &LineStyle{
		LineWidth: 1,
		LineType:  1,
		LineColor: "#000000",
		PointType: 1,
		PointSize: 1,
		DashType:  "",
	}
}

func NewLineStyleFromMap(values map[string]interface{}) (*LineStyle, error) {
	ls := NewLineStyle()
	if values != nil {
		for key, value := range values {
			if err := ls.Set(key, value); err != nil {
				return nil, err
			}
		}
	}
	return ls, nil
}

func (ls *LineStyle) String() string {
	return fmt.Sprintf(`lw %d lt %d lc rgb "%s" pt %d ps %d dt "%s"`,
		ls.LineWidth, ls.LineType, ls.LineColor, ls.PointType, ls.PointSize, ls.DashType)
}

func (ls *LineStyle) Set(key string, value interface{}) error {
	switch key {
	case `LineWidth`, `lineWidth`, `linewidth`, `lw`:
		ls.LineWidth = value.(int)
	case `LineType`, `lineType`, `linetype`, `lt`:
		ls.LineType = value.(int)
	case `LineColor`, `lineColor`, `linecolor`, `lc`:
		ls.LineColor = value.(string)
	case `PointType`, `pointType`, `pointtype`, `pt`:
		ls.PointType = value.(int)
	case `PointSize`, `pointSize`, `pointsize`, `ps`:
		ls.PointSize = value.(int)
	case `DashType`, `dashType`, `dashtype`, `dt`:
		ls.DashType = value.(string)
	default:
		return fmt.Errorf("Unknown key %v", value)
	}
	return nil
}
