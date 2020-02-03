package main

import (
	"reflect"
	"testing"
)

func TestSolequa(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"first", args{5}, [][]int{{3,1}}},
		{"second", args{12}, [][]int{{4, 1}}},
		{"third", args{13}, [][]int{{7, 3}}},
		{"fourth", args{9005}, [][]int{{4503, 2251}, {903, 449}}},
		{"fifth", args{9008}, [][]int{{1128, 562}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solequa(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solequa() = %v, want %v", got, tt.want)
			}
		})
	}
}
