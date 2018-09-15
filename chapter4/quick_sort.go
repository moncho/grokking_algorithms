package chapter4

func QuickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	var less []int
	var greater []int
	pivot := numbers[0]
	for _, number := range numbers[1:] {
		if number < pivot {
			less = append(less, number)
		} else {
			greater = append(greater, number)
		}
	}
	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
