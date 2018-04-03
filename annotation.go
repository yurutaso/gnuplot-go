package goplot

import (
	"fmt"
)

type AnnotationLabel struct {
	Name     string    `xml:"name,attr"`
	ID       int       `xml:"id,attr"` // No need to set ID by yourself
	Label    *string   `xml:"text"`
	Font     *Font     `xml:"font"`
	Location *Location `xml:"location"`
}

func (ann *AnnotationLabel) String() string {
	return fmt.Sprintf("set label %d \"%s\" at %s font \"%s,%d\"\n", ann.ID, *ann.Label, ann.Location, ann.Font.Name, ann.Font.Size)
}

func (ann *AnnotationLabel) Clear() string {
	return fmt.Sprintf("unset label %d\n", ann.ID)
}

func (ann *AnnotationLabel) SetID(num int) {
	ann.ID = num
}

func (ann *AnnotationLabel) Apply(new *AnnotationLabel) {
	if ann.Label == nil {
		ann.Label = new.Label
	}
	if ann.Location == nil {
		ann.Location = new.Location
	}
	if ann.Font == nil {
		ann.Font = new.Font
	}
}

func NewAnnotationLabel(label string, loc *Location) *AnnotationLabel {
	s := label
	return &AnnotationLabel{
		Label:    &s,
		Location: loc,
		Font:     &Font{Name: "Helvetica", Size: 12},
	}
}

type AnnotationArrow struct {
	Name      string `xml:"name,attr"`
	ID        int    `xml:"id,attr"` // No need to set ID by yourself
	HideHead  *bool  `xml:"hideHead,attr"`
	LineStyle *LineStyle
	Location1 *Location `xml:"from"`
	Location2 *Location `xml:"to"`
}

func NewAnnotationArrow(loc1, loc2 *Location, ls *LineStyle) *AnnotationArrow {
	hide := false
	return &AnnotationArrow{
		HideHead:  &hide,
		Location1: loc1,
		Location2: loc2,
		LineStyle: ls,
	}
}

func (ann *AnnotationArrow) Apply(new *AnnotationArrow) {
	if ann.HideHead == nil {
		ann.HideHead = new.HideHead
	}
	if ann.LineStyle == nil {
		ann.LineStyle = new.LineStyle
	}
	if ann.Location1 == nil {
		ann.Location1 = new.Location1
	}
	if ann.Location2 == nil {
		ann.Location2 = new.Location2
	}
}

func (ann *AnnotationArrow) String() string {
	sls := ""
	if ann.LineStyle != nil {
		sls = ann.LineStyle.String()
	}

	if !*ann.HideHead {
		return fmt.Sprintf("set arrow %d from %s to %s %s\n", ann.ID, ann.Location1, ann.Location2, sls)
	}
	return fmt.Sprintf("set arrow %d from %s to %s nohead %s\n", ann.ID, ann.Location1, ann.Location2, sls)
}

func (ann *AnnotationArrow) NoHead() {
	*ann.HideHead = true
}

func (ann *AnnotationArrow) ShowHead() {
	*ann.HideHead = false
}

func (ann *AnnotationArrow) Clear() string {
	return fmt.Sprintf("unset arrow %d\n", ann.ID)
}

func (ann *AnnotationArrow) SetID(num int) {
	ann.ID = num
}
