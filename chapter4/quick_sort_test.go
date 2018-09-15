package chapter4

import (
	"reflect"
	"testing"

	grokkTesting "github.com/moncho/grokking/testing"
)

func Test_QuickSort(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"quick sort - empty array",
			args{
				[]int{},
			},
			[]int{},
		},
		{
			"quick sort - unsorted array",
			args{
				[]int{0, 5, 2, 3, 1, 4},
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"quick sort - unsorted array",
			args{
				[]int{0, 5, 2, 3, 1, 4, 6, 9, 8, 7},
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"quick sort - desc sorted array",
			args{
				[]int{5, 4, 3, 2, 1, 0},
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"quick sort - asc sorted array",
			args{
				[]int{0, 1, 2, 3, 4, 5},
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickSort(tt.args.ints)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	grokkTesting.RunSortBenchmarks(b, QuickSort)
}
