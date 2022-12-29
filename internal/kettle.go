package gokettle

type Kettle struct {
	Height, Diameter float64
}

func (k Kettle) TotalVolume() float64 {
	return VolumeCylinder(Radius(k.Diameter), k.Height)
}

func (k Kettle) ContentsVolumeByEmptyHeight(height float64) float64 {
	return k.TotalVolume() - VolumeCylinder(Radius(k.Diameter), height)
}
