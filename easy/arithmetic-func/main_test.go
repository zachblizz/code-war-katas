package main

import "testing"

func TestSquaresInRect(t *testing.T) {
	cases := []struct{
		a int
		b int
		want int
		op string
	}{
		{5,2,7,"add"},
		{5,2,3,"subtract"},
		{5,2,10,"multiply"},
		{5,2,2,"divide"},
	}

	for _, c := range cases {
		got := Arithmetic(c.a, c.b, c.op)

		if got != c.want {
			t.Errorf("%v does not equal %v (op: %v)", got, c.want, c.op)
		}
	}
}
