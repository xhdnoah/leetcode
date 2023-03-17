package main

import . "leetcode/utils"

// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7] Output: [3,9,20,null,null,15,7]
// 	  3
// 	 / \
//  9  20
// 	  / \
// 	 15  7
// 对任意树 preorder 形式总是 [root, 左子树 preorder, 右子树 preorder] 
// inorder 形式总是为 [左子树 inorder, root, 右子树 inorder]
func buildTree_recursive(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历的第一个即为根节点
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	// 在中序遍历中定位根节点以获取左右子树的节点数
	// 而同一棵子树的 preorder inorder 长度相同 使用该长度对应到 preorder 中定位左右子树
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	leftTree, rightTree := inorder[:i], inorder[i+1:]
	root.Left = buildTree_recursive(preorder[1:i+1], leftTree)
	root.Right = buildTree_recursive(preorder[i+1:], rightTree)
	return root
}

// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7] Output: [3,9,20,null,null,15,7]
// 	  3
// 	 / \
//  9  20
// 	  / \
// 	 15  7
// 一个 stack 维护「当前节点的所有还没有考虑右儿子的祖先节点」栈顶就是当前节点 即只有栈中的节点才能连接一个新的右儿子
// 一个指针指向 inorder 第一个节点 依次枚举 preorder 中除第一个节点外的每一个 如果 index 恰好指向栈顶 那么不断弹出栈顶并右移 index
func buildTree_iterative(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			ls := len(stack)
			for ls > 0 && stack[ls-1].Val == inorder[inorderIndex] {
				node = stack[ls-1]
				stack = stack[:ls-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
