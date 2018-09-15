package chapter4

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 + 1",
			args{
				[]int{1, 1},
			},
			2,
		},

		{"1 + 2 + 3",
			args{
				[]int{1, 2, 3},
			},
			6,
		},
		{"5 + 0 + 4 +90 ",
			args{
				[]int{5, 0, 4, 90},
			},
			99,
		},
		{"empty sum",
			args{
				[]int{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.numbers); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
