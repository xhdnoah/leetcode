package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: [ 1->4->5, 1->3->4, 2->6 ] Output: 1->1->2->3->4->4->5->6
// 分而治之，最终将问题归为合并两个链表 Time Complexity O(log n)
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	mid := length / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoList(left, right)
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoList(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoList(l1, l2.Next)
	return l2
}
