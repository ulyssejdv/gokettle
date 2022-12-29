package gokettle

const PI = 3.14

func VolumeCylinder(radius float64, height float64) float64 {
	return (PI * (radius * radius) * height) / 1000
}

func Radius(diameter float64) float64 {
	return diameter / 2.0
}
