package main

// Input: head = [3,2,0,-4], pos = 1 Output: true
// pos is used to denote the index of the node that tail's next pointer is connected to
// 快慢指针：快指针一次走 2 格，慢指针一次走 1 格。如果存在环，那么前一个指针一定会经过若干圈之后追上慢的指针
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { // 快指针永远走在前所以防御性判断只考虑 fast
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
