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
