package main

// func Soma(numeros [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numeros[i]
// 	}
// 	return sum
// }

// func Soma(numeros [5]int) int {
// 	sum := 0
// 	// Range é tipo a farofa, primeiro parametro é o indice e o segundo o valor
// 	for _, numero := range numeros {
// 		sum += numero
// 	}
// 	return sum
// }

// Agora ele recebe um slice ao inves de um array
func Soma(numeros []int) int {
	sum := 0
	for _, numero := range numeros {
		sum += numero
	}
	return sum
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	// len === .length
	quantidadeDeNumeros := len(numerosParaSomar)
	// Make cria um array com a quantidade inicial de "quantidade de numeros"
	somas = make([]int, quantidadeDeNumeros)

	// fmt.Println(numerosParaSomar)
	for i, numeros := range numerosParaSomar {
		// Numeros === array
		somas[i] = Soma(numeros)
	}

	return
}

func main() {
	SomaTudo([]int{1, 2}, []int{0, 9})
}
