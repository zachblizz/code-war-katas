package main

import "testing"

// func TestSquaresInRect(t *testing.T) {
// 	cases := []struct{
// 		l int
// 		w int
// 		want []int
// 	}{
// 		{5,3,[]int{3,2,1,1}},
// 		{5,5,nil},
// 	}

// 	for _, c := range cases {
// 		got := SquaresInRect(c.l, c.w)

// 		if len(got) != len(c.want) {
// 			t.Errorf("length of SquaresInRect(%v,%v) == %d, want %d", c.l, c.w, len(got), len(c.want))
// 		} else {
// 			for i, g := range got {
// 				if g != c.want[i] {
// 					t.Errorf("SquaresInRect(%v,%v) index %d of got (%d) != %d", c.l, c.w, i, g, c.want[i])
// 				}
// 			}
// 		}
// 	}
// }

func TestGetLargestPossibleSquare(t *testing.T) {
	cases := []struct {
		l int
		w int
		want int
	}{
		{5,3,3},
		{3,5,3},
		{20,14,14},
	}

	for _, c := range cases {
		got := getLargestPossibleSquare(c.l, c.w)

		if got != c.want {
			t.Errorf("getLargestPossibleSquare(%d, %d) == %d, want %d", c.l, c.w, got, c.want)
		}
	}
}

func TestFillOutSquare(t *testing.T) {
	cases := []struct {
		square int
		init [][]int
		want [][]int
	}{
		{
			3,
			[][]int{
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
			},
			[][]int{
				{1,1,1,0,0},
				{1,1,1,0,0},
				{1,1,1,0,0},
			},
		},
		{
			2,
			[][]int{
				{1,1,1,0,0},
				{1,1,1,0,0},
				{1,1,1,0,0},
			},
			[][]int{
				{1,1,1,1,1},
				{1,1,1,1,1},
				{1,1,1,0,0},
			},
		},
		{
			1,
			[][]int{
				{1,1,1,1,1},
				{1,1,1,1,1},
				{1,1,1,0,0},
			},
			[][]int{
				{1,1,1,1,1},
				{1,1,1,1,1},
				{1,1,1,1,0},
			},
		},
		{
			1,
			[][]int{
				{1,1,1,1,1},
				{1,1,1,1,1},
				{1,1,1,1,0},
			},
			[][]int{
				{1,1,1,1,1},
				{1,1,1,1,1},
				{1,1,1,1,1},
			},
		},
	}

	for _, c := range cases {
		fillOutSquare(c.square, c.init)

		for i, inner := range c.want {
			for j := range inner {
				w := inner[j]
				g := c.init[i][j]

				if w != g {
					t.Errorf("index %d %d of got = %d, but want %d", i, j, g, w)
				}
			}
		}
	}
}

func TestGetStartRowAndCol(t *testing.T) {
	cases := []struct {
		rect [][]int
		square int
		row int
		col int
	}{
		{
			[][]int{
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
			},
			3,
			0,
			0,
		},
		{
			[][]int{
				{1,1,1,0,0},
				{1,1,1,0,0},
				{1,1,1,0,0},
			},
			2,
			0,
			3,
		},
		{
			[][]int{
				{1,1,1,1,1},
				{1,1,1,1,1},
				{1,1,1,0,0},
			},
			1,
			2,
			3,
		},
	}

	for i, c := range cases {
		sR, sC, _, _ := getStartRowAndCol(c.square, c.rect)

		if sR != c.row {
			t.Errorf("%v startRow %v != %v", i, sR, c.row)
		}

		if sC != c.col {
			t.Errorf("%v startCol %v != %v", i, sC, c.col)
		}
	}
}
