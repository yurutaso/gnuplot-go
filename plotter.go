package goplot

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

/* type plotter*/
/*
Structure:
  Plotter --- FontConfig
           |_ Panel --- PanelOption --- AxisOption
                     |_ PanelData --- DataOption --- LineStyle
*/

type Plotter struct {
	panels   []*Panel
	font     FontConfig
	terminal string
	figname  string
}

func NewPlotter(font FontConfig) (*Plotter, error) {
	var err error
	if font == nil {
		font, err = NewFontConfig(nil)
		if err != nil {
			return nil, err
		}
	}
	return &Plotter{
		panels:   make([]*Panel, 0, 0),
		font:     font,
		terminal: `postscript eps enhanced color`,
		figname:  `output.eps`,
	}, nil
}

func (plotter *Plotter) String() string {
	s := fmt.Sprintf(`#!/usr/bin/env/gnuplot
%s
set terminal %s
set output "%s"`,
		plotter.font,
		plotter.terminal,
		plotter.figname,
	)
	for _, panel := range plotter.panels {
		s += panel.String()
	}
	return s
}

func (plotter *Plotter) SetOutput(figname string) {
	plotter.figname = figname
}

func (plotter *Plotter) SetTerminal(terminal string) {
	plotter.terminal = terminal
}

func (plotter *Plotter) AddPanel(panel *Panel) {
	plotter.panels = append(plotter.panels, panel)
}

func (plotter *Plotter) Plot() error {
	tmp, err := ioutil.TempFile("", "goplot_exec_temp")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	tmp.WriteString(plotter.String())
	err = tmp.Close()
	if err != nil {
		return err
	}
	cmd := exec.Command("gnuplot", tmp.Name())
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
