package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	// Part One
	println("Answer 1:", PartOne(lines))

	// Part Two
	println("Answer 2:", PartTwo(lines))
}
