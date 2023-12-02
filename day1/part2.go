package main

import (
	"strconv"
	"unicode"
)

func PartTwo(lines []string) int {
	var sum int

	spelledDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range lines {
		var numbersFound []string
		var calibrationValue string

		for i, char := range line {
			if unicode.IsDigit(char) {
				numbersFound = append(numbersFound, string(char))
				continue
			}

			for number := range spelledDigits {
				if len(line)-(i) < len(number) {
					continue
				}

				if line[i:i+len(number)] == number {
					numbersFound = append(numbersFound, spelledDigits[number])
				}
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
