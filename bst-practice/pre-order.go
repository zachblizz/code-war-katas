package main

func iterativePreOrder(root *TreeNode) []int {
	list := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			list = append(list, root.Val)
			root = root.Left
		}

		tmp := stack[len(stack)-1]      // pop head
		stack = stack[0 : len(stack)-1] // remove head from stack...
		root = tmp.Right
	}

	return list
}

func recursivePreOrder(root *TreeNode) []int {
	return _recursivePreOrder(root, []int{})
}

func _recursivePreOrder(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}

	list = append(list, root.Val)
	list = _recursivePreOrder(root.Left, list)
	list = _recursivePreOrder(root.Right, list)

	return list
}
