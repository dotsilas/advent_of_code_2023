package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// extract string of calibration values
func extractNumbers(text string) (numbers string) {

	numbersList := []string{}
	for _, char := range text {
		if unicode.IsDigit(char) {
			numbersList = append(numbersList, string(char))
		}
	}

	if len(numbersList) == 1 {
		numbers = numbersList[0] + numbersList[0]
	} else if len(numbersList) > 1 {
		numbers = numbersList[0] + numbersList[len(numbersList)-1]
	} else {
		numbers = ""
	}
	return numbers
}

func main() {
	// read file
	filePath := "./input1.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// split content into lines
	lines := strings.Split(string(content), "\n")

	answer1 := 0

	// iterate over lines
	for _, line := range lines {
		numbersString := extractNumbers(line)
		if numbersString == "" {
			continue
		}

		numberInt, _ := strconv.ParseInt(numbersString, 10, 64)
		answer1 += int(numberInt)
	}

	fmt.Println(answer1)
}
