package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("Failed to read data: %v", err)
	}
	err = plotData("out.png", xys)

}

func plotData(path string, xys plotter.XYs) error {
	f, err := os.Create("out.png")
	if err != nil {
		return fmt.Errorf("Could not create png file: %v", err)
	}
	//create plot
	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("Could not create a plot: %v", err)
	}

	//create scatter
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("Could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	var x, c float64
	x = 1.2
	c = -3

	// create fake linear regression result
	l, err := plotter.NewLine(plotter.XYs{
		{3, 3*x + c}, {20, 20*x + c},
	})
	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("Could not create Writer: %v", err)
	}
	if _, err := wt.WriteTo(f); err != nil {
		return fmt.Errorf("Error Writing to plotter to file: %v", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("Error Closing file: %v", err)
	}

	return nil
}

type xy struct{ x, y float64 }

func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys plotter.XYs
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("Discarding bad data point %q:%v", s.Text(), err)
		}
		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("Could not scan: %v", err)
	}
	return xys, nil
}
