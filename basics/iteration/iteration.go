package iteration

func Repeat(value string, iteration int) string {
	var result string
	if iteration == 0 {
		iteration = 5
	}
	for i := 0; i < iteration; i++ {
		result += value
	}

	return result
}
