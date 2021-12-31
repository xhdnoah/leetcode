package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given this linked list: 1->2->3->4->5
// For k = 2, you should return: 2->1->4->3->5 For k = 3, return: 3->2->1->4->5
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ { // 找 k-Group 尾结点
		if node == nil { // 不足 k 个结点返回 head
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k) // 此时 head 为上轮 reverse 的尾结点
	return newHead
}

// 反转 [first, last) 区间(注意左闭右开)元素
func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	// prev: 4 first: 1->2->3->4 => prev: 1->4 first: 2->3->4 => prev: 2->1->4 first: 3->4 => prev: 3->2->1->4 first: 4(last)
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first // prev 不断接收 first 结点指向 last 构造出反转链表
		first = tmp  // first 不断往 last 方向移动
	}
	return prev
}

func main() {
	fourth := &ListNode{Val: 5, Next: nil}
	third := &ListNode{Val: 4, Next: fourth}
	second := &ListNode{Val: 3, Next: third}
	first := &ListNode{Val: 2, Next: second}
	head := &ListNode{Val: 1, Next: first}
	reverseKGroup(head, 3)
}
