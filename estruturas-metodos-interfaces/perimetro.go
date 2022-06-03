package perimetro

type Retangulo struct {
	Largura float64
	Altura  float64
}

func Perimetro(value Retangulo) (res float64) {
	return (value.Altura + value.Largura) * 2
}

func Area(value Retangulo) (res float64) {
	return value.Altura * value.Largura
}
