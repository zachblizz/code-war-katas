package main

import "fmt"

// kata - https://www.codewars.com/kata/526dbd6c8c0eb53254000110/train/go

func alphanumeric(str string) bool {
	result := false

	for _, x := range str {
		xInt := int(x)
		if (xInt >= 65 && xInt <= 90) || (xInt >= 97 && xInt <= 122) || (xInt >= 48 && xInt <= 57) {
			result = true
		} else {
			result = false
			break
		}
	}
	
	return result
}

func main() {
	fmt.Println(alphanumeric(".*?"))
}
