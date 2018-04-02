package goplot

import (
	"encoding/xml"
	"os"
	"strings"
	"fmt"
)

type XMLTemplate struct {
	Fonts []*Font `xml:"Fonts>Font"`
	Functions []*Function `xml:"Functions>Function"`
	LineStyles []*LineStyle `xml:"LineStyles>LineStyle"`
	Axes []*Axis `xml:"Axes>Axis"`
	PlotOptions []*PlotOption `xml:"PlotOptions>PlotOption"`
	Annotation struct{
		Arrows []*AnnotationArrow `xml:"arrow"`
		Labels []*AnnotationLabel `xml:"label"`
	}
}

func (template *XMLTemplate) Find(entry, name string) (interface{}, bool) {
	switch strings.ToLower(entry) {
	case `plotoption`:
		for _, entity := range template.PlotOptions {
			if entity.Name == name {
				return entity, true
			}
		}
	case `function`:
		for _, entity := range template.Functions {
			if entity.Name == name {
				return entity, true
			}
		}
	case `linestyle`:
		for _, entity := range template.LineStyles {
			if entity.Name == name {
				return entity, true
			}
		}
	case `axis`:
		for _, entity := range template.Axes {
			if entity.Name == name {
				return entity, true
			}
		}
	}
	return nil, false
}

type XML struct {
	Figures []*Figure `xml:"Figure"`
	TemplateFile string `xml:"template,attr"`
	Template *XMLTemplate `xml:"-"`
}

func ReadXMLTemplate(filename string) (*XMLTemplate, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	var template XMLTemplate

	decoder := xml.NewDecoder(xmlFile)
	if err:= decoder.Decode(&template); err != nil {
		return nil, err
	}
	return &template, nil
}

func ReadXML(filename string) error {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	var data XML

	decoder := xml.NewDecoder(xmlFile)
	if err:= decoder.Decode(&data); err != nil {
		return err
	}

	if data.TemplateFile != "" {
		data.Template, err = ReadXMLTemplate(data.TemplateFile)
		if err != nil {
			return err
		}
		// Apply templates to Figures
		for _, fig := range data.Figures {
			fig.ApplyTemplate(data.Template)
			fmt.Println(fig)
		}
	}
	return nil
}