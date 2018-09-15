package chapter4

func countElements(i []int) int {
	if len(i) == 0 {
		return 0
	}
	return 1 + countElements(i[1:])
}
