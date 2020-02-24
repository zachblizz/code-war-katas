package main

import (
	"fmt"

	"github.com/zachblizz/code-war-katas/bst-practice/utils"
)

func iterativeInOrder(root *utils.TreeNode) []int {
	list := []int{}
	stack := []*utils.TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			list = append(list, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		tmp := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		root = tmp.Right
	}

	return list
}

func recursiveInOrder(root *utils.TreeNode) []int {
	return _recursiveInOrder(root, []int{})
}

func _recursiveInOrder(root *utils.TreeNode, list []int) []int {
	if root == nil {
		return list
	}

	list = _recursiveInOrder(root.Left, list)
	list = append(list, root.Val)
	list = _recursiveInOrder(root.Right, list)

	return list
}

func main() {
	root := utils.BuildTree([]int{1, 2, 3, 4, 5, 6, 6, 6, 6}, utils.TreeNode{}, 0)
	fmt.Println(recursiveInOrder(root))
}
