package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectSibling(root.Left, root.Right)
	return root
}

func connectSibling(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right
	connectSibling(left.Left, left.Right)
	connectSibling(right.Left, right.Right)
	connectSibling(left.Right, right.Left)
}
