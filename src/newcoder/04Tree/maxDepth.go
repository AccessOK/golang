package Tree

func maxDepth(root *TreeNode) int {
	// write code here
	res := traversql4(root, 1)
	return res
}

func traversql4(root *TreeNode, res int) int {
	if root == nil {
		return res - 1
	}
	left := traversql4(root.Left, res+1)
	right := traversql4(root.Right, res+1)
	if left > right {
		return left
	} else {
		return right
	}
}
