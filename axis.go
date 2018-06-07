package goplot

import (
	"fmt"
)

/* type Axis */
type AxisFormat struct {
	Format *string  `xml:",chardata"`
	Tics   *float64 `xml:"tics,attr"`
	Mtics  *float64 `xml:"mtics,attr"`
}

func NewAxisFormat() *AxisFormat {
	format := "%g"
	tics := 10.0
	mtics := 1.0
	return &AxisFormat{
		Format: &format,
		Tics:   &tics,
		Mtics:  &mtics,
	}
}

type AxisLabel struct {
	Text   *string   `xml:"text"`
	Offset *Location `xml:"offset"`
}

func NewAxisLabel() *AxisLabel {
	text := ""
	offset := NewLocation("character", 0., 0.)
	return &AxisLabel{
		Text:   &text,
		Offset: &offset,
	}
}

type Axis struct {
	Name      string `xml:"name,attr"`
	Coord     string
	Min       *float64 `xml:"min,attr"`
	Max       *float64 `xml:"max,attr"`
	Log       *bool    `xml:"log,attr"`
	ShowLabel *bool    `xml:"show,attr"`
	Format    *AxisFormat
	Label     *AxisLabel
}

func NewAxis(coord string) *Axis {
	min := 0.0
	max := 10.0
	logscale := false
	showlabel := true
	return &Axis{
		Coord:     coord,
		Min:       &min,
		Max:       &max,
		Log:       &logscale,
		ShowLabel: &showlabel,
		Format:    NewAxisFormat(),
		Label:     NewAxisLabel(),
	}
}

func (axis *Axis) Apply(new *Axis) {
	axis.Coord = new.Coord
	if axis.Min == nil {
		axis.Min = new.Min
	}
	if axis.Max == nil {
		axis.Max = new.Max
	}
	if axis.Log == nil {
		axis.Log = new.Log
	}
	if axis.ShowLabel == nil {
		axis.ShowLabel = new.ShowLabel
	}
	if axis.Format == nil {
		axis.Format = new.Format
	}
	if axis.Label == nil {
		axis.Label = new.Label
	}
}

func (axis *Axis) Show() {
	*axis.ShowLabel = true
}

func (axis *Axis) Hide() {
	*axis.ShowLabel = false
}

func (axis *Axis) Set(key string, value interface{}) error {
	switch key {
	case `coord`:
		axis.Coord = value.(string)
	case `min`:
		axis.Min = value.(*float64)
	case `max`:
		axis.Max = value.(*float64)
	case `tics`:
		axis.Format.Tics = value.(*float64)
	case `mtics`:
		axis.Format.Mtics = value.(*float64)
	case `label`:
		axis.Label.Text = value.(*string)
	case `labelOffset`, `labeloffset`:
		axis.Label.Offset = value.(*Location)
	case `log`:
		axis.Log = value.(*bool)
	default:
		return fmt.Errorf(`Unknow key %v`, key)
	}
	return nil
}

func (axis *Axis) String() string {
	s := fmt.Sprintf(`
set %srange [%f:%f]
set %stics %f
set m%stics %f
`,
		axis.Coord, *axis.Min, *axis.Max,
		axis.Coord, *axis.Format.Tics,
		axis.Coord, *axis.Format.Mtics,
	)
	if *axis.ShowLabel {
		s += fmt.Sprintf(`set format %s "%s"
set %slabel "%s" offset %s
`,
			axis.Coord, *axis.Format.Format,
			axis.Coord, *axis.Label.Text, axis.Label.Offset,
		)
	} else {
		s += fmt.Sprintf("set format %s \"\"\nset %slabel \"\"\n", axis.Coord, axis.Coord)
	}
	if *axis.Log {
		s += fmt.Sprintf("set log %s\n", axis.Coord)
	} else {
		s += fmt.Sprintf("unset log %s\n", axis.Coord)
	}
	return s
}

func (axis *Axis) Copy() *Axis {
	axis2 := &Axis{}
	*axis2 = *axis
	return axis2
}
