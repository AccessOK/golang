package Tree

func inorderTraversal(root *TreeNode) []int {
	// write code here
	res := &[]int{}
	traversal2(root, res)
	return *res
}
func traversal2(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	traversal2(root.Left, res)
	*res = append(*res, root.Val)
	traversal2(root.Right, res)
}
