package main

import . "leetcode/utils"

// 深度优先递归 后入先出
func reversePrint(head *ListNode) []int {
	var dfs func(*ListNode) []int
	dfs = func(node *ListNode) []int {
		if node == nil {
			return []int{}
		}
		return append(dfs(node.Next), node.Val)
	}
	return dfs(head)
}

func reversePrint_stack(head *ListNode) (res []int) {
	stack := []int{}
	for node := head; node != nil; node = node.Next {
		stack = append(stack, node.Val)
	}
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	return
}
