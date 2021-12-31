package main

// 给定一个数 x，比 x 大或等于的数字都要排列在比 x 小的数字后面，并且相对位置不能发生变化
// Input: head = 1->4->3->2->5->2, x = 3 Output: 1->2->2->4->3->5
// 简单做法是构造双向链表，时间复杂度 O(n^2) 更优是新构造 2 个链表，一个存储 < x 的结点，另一个存储 >= x 的结点

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{Val: 0, Next: nil}
	before := beforeHead // before 代表 < x 的结点 beforeHead.Next 指向第一个 before 结点; 初始化时两个指针变量指向同一块内存地址
	afterHead := &ListNode{Val: 0, Next: nil}
	after := afterHead // after 代表 >= x 的结点

	for head != nil {
		if head.Val < x {
			before.Next = head   // 这里的目的是将 beforeHead.Next 或上一个 "before"(指向的) 结点的 Next 指针指向当前结点以构造出链表
			before = before.Next // 再将 before 指向当前结点 此时 before head 指针指向同一块地址  before.Next 指向当前结点的下一个结点地址
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next // 最后一个 "before" 结点指向第一个 "after" 结点
	return beforeHead.Next
}

func main() {
	fifth := &ListNode{Val: 2, Next: nil}
	fourth := &ListNode{Val: 5, Next: fifth}
	third := &ListNode{Val: 2, Next: fourth}
	second := &ListNode{Val: 3, Next: third}
	first := &ListNode{Val: 4, Next: second}
	head := &ListNode{Val: 1, Next: first}
	partition(head, 3)
}
