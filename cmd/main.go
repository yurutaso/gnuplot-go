package main

import (
	"fmt"
	"github.com/yurutaso/gnuplot-go"
)

func main() {
	opt1, _ := goplot.NewDataOption(nil)
	opt2, _ := goplot.NewDataOption(nil)
	opt2.Set(`using`, `1:3`)

	data1, _ := goplot.NewPanelData(`test.dat`, opt1)
	data2, _ := goplot.NewPanelData(`test.dat`, opt2)

	panel1, _ := goplot.NewPanel(nil)
	panel1.AddData(data1)
	panel1.Opt.Xaxis.Hide()
	panel1.Opt.Xaxis.Set(`label`, `xaxis`)
	panel1.Opt.Yaxis.Set(`label`, `yaxis`)

	panel2, _ := goplot.NewPanel(nil)
	panel2.AddData(data2)
	panel2.Opt.Xaxis.Set(`label`, `xaxis`)
	panel2.Opt.Yaxis.Set(`label`, `yaxis`)

	plotter, _ := goplot.NewPlotter(nil)
	plotter.AddPanel(panel1)
	plotter.AddPanel(panel2)
	plotter.SetLayout(2, 1)
	plotter.SetInMargins(0., 0.)
	plotter.SetOutput(`output.eps`)
	plotter.Plot()
	fmt.Println(plotter)
}
