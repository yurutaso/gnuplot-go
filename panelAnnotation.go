package goplot

import (
	"fmt"
)

type PanelAnnotation struct {
	num      int
	name     string
	arg      string
	location *Location
}

func (ann *PanelAnnotation) String() string {
	return fmt.Sprintf("set %s %d %s at %s\n", name, num, arg, location)
}

func NewAnnLabel(num int, label string, loc *Location) *PanelAnnotation {
	return &PanelAnnotation{
		num:      num,
		name:     `label`,
		arg:      label,
		location: loc,
	}
}

func NewAnnLine(num int, loc1, loc2 *Location, ls *LineStyle) *PanelAnnotation {
	return &PanelAnnotation{
		num:      num,
		name:     `arrow`,
		location: loc,
	}
}
