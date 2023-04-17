package main

import . "leetcode/utils"

// Input: [1,null,2,3]
//
//	1
//	 \
//	  2
//	 /
//	3
//
// Output: [1,3,2]
// Inorder Traversal: Left, Root, Right
func inorderTraversal_recursive(root *TreeNode) []int {
	ans := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return ans
}

// 迭代法 模拟实现递归栈
func inorderTraversal_iterative(root *TreeNode) (ans []int) {
	stack := []*TreeNode{}
	// root 为空且 stack 为空，遍历结束
	for root != nil || len(stack) > 0 {
		// 每到一个节点 A，因为根的访问在中间，入栈 A 然后遍历左子树，接着访问 A 最后遍历右子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 访问完 A 后 A 就可以出栈了 因为 A 和其左子树都已经访问完成
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 弹出栈
		ans = append(ans, root.Val)  // 加入结果集
		root = root.Right // 遍历右子树 如果不存在则在下一轮循环取父结点弹出栈
	}
	return
}
