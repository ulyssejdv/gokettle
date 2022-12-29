package gokettle

import (
	"testing"

	gokettle "github.com/ulyssejdv/gokettle/pkg"
)

func TestShouldReturnVolumeCylinder(t *testing.T) {
	// Given
	radius := 10.
	height := 10.
	expected := 3.14

	// When
	result := gokettle.VolumeCylinder(radius, height)

	// Then
	if result != expected {
		t.Fatalf(`expected %f, actual %f`, expected, result)
	}
}
