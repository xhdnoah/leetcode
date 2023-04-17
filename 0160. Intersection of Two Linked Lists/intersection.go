package main

import . "leetcode/utils"

// 找到两个链表的交叉点，思路类似链表找环
// 因为相交点之后的长度是相同的 拼接成 A->B B->A 两长度相同链表
// 双指针同时从头往后扫描到相同的结点则为相交点 其之后的长度也相同
// A: 0->9->1->2->4 B: 3->2->4
// A+B: 0->9->1->2->4->3->2->4
// B+A: 3->2->4->0->9->1->2->4
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { // boundary check
		return nil
	}

	a, b := headA, headB
	// if a & b have different len, then we'll stop the loop after second iteration
	for a != b { // 类似于 K8s 的调协循环: for 条件判断是不符合预期的等式 不符合则需进一步调整到预期结果
		// for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB // 表 A 走到头接上 B: A->B
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA // 表 B 走到头接上 A: B->A
		} else {
			b = b.Next
		}
	}
	return a
}
