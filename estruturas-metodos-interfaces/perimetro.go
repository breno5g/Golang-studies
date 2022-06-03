package main

import "math"

type Retangulo struct {
	Largura float64
	Altura  float64
}

func Perimetro(value Retangulo) (res float64) {
	return (value.Altura + value.Largura) * 2
}

// func Area(value Retangulo) (res float64) {
// 	return value.Altura * value.Largura
// }

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}
