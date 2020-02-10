package main

// lc - https://leetcode.com/problems/integer-to-roman/

import (
	"bytes"
	"fmt"

)

func intToRoman(num int) string {
	sym := map[int]string{
		1: "I", 2: "II", 3: "III", 5: "V", 6: "VI",
		7: "VII", 8: "VIII", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M",
		4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM",
	}

	m := 1
	arr := []string{}
	var ret bytes.Buffer

	for num > 0 {
		r := num % 10 * m
		num /= 10
		m *= 10

		if _, ok := sym[r]; ok {
			arr = append(arr, sym[r])
		} else if r > 10 {
			arr = append(arr, calculateNumber(r, sym))
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		ret.WriteString(arr[i])
	}

	return ret.String()
}

func calculateNumber(r int, sym map[int]string) string {
	var ret bytes.Buffer
	sub := 10
	m := make(map[int]int)

	m[10] = 0
	m[100] = 0
	m[1000] = 0

	for (r > 10 && r < 100) || (r > 100 && r < 1000) || r > 1000 {
		sub = 10
		if r >= 100 && r < 1000 {
			sub = 100
		} else if r >= 1000 {
			sub = 1000
		}

		n := r - sub

		if s, ok := sym[n]; ok {
			ret.WriteString(s)
			r -= n
		} else {
			m[sub]++
			r -= sub
		}
	}

	for i := m[sub]; i >= 0; i-- {
		ret.WriteString(sym[sub])
	}

	return ret.String()
}

// better solution yoinked from someone...
func intToRoman2(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	var result bytes.Buffer
	for i := range nums {
		for num >= nums[i] {
			num -= nums[i]
			result.WriteString(numMap[nums[i]])
		}
	}
	return result.String()
}

func main() {
	fmt.Println(intToRoman(109993))
	fmt.Println(intToRoman2(109993))
}
