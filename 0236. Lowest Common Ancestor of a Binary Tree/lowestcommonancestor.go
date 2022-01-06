package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先
// 递归+后序遍历(需要先知道左右子树的情况，然后决定向上返回什么
// 若 root 是 p, q 的 最近公共祖先 ，则只可能为以下情况之一：
// p 和 q 在 root 的子树中，且分列 root 的 异侧（即分别在左、右子树中）；
// p = root 且 q 在 root 的左或右子树中; q = root 且 p 在 root 的左或右子树中；
// 考虑通过递归对二叉树进行遍历，当遇到节点 p 或 q 时返回。从底至顶回溯
// 当节点 p, q 在节点 root 的异侧时，节点 root 即为最近公共祖先，则向上返回 root
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
