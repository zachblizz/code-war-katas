package main

import "testing"

func TestFindNb(t *testing.T) {
	cases := []struct {
		in int
		want int
	}{
		{4183059834009, 2022,},
		{24723578342962, -1,},
	}

	for _, c := range cases {
		got := FindNb(c.in)

		if got != c.want {
			t.Errorf("got %v but wanted %v", got, c.want)
		}
	}
}
