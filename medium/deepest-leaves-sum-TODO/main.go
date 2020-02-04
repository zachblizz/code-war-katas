package main 

import "fmt"

type treeNode struct {
	Val int
	Left *treeNode
	Right *treeNode
}

func newNode(val int) *treeNode {
	node := treeNode{val, nil, nil}
	return &node
}

func buildTree(nums []int, start, end int) *treeNode {
	if start > end {
		return nil
	}

	mid := (start+end)/2
	root := newNode(nums[mid])

	root.Left = buildTree(nums, start, mid-1)
	root.Right = buildTree(nums, mid+1, end)

	return root
}

func printTree(node *treeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Val)
	printTree(node.Left)
	printTree(node.Right)
}

func main() {
	nums := []int{1,2,3,4,5,-1,6,7,-1,-1,-1,-1,8}
	root := buildTree(nums, 0, len(nums)-1)

	printTree(root)
}
