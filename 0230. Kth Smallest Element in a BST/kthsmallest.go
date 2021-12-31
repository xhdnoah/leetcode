package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	inorderSort(root, k, &res, &count)
	return res
}

func inorderSort(node *TreeNode, k int, ans, count *int) {
	if node == nil {
		return
	}

	inorderSort(node.Left, k, ans, count)
	*count++
	if *count == k {
		*ans = node.Val
		return
	}
	inorderSort(node.Right, k, ans, count)
}
