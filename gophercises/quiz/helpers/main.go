package helpers

import (
	"encoding/csv"
	"log"
	"os"
)

type Question struct {
	Question string
	Answer   string
}

func ReadCsv(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal("Failed to open the CSV file")
	}

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Failed to parse the provided CSV file")
	}

	return data
}

func ParseLines(lines [][]string) []Question {
	arr := make([]Question, len(lines))

	for i, line := range lines {
		arr[i] = Question{
			Question: line[0],
			Answer:   line[1],
		}
	}

	return arr
}
