package solutions

func Day01Part1(input []int) int {
	count := 0
	for i := 0; i < len(input) - 1; i++ {
		if input[i + 1] > input [i] {
			count++
		}
	}
	return count
}

func Day01Part2(input []int) int {
	count := 0
	for i := 0; i < len(input) - 3; i++ {
		firstWindow := input [i] + input [i + 1] + input [i + 2]
		secondWindow := input [i + 1] + input [i + 2] + input [i + 3]
		if secondWindow > firstWindow {
			count++
		}
	}
	return count
}
