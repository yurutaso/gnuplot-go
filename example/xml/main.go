package main

import (
	"fmt"
	"github.com/yurutaso/gnuplot-go"
	"log"
)

func main() {
	xml, err := goplot.ReadXML(`style.xml`)
	if err != nil {
		log.Fatal(err)
	}
	for _, fig := range xml.Figures {
		fmt.Println(fig)
	}
}
