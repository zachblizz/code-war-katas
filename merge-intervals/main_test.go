package main

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"first",
			args{
				[][]int{
					[]int{1,4},
					[]int{0,1},
				},
			},
			[][]int{[]int{0,4}},
		},
		{
			"second",
			args{
				[][]int{
					[]int{1,3},
					[]int{2,8},
					[]int{8,10},
					[]int{10,18},
				},
			},
			[][]int{[]int{1,18}},
		},
		{
			"third",
			args{
				[][]int{
					[]int{1,4},
					[]int{0,0},
				},
			},
			[][]int{[]int{0,0},[]int{1,4}},
		},
		{
			"fourth",
			args{
				[][]int{
					[]int{2,3},
					[]int{4,5},
					[]int{6,7},
					[]int{8,9},
					[]int{1,10},
				},
			},
			[][]int{[]int{1,10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
