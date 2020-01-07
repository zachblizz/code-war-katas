package main

import (
	"math"
	"testing"
)

func TestSquaresInRect(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 4, 0, 100, 4, 11, 2602, 36}, 11},
		{[]int{160, 3, 1719, 19, 11, 13, -21}, 160},
		{[]int{9, 3, 1719, 19, 11, 13, -21}, 0},
		{[]int{math.MaxInt32, 0, 1}, 0},
	}

	for _, c := range cases {
		got := FindOutlier(c.in)

		if got != c.want {
			t.Errorf("%v does not equal %v", got, c.want)
		}
	}
}
