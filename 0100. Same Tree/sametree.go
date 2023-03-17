package main

import . "leetcode/utils"

// 对称性递归：特殊处理 + 返回值
// 特殊：左右都 nil；返回：左右都不空 + root,left,right 相同
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	return p != nil && q != nil && p.Val == q.Val &&
		isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
