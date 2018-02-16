package main

import (
	"fmt"
	"github.com/yurutaso/gnuplot-go"
)

func main() {
	opt1, _ := goplot.NewDataOption(nil)

	opt2, _ := goplot.NewDataOption(nil)
	opt2.Set(`u`, `1:3`)

	data1, _ := goplot.NewPanelData(`test.dat`, opt1)
	data2, _ := goplot.NewPanelData(`test.dat`, opt2)

	panel, _ := goplot.NewPanel(nil)
	panel.AddData(data1)
	panel.AddData(data2)

	plotter, _ := goplot.NewPlotter(nil)
	plotter.AddPanel(panel)
	plotter.AddPanel(panel)
	fmt.Println(plotter)
	/*
		plotter.SetOutput(`output.eps`)
		plotter.Plot()
	*/
}
