package main

import "math/rand"

func partition(arr []int, low, high int) int {
	// 循环不变量：[low+1, i] 元素小于 pivot, [i+1, high] 元素大于 pivot
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

func partition_random(arr []int, lo, hi int) int {
	pi := rand.Intn(hi-lo+1) + lo
	arr[lo], arr[pi] = arr[pi], arr[lo]
	return partition(arr, lo, hi)
}

// 快排 二叉树的前序遍历
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivot := partition(arr, low, high)
	quickSort(arr, low, pivot-1)
	quickSort(arr, pivot+1, high)
}

// 对于大量 <v 的重复元素 会导致左右两边不均衡 性能下降
// 双路快排将 <v 和 >v 两部分放在数组两端 双指针 i 指向 <v 部分下一元素、j 指向 >v 部分前一元素
// i 向后遍历 如果元素 e<v 则继续直到 e>=v 停止; j 向前遍历 如果元素 e>v 则继续直到 e <=v 停止
// 交换 i j 指向的元素 然后 i++ j-- 继续比较下一个
func partition_2Ways(arr []int, lo, hi int) int {
	i, j, pi := lo, hi, arr[lo]
	for i < j {
		for i < j && arr[i] <= pi {
			i++
		}
		for i < j && arr[j] > pi {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[lo] = arr[lo], arr[i]
	return i
}

// 三路快排 划分三块: 小于 k、等于 k、大于 k 处理大量重复元素的用例
// 相比双路快排减少了对重复元素的比较操作，因为重复元素在一次排序中就已经作为单独一部分排好了
func quickSort_3Ways(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	k := arr[lo]
	// 递归假定前提：数组其他部分已经排好序
	// 循环不变: lk 指向小于 k 部分的最后一个元素、gk 指向大于 k 部分的第一个元素
	i, lk, gk := lo+1, lo, hi+1
	for i < gk {
		switch {
		case arr[i] > k:
			arr[i], arr[gk-1] = arr[gk-1], arr[i]
			gk-- // i 不需要继续 因为此时 i 指向的元素是和 gk 前的空白元素交换过来的
		case arr[i] < k:
			arr[lk+1], arr[i] = arr[i], arr[lk+1] // 遍历到小于 k 的元素和 =v 部分的第一个交换
			lk++                                  // 此时 i 指向的元素是和已知排好序的元素交换得来因此需要继续前进
			i++
		default:
			i++ // 遍历元素 e=k 则直接合并到 =k 的部分 双指针不动 i 继续向前
		}
	}
	arr[lk], arr[lo] = arr[lo], arr[lk]
	quickSort_3Ways(arr, lo, lk-1)
	quickSort_3Ways(arr, gk, hi)
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
