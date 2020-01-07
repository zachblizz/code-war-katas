package main

// import "fmt"
import "reflect"

// problem - https://leetcode.com/problems/minimum-time-visiting-all-points

func minTimeToVisitAllPoints(points [][]int) int {
	sum := 0

	for i, p1 := range points {
		if i+1 < len(points) {
			p2 := points[i+1]
			sum += getTimeFromPoints(p1, p2)
		}
	}

	return sum
}

func getTimeFromPoints(p1, p2 []int) int {
	time := 0

	for !reflect.DeepEqual(p1, p2) {
		if p1[0] < p2[0] {
			p1[0]++
		}

		if p1[1] < p2[1] {
			p1[1]++
		}

		time++
	}

	return time
}
