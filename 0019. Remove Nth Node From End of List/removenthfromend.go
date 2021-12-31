package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: head = [1,2,3,4,5], n = 2 Output: [1,2,3,5]
// 快慢双指针：fast 移动 n 步再让双指针同时移动直到 fast 指向表尾此时删除 slow 所指结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead
	for n != 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
