package main

import "testing"

func TestJosephusSurvivor(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{7, 3}, 4},
		{"second", args{11, 19}, 10},
		{"third", args{40, 3}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JosephusSurvivor(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("JosephusSurvivor() = %v, want %v", got, tt.want)
			}
		})
	}
}
