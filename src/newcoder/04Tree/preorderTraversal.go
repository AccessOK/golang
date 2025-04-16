package Tree

func preorderTraversal(root *TreeNode) []int {
	// write code here
	res := &[]int{}
	traversal1(root, res)
	return *res
}
func traversal1(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	//存储当前节点
	*res = append(*res, root.Val)
	//访问左子树
	traversal1(root.Left, res)
	//访问右子树
	traversal1(root.Right, res)
}
