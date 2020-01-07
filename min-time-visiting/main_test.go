package main

import (
	"testing"
)

func Test_minTimeToVisitAllPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"first test",
			args{[][]int{
				{1, 1},
				{3, 4},
				{-1, 0},
			}},
			7,
		},
		// {
		// 	"second test",
		// 	args{[][]int{
		// 		{3,2},
		// 		{-2,2},
		// 	}},
		// 	5,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitAllPoints(tt.args.points); got != tt.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTimeFromPoints(t *testing.T) {
	type args struct {
		p1 []int
		p2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"first test",
			args{[]int{1,1},[]int{3,4}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTimeFromPoints(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("getTimeFromPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
