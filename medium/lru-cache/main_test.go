package main

import (
	"reflect"
	"testing"
)

func Test_setup(t *testing.T) {
	type args struct {
		ops  []string
		vals [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"first",
			args{
				[]string{"LRUCache", "get", "put", "get", "put", "put", "get", "get"},
				[][]int{[]int{2}, []int{2}, []int{2, 6}, []int{1}, []int{1, 5}, []int{1, 2}, []int{1}, []int{2}},
			},
			[]int{-1, -1, 2, 6},
		},
		{
			"second",
			args{
				[]string{"LRUCache", "put", "put", "put", "put", "get", "get"},
				[][]int{[]int{2}, []int{2, 1}, []int{1, 1}, []int{2, 3}, []int{4, 1}, []int{1}, []int{2}},
			},
			[]int{-1, 3},
		},
		{
			"third",
			args{
				[]string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				[][]int{[]int{2}, []int{1, 1}, []int{2, 2}, []int{1}, []int{3, 3}, []int{2}, []int{4, 4}, []int{1}, []int{3}, []int{4}},
			},
			[]int{1, -1, -1, 3, 4},
		},
		{
			"fourth",
			args{
				[]string{"LRUCache", "put", "put", "put", "put", "get", "get", "get", "get", "put", "get", "get", "get", "get", "get"},
				[][]int{[]int{3}, []int{1, 1}, []int{2, 2}, []int{3, 3}, []int{4, 4}, []int{4}, []int{3}, []int{2}, []int{1}, []int{5, 5}, []int{1}, []int{2}, []int{3}, []int{4}, []int{5}},
			},
			[]int{4, 3, 2, -1, -1, 2, 3, -1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setup(tt.args.ops, tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setup() = %v, want %v", got, tt.want)
			}
		})
	}
}
