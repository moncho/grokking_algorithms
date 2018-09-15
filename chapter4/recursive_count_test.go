package chapter4

import "testing"

func Test_countElements(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"slice of 5",
			args{
				[]int{0, 0, 0, 0, 0},
			},
			5,
		},
		{
			"slice of 1",
			args{
				[]int{0},
			},
			1,
		},
		{
			"empty",
			args{
				[]int{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.i); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
