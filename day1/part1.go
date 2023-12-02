package main

import (
	"strconv"
	"unicode"
)

func PartOne(lines []string) int {
	var sum int

	for _, line := range lines {
		var numbersFound []string
		var calibrationValue string

		for _, char := range line {
			if unicode.IsDigit(char) {
				numbersFound = append(numbersFound, string(char))
			}
		}

		if len(numbersFound) == 0 {
			continue
		}

		calibrationValue = numbersFound[0] + numbersFound[len(numbersFound)-1]

		asInt, err := strconv.Atoi(calibrationValue)

		if err != nil {
			panic(err)
		}

		sum += asInt
	}
	return sum
}
