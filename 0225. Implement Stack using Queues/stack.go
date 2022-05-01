package main

// 为实现栈的特性-后进先出，应满足队列前端元素是最后入栈的元素
// 双队列：queue1 存储栈内元素，queue2 作为入栈操作的辅助队列
// 入栈操作时，先将元素入队到 queue2，再将 queue1 全部元素依次出队并入队到 queue2
// 此时 queue2 前端元素即为新入栈元素，再将 queue1、2 互换，即 queue1 元素为栈内元素
// queue1 的前端和后端分别对应栈顶和栈底
type MyStack struct {
	queue1, queue2 []int
}

func Constructor() (s MyStack) {
	return
}

func (s *MyStack) Push(x int) {
	s.queue2 = append(s.queue2, x)
	for len(s.queue1) > 0 {
		s.queue2 = append(s.queue2, s.queue1[0])
		s.queue1 = s.queue1[1:]
	}
	s.queue1, s.queue2 = s.queue2, s.queue1
}

func (s *MyStack) Pop() int {
	v := s.queue1[0]
	s.queue1 = s.queue1[1:]
	return v
}

func (s *MyStack) Top() int {
	return s.queue1[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0
}

// 单队列：同样需要满足队列前端的元素是最后入栈的元素
// 入栈操作时，首先获得入栈前的元素个数 n 后将新元素入队
// 再将队列中的前 n 个元素依次出队并入队到队列，此时队列前端为新元素，且前端和后端分别对应栈顶和栈底
type Stack struct {
	queue []int
}

func Builder() (s Stack) {
	return
}

func (s *Stack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for ; n > 0; n-- {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

func (s *Stack) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

func (s *Stack) Top() int {
	return s.queue[0]
}

func (s *Stack) Empty() bool {
	return len(s.queue) == 0
}
