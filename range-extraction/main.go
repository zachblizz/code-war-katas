package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// kata - https://www.codewars.com/kata/51ba717bb08c1cd60f00002f/train/go

// SolutionTakeOne - takes a list of integers in increasing order
// and returns a correctly formatted string in the range format
// For example:
// 	given: [12, 13, 15, 16, 17]
//	result: "12, 13, 15-17")
func SolutionTakeOne(list []int) string {
	var start string
	var result bytes.Buffer
	rangeCount := 1
	prev := list[0]
	haveStart := false
	
	for i := 1; i < len(list); i++ {
		curr := list[i]

		if prev+1 != curr || i == len(list) - 1 {
			strPrev := strconv.Itoa(prev)

			if i == len(list) - 1 && prev+1 == curr {
				strPrev = strconv.Itoa(curr)
			}

			if haveStart && rangeCount >= 3 {
				result.WriteString(fmt.Sprintf("%v-%v", start, strPrev))
			} else {
				if haveStart {
					result.WriteString(fmt.Sprintf("%v,", start))
				}
				
				result.WriteString(strPrev)
			}

			if i+1 < len(list) {
				result.WriteString(",")
			} else if !haveStart {
				result.WriteString(fmt.Sprintf(",%v", curr))
			}

			rangeCount = 1
			haveStart = false
		} else if prev != 0 && curr != 0 && !haveStart {
			rangeCount++
			haveStart = true
			start = strconv.Itoa(prev)
		} else if haveStart {
			rangeCount++
		}

		prev = curr
	}

	return result.String()
}

// Solution - takes a list of integers in increasing order
// and returns a correctly formatted string in the range format
func Solution(list []int) string {
	var start int
	var result bytes.Buffer
	prev := list[0]
	haveStart := false

	for i := 1; i < len(list); i++ {
		curr := list[i]

		if prev+1 != curr && i != len(list) - 1 {
			if haveStart {
				rangeCount := 1+prev-start

				if rangeCount >= 3 {
					// might not want to do %v-%v,
					// might want to just print prev
					// and have "start-" in the result already
					result.WriteString(fmt.Sprintf("%v-%v,", start, prev))
				} else {
					// if we put "start-" in the result already,
					// we'll need to replace the "-" with a ","
					result.WriteString(fmt.Sprintf("%v,%v,", start, prev))
				}

				haveStart = false
			} else {
				result.WriteString(fmt.Sprintf("%v,", prev))
			}
		} else if prev+1 == curr && !haveStart && i != len(list) - 1 {
			// might want to put "start-" here...
			start = prev
			haveStart = true
		} else if i == len(list) - 1 {
			// TODO: this doesn't work for negative numbers...
			if haveStart && curr-prev == 1 {
				rangeCount := 1+curr-start

				if rangeCount >= 3 {
					result.WriteString(fmt.Sprintf("%v-%v", start, curr))
				} else {
					result.WriteString(fmt.Sprintf("%v,%v", start, curr))
				}
			} else if haveStart && curr-prev != 1 {
				result.WriteString(fmt.Sprintf("%v,%v,%v", start, prev, curr))
			} else {
				result.WriteString(fmt.Sprintf("%v,%v", prev, curr))
			}
		}

		prev = curr
	}

	return result.String()
}
