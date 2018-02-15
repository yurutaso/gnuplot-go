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

func NewLineStyle() *LineStyle {
	return &LineStyle{
		lineWidth: 1,
		lineType:  1,
		lineColor: "#000000",
		pointType: 1,
		pointSize: 1,
		dashType:  "",
	}
}

func (ls *LineStyle) String() string {
	return fmt.Sprintf(`lw %d lt %d lc rgb "%s" pt %d ps %d dt "%s"`,
		ls.lineWidth, ls.lineType, ls.lineColor, ls.pointType, ls.pointSize, ls.dashType)
}

func (ls *LineStyle) SetLineWidth(lw int) {
	ls.lineWidth = lw
}

func (ls *LineStyle) SetLineType(lt int) {
	ls.lineType = lt
}

func (ls *LineStyle) SetLineColor(lc string) {
	ls.lineColor = lc
}

func (ls *LineStyle) SetPointType(pt int) {
	ls.pointType = pt
}

func (ls *LineStyle) SetPointSize(ps int) {
	ls.pointSize = ps
}

func (ls *LineStyle) SetDashType(dt string) {
	ls.dashType = dt
}
