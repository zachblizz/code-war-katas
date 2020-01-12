package main

import (
	"reflect"
	"testing"
)

func Test_numSmallerByFrequence(t *testing.T) {
	type args struct {
		queries []string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"first",
			args{[]string{"cbd"}, []string{"zaaaz"}},
			[]int{1},
		},
		{
			"second",
			args{
				[]string{"bbb", "cc"},
				[]string{"a", "aa", "aaa", "aaaa"},
			},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSmallerByFrequency(tt.args.queries, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSmallerByFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
