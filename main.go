package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
)

func main() {
	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("Failed to read data: %v", err)
	}
	for _, xy := range xys {
		fmt.Println(xy)
	}

	//create plot
	p, err := plot.New()
	if err != nil {
		log.Fatalf("Could not create a plot: %v", err)
	}
	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		log.Fatalf("Could not create Writer: %v", err)
	}
	f, err := os.Open("out.png")
}

type xy struct{ x, y float64 }

func readData(path string) ([]xy, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys []xy
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("Discarding bad data point %q:%v", s.Text(), err)
		}
		xys = append(xys, xy{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("Could not scan: %v", err)
	}
	return xys, nil
}
