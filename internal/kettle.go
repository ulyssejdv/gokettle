package gokettle

type Kettle struct {
	TotalHeight, EmptyHeight, Diameter string
}

func (k Kettle) TotalVolume() float64 {
	cylinder := CylinderFromString(k.Diameter, k.TotalHeight)
	return cylinder.Volume()
}

func (k Kettle) EmptyVolume() float64 {
	cylinder := CylinderFromString(k.Diameter, k.EmptyHeight)
	return cylinder.Volume()
}

func (k Kettle) FilledVolume() float64 {
	return k.TotalVolume() - k.EmptyVolume()
}
