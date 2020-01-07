package main

import "testing"

func Test_BalancedStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{"RLRRLLRLRL"}, 4},
		{"3", args{"RLLLLRRRLR"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BalancedStringSplit(tt.args.s); got != tt.want {
				t.Errorf("BalancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
