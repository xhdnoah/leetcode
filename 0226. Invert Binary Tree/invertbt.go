package main

import . "leetcode/utils"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 前序遍历位置 交换左右子结点
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
