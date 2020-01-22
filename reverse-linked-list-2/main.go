package main

import "fmt"

// LC problem - https://leetcode.com/problems/reverse-linked-list-ii

// ListNode - the list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func createNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func createList(input []int) *ListNode {
	head := createNode(input[0])
	node := head

	for i := 1; i < len(input); i++ {
		node.Next = createNode(input[i])
		node = node.Next
	}

	return head
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Printf("%v->", node.Val)
		node = node.Next
	}

	fmt.Println()
}

func _reverseBetween2(stack []*ListNode, start, end int) {
	if start == end {
		return
	}

	if start == 0 {
		stack[start].Next = stack[end]
		start++
	}

	tmp := stack[start].Next
	stack[start].Next = stack[end].Next
	stack[end].Next = tmp
	tmp.Next = stack[start]

	_reverseBetween2(stack, start+1, end-1)
}

func _reverseBetween(stack []*ListNode) {
	// update stack[0].Next = stack[len(stack)-1]
	start := 1
	end := len(stack) - 1
	stack[0].Next = stack[end]

	for ; start < end; start++ {
		tmp := stack[start].Next
		stack[start].Next = stack[end].Next
		stack[end].Next = tmp
		tmp.Next = stack[start]
		end--
	}
}

func _reverseBetween3(head *ListNode, stack []*ListNode, start, end int) {
	if end < start {
		return
	}

	p := stack[start-1]
	s := stack[start]
	pe := stack[end-1]
	e := stack[end]

	tmp := s.Next
	s.Next = e.Next
	e.Next = tmp
	pe.Next = s
	p.Next = e

	printList(head)

	_reverseBetween3(head, stack, start+1, end-1)
}

func _reverseBetweenMine(stack []*ListNode) {
	s, e := 0, len(stack)-1

	for ; s < e; s++ {
		st := stack[s]
		end := stack[e]
		tmp := end.Val

		end.Val = st.Val
		st.Val = tmp

		e--
	}
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	node := head
	stack := []*ListNode{}
	idx := 1

	for node != nil {
		if idx >= m && idx <= n {
			stack = append(stack, node)
		} else if idx == n {
			break
		}

		node = node.Next
		idx++
	}

	_reverseBetweenMine(stack)

	return head
}

// var (
// 	l *ListNode
// 	stop bool
// )

// // it's position not value.......
// func reverseBetween(head *ListNode, m, n int) *ListNode {
// 	r := &ListNode{head.Val, head.Next}
// 	l = &ListNode{head.Val, head.Next}
// 	stop = false

// 	_rb(r, m, n)
// 	return head
// }

// func _rb(r *ListNode, m, n int) {
// 	if n == 1 {
// 		return
// 	}

// 	r = r.Next
	
// 	if m > 1 {
// 		l = l.Next
// 	}

// 	_rb(r, m-1, n-1)

// 	if l == r || r.Next == l {
// 		stop = true
// 	}

// 	if !stop {
// 		tmp := l.Val
// 		l.Val = r.Val
// 		r.Val = tmp
// 		l = l.Next
// 	}
// }

func main() {
	input := []int{1, 2, 3, 4, 5}
	// input = []int{3,5}
	head := createList(input)
	reverseBetween(head, 2, 4)

	printList(head)
}
