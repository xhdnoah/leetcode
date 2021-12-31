package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given binary tree [3,9,20,null,null,15,7],
// 3
// / \
// 9  20
//  /  \
// 15   7
// return its depth = 3
// 遍历根节点的左孩子的高度和根节点右孩子的高度，取出两者的最大值+1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
