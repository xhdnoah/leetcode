package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: head = [1,2,3,4,5] Output: [5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	var behind *ListNode // 新构造链表的头结点
	for head != nil {
		next := head.Next  // 缓存当前头节点下一个结点
		head.Next = behind // head 脱离原链表 head.Next 指向新链表
		behind = head      // 头结点作为新链表表头
		head = next        // 下一个结点作为原表头
	}
	return behind
}

// behind: 1, head: 2->3->4->5
// behind: 2->1, head: 3->4->5
// behind: 3->2->1, haed: 4->5
// ...
