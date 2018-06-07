package main

import (
	"github.com/yurutaso/gnuplot-go"
	"log"
	"os"
)

func main() {
	args := os.Args
	xmlfile := args[1]
	if len(args) != 2 {
		log.Fatal(`Usage: plotxml filename`)
	}
	xml, err := goplot.ReadXML(xmlfile)
	if err != nil {
		log.Fatal(err)
	}
	for _, fig := range xml.Figures {
		//fmt.Println(fig)
		fig.Plot()
	}
}
