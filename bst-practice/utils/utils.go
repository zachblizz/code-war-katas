package utils

// TreeNode - the tree node...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree - builds the tree given the input array
func BuildTree(input []int, root TreeNode, i int) *TreeNode {
	if i < len(input) {
		root = createNode(input[i])

		BuildTree(input, root.Left, 2*i+1)
		BuildTree(input, root.Right, 2*i+2)
	}
	return &root
}

func createNode(val int) *TreeNode {
	if val == -1 {
		return nil
	}

	return &TreeNode{val, nil, nil}
}
