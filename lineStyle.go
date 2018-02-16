package goplot

import (
	"fmt"
)

/* type LineStyle */
type LineStyle struct {
	lineWidth int
	lineType  int
	lineColor string
	pointType int
	pointSize int
	dashType  string
}

func NewLineStyle(values map[string]interface{}) (*LineStyle, error) {
	ls := &LineStyle{
		lineWidth: 1,
		lineType:  1,
		lineColor: "#000000",
		pointType: 1,
		pointSize: 1,
		dashType:  "",
	}
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
		ls.lineWidth, ls.lineType, ls.lineColor, ls.pointType, ls.pointSize, ls.dashType)
}

func (ls *LineStyle) Set(key string, value interface{}) error {
	switch key {
	case `LineWidth`, `lineWidth`, `linewidth`, `lw`:
		ls.lineWidth = value.(int)
	case `LineType`, `lineType`, `linetype`, `lt`:
		ls.lineType = value.(int)
	case `LineColor`, `lineColor`, `linecolor`, `lc`:
		ls.lineColor = value.(string)
	case `PointType`, `pointType`, `pointtype`, `pt`:
		ls.pointType = value.(int)
	case `PointSize`, `pointSize`, `pointsize`, `ps`:
		ls.pointSize = value.(int)
	case `DashType`, `dashType`, `dashtype`, `dt`:
		ls.dashType = value.(string)
	default:
		return fmt.Errorf("Unknown key %v", value)
	}
	return nil
}
