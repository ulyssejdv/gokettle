package gokettle

import (
	"log"
	"math"
	"strconv"
)

type Cylinder struct {
	diameter, height float64
}

func CylinderFromString(d string, h string) Cylinder {
	diameter, errDiameter := strconv.ParseFloat(d, 64)
	height, errHeight := strconv.ParseFloat(h, 64)

	if errDiameter != nil || errHeight != nil {
		log.Fatal("Can't convert string to float during cylinder creation.")
	}
	return Cylinder{diameter: diameter, height: height}
}

func (c Cylinder) Volume() float64 {
	return (math.Pi * (c.Radius() * c.Radius()) * c.height) / 1000.
}

func (c Cylinder) Radius() float64 {
	return c.diameter / 2.
}
