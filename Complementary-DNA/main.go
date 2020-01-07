package main

import (
	"bytes"
	"fmt"
)

// DNAStrand - gives the complement on the given dna strand
func DNAStrand(dna string) string {
	var compDNA bytes.Buffer

	for _, x := range dna {
		switch x {
		case 'A':
			compDNA.WriteString("T")
		case 'T':
			compDNA.WriteString("A")
		case 'C':
			compDNA.WriteString("G")
		case 'G':
			compDNA.WriteString("C")
		}
	}

	return compDNA.String()
}

func main() {
	fmt.Println(DNAStrand("AAAA"))
}
