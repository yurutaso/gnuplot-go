package goplot

import (
	"fmt"
)

/* type font and fontconfig */
type Font struct {
	name string
	size int
}

func (font *Font) String() string {
	return fmt.Sprintf(`font "%s,%d"`, font.name, font.size)
}

type FontConfig map[string]*Font

func NewFontConfig() FontConfig {
	return FontConfig{
		`xlabel`: &Font{name: `Helvetica`, size: 16},
		`ylabel`: &Font{name: `Helvetica`, size: 16},
		`title`:  &Font{name: `Helvetica`, size: 16},
		`key`:    &Font{name: `Helvetica`, size: 12},
		`xtics`:  &Font{name: `Helvetica`, size: 12},
		`ytics`:  &Font{name: `Helvetica`, size: 12},
	}
}

func (conf FontConfig) Set(key string, font *Font) error {
	switch key {
	case `xlabel`, `ylabel`, `title`, `key`, `xtics`, `ytics`:
		conf[key] = font
	default:
		return fmt.Errorf(`Invalid key %s`, key)
	}
	return nil
}

func (conf FontConfig) String() string {
	return fmt.Sprintf(`set xlabel %s
set ylabel %s
set title %s
set key %s
set xtics %s
set ytics %s`,
		conf[`xlabel`],
		conf[`ylabel`],
		conf[`title`],
		conf[`key`],
		conf[`xtics`],
		conf[`ytics`],
	)
}
