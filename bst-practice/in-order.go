package main

import "github.com/zachblizz/code-war-katas/utils"

func iterativeInOrder(root *utils.TreeNode) []int {
	list := []int{}
	stack := []int{}

	for root || len(stack) > 0 {
		for root {
			list = append(list, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		tmp := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		root = root.Right
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
	root := &utils.TreeNode{}
	utils.BuildTree([]int{1, -1, 2, 3}, root, 0)
	recursiveInOrder(root)
}
