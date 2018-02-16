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

func NewFontConfig(values map[string]*Font) (FontConfig, error) {
	if values == nil {
		return FontConfig{
			`xlabel`: &Font{name: `Helvetica`, size: 16},
			`ylabel`: &Font{name: `Helvetica`, size: 16},
			`title`:  &Font{name: `Helvetica`, size: 14},
			`key`:    &Font{name: `Helvetica`, size: 12},
			`xtics`:  &Font{name: `Helvetica`, size: 12},
			`ytics`:  &Font{name: `Helvetica`, size: 12},
		}, nil
	}
	conf := FontConfig{}
	for key, value := range values {
		if err := conf.SetFont(key, value); err != nil {
			return nil, err
		}
	}
	return conf, nil
}

func (conf FontConfig) SetFont(key string, font *Font) error {
	switch key {
	case `xlabel`, `ylabel`, `title`, `key`, `xtics`, `ytics`:
		conf[key] = font
	default:
		return fmt.Errorf(`Unknown key %v`, key)
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
