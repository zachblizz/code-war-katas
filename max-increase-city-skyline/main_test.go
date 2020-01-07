package main

import (
	"reflect"
	"testing"
)

func Test_MaxIncreaseKeepingSkyline(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"should return 35",
			args{[][]int{
				{3,0,8,4},
				{2,4,5,7},
				{9,2,6,3},
				{0,3,1,0},
			}},
			35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxIncreaseKeepingSkyline(tt.args.grid); got != tt.want {
				t.Errorf("maxIncreaseKeepingSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaxes(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []Max
	}{
		{
			"should return 35",
			args{[][]int{
				{3,0,8,4}, // {8,9}
				{2,4,5,7}, // {7,4}
				{9,2,6,3}, // {9,8}
				{0,3,1,0}, // {3,9}
			}},
			[]Max{Max{8, 9}, Max{7, 4}, Max{9, 8}, Max{3, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxes(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaxes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reconstructGrid(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"should return 35",
			args{[][]int{
				{3,0,8,4}, // {8,9}
				{2,4,5,7}, // {7,4}
				{9,2,6,3}, // {9,8}
				{0,3,1,0}, // {3,9}
			}},
			[][]int{
				{8,4,8,7},
				{7,4,7,7},
				{9,4,8,7},
				{3,3,3,3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructGrid(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
