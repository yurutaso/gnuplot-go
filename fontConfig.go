package goplot

import (
	"fmt"
)

/* type font and fontconfig */
type Font struct {
	Field string `xml:",chardata"`
	Name string `xml:"name,attr"`
	Size int `xml:"size,attr"`
}

func (font Font) String() string {
	return fmt.Sprintf(`set %s font "%s,%d"`, font.Field, font.Name, font.Size)
}