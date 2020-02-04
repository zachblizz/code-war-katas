package main

// problem  - https://leetcode.com/problems/di-string-match

func diStringMatch(S string) []int {
	var match []int
	match = append(match, helper(S, 0, 0, len(S), &match))

	return match
}

func helper(S string, i, min, max int, m *[]int) int {
	if i == len(S) {
		if string(S[i-1]) == "I" {
			return min
		}

		return max
	}

	if string(S[i]) == "I" {
		*m = append(*m, min)
		min++
	} else {
		*m = append(*m, max)
		max--
	}

	i++

	return helper(S, i, min, max, m)
}
