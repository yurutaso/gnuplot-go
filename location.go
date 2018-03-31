package goplot

import (
	"fmt"
)

type Coordinate struct {
	System string `xml:"system,attr"`
	Value float64 `xml:",chardata"`
}

func (coord Coordinate) String() string {
	return fmt.Sprintf("%s %f", coord.System, coord.Value)
}

type Location struct {
	X Coordinate `xml:"x"`
	Y Coordinate `xml:"y"`
	Z Coordinate `xml:"z"`
}

func (loc Location) String() string {
	if len(loc.Z.System) == 0 {
		return fmt.Sprintf("%s, %s", loc.X, loc.Y)
	}
	return fmt.Sprintf("%s, %s, %s", loc.X, loc.Y, loc.Z)
}

func NewLocation(system string, x, y float64) Location {
	return Location{
		X: Coordinate{System: system, Value: x},
		Y: Coordinate{System: system, Value: y},
	}
}

func NewLocationXYZ(system string, x, y, z float64) Location {
	return Location{
		X: Coordinate{System: system, Value: x},
		Y: Coordinate{System: system, Value: y},
		Z: Coordinate{System: system, Value: z},
	}
}

func (loc *Location) SetX(system string, x float64) error {
	if checkSystem(system) {
		return fmt.Errorf("Unknown system %s\n", system)
	}
	loc.X.System = system
	loc.X.Value = x
	return nil
}

func (loc *Location) SetY(system string, y float64) error {
	if checkSystem(system) {
		return fmt.Errorf("Unknown system %s\n", system)
	}
	loc.Y.System = system
	loc.Y.Value = y
	return nil
}

func (loc *Location) SetZ(system string, z float64) error {
	if checkSystem(system) {
		return fmt.Errorf("Unknown system %s\n", system)
	}
	loc.Z.System = system
	loc.Z.Value = z
	return nil
}

func checkSystem(system string) bool {
	switch system {
	case `first`, `second`, `graph`, `screen`, `character`:
		return true
	default:
		return false
	}
}
