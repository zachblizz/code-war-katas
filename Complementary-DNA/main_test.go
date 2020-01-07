package main

import "testing"

func TestDNAStrand_test(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"given AAAA should return TTTT",
			"AAAA",
			"TTTT",
		},
		{
			"given ATTGC should return TAACG",
			"ATTGC",
			"TAACG",
		},
		{
			"given GTAT should return CATA",
			"GTAT",
			"CATA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DNAStrand(tt.input)

			if result != tt.expected {
				t.Errorf("DNAStrand returned unexpected complement: got %v expected %v",
					result, tt.expected)
			}
		})
	}
}
