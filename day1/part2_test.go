package main

import (
	"testing"
)

func TestPart2(t *testing.T) {
	line := []string{"two1nine"}
	want := 29

	actual := PartTwo(line)

	if actual != want {
		t.Fatalf("got %d, but wanted %d", actual, want)
	}

}
