package main

import . "leetcode/utils"

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	// 后序遍历位置 此时左右子树已经拉平
	left := root.Left
	right := root.Right

	// 左子树作为右子树
	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil { // for 循环直至现右子树末端
		p = p.Right
	}
	p.Right = right // 原右子树接至现右子树末端
}
