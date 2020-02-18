package utils

// TreeNode - the tree node...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree - builds the tree given the input array
/* e.g.:
  [1,nil,2,3]

	1
	 \
	  2
	 /
   3
*/
func BuildTree(input []int, root *TreeNode, i int) {
	if i < len(input) {
		tmp := createNode(input[i])
		root = tmp

		BuildTree(input, root.Left, 2*i+1)
		BuildTree(input, root.Right, 2*i+2)
	}
}

func createNode(val int) *TreeNode {
	if val == -1 {
		return nil
	}

	return &TreeNode{val, nil, nil}
}

//func main() {
//BuildTree([]int{1, -1, 2, 3})
//}
