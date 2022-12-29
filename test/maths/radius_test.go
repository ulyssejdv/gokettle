package gokettle

import (
	"testing"

	gokettle "github.com/ulyssejdv/gokettle/pkg"
)

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
