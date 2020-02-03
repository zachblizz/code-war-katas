package main

// kata - https://www.codewars.com/kata/555615a77ebc7c2c8a0000b8/train/go

// Tickets - returns YES if we can NO if we can't
func Tickets(peopleInLine []int) string {
	remaining := 0

	for _, amount := range peopleInLine {
		if amount == 25 {
			remaining += amount
		} else {
			subAmt := amount - 25
			remaining -= subAmt
			remaining += amount
		}
	}

	if remaining >= 0 {
		return "YES"
	}

	return "NO"
}
