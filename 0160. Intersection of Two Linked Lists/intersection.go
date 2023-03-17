package main

import . "leetcode/utils"

// 找到两个链表的交叉点，思路类似链表找环
// 给定两个链表如果长度一样，从头往后扫即可
// 不同则先拼成 A->B 和 B->A 两个长度相同的链表再扫描
// A: 0->9->1->2->4 B: 3->2->4
// A+B: 0->9->1->2->4->3->2->4
// B+A: 3->2->4->0->9->1->2->4
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	// if a & b have different len, then we'll stop the loop after second iteration
	for a != b {
		// for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB // 表 A 走到头接上 B: A+B
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA // 表 B 走到头接上 A: B+A
		} else {
			b = b.Next
		}
	}
	return a
}
