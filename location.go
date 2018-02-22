package goplot

import (
	"fmt"
)

type Location struct {
	systemX string
	systemY string
	systemZ string
	x       float64
	y       float64
	z       float64
}

func (loc *Location) String() string {
	if len(loc.systemZ) == 0 {
		return fmt.Sprintf("%s %f, %s %f", loc.systemX, loc.x, loc.systemY, loc.y)
	}
	return fmt.Sprintf("%s %f, %s %f, %s %f", loc.systemX, loc.x, loc.systemY, loc.y, loc.systemZ, loc.z)
}

func NewLocation(system string, x, y float64) *Location {
	return &Location{
		systemX: system,
		systemY: system,
		x:       x,
		y:       y,
	}
}

func NewLocationXYZ(system string, x, y, z float64) *Location {
	return &Location{
		systemX: system,
		systemY: system,
		systemZ: system,
		x:       x,
		y:       y,
		z:       z,
	}
}

func (loc *Location) SetX(system string, x float64) error {
	if checkSystem(system) {
		return fmt.Errorf("Unknown system %s\n", system)
	}
	loc.systemX = system
	loc.x = x
	return nil
}

func (loc *Location) SetY(system string, y float64) error {
	if checkSystem(system) {
		return fmt.Errorf("Unknown system %s\n", system)
	}
	loc.systemY = system
	loc.y = y
	return nil
}

func (loc *Location) SetZ(system string, z float64) error {
	if checkSystem(system) {
		return fmt.Errorf("Unknown system %s\n", system)
	}
	loc.systemZ = system
	loc.z = z
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
