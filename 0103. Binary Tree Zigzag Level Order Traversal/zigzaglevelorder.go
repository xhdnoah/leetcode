package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 按照 Z 字型层序遍历一棵树
//     3
//    / \
//   9  20
//  /    \
// 15    7
// return its zigzag level order traversal as:
// [
//   [3],
//   [20,9],
//   [15,7]
// ]
// 按层序从上到下遍历一颗树，但是每一层的顺序是相互反转的
// 即上一层从左往右，下一层就从右往左，以此类推。可用队列实现
func zigzagLevelOrder_recursive(root *TreeNode) [][]int {
	var res [][]int
	search(root, 0, &res)
	return res
}

// 前序遍历
func search(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	for len(*res) < depth+1 {
		*res = append(*res, []int{})
	}
	if depth%2 == 0 { // 偶数层 往后 append
		(*res)[depth] = append((*res)[depth], root.Val)
	} else { // 奇数层 往前 append
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	}
	search(root.Left, depth+1, res)
	search(root.Right, depth+1, res)
}
