package main

import "testing"

func Test_alphanumeric(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"should give false with '.*?'", args{".*?"}, false},
		{"should give true with 'a'", args{"a"}, true},
		{"should give true with 'Mazinkaiser'", args{"Mazinkaiser"}, true},
		{"should give false with 'hello world_'", args{"hello world_"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphanumeric(tt.args.str); got != tt.want {
				t.Errorf("alphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
