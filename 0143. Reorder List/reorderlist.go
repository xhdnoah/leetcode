package main

import . "leetcode/utils"

// Given a singly linked list L: L0→L1→…→Ln-1→Ln
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
// Given 1->2->3->4, reorder it to 1->4->2->3.
// Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
// 先找到链表的中间结点，再利用逆序区间的操作，如 No.92 reverseBetween 只不过这里的反转区间是从中点一直到末尾
// 最后用双指针，一个指向头结点，一个指向中间结点，开始拼接最终结果。这种做法的时间复杂度是 O(n)，空间复杂度是 O(1)
func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针寻找中间结点
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// 反转链表后半 1->2->3->4->5->6 to 1->2->3->6->5->4
	preMiddle := p1              // 3
	preCurrent := p1.Next        // 4
	for preCurrent.Next != nil { // preCurrent 充当游标往右移动不断将后面的元素插入 preMiddle 后面
		current := preCurrent.Next     // 5, 6
		preCurrent.Next = current.Next // 4->6, 4->nil
		current.Next = preMiddle.Next  // 5->4, 6->5
		preMiddle.Next = current       // 3->5(->4->6), 3->6(->5->4)
	}

	// 重新拼接链表 1->2->3->6->5->4 to 1->6->2->5->3->4
	p1 = head           // 1
	p2 = preMiddle.Next // 6
	for p1 != preMiddle {
		preMiddle.Next = p2.Next // 3->5
		p2.Next = p1.Next        // 6->2
		p1.Next = p2             // 1->6(->2->3->5->4)
		p1 = p2.Next             // 2
		p2 = preMiddle.Next      // 5
	}
	return head
}

func reorderList_composition(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 1->2, 3->4 => 1->3->2->4
func mergeList(l1, l2 *ListNode) {
	var next1, next2 *ListNode
	for l1 != nil && l2 != nil {
		next1 = l1.Next
		next2 = l2.Next

		l1.Next = l2
		l1 = next1

		l2.Next = l1
		l2 = next2
	}
}
