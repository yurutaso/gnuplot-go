package main

import (
	"fmt"
	"github.com/yurutaso/gnuplot-go"
)

func main() {
	opt1 := goplot.NewDataOption()

	opt2 := goplot.NewDataOption()
	opt2.SetUsing(`1:3`)

	data1 := goplot.NewPanelData(`test.dat`, opt1)
	data2 := goplot.NewPanelData(`test.dat`, opt2)

	panel := goplot.NewPanel()
	panel.AddData(data1)
	panel.AddData(data2)

	plotter := goplot.NewPlotter()
	plotter.AddPanel(panel)
	plotter.AddPanel(panel)
	fmt.Println(plotter)
	/*
		plotter.SetOutput(`output.eps`)
		plotter.Plot()
	*/
}
