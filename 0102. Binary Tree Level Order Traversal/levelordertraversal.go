package main

import . "leetcode/utils"

// Input: [3,9,20,null,null,15,7],
//
//	  3
//	 / \
//	9  20
//	  /  \
//	 15   7
//
// Output: [
//
//	[3],
//	[9,20],
//	[15,7]
//
// ]
// 按层序遍历 广度优先并使用队列记录每一层的结点
func levelOrder_bfs(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		tmp := make([]int, 0, n)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		queue = queue[n:] // 进入下一层之前去除当前层结点
		ans = append(ans, tmp)
	}
	return ans
}

func levelOrder_dfs(root *TreeNode) [][]int {
	var ans [][]int
	var dfsLevel func(node *TreeNode, level int)
	dfsLevel = func(node *TreeNode, level int) { // 递归传递层级信息
		if node == nil { // 递归终止条件：到达叶子结点
			return
		}
		if len(ans) == level { // 到达新一层
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], node.Val)
		dfsLevel(node.Left, level+1)
		dfsLevel(node.Right, level+1)
	}
	dfsLevel(root, 0)
	return ans
}
