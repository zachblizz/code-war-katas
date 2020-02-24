package main

import "fmt"

func isPalendrom(str string) bool {
	j := len(str) - 1

	for i := 0; i < j; i++ {
		if str[i] != str[j] {
			return false
		}

		j--
	}
	return true
}

func main() {
	fmt.Println(isPalendrom("dod"))
	fmt.Println(isPalendrom("racecar"))
	fmt.Println(isPalendrom("test"))
	fmt.Println(isPalendrom("pizzazzip"))
}
