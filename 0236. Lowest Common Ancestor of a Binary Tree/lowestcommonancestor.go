package main

import . "leetcode/utils"

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先
// 递归 + 后序遍历(需要先知道左右子树的情况，然后决定向上返回什么
// 若 root 是 p, q 的最近公共祖先 ，则只可能为以下三种情况之一：
// 1. p 和 q 在 root 的子树中，且分列 root 的异侧（即分别在左、右子树中）；
// 2. p = root 且 q 在 root 的左或右子树中; 3. q = root 且 p 在 root 的左或右子树中；
// 考虑通过递归对二叉树进行遍历，当遇到节点 p 或 q 时返回，从底至顶回溯
// 当节点 p, q 在节点 root 的异侧时，节点 root 即为最近公共祖先，则向上返回 root
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 越过叶节点返回 nil 遇到 pq 返回因为不能更深了
	if root == nil || root == p || root == q {
		return root
	}
	// root 非 pq 则继续分别往左右子树查找 pq
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 因为递归，使用函数后可认为左右子树已经算出结果，用 left 和 right 表示
	// left 子树没有查找到 pq 返回 right 子树的结果
	if left == nil {
		return right
	}
	// 同上
	if right == nil {
		return left
	}
	// 两边都查找到 pq 说明分别在左右子树中 此时最近公共祖先为 root
	return root
}
