package iteracao

func Repetir(value string, quantidadeDeRepeticoes int) string {
	var resultado string = ""

	if quantidadeDeRepeticoes == 0 {
		quantidadeDeRepeticoes = 5
	}
	for i := 0; i < quantidadeDeRepeticoes; i++ {
		resultado = resultado + value
	}

	return resultado
}

// exemplos extras de for

// import "fmt"

// func exeplos() {
// 	i := 1
// 	for i <= 3 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}

// 	for j := 7; j <= 9; j++ {
// 		fmt.Println(j)
// 	}

// 	for {
// 		fmt.Println("loop")
// 		break
// 	}

// 	for n := 0; n <= 5; n++ {
// 		if n%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(n)
// 	}
// }
