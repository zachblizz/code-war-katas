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
func BuildTree([]interface{} input) *TreeNode {
	var root TreeNode

	for i := 0; i < len(input); i++ {
		j := 0
		var tmp TreeNode

		// not there yet...
		while (j <= 2) {
			switch(j) {
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
		}
	}

	return root
}

func createNode(int val) TreeNode {
	return TreeNode{val, nil, nil}
}

