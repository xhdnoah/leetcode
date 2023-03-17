package main

// 队列是一种 先进先出 FIFO 数据结构 实现队列最直观的方法是用链表
// 栈是一种 后进先出 LIFO 数据结构 为满足队列特性，需用到两个栈:
// 一个用来反转元素的入队顺序，一个存储元素的最终顺序
// 将一个栈当作输入栈，用于压入 push 传入的数据；另一个栈当作输出栈，用于 pop 和 peek 操作
// 每次 pop 或 peek 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈，这样输出栈从栈顶往栈底的顺序就是队列从队首往队尾的顺序
// in: 3, 2, 1] => out: 1, 2, 3]
type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	q.in2out()
	popped := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return popped
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	q.in2out()
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack)+len(q.outStack) == 0
}

func (q *MyQueue) in2out() {
	if len(q.outStack) > 0 {
		return
	}
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1]) // 输出栈接受元素
		q.inStack = q.inStack[:len(q.inStack)-1]                     // 输入栈弹出元素
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
