package main

// problem - https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to
import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	resMap := make(map[int][][]int)
	res := [][]int{}

	// p - the person (idx)
	// gs - group size they belong to
	for p, gs := range groupSizes {
		_groupThePeople(resMap, p, gs)
	}

	for _, v := range resMap {
		for _, g := range v {
			res = append(res, g)
		}
	}

	return res
}

func _groupThePeople(m map[int][][]int, p, gs int) {
	if _, ok := m[gs]; ok {
		formNewGroup := true
		for i, group := range m[gs] {
			if len(group) < gs {
				formNewGroup = false
				group = append(group, p)
				m[gs][i] = group
			}
		}

		if formNewGroup {
			m[gs] = append(m[gs], []int{p})
		}
	} else {
		m[gs] = [][]int{[]int{p}}
	}
}

// stole from someone's solution
func groupThePeopleNiceSolution(groupSizes []int) [][]int {
	res := [][]int{}
	x := make(map[int][]int)

	for i, groupSize := range groupSizes {
		if _, ok := x[groupSize]; !ok {
			x[groupSize] = []int{}
		}

		x[groupSize] = append(x[groupSize], i)

		if len(x[groupSize]) == groupSize {
			res = append(res, x[groupSize])
			x[groupSize] = []int{}
		}
	}

	return res
}

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}
