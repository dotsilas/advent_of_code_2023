package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	line := []string{"two1nine", "abcone2threexyz", "4nineeightseven2"}
	want := 11 + 22 + 42

	actual := PartOne(line)

	if actual != want {
		t.Fatalf("got %d, but wanted %d", actual, want)
	}

}
