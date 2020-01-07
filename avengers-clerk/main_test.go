package main

import "testing"

func TestTickets(t *testing.T) {
	cases := []struct {
		in []int
		name string
		want string
	}{
		{[]int{25,25,50}, "first", "YES"},
		{[]int{25,100}, "second", "NO"},
		{[]int{25,25,50,50,100}, "third", "NO"},
		{[]int{25,25}, "fourth", "YES"},
		{[]int{25}, "fifth", "YES"},
		{[]int{50}, "sixth", "NO"},
		{[]int{100,100,100,100,100,100,100,100,100,100}, "seventh", "YES"},
		{[]int{25,25,25,25,50,100,50}, "eigth", "YES"},
	}

	for _, c := range cases {
		got := Tickets(c.in)

		if got != c.want {
			t.Errorf("[%v case] %v != %v", c.name, got, c.want)
		}
	}
}
