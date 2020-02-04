package main

import (
	"bytes"
	"fmt"
)

import . "strings"

// CamelCase - camel cases the string provided
func CamelCase(s string) string {
	var ret bytes.Buffer
	prevSpace := false

	for i, c := range s {
		if c != 32 && (prevSpace || i == 0) {
			ret.WriteString(string(c - 32))
		} else if c != 32 {
			ret.WriteString(string(c))
		}

		prevSpace = c == 32
	}

	return ret.String()
}

func camelCaseSpecial(s string) string {
	return Replace(Title(s)," ","",-1)
}

func anotherCamelCase(s string) string {
	return Join(Fields(Title(s)), "")
}

func main() {
	fmt.Printf(":%v:\n", anotherCamelCase(" camel case word"))
}
