package main

func preorderTraversal(root *TreeNode) []int {
	// write code here
	res := []int{}
	p := root
	for p != nil {
		res = append(res, p.Val)
		if(p.Left.Val!=nil)p = p.Left
	}

	return res
}
