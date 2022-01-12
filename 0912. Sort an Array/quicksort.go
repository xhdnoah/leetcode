package main

func partition(arr []int, low, high int) int {
	i, pivot := low, arr[low]
	for j := low + 1; j <= high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[i] = arr[i], arr[low]
	return i // 最终基准点
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
	return arr
}

type Stack []int

func (s *Stack) Push(el int) {
	*s = append(*s, el)
}

func (s *Stack) Pop() int {
	i := len(*s) - 1
	el := (*s)[i]
	*s = (*s)[:i]
	return el
}

// 非递归形式 使用栈存放左右下标
func quickSort_stack(arr []int) []int {
	start, end := 0, len(arr)-1
	var stack Stack
	stack.Push(end)      // 右指针入栈
	stack.Push(start)    // 左指针入栈
	for len(stack) > 0 { // 栈不为空，排序尚未完成
		l := stack.Pop()          // 后进先出 待排序列最左下标
		r := stack.Pop()          // 待排序列最右下标
		p := partition(arr, l, r) // 最终有一个数获得最终位置 i
		if l < p-1 {              // i(已排好序)将待排序列一分为二
			stack.Push(p - 1) // 左边部分的右指针入栈
			stack.Push(l)     // 左边部分的左指针入栈
		}
		if r > p+1 {
			stack.Push(r)
			stack.Push(p + 1)
		}
	}
	return arr
}

func main() {
	quickSort_stack([]int{0, 1, 1, 2, 5, 0})
}
