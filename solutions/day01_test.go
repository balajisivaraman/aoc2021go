package solutions

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7
	actual := Day01Part1(input)
	if !(actual == want) {
		t.Errorf("Day01: Part 1(%d) = %d, want %d", input, actual, want)
	}
}

func TestDay01Part2(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 5
	actual := Day01Part2(input)
	if !(actual == want) {
		t.Errorf("Day01: Part 2(%d) = %d, want %d", input, actual, want)
	}
}
