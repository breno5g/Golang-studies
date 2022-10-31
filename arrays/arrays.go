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

func SumMultiplesArrays(numbers ...[]int) []int {
	// // Arrays quantity
	// numbersQtd := len(numbers)
	// // make create array with 0 as default value and with x positions
	// sums = make([]int, numbersQtd)

	// // numberArr is each array
	// for i, numberArr := range numbers {
	// 	// We used the sum array to sum individual array
	// 	// And save in sums in index position
	// 	sums[i] = SumArray(numberArr)
	// }

	// Declare empty array
	var sums []int

	for _, numbersArr := range numbers {
		// append sum to sums array
		sums = append(sums, SumArray(numbersArr))
	}

	return sums
}
