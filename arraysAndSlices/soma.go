package main

// func Soma(numeros [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numeros[i]
// 	}
// 	return sum
// }

func Soma(numeros [5]int) int {
	sum := 0
	// Range é tipo a farofa, primeiro parametro é o indice e o segundo o valor
	for _, numero := range numeros {
		sum += numero
	}
	return sum
}
