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
	panels     []*Panel
	font       FontConfig
	terminal   string
	figname    string
	marginsIn  [2]float64
	marginsOut [4]float64
	row        int
	col        int
}

func NewPlotter() *Plotter {
	return &Plotter{
		panels:     make([]*Panel, 0, 0),
		font:       NewFontConfig(),
		terminal:   `postscript eps enhanced color`,
		figname:    `output.eps`,
		marginsOut: [4]float64{0.1, 0.2, 0.2, 0.1},
		marginsIn:  [2]float64{0., 0.},
		row:        1,
		col:        1,
	}
}

func (plotter *Plotter) SetOutMargins(r, b, l, t float64) {
	plotter.marginsOut[0] = r
	plotter.marginsOut[1] = b
	plotter.marginsOut[2] = l
	plotter.marginsOut[3] = t
}

func (plotter *Plotter) SetInMargins(h, v float64) {
	plotter.marginsIn[0] = h
	plotter.marginsIn[1] = v
}

func (plotter *Plotter) SetFont(font FontConfig) {
	plotter.font = font
}

func (plotter *Plotter) getSize() (float64, float64) {
	xsize := 1.0 / (float64(plotter.row) + plotter.marginsOut[0] + plotter.marginsOut[2]) * (1.0 - plotter.marginsIn[0]/2.)
	ysize := 1.0 / (float64(plotter.col) + plotter.marginsOut[1] + plotter.marginsOut[3]) * (1.0 - plotter.marginsIn[1]/2.)
	return xsize, ysize
}

func (plotter *Plotter) getOrigin(col, row int) (float64, float64) {
	x := 1.0 / (float64(plotter.row) + plotter.marginsOut[0] + plotter.marginsOut[2]) * (float64(row) + plotter.marginsOut[2])
	y := 1.0 / (float64(plotter.col) + plotter.marginsOut[1] + plotter.marginsOut[3]) * (float64(plotter.col) - float64(col+1) + plotter.marginsOut[1])
	return x, y
}

func (plotter *Plotter) getColumn(i int) (col, row int) {
	row = i % plotter.row
	col = int((i - row) / plotter.row)
	return col, row
}

func (plotter *Plotter) SetLayout(col, row int) error {
	if col < 1 || row < 1 {
		return fmt.Errorf("col and row must be positive integer.")
	}
	plotter.col = col
	plotter.row = row
	return nil
}

func (plotter *Plotter) String() string {
	xsize, ysize := plotter.getSize()

	s := fmt.Sprintf(`#!/usr/bin/env/gnuplot
%s
set bmargin 0
set tmargin 0
set lmargin 0
set rmargin 0
set terminal %s
set output "%s"
`,
		plotter.font,
		plotter.terminal,
		plotter.figname,
	)

	s += fmt.Sprintf("set multiplot layout %d, %d\n", plotter.col, plotter.row)

	for i, panel := range plotter.panels {
		col, row := plotter.getColumn(i)
		x, y := plotter.getOrigin(col, row)
		s += fmt.Sprintf("set size %f, %f\nset origin %f, %f\n%s\n", xsize, ysize, x, y, panel)
	}

	s += "unset multiplot\n"
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

func (plotter *Plotter) Plot() (string, error) {
	tmp, err := ioutil.TempFile("", "goplot_exec_temp")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())
	tmp.WriteString(plotter.String())
	err = tmp.Close()
	if err != nil {
		return "", err
	}
	cmd := exec.Command("gnuplot", tmp.Name())
	b, err := cmd.CombinedOutput()
	return string(b), err
}

func (plotter *Plotter) Close() {
	/* Clear all tmpfiles generated by NewPanelDataFromArray() */
	for _, panel := range plotter.panels {
		for _, data := range panel.Data {
			data.atexit()
		}
	}
}
