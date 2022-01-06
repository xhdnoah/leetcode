package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: head = [1,2,3,4,5], left = 2, right = 4 Output: [1,4,3,2,5]
// 找到第一个需要反转的结点的前一个结点 p，从这个结点开始，依次把后面的结点用“头插”法，插入到 p 结点的后面。循环次数用 n-m 来控制
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < left-1; count++ {
		pre = pre.Next // 找到第一个需要反转结点的前结点
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next // 2->3->4
	// 这一题结点可以原地变化，更改各个结点的 next 指针就可以。不需要游标 p 指针
	// 因为每次逆序以后，原有结点的相对位置就发生变化，相当于游标指针移动，不需要再有 p = p.Next 的操作
	for i := 0; i < right-left; i++ {
		tmp := pre.Next          // 2
		pre.Next = cur.Next      // 把当前下个结点指给 pre 后面 1->3
		cur.Next = cur.Next.Next // 把下下个结点指给当前结点后面 2->4
		pre.Next.Next = tmp      // 当前结点再接上链表 (1->)3->2(->4)
	}
	return newHead.Next
}
