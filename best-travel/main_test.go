package main

import "testing"

func TestChooseBestSum(t *testing.T) {
	type args struct {
		t  int
		k  int
		ls []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"should return 163", args{163, 3, []int{50, 55, 56, 57, 58}}, 163},
		// {"should return -1", args{163, 3, []int{50}}, -1},
		// {"should return 228", args{230, 3, []int{91, 74, 73, 85, 73, 81, 87}}, 228},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChooseBestSum(tt.args.t, tt.args.k, tt.args.ls); got != tt.want {
				t.Errorf("ChooseBestSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
