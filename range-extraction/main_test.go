package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"should be good...",
			args{[]int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20}},
			"-6,-3-1,3-5,7-11,14,15,17-20",
		},
		{
			"should be good...",
			args{[]int{40,44,48,51,52,54,55,58,67,73}},
			"40,44,48,51,52,54,55,58,67,73",
		},
		{
			"should be good...",
			args{[]int{21,26,27,28,29,30,31,36,45,46,47}},
			"21,26-31,36,45-47",
		},
		{
			"should be good...",
			args{[]int{-32,-31,-30,-29,-22,-21,-20,-16,-15,-14,-8,1,7,8,10,11,16}},
			"-32--29,-22--20,-16--14,-8,1,7,8,10,11,16",
		},
		{
			"should be good",
			args{[]int{-3,3,8,9,17,18,19,24,28,29}},
			"-3,3,8,9,17-19,24,28,29",
		},
		{
			"should be good...",
			args{[]int{-24,-16,-15,-10,-9,-8,-7,-6,-5,1}},
			"-24,-16,-15,-10--5,1",
		},
		{
			"interesting...",
			args{[]int{20,25,29,30,35,36,37,38,39,44}},
			"20,25,29,30,35-39,44",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.list); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
