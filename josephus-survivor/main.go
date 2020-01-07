package main

// kata - https://www.codewars.com/kata/josephus-survivor/train/go

// JosephusSurvivor - garbage...
func JosephusSurvivor(n, k int) int {
	list := make([]int, n)
	count := 1
	i := 0

	for i := 1; i <= n; i++ {
		list[i-1] = i
	}

	for len(list) > 1 {
		if count < k {
			count++

			i = (i+1)%len(list)
		} else {
			count = 1
			list = remove(list, i)
		}
	}

	return list[0]
}

// JoeTwo - WOW
func JoeTwo(n, k int) int {
	if n == 1 {return 1}
  	return (JoeTwo(n-1, k)+k-1)%n + 1
}

func remove(s []int, x int) []int {
	return append(s[:x], s[x+1:]...)
}
