package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7] Output: [3,9,20,null,null,15,7]
// 	  3
// 	 / \
//  9  20
// 	  / \
// 	 15  7
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, 0, len(preorder)-1, len(inorder)-1)
}

func build(preorder, inorder []int, preStart, inStart, preEnd, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	index, rootVal := 0, preorder[preStart]
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
		}
	}
	leftSize := index - inStart
	root := &TreeNode{Val: rootVal}
	root.Left = build(preorder, inorder, preStart+1, inStart, preStart+leftSize, index-1)
	root.Right = build(preorder, inorder, preStart+leftSize+1, index+1, preEnd, inEnd)

	return root
}
