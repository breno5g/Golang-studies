package main

import (
	"fmt"

	"github.com/breno5g/quiz/helpers"
)

// V1
// func main() {
// 	csvFileName := "problems.csv"

// 	questions, err := os.Open(csvFileName)

// 	if err != nil {
// 		log.Fatal("Failed to open the CSV file")
// 	}

// 	defer questions.Close()

// 	csvReader := csv.NewReader(questions)
// 	data, err := csvReader.ReadAll()

// 	if err != nil {
// 		log.Fatal("Failed to parse the provided CSV file")
// 	}

// 	quiz(data)

// }

// func quiz(lines [][]string) {
// 	var total int
// 	for _, line := range lines {
// 		fmt.Printf("What is %s?\n", line[0])
// 		var answer string
// 		fmt.Scanf("%s\n", &answer)
// 		if answer == line[1] {
// 			fmt.Println("Correct!")
// 			total++
// 		} else {
// 			fmt.Println("Wrong!")
// 		}
// 	}

// 	fmt.Printf("You scored %d out of %d", total, len(lines))
// }

// V2
func main() {
	file := helpers.ReadCsv("problems.csv")

	parsedLines := helpers.ParseLines(file)

	quiz(parsedLines)
}

func quiz(questions []helpers.Question) {
	var total int

	for _, question := range questions {
		fmt.Printf("What is %s?\n", question.Question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == question.Answer {
			fmt.Println("Correct!")
			total++
		} else {
			fmt.Println("Wrong!")
		}
	}

	fmt.Printf("You scored %d out of %d", total, len(questions))
}
