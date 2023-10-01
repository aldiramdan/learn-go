package src

import "math"

type hitung2d interface {
	Luas() float64
	Keliling() float64
}

type hitung3d interface {
	Volume() float64
}

type Hitung interface {
	hitung2d
	hitung3d
}

type Kubus struct {
	Sisi *float64
}

func (k *Kubus) Volume() float64 {
	return math.Pow(*k.Sisi, 3)
}

func (k *Kubus) Luas() float64 {
	return math.Pow(*k.Sisi, 2) * 6
}

func (k *Kubus) Keliling() float64 {
	return *k.Sisi * 12
}
