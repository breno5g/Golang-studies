package iteration

func Repeat(value string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += value
	}

	return result
}
