package main

import "fmt"

import "strings"

func getNumLetters(s string) (int, []rune) {
	for _, c := range s {
		ic := int(c)
		if ic >= 97 && ic <= 122 {
			fmt.Println(strings.ToUpper(string(c)))
		} else if ic >= 65 && ic <= 90 {
			fmt.Println(strings.ToLower(string(c)))
		}
	}
	return 1, []rune{}
}

func letterCasePermutation(S string) []string {
	getNumLetters(S)
	// m := make(map[string][]rune)

	return []string{}
}

func main() {
	fmt.Println(letterCasePermutation("a1zAB2"))
}
