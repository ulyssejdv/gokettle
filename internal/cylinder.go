package gokettle

import (
	"math"
	"strconv"
)

type Cylinder struct {
	diameter float64
	height   float64
}

func (c *Cylinder) FromString(d string, h string) {

	if r, e := strconv.ParseFloat(d, 64); e == nil {
		c.diameter = r
	}

	if r, e := strconv.ParseFloat(h, 64); e == nil {
		c.height = r
	}
}

func (c Cylinder) Volume() float64 {
	return (math.Pi * (c.Radius() * c.Radius()) * c.height) / 1000
}

func (c Cylinder) Radius() float64 {
	return c.diameter / 2.0
}
