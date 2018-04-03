package goplot

import (
	"encoding/xml"
	"os"
	"strings"
)

type XMLTemplate struct {
	Fonts        []*Font        `xml:"Font"`
	Functions    []*Function    `xml:"Function"`
	LineStyles   []*LineStyle   `xml:"LineStyle"`
	Axes         []*Axis        `xml:"Axis"`
	PlotOptions  []*PlotOption  `xml:"PlotOption"`
	PanelOptions []*PanelOption `xml:"PanelOption"`
	Annotation   struct {
		Arrows []*AnnotationArrow `xml:"arrow"`
		Labels []*AnnotationLabel `xml:"label"`
	}
}

func (template *XMLTemplate) Find(entry, name string) (interface{}, bool) {
	switch strings.ToLower(entry) {
	case `paneloption`:
		for _, entity := range template.PanelOptions {
			if entity.Name == name {
				return entity, true
			}
		}
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
	case `label`:
		for _, entity := range template.Annotation.Labels {
			if entity.Name == name {
				return entity, true
			}
		}
	case `arrow`:
		for _, entity := range template.Annotation.Arrows {
			if entity.Name == name {
				return entity, true
			}
		}
	}
	return nil, false
}

type XML struct {
	Figures      []*Figure    `xml:"Figure"`
	TemplateFile string       `xml:"template,attr"`
	Template     *XMLTemplate `xml:"-"`
}

func ReadXMLTemplate(filename string) (*XMLTemplate, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	var template XMLTemplate

	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&template); err != nil {
		return nil, err
	}
	return &template, nil
}

func ReadXML(filename string) (*XML, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	var data XML

	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	if data.TemplateFile != "" {
		data.Template, err = ReadXMLTemplate(data.TemplateFile)
		if err != nil {
			return nil, err
		}
		// Apply templates to Figures
		for _, fig := range data.Figures {
			if err := fig.ApplyTemplate(data.Template); err != nil {
				return nil, err
			}
		}
	}
	return &data, nil
}
