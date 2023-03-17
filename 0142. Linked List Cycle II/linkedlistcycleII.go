package main

import . "leetcode/utils"

// Input: head = [3,2,0,-4], pos = 1 Output: tail connects to node index 1
// 快慢指针相遇以后，如果 slow 继续往前走，fast 指针回到起点 head，两者都每次走一步，那么必定会在环的起点相遇
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	hasCycle, slow := hasCycle(head)
	if !hasCycle {
		return nil
	}
	fast := head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func hasCycle(head *ListNode) (bool, *ListNode) {
	slow := head
	fast := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}
