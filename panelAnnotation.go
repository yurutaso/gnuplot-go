package goplot

import (
	"fmt"
)

type PanelAnnotation interface {
	String() string
	Clear() string
	SetID(int)
}

type PanelAnnotationLabel struct {
	num      int
	label    string
	location *Location
}

func (ann *PanelAnnotationLabel) String() string {
	return fmt.Sprintf("set label %d \"%s\" at %s\n", ann.num, ann.label, ann.location)
}

func (ann *PanelAnnotationLabel) Clear() string {
	return fmt.Sprintf("unset label %d\n", ann.num)
}

func (ann *PanelAnnotationLabel) SetID(num int) {
	ann.num = num
}

func NewAnnotationLabel(label string, loc *Location) *PanelAnnotationLabel {
	return &PanelAnnotationLabel{
		label:    label,
		location: loc,
	}
}

type PanelAnnotationArrow struct {
	num       int
	showHead  bool
	lineStyle *LineStyle
	location1 *Location
	location2 *Location
}

func NewAnnotationArrow(loc1, loc2 *Location, ls *LineStyle) *PanelAnnotationArrow {
	return &PanelAnnotationArrow{
		showHead:  true,
		location1: loc1,
		location2: loc2,
		lineStyle: ls,
	}
}

func (ann *PanelAnnotationArrow) String() string {
	sls := ""
	if ann.lineStyle != nil {
		sls = ann.lineStyle.String()
	}

	if ann.showHead {
		return fmt.Sprintf("set arrow %d from %s to %s %s\n", ann.num, ann.location1, ann.location2, sls)
	}
	return fmt.Sprintf("set arrow %d from %s to %s nohead %s\n", ann.num, ann.location1, ann.location2, sls)
}

func (ann *PanelAnnotationArrow) NoHead() {
	ann.showHead = false
}

func (ann *PanelAnnotationArrow) ShowHead() {
	ann.showHead = true
}

func (ann *PanelAnnotationArrow) Clear() string {
	return fmt.Sprintf("unset arrow %d\n", ann.num)
}

func (ann *PanelAnnotationArrow) SetID(num int) {
	ann.num = num
}
