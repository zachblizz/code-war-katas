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

func swapNode(c, n *ListNode) {
	c.Next = n.Next
	n = c
}

func _reverseBetween(stack []*ListNode) {
	// update stack[0].Next = stack[len(stack)-1]
	start := 1
	end := len(stack)-1
	stack[0].Next = stack[end]

	for ; start < end; start++ {
		tmp := stack[start].Next
		stack[start].Next = stack[end].Next
		stack[end].Next = tmp
		tmp.Next = stack[start]
		end--
	}
	// if s == nil {
	// 	returncf4fc4d7238f982253543d0626b3ec7a4afe2243
	// }

	// tmp := s.Next
	// swapNode(s, s.Next)
	// _reverseBetween(tmp, e)
}

cf4fc4d7238f982253543d0626b3ec7a4afe2243

func reverseBetween(head *ListNode, m, n int) *ListNode {
	node := head
	start := false
	tmp := head
	stack := []*ListNode{}

	for node != nil {
		if node.Val == m {
			stack = append(stack, tmp)
			start = true
		} else if node.Val == n {
			start = false
		} 
		
		if (start || node.Val == n) {
			stack = append(stack, node)
		}

		tmp = node
		node = node.Next
	}

	_reverseBetween(stack)

	return head
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	head := createList(input)
	reverseBetween(head, 2, 5)

	printList(head)
}
