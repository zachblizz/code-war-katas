package bst

// TreeNode - the tree node...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
  [1,nil,2,3]

       1
	    \
		 2
	   /
	  3
*/
// BuildTree - builds the tree given the input array
func BuildTree(input []int) *TreeNode {
	var root TreeNode

	for i := 0; i < len(input); i++ {
		j := 0
		var tmp TreeNode

		// not there yet...
		for ; j <= 2 && i < len(input); j++ {
			switch j {
			case 0:
				tmp = createNode(input[i])
				break
			case 1:
				tmp.Left = createNode(input[i])
				break
			case 2:
				tmp.Right = createNode(input[i])
				break
			}
			i++
		}
	}

	return root
}

func createNode(int val) TreeNode {
	if val == -1 {
		return nil
	}
	return TreeNode{val, nil, nil}
}

func main() {
	BuildTree([]int{1, -1, 2, 3})
}
