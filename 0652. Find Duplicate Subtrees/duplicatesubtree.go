package main

import (
	"strconv"

	. "leetcode/utils"
)

// 后序遍历 序列化子树 寻找重复序列
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count := make(map[string]int, 1024)
	res := make([]*TreeNode, 0)

	helper(root, count, &res)
	return res
}

func helper(root *TreeNode, count map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "*"
	}
	l := helper(root.Left, count, res)
	r := helper(root.Right, count, res)

	key := strconv.Itoa(root.Val) + "," + l + r
	count[key]++

	if count[key] == 2 {
		*res = append(*res, root)
	}
	return key
}
