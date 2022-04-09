package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Sort a linked list in O(n log n) time using constant space complexity.
// Input: 4->2->1->3 Output: 1->2->3->4 Input: -1->5->3->4->0 Output: -1->0->3->4->5
// 归并排序符合条件 (O(nlogn) time; O(1) space): No.876 取中点 + No.21 合并两个有序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middleNode := getMiddleNode(head)
	temp := middleNode.Next // 保存后半段起点
	middleNode.Next = nil   // 从中点断开链表

	left := sortList(head)
	right := sortList(temp)
	return mergeTwoLists(left, right)
}

// 快慢指针法（偶数链表时 p1 会停在中点前一个结点
func getMiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

// 合并有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
