package main

import . "leetcode/utils"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量
// BFS 利用队列记录广度优先扫描的当前层结点
func minDepth_BFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 1

	for len(q) > 0 {
		s := len(q)
		for i := 0; i < s; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil { // 叶子结点
				return depth
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return -1
}

// DFS 利用递归栈记录路径
func minDepth_DFS(root *TreeNode) int {
	if root == nil { // base case
		return 0
	}
	// 对于某一结点会面临三种情形 1 2 两种包含了都不存在的情形
	// 1. 左结点不存在
	if root.Left == nil {
		return minDepth_DFS(root.Right) + 1
	}
	// 2. 右结点不存在
	if root.Right == nil {
		return minDepth_DFS(root.Left) + 1
	}
	// 3. 左右子结点都存在
	return Min(minDepth_DFS(root.Left), minDepth_DFS(root.Right))
}
