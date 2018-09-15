package main

import (
	"reflect"
	"testing"

	grokkTesting "github.com/moncho/grokking/testing"
)

func Test_selectionSort(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"selection sort - empty array",
			args{
				[]int{},
			},
			[]int{},
		},
		{
			"selection sort - unsorted array",
			args{
				[]int{0, 5, 2, 3, 1, 4},
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"selection sort - sorted array",
			args{
				[]int{0, 1, 2, 3, 4, 5},
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectionSort(tt.args.ints)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	grokkTesting.RunSortBenchmarks(b, SelectionSort)
}
