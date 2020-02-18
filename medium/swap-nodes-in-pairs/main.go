package main

import "fmt"

// lc - https://leetcode.com/problems/swap-nodes-in-pairs/

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
		fmt.Printf("%v", node.Val)
		if node.Next != nil {
			fmt.Printf("->")
		}

		node = node.Next
	}

	fmt.Println()
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := head.Next
	for head != nil {
		n := head.Next
		if n != nil {
			nn := n.Next

			if nn != nil && nn.Next != nil {
				head.Next = nn.Next
			} else {
				head.Next = nn
			}

			n.Next = head
			head = nn
		} else {
			return ret
		}
	}

	return ret
}

func main() {
	head := createList([]int{1, 2, 3, 4})
	printList(swapPairs(head))

	head = createList([]int{1, 2})
	printList(swapPairs(head))

	head = createList([]int{1, 2, 3})
	printList(swapPairs(head))

	head = createList([]int{1, 2, 3, 4, 5})
	printList(swapPairs(head))
}
