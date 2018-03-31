package goplot

import (
	"encoding/xml"
	"os"
)

func GetFigureFromXML(filename string) (*Figure, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	var fig Figure

	decoder := xml.NewDecoder(xmlFile)
	if err:= decoder.Decode(&fig); err != nil {
		return nil, err
	}
	return &fig, nil
}