package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given 1->2->3->4, you should return the list as 2->1->4->3
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		pre = head
		head = next
	}
	return dummy.Next
}

func swapPairs_recursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs_recursive(next.Next)
	next.Next = head
	return next
}

func swapPairs_inline(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		// The assignment proceeds in two phases:
		// First, the operands of index expressions and pointer indirections (including implicit pointer indirections in selectors) on the left
		// and the expressions on the right are all evaluated in the usual order. Second, the assignments are carried out in left-to-right order.
		// 0 -> 1 -> 2 -> 3 -> 4 => 0 -> 2 -> 1 -> 3 -> 4 => 0 -> 2 -> 1 -> 4 -> 3
		// 每三个一组置换结点 1 和 2 并将下一个 pt 指向 1:
		// 左边的指针指向先算好后再赋值，如 pt.Next.Next 已经算好为第一个结点 Next 指针地址 整个赋值过程是静态的
		// pt 被指向结点 1, pt.Next 原指向头结点被指向结点 2, pt.Next.Next 原结点 1 的 Next 指针被指向结点 3
		// pt.Next.Next.Next 原为结点 2 的 Next 指针现指向结点 1 最终得到 2 -> 1 -> 3
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}

func main() {
	third := &ListNode{Val: 4, Next: nil}
	second := &ListNode{Val: 3, Next: third}
	first := &ListNode{Val: 2, Next: second}
	head := &ListNode{Val: 1, Next: first}
	swapPairs(head)
}
