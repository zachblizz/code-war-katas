package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"first", args{"()"}, true},
		{"second", args{"()[]{}"}, true},
		{"third", args{"(]"}, false},
		{"third and a half", args{"([)]"}, false},
		{"fourth", args{"{[]}"}, true},
		{"fifth", args{"("}, false},
		{"sixth", args{"(("}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidStack(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
