package main

import "testing"

func TestCamelCase(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"test case", "TestCase"},
		{"camel case method", "CamelCaseMethod"},
		{"say hello ","SayHello"},
        {" camel case word","CamelCaseWord"},
        {"",""},
	}

	for _, c := range cases {
		got := CamelCase(c.in)

		if got != c.want {
			t.Errorf("CamelCase(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
