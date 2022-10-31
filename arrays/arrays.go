package arrays

func SumArray(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	// for of?
	for _, number := range numbers {
		//params are index and value
		sum += number
	}
	return sum
}
