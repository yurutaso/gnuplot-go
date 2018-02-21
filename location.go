package goplot

import (
	"fmt"
)

type Location struct {
	mode string
	x    float64
	y    float64
}

func (loc *Location) String() string {
	return fmt.Sprintf("%s %f, %f", mode, x, y)
}

func NewLocation(mode string, x, y float64) *Location {
	return &Location{
		mode: mode,
		x:    x,
		y:    y,
	}
}

func NewLocationAbs(x, y float64) *Location {
	return &Location{
		mode: `graph`,
		x:    x,
		y:    y,
	}
}

func NewLocationRelative(x, y float64) *Location {
	return &Location{
		mode: `panel`,
		x:    x,
		y:    y,
	}
}
