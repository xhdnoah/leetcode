package main

import . "leetcode/utils"

// Input: head = [1,2,3,4,5], left = 2, right = 4 Output: [1,4,3,2,5]
// 找到第一个需要反转的结点的前一个结点 p，从这个结点开始，依次把后面的结点用“头插”法，插入到 p 结点的后面。循环次数用 n-m 来控制
// 新建一个虚结点的作用将 left 为 0
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
	cur := pre.Next // (1->)2->3->4
	// 这一题结点可以原地变化，更改各个结点的 next 指针就可以。不需要游标 p 指针
	// 因为每次逆序以后，原有结点的相对位置就发生变化，cur 相当于游标指针移动，不需要再有 p = p.Next 的操作
	// 单向链表的重排操作上应该是*单向*顺序 1->3 再 2->4 错误的做法如 2->4 再 3->2 这样引用关系就乱了
	for i := 0; i < right-left; i++ {
		tmp := pre.Next          // 循环过程中依次为 2, 3
		pre.Next = cur.Next      // 把当前下个结点指给 pre 后面 1->3, 1->4
		cur.Next = cur.Next.Next // 把下下个结点指给当前结点后面 2->4, 2->5
		pre.Next.Next = tmp      // 当前结点再接上链表 (1->)3->2(->4), (1->)4->3(->2->5)
	}
	return newHead.Next
}

// Better: 头插法
func reverseBetween_head(head *ListNode, left, right int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	pre := dummyHead
	// pre 指向 left 的前置节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// Step1: p1, p2 用于区间内两两一翻转，执行 right-left 次
	p1 := pre.Next // 2
	p2 := p1.Next  // 3
	for i := 0; i < right-left; i++ {
		tmp := p2.Next // p2 断开后连接 4, 5
		p2.Next = p1   // p2 节点头插前指 3->2, 4->3->2
		p1 = p2        // 更新 p1 p2 (右移) 3, 4
		p2 = tmp       // 4, 5
	}
	// Step2: 连接，注意顺序
	pre.Next.Next = p2 // 2->5
	pre.Next = p1      // 1->4(->3->2->5)
	return dummyHead.Next
}
