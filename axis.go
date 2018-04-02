package goplot

import (
	"fmt"
)

/* type Axis */
type AxisFormat struct {
	Format string `xml:",chardata"`
	Tics float64  `xml:"tics,attr"`
	Mtics float64 `xml:"mtics,attr"`
}

type AxisLabel struct {
	Text string `xml:"text"`
	Offset Location `xml:"offset"`
}

type Axis struct {
	Name        string `xml:"name,attr"`
	Coord       string
	Min         float64 `xml:"min,attr"`
	Max         float64 `xml:"max,attr"`
	Log         bool `xml:"log,attr"`
	ShowLabel   bool `xml:"show,attr"`
	Format      AxisFormat
	Label       AxisLabel
}

func NewAxis(coord string) *Axis {
	return &Axis{
		Coord:        coord,
		Min:         0,
		Max:         10,
		Log:         false,
		ShowLabel:   true,
		Format:      AxisFormat{
			Tics: 10,
			Mtics: 2,
			Format: `%g`,
		},
		Label:       AxisLabel{
			Text: "label",
			Offset: NewLocation(`character`, 0., 0.),
		},
	}
}

func NewAxisFromMap(name string, values map[string]interface{}) (*Axis, error) {
	axis := NewAxis(name)
	if values != nil {
		for key, value := range values {
			if err := axis.Set(key, value); err != nil {
				return nil, err
			}
		}
	}
	return axis, nil
}

func (axis *Axis) Show() {
	axis.ShowLabel = true
}

func (axis *Axis) Hide() {
	axis.ShowLabel = false
}

func (axis *Axis) Set(key string, value interface{}) error {
	switch key {
	case `coord`:
		axis.Coord = value.(string)
	case `min`:
		axis.Min = value.(float64)
	case `max`:
		axis.Max = value.(float64)
	case `tics`:
		axis.Format.Tics = value.(float64)
	case `mtics`:
		axis.Format.Mtics = value.(float64)
	case `label`:
		axis.Label.Text = value.(string)
	case `labelOffset`, `labeloffset`:
		axis.Label.Offset = value.(Location)
	case `log`:
		axis.Log = value.(bool)
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
		axis.Coord, axis.Min, axis.Max,
		axis.Coord, axis.Format.Tics,
		axis.Coord, axis.Format.Mtics,
	)
	if axis.ShowLabel {
		s += fmt.Sprintf(`set format %s "%s"
set %slabel "%s" offset %s
`,
			axis.Coord, axis.Format.Format,
			axis.Coord, axis.Label.Text, axis.Label.Offset,
		)
	} else {
		s += fmt.Sprintf("set format %s \"\"\nset %slabel \"\"\n", axis.Coord, axis.Coord)
	}
	if axis.Log {
		s += fmt.Sprintf("set log %s\n", axis.Coord)
	}
	return s
}

func (axis *Axis) Copy() *Axis {
	axis2 := &Axis{}
	*axis2 = *axis
	return axis2
}
