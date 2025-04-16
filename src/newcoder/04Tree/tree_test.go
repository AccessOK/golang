package Tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 1
	tree.Left = new(TreeNode)
	tree.Left.Val = 2
	tree.Right = new(TreeNode)
	tree.Right.Val = 3
	//tree.Left.Left = new(TreeNode)
	//tree.Left.Left.Val = 4
	//tree.Left.Right = new(TreeNode)
	//tree.Left.Right.Val = 5
	res := maxDepth(tree)
	fmt.Printf("%v\n", res)
}
