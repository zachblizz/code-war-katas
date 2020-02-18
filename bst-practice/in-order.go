package main

import "github.com/zachblizz/code-war-katas/utils"

func iterativeInOrder(root *TreeNode) []int {
	list := []int{}
	// TODO

	return list
}

func recursiveInOrder(root *TreeNode) []int {
	return _recursiveInOrder(root, []int{})
}

func _recursiveInOrder(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}

	list = _recursiveInOrder(root.Left, list)
	list = append(list, root.Val)
	list = _recursiveInOrder(root.Right, list)

	return list
}

func main() {
	root := &TreeNode{}
	BuildTree([]int{1, -1, 2, 3}, root, 0)
	recursiveInOrder(root)
}
