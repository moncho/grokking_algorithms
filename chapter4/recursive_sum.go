package chapter4

func sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	return numbers[0] + sum(numbers[1:])
}
