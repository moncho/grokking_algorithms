package main

func SelectionSort(ints []int) []int {
	if len(ints) == 0 {
		return ints
	}
	sorted := make([]int, len(ints))
	smallestIndex := 0
	unsorted := ints
	for i := range ints {
		smallestIndex = findIndexOfSmallest(unsorted)
		sorted[i] = unsorted[smallestIndex]
		unsorted = append(unsorted[:smallestIndex], unsorted[smallestIndex+1:]...)
	}
	return sorted
}

func findIndexOfSmallest(ints []int) int {
	smallest := ints[0]
	smallestIndex := 0
	for i, v := range ints {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex
}
