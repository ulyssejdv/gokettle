package gokettle

import (
	"math"
	"testing"
)

func TestShouldReturnVolumeCylinder(t *testing.T) {
	// Given
	expected := 3.141593
	c := Cylinder{diameter: 20., height: 10.}

	// When
	result := c.Volume()

	// Then
	if math.Round(result) != math.Round(expected) {
		t.Fatalf(`expected %f ; actual %f`, expected, result)
	}
}

func TestShouldReturnRadius(t *testing.T) {
	// Given
	expected := 10.
	cylinder := Cylinder{diameter: 20., height: 10.}

	// When
	result := cylinder.Radius()

	// Then
	if result != expected {
		t.Fatalf(`expected %f ; actual %f`, expected, result)
	}
}
