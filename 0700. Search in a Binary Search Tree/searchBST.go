package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		searchBST(root.Left, val)
	}
	if root.Val < val {
		searchBST(root.Right, val)
	}
	return root
}
