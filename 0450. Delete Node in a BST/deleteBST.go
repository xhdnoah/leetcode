package main

import . "leetcode/utils"

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	minNode := getMin(root.Right)
	root.Right = deleteNode(root.Right, minNode.Val)
	minNode.Left = root.Left
	minNode.Right = root.Right
	root = minNode

	return root
}

func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
