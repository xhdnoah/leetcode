package main

import . "leetcode/utils"

// Input: l1 = [2,4,3], l2 = [5,6,4] Output: [7,0,8] Explanation: 342 + 465 = 807.
// 2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出 需要注意各种进位问题 极端情况如：
// Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> ) Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1
// 为了处理方法统一，可以先建立一个虚拟头结点，这样 head 不需要单独处理，注意循环终止条件是 p != nil
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	vHead := &ListNode{}
	n1, n2, carry, current := 0, 0, 0, vHead // carry 表示进位
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next // 游标右移
		carry = (n1 + n2 + carry) / 10
	}
	return vHead.Next
}
