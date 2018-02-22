package main

import (
	"fmt"
	"github.com/yurutaso/gnuplot-go"
	"log"
)

func main() {
	opt1, _ := goplot.NewDataOption(nil)
	opt2, _ := goplot.NewDataOption(nil)
	opt2.Set(`using`, `1:3`)

	data1, _ := goplot.NewPanelData(`test.dat`, opt1)
	data2, _ := goplot.NewPanelData(`test.dat`, opt2)

	label1 := goplot.NewAnnotationLabel(`label1`, goplot.NewLocation(`screen`, 0.5, 0.9))
	arrow1 := goplot.NewAnnotationArrow(
		goplot.NewLocation(`screen`, 0.2, 1.0),
		goplot.NewLocation(`screen`, 0.2, 0),
		nil)

	opt, err := goplot.NewPanelOption(nil)
	if err != nil {
		log.Fatal(err)
	}
	opt.Xaxis.Hide()
	opt.Xaxis.Set(`label`, `xaxis`)
	opt.Xaxis.Set(`tics`, 1.)
	opt.Xaxis.Set(`mtics`, 5.)
	opt.Yaxis.Set(`label`, `yaxis2`)

	panel1, _ := goplot.NewPanel(opt.Copy())
	panel1.AddData(data1)
	panel1.AddAnnotation(label1)

	panel2, _ := goplot.NewPanel(opt.Copy())
	panel2.AddData(data2)
	panel2.AddAnnotation(arrow1)
	panel2.Opt.Xaxis.Show()

	plotter, _ := goplot.NewPlotter(nil)
	plotter.AddPanel(panel1)
	plotter.AddPanel(panel2)
	plotter.SetLayout(2, 1)
	plotter.SetInMargins(0., 0.)
	plotter.SetOutput(`output.eps`)
	output, err := plotter.Plot()
	fmt.Println(plotter)
	fmt.Println(output)
	if err != nil {
		log.Fatal(err)
	}
}
