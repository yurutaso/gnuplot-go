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
	ID      int `xml:"id,attr"` // No need to set ID by yourself
	Label    string `xml:"text"`
	Location Location `xml:"location"`
}

func (ann *PanelAnnotationLabel) String() string {
	return fmt.Sprintf("set label %d \"%s\" at %s\n", ann.ID, ann.Label, ann.Location)
}

func (ann *PanelAnnotationLabel) Clear() string {
	return fmt.Sprintf("unset label %d\n", ann.ID)
}

func (ann *PanelAnnotationLabel) SetID(num int) {
	ann.ID = num
}

func NewAnnotationLabel(num int, label string, loc Location) *PanelAnnotationLabel {
	return &PanelAnnotationLabel{
		Label:    label,
		Location: loc,
	}
}

type PanelAnnotationArrow struct {
	ID       int `xml:"id,attr"` // No need to set ID by yourself
	HideHead  bool `xml:"hideHead,attr"`
	LineStyle *LineStyle `xml:"lineStyle"`
	Location1 Location `xml:"from"`
	Location2 Location `xml:"to"`
}

func NewAnnotationArrow(num int, loc1, loc2 Location, ls *LineStyle) *PanelAnnotationArrow {
	return &PanelAnnotationArrow{
		HideHead:  false,
		Location1: loc1,
		Location2: loc2,
		LineStyle: ls,
	}
}

func (ann *PanelAnnotationArrow) String() string {
	sls := ""
	if ann.LineStyle != nil {
		sls = ann.LineStyle.String()
	}

	if !ann.HideHead {
		return fmt.Sprintf("set arrow %d from %s to %s %s\n", ann.ID, ann.Location1, ann.Location2, sls)
	}
	return fmt.Sprintf("set arrow %d from %s to %s nohead %s\n", ann.ID, ann.Location1, ann.Location2, sls)
}

func (ann *PanelAnnotationArrow) NoHead() {
	ann.HideHead = true
}

func (ann *PanelAnnotationArrow) ShowHead() {
	ann.HideHead = false
}

func (ann *PanelAnnotationArrow) Clear() string {
	return fmt.Sprintf("unset arrow %d\n", ann.ID)
}

func (ann *PanelAnnotationArrow) SetID(num int) {
	ann.ID = num
}
