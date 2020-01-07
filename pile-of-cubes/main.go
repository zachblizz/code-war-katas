package main

// kata - https://www.codewars.com/kata/5592e3bd57b64d00f3000047/train/go

// FindNb - foobar
func FindNb(m int) int {
	// your code
	n := 0
	count := 0

	for n < m {
		n += count * count * count
		count++
	}

	if n != m {
		return -1
	}

	return count - 1
}
