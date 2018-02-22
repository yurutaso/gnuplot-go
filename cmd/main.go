package main

import (
	"github.com/yurutaso/gnuplot-go"
)

func main() {
	data1 := goplot.NewPanelData(`test.dat`, nil)
	data2 := goplot.NewPanelData(`test.dat`, nil)
	data2.Opt.Set(`using`, `1:3`)

	label1 := goplot.NewAnnotationLabel(`label1`, goplot.NewLocation(`screen`, 0.5, 0.9))
	arrow1 := goplot.NewAnnotationArrow(
		goplot.NewLocation(`screen`, 0.2, 1.0),
		goplot.NewLocation(`screen`, 0.2, 0),
		nil)
	arrow1.NoHead()

	template := goplot.NewPanelOption()
	template.Xaxis.Set(`label`, `xaxis`)
	template.Xaxis.Set(`tics`, 1.)
	template.Xaxis.Set(`mtics`, 5.)
	template.Xaxis.Set(`min`, 1.)
	template.Xaxis.Set(`max`, 6.)
	template.Yaxis.Set(`label`, `yaxis`)
	template.Yaxis.Set(`tics`, 2.)
	template.Yaxis.Set(`mtics`, 2.)
	template.Yaxis.Set(`min`, 1.)
	template.Yaxis.Set(`max`, 10.)
	template.Yaxis.Set(`labelOffset`, goplot.NewLocation(`character`, 2., 0.))

	panel1 := goplot.NewPanel(template.Copy())
	panel1.Opt.Xaxis.Hide()
	panel1.AddData(data1)
	panel1.AddAnnotation(label1)

	panel2 := goplot.NewPanel(template.Copy())
	panel2.AddData(data2)
	panel2.AddAnnotation(arrow1)

	plotter := goplot.NewPlotter()
	defer plotter.Close()

	plotter.AddPanel(panel1)
	plotter.AddPanel(panel2)
	plotter.SetLayout(2, 1)
	plotter.SetInMargins(0., 0.2)
	plotter.SetOutMargins(0.1, 0.2, 0.1, 0.05)
	plotter.SetOutput(`output.eps`)
	plotter.Plot()
}
