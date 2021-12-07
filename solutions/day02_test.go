package solutions

import (
	"testing"
)

func TestDay02Part1(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	want := 150
	actual := Day02Part1(input)
	if !(actual == want) {
		t.Errorf("Day02: Part 1(%s) = %d, want %d", input, actual, want)
	}
}
