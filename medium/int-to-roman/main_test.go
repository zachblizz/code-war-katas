package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"first",
			args{20},
			"XX",
		},
		{
			"second",
			args{1994},
			"MCMXCIV",
		},
		{
			"third",
			args{70},
			"LXX",
		},
		{
			"fourth",
			args{200},
			"CC",
		},
		{
			"fifth",
			args{2000},
			"MM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
