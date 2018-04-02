package goplot

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

/* type plotter*/
/*
Layers:
	Figure----Fonts
		|---- Panels----PanelOption---AxisOption
				|-------PanelData/Function---PlotOption---LineStyle
				|-------PanelAnnotation
*/

type Figure struct {
	Panels     []*Panel `xml:"Panels>Panel"`
	Fonts      []*Font `xml:"Fonts>Font"`
	Terminal   string
	Figname    string `xml:"Name"`
	Margins    *Margin
	Row        int
	Col        int
}

func (fig *Figure) IsFontSet(field string) bool {
	for _, font := range fig.Fonts {
		if field == font.Field {
			return true
		}
	}
	return false
}

func (fig *Figure) ApplyTemplate(template *XMLTemplate) error {
	for _, font := range template.Fonts {
		if !fig.IsFontSet(font.Field) {
			fig.Fonts = append(fig.Fonts, font)
		}
	}
	for _, panel := range fig.Panels {
		// Xaxis
		if panel.Xaxis.Name != "" {
			if axis, found := template.Find(`axis`, panel.Xaxis.Name); found {
				panel.Xaxis = axis.(*Axis)
			} else {
				return fmt.Errorf("Axis named %s not found in the template file\n", panel.Xaxis.Name)
			}
		}
		// Yaxis
		if panel.Yaxis.Name != "" {
			if axis, found := template.Find(`axis`, panel.Yaxis.Name); found {
				panel.Yaxis = axis.(*Axis)
			} else {
				return fmt.Errorf("Axis named %s not found in the template file\n", panel.Yaxis.Name)
			}
		}
		// Data
		for i, data := range panel.Plot.Data {
			if data.Opt != nil && data.Opt.Name != "" {
				if opt, found := template.Find(`plotOption`, data.Opt.Name); found {
					panel.Plot.Data[i].Opt = opt.(*PlotOption)
				} else {
					return fmt.Errorf("PlotOption named %s not found in the template file\n", data.Opt.Name)
				}
			}
		}
		// Function
		for i, data := range panel.Plot.Func {
			if data.Opt != nil && data.Opt.Name != "" {
				if opt, found := template.Find(`plotOption`, data.Opt.Name); found {
					panel.Plot.Func[i].Opt = opt.(*PlotOption)
				} else {
					return fmt.Errorf("PlotOption named %s not found in the template file\n", data.Opt.Name)
				}
			}
			if data.Function != nil && data.Function.Name != "" {
				if f, found := template.Find(`function`, data.Function.Name); found {
					panel.Plot.Func[i].Function = f.(*Function)
				} else {
					return fmt.Errorf("Function named %s not found in the template file\n", data.Function.Name)
				}
			}
		}
	}
	return nil
}

type Margin struct {
	Right float64 `xml:"right,attr"`
	Left float64 `xml:"left,attr"`
	Top float64 `xml:"top,attr"`
	Bottom float64 `xml:"bottom,attr"`
	Horizontal float64 `xml:"h,attr"`
	Vertical float64 `xml:"v,attr"`
}

func NewFigure() *Figure {
	return &Figure{
		Panels:     make([]*Panel, 0, 0),
		Fonts:       make([]*Font, 0, 0),
		Terminal:   `postscript eps enhanced color`,
		Figname:    `output.eps`,
		Margins: &Margin{
			Right: 0.1,
			Bottom: 0.2,
			Left: 0.2,
			Top: 0.1,
			Horizontal: 0.,
			Vertical: 0.,
		},
		Row:        1,
		Col:        1,
	}
}


func (fig *Figure) SetOutMargins(r, b, l, t float64) {
	fig.Margins.Right = r
	fig.Margins.Bottom = b
	fig.Margins.Left = l
	fig.Margins.Top = t
}

func (fig *Figure) SetInMargins(h, v float64) {
	fig.Margins.Horizontal = h
	fig.Margins.Vertical = v
}

func (fig *Figure) SetFont(font []*Font) {
	fig.Fonts = font
}

func (fig *Figure) getSize() (float64, float64) {
	xsize := 1.0 / (float64(fig.Row) + fig.Margins.Right + fig.Margins.Left) * (1.0 - fig.Margins.Horizontal/2.)
	ysize := 1.0 / (float64(fig.Col) + fig.Margins.Bottom + fig.Margins.Top) * (1.0 - fig.Margins.Vertical/2.)
	return xsize, ysize
}

func (fig *Figure) getOrigin(col, row int) (float64, float64) {
	x := 1.0 / (float64(fig.Row) + fig.Margins.Right + fig.Margins.Left) * (float64(row) + fig.Margins.Left)
	y := 1.0 / (float64(fig.Col) + fig.Margins.Bottom + fig.Margins.Top) * (float64(fig.Col) - float64(col+1) + fig.Margins.Bottom)
	return x, y
}

func (fig *Figure) getColumn(i int) (col, row int) {
	row = i % fig.Row
	col = int((i - row) / fig.Row)
	return col, row
}

func (fig *Figure) SetLayout(col, row int) error {
	if col < 1 || row < 1 {
		return fmt.Errorf("col and row must be positive integer.")
	}
	fig.Col = col
	fig.Row = row
	return nil
}

func (fig *Figure) String() string {
	xsize, ysize := fig.getSize()

	s := "#!/usr/bin/env/gnuplot\n"
	for _, font := range fig.Fonts {
		s += fmt.Sprintf("%s\n", font)
	}
	s += fmt.Sprintf(`set bmargin 0
set tmargin 0
set lmargin 0
set rmargin 0
set terminal %s
set output "%s"
`,
		fig.Terminal,
		fig.Figname,
	)

	s += fmt.Sprintf("set multiplot layout %d, %d\n", fig.Col, fig.Row)

	for i, panel := range fig.Panels {
		col, row := fig.getColumn(i)
		x, y := fig.getOrigin(col, row)
		s += fmt.Sprintf("set size %f, %f\nset origin %f, %f\n%s\n", xsize, ysize, x, y, panel)
	}

	s += "unset multiplot\n"
	return s
}

func (fig *Figure) SetOutput(figname string) {
	fig.Figname = figname
}

func (fig *Figure) SetTerminal(terminal string) {
	fig.Terminal = terminal
}

func (fig *Figure) AddPanel(panel *Panel) {
	fig.Panels = append(fig.Panels, panel)
}

func (fig *Figure) Plot() (string, error) {
	tmp, err := ioutil.TempFile("", "goplot_exec_temp")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())
	tmp.WriteString(fig.String())
	err = tmp.Close()
	if err != nil {
		return "", err
	}
	cmd := exec.Command("gnuplot", tmp.Name())
	b, err := cmd.CombinedOutput()
	return string(b), err
}

func (fig *Figure) Close() {
	/* Clear all tmpfiles generated by NewPanelDataFromArray() */
	for _, panel := range fig.Panels {
		for _, data := range panel.Plot.Data {
			data.atexit()
		}
	}
}