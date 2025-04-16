package Tree

func postorderTraversal(root *TreeNode) []int {
	// write code here
	res := &[]int{}
	traversal3(root, res)
	return *res
}
func traversal3(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	traversal3(root.Left, res)
	traversal3(root.Right, res)
	*res = append(*res, root.Val)
}
