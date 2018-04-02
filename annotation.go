package goplot

import (
	"fmt"
)

type AnnotationLabel struct {
	Name    string `xml:"name,attr"`
	ID      int `xml:"id,attr"` // No need to set ID by yourself
	Label    string `xml:"text"`
	Location Location `xml:"location"`
}

func (ann *AnnotationLabel) String() string {
	return fmt.Sprintf("set label %d \"%s\" at %s\n", ann.ID, ann.Label, ann.Location)
}

func (ann *AnnotationLabel) Clear() string {
	return fmt.Sprintf("unset label %d\n", ann.ID)
}

func (ann *AnnotationLabel) SetID(num int) {
	ann.ID = num
}

func NewAnnotationLabel(num int, label string, loc Location) *AnnotationLabel {
	return &AnnotationLabel{
		Label:    label,
		Location: loc,
	}
}

type AnnotationArrow struct {
	Name     string `xml:"name,attr"`
	ID       int `xml:"id,attr"` // No need to set ID by yourself
	HideHead  bool `xml:"hideHead,attr"`
	LineStyle *LineStyle `xml:"lineStyle"`
	Location1 Location `xml:"from"`
	Location2 Location `xml:"to"`
}

func NewAnnotationArrow(num int, loc1, loc2 Location, ls *LineStyle) *AnnotationArrow {
	return &AnnotationArrow{
		HideHead:  false,
		Location1: loc1,
		Location2: loc2,
		LineStyle: ls,
	}
}

func (ann *AnnotationArrow) String() string {
	sls := ""
	if ann.LineStyle != nil {
		sls = ann.LineStyle.String()
	}

	if !ann.HideHead {
		return fmt.Sprintf("set arrow %d from %s to %s %s\n", ann.ID, ann.Location1, ann.Location2, sls)
	}
	return fmt.Sprintf("set arrow %d from %s to %s nohead %s\n", ann.ID, ann.Location1, ann.Location2, sls)
}

func (ann *AnnotationArrow) NoHead() {
	ann.HideHead = true
}

func (ann *AnnotationArrow) ShowHead() {
	ann.HideHead = false
}

func (ann *AnnotationArrow) Clear() string {
	return fmt.Sprintf("unset arrow %d\n", ann.ID)
}

func (ann *AnnotationArrow) SetID(num int) {
	ann.ID = num
}
