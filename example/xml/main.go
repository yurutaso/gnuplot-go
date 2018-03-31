package main

import (
	"github.com/yurutaso/gnuplot-go"
	"log"
)

func main() {
	fig, err := goplot.GetFigureFromXML(`style.xml`)
	if err != nil {
		log.Fatal(err)
	}
	fig.Plot()
}