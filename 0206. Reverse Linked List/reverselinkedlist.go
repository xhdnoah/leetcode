package main

import . "leetcode/utils"

// Input: head = [1,2,3,4,5] Output: [5,4,3,2,1]
// 涉及到链表的操作，一定要在纸上把过程先画出来，再写程序
// 双指针法：head 在前 cur 在后，每次让 head.next 指向 cur 实现一次局部反转
// 反转完成后 head cur 在链表上同时前移一个单位 循环反转过程直到 head 到达尾部
func reverseList_iterative(head *ListNode) *ListNode {
	var cur *ListNode
	for head != nil {
		next := head.Next // 缓存
		head.Next = cur
		cur = head
		head = next
	}
	return cur
}

// behind: 1, head: 2->3->4->5
// behind: 2->1, head: 3->4->5
// behind: 3->2->1, haed: 4->5
// ...

// 一直递归到链表结尾 即为反转后的头结点 记为 newHead 作为最终结果返回
// 此后每次在函数返回过程中让当前结点的 next.next 指向当前结点
// 同时让当前结点的 next 指向 nil 实现从结尾开始的局部反转 当函数全部出栈后反转完成
func reverseList_recursive(head *ListNode) *ListNode {
	// 递归终止条件 空结点或
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList_recursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
