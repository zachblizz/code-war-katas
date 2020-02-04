package main

// lc - https://leetcode.com/problems/add-two-numbers/

import "fmt"

// ListNode - the list node...
type ListNode struct {
	Val int
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

func findLarger(l1, l2 *ListNode) (*ListNode, *ListNode) {
	l1S := 0
	l2S := 0
	tmp1 := l1
	tmp2 := l2

	for tmp1 != nil {
		l1S++
		tmp1 = tmp1.Next
	}

	for tmp2 != nil {
		l2S++
		tmp2 = tmp2.Next
	}

	if l1S > l2S {
		return l1, l2
	}

	return l2, l1
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	of := 0
	head := &ListNode{Val: -1, Next: nil}
	node := &ListNode{Val: -1, Next: nil}

	o, t := findLarger(l1, l2)

	for o != nil {
		if t != nil {
			val := t.Val + o.Val + of

			if val >= 10 {
				of = 1
				val %= 10
			} else {
				of = 0
			}

			if head.Val == -1 {
				head = &ListNode{Val: val, Next: nil}
				node = head
			} else {
				node.Next = &ListNode{Val: val, Next: nil}
				node = node.Next
			}

			t = t.Next
		} else {
			val := o.Val + of

			if val >= 10 {
				val %= 10
				of = 1
			} else {
				of = 0
			}

			node.Next = &ListNode{Val: val, Next: nil}
			node = node.Next
		}

		o = o.Next
		// node = node.Next
	}

	if of == 1 {
		node.Next = &ListNode{Val: of, Next: nil}
	}
	
	return head
}

func main() {
	l1 := createList([]int{5,6})
	l2 := createList([]int{2,4,3})

	printList(addTwoNumbers(l1, l2))
}
