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

// func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
// 	// len === .length
// 	quantidadeDeNumeros := len(numerosParaSomar)
// 	// Make cria um array com a quantidade inicial de "quantidade de numeros"
// 	somas = make([]int, quantidadeDeNumeros)

// 	fmt.Println(numerosParaSomar)
// 	for i, numeros := range numerosParaSomar {
// 		// Numeros === array
// 		somas[i] = Soma(numeros)
// 	}

// 	return
// }

func SomaTudo(numerosParaSomar ...[]int) []int {
	var somas []int // Cria um slice vazio
	for _, numeros := range numerosParaSomar {
		// Dá um append em "somas" com o retorno de "Soma()"
		somas = append(somas, Soma(numeros))
	}

	return somas
}

// func SomaTodoOResto(numerosParaSomar ...[]int) []int {
// 	var somas []int
// 	for _, numeros := range numerosParaSomar {
// 		// isso é dar um slice no array, ele pega da posição passada como primeiro parametro até a do segundo
// 		// Se você omitir o segundo parametro, ele vai do primeiro até o final
// 		// O inverso também vale
// 		final := numeros[1:]
// 		somas = append(somas, Soma(final))
// 	}

// 	return somas
// }

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 { // Verifica se o array é maior que zero
			somas = append(somas, 0) // adiciona zero a soma
		} else { // se não
			final := numeros[1:] // Procedimento normal
			somas = append(somas, Soma(final))
		}
	}

	return somas
}

func main() {
	SomaTudo([]int{1, 2}, []int{0, 9})
	SomaTodoOResto([]int{1, 2, 3, 4}, []int{10, 20, 30})
}
