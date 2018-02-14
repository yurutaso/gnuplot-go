package main

import (
	"github.com/yurutaso/gnuplot-go"
)

func main() {
	data := goplot.NewPlotData(`test.dat`, nil)

	panel := goplot.NewPanel()
	panel.AddData(data)

	plotter := goplot.NewPlotter()
	plotter.AddPanel(panel)
	plotter.SetOutput(`output.eps`)
	plotter.Plot()
}
