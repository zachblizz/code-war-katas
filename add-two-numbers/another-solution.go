package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	root := &ListNode{}

	curr := root
	left := l1
	right := l2

	for {
		if left != nil {
			curr.Val += left.Val
			left = left.Next
		}

		if right != nil {
			curr.Val += right.Val
			right = right.Next
		}

		if curr.Val >= 10 {
			curr.Val = curr.Val - 10
			if left != nil {
				left.Val++
			} else if right != nil {
				right.Val++
			} else {
				left = &ListNode{
					Val: 1,
				}
			}
		}

		if (left != nil) || (right != nil) {
			curr.Next = &ListNode{}
			curr = curr.Next
			continue
		}

		return root
	}
}
