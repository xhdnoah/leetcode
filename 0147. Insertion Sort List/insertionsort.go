package main

import . "leetcode/utils"

// Input: -1->5->3->4->0 Output: -1->0->3->4->5
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 为在循环体内统一，此处初始化不直接指向 head
	newHead := &ListNode{Val: 0, Next: nil}
	cur, pre := head, newHead
	for cur != nil {
		next := cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next // 插入时比对大小，不满足排序则右移 pre 直至满足
		}
		cur.Next = pre.Next // cur 结点插入新链表
		pre.Next = cur      // 接上 pre 结点
		pre = newHead       // 归位重头开始
		cur = next          // cur 游标右移
	}
	return newHead.Next
}
