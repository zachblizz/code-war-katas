package main

import "testing"

func TestLRUCache_Get(t *testing.T) {
	cache := Constructor(1)

	type args struct {
		key int
	}
	tests := []struct {
		name string
		l    *LRUCache
		args args
		want int
	}{
		{
			"Get 1",
			cache,
			args{1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Get(tt.args.key); got != tt.want {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
