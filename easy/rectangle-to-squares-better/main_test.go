package main

import "testing"

func TestSquaresInRect(t *testing.T) {
	cases := []struct{
		l int
		w int
		want []int
	}{
		{5,3,[]int{3,2,1,1}},
		{5,5,nil},
	}

	for _, c := range cases {
		got := SquaresInRect(c.l, c.w)

		if len(got) != len(c.want) {
			t.Errorf("length of SquaresInRect(%v,%v) == %d, want %d", c.l, c.w, len(got), len(c.want))
		} else {
			for i, g := range got {
				if g != c.want[i] {
					t.Errorf("SquaresInRect(%v,%v) index %d of got (%d) != %d", c.l, c.w, i, g, c.want[i])
				}
			}
		}
	}
}

type w struct {
	goodSquare bool
	newP Point
}

func TestCanMakeSquare(t *testing.T) {
	cases := []struct{
		square int
		p Point
		rectL int
		rectW int
		want w
	}{
		{3,Point{0,0},5,3,w{true, Point{0,3}}},
		{2,Point{0,3},5,3,w{true, Point{2,3}}},
		{3,Point{0,0},3,5,w{true, Point{3,0}}},
	}

	for _, c := range cases {
		goodSquare, gotP := canMakeSquare(c.square, c.rectL, c.rectW, c.p)

		if goodSquare != c.want.goodSquare {
			t.Errorf("p: %v, s: %v; %v != %v", c.p, c.square, goodSquare, c.want.goodSquare)
		}

		if gotP.r != c.want.newP.r {
			t.Errorf("wrong r %v wanted %v", gotP.r, c.want.newP.r)
		}

		if gotP.c != c.want.newP.c {
			t.Errorf("wrong c %v wanted %v", gotP.c, c.want.newP.c)
		}
	}
}
