package main

import . "leetcode/utils"

func isValidBST(root *TreeNode) bool {
	return validateWithRestriction(root, nil, nil)
}

// 限定以 root 为根的子树结点必须满足 max.Val > root.Val > min.Val
func validateWithRestriction(root, min, max *TreeNode) bool {
	if root == nil { // base case
		return true
	}
	// 先序遍历位置判断是否合法
	if min != nil && root.Val <= min.Val ||
		max != nil && root.Val >= max.Val {
		return false
	}
	// 限定左子树最大值和右子树的最小值
	return validateWithRestriction(root.Left, min, root) &&
		validateWithRestriction(root.Right, root, max)
}
