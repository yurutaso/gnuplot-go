package main

import (
	"github.com/yurutaso/gnuplot-go"
	"log"
)

func main() {
	err := goplot.ReadXML(`style.xml`)
	if err != nil {
		log.Fatal(err)
	}
}