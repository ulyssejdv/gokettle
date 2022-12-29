package gokettle_test

import (
	"testing"

	gokettle "github.com/ulyssejdv/gokettle/internal"
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

func TestShouldReturnRadius(t *testing.T) {
	// Given
	diamter := 10.
	expected := 5.

	// When
	result := gokettle.Radius(diamter)

	// Then
	if result != expected {
		t.Fatalf(`expected %f, actual %f`, expected, result)
	}
}
