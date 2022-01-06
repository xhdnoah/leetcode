package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Input: root = [1,2,3,null,5,null,4] Output: [1,3,4]
// Explanation:
//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---
// 按层序遍历元素，依次取每层最右元素。用一个队列实现
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[n-1].Val)
		queue = queue[n:]
	}
	return res
}

// 前序遍历+递归，用 depth 记录深度过滤左侧元素
func rightSideView_recursive(root *TreeNode) (res []int) {
	var traverse func(depth int, node *TreeNode)
	traverse = func(depth int, node *TreeNode) {
		if node == nil {
			return
		}
		if depth > len(res) {
			res = append(res, node.Val)
		}
		traverse(depth+1, node.Right)
		traverse(depth+1, node.Left)
	}

	traverse(1, root)
	return res
}
