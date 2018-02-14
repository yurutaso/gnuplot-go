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

type FontConfig struct {
	xlabel *Font
	ylabel *Font
	title  *Font
	key    *Font
	xtics  *Font
	ytics  *Font
}

func NewFontConfig() *FontConfig {
	return &FontConfig{
		xlabel: &Font{name: `Helvetica`, size: 16},
		ylabel: &Font{name: `Helvetica`, size: 16},
		title:  &Font{name: `Helvetica`, size: 14},
		key:    &Font{name: `Helvetica`, size: 12},
		xtics:  &Font{name: `Helvetica`, size: 12},
		ytics:  &Font{name: `Helvetica`, size: 12},
	}
}

func (conf *FontConfig) String() string {
	return fmt.Sprintf(`set xlabel %s
set ylabel %s
set title %s
set key %s
set xtics %s
set ytics %s`,
		conf.xlabel,
		conf.ylabel,
		conf.title,
		conf.key,
		conf.xtics,
		conf.ytics,
	)
}

func (conf *FontConfig) SetFontXLabel(font *Font) {
	conf.xlabel = font
}

func (conf *FontConfig) SetFontYLabel(font *Font) {
	conf.ylabel = font
}

func (conf *FontConfig) SetFontTitle(font *Font) {
	conf.title = font
}

func (conf *FontConfig) SetFontKey(font *Font) {
	conf.key = font
}

func (conf *FontConfig) SetFontXTics(font *Font) {
	conf.xtics = font
}

func (conf *FontConfig) SetFontYTics(font *Font) {
	conf.ytics = font
}
