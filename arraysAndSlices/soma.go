package main

func Soma(numeros [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numeros[i]
	}
	return sum
}
