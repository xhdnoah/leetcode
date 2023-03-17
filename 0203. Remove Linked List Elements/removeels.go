package main

import . "leetcode/utils"

// 直接从原 head 开始遍历需区分头结点和其他结点；设置虚拟头结点以统一操作逻辑
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
