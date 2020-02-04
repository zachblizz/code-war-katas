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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setup(tt.args.ops, tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setup() = %v, want %v", got, tt.want)
			}
		})
	}
}
