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
		// {
		// 	"first",
		// 	args{
		// 		[]string{"LRUCache", "get", "put", "get", "put", "put", "get", "get"},
		// 		[][]int{[]int{2}, []int{2}, []int{2, 6}, []int{1}, []int{1, 5}, []int{1, 2}, []int{1}, []int{2}},
		// 	},
		// 	[]int{-1, -1, 2, 6},
		// },
		{
			"second",
			args{
				[]string{"LRUCache", "put", "put", "put", "put", "get", "get"},
				[][]int{[]int{2}, []int{2, 1}, []int{1, 1}, []int{2, 3}, []int{4, 1}, []int{1}, []int{2}},
			},
			[]int{-1, 3},
		},
		// {
		// 	"third",
		// 	args{
		// 		[]string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
		// 		[][]int{[]int{2}, []int{1, 1}, []int{2, 2}, []int{1}, []int{3, 3}, []int{2}, []int{4, 4}, []int{1}, []int{3}, []int{4}},
		// 	},
		// 	[]int{1, -1, -1, 3, 4},
		// },
		// {
		// 	"fourth",
		// 	args{
		// 		[]string{"LRUCache", "put", "put", "put", "put", "get", "get", "get", "get", "put", "get", "get", "get", "get", "get"},
		// 		[][]int{[]int{3}, []int{1, 1}, []int{2, 2}, []int{3, 3}, []int{4, 4}, []int{4}, []int{3}, []int{2}, []int{1}, []int{5, 5}, []int{1}, []int{2}, []int{3}, []int{4}, []int{5}},
		// 	},
		// 	[]int{4, 3, 2, -1, -1, 2, 3, -1, 5},
		// },
		// {
		// 	"fifth",
		// 	args{
		// 		[]string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"},
		// 		[][]int{[]int{10}, []int{10, 13}, []int{3, 17}, []int{6, 11}, []int{10, 5}, []int{9, 10}, []int{13}, []int{2, 19}, []int{2}, []int{3}, []int{5, 25}, []int{8}, []int{9, 22}, []int{5, 5}, []int{1, 30}, []int{11}, []int{9, 12}, []int{7}, []int{5}, []int{8}, []int{9}, []int{4, 30}, []int{9, 3}, []int{9}, []int{10}, []int{10}, []int{6, 14}, []int{3, 1}, []int{3}, []int{10, 11}, []int{8}, []int{2, 14}, []int{1}, []int{5}, []int{4}, []int{11, 4}, []int{12, 24}, []int{5, 18}, []int{13}, []int{7, 23}, []int{8}, []int{12}, []int{3, 27}, []int{2, 12}, []int{5}, []int{2, 9}, []int{13, 4}, []int{8, 18}, []int{1, 7}, []int{6}, []int{9, 29}, []int{8, 21}, []int{5}, []int{6, 30}, []int{1, 12}, []int{10}, []int{4, 15}, []int{7, 22}, []int{11, 26}, []int{8, 17}, []int{9, 29}, []int{5}, []int{3, 4}, []int{11, 30}, []int{12}, []int{4, 29}, []int{3}, []int{9}, []int{6}, []int{3, 4}, []int{1}, []int{10}, []int{3, 29}, []int{10, 28}, []int{1, 20}, []int{11, 13}, []int{3}, []int{3, 12}, []int{3, 8}, []int{10, 9}, []int{3, 26}, []int{8}, []int{7}, []int{5}, []int{13, 17}, []int{2, 27}, []int{11, 15}, []int{12}, []int{9, 19}, []int{2, 15}, []int{3, 16}, []int{1}, []int{12, 17}, []int{9, 1}, []int{6, 19}, []int{4}, []int{5}, []int{5}, []int{8, 1}, []int{11, 7}, []int{5, 2}, []int{9, 28}, []int{1}, []int{2, 2}, []int{7, 4}, []int{4, 22}, []int{7, 24}, []int{9, 26}, []int{13, 28}, []int{11, 26}},
		// 	},
		// 	[]int{-1, 19, 17, -1, -1, -1, 5, -1, 12, 3, 5, 5, 1, -1, 30, 5, 30, -1, -1, 24, 18, -1, 18, -1, 18, -1, 4, 29, 30, 12, -1, 29, 17, 22, 18, -1, 20, -1, 18, 18, 20},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setup(tt.args.ops, tt.args.vals, tt.want); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setup() = %v, want %v", got, tt.want)
			}
		})
	}
}
