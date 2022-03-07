package main

// Given nums = [0,0,1,1,1,2,2,3,3,4], return length = 5
// 将重复元素依次移至数组后面的空间，last 遍历指针，finder 搜寻指针
// finder 从数组起点查找第一个不同于 last 的元素，然后 push 到 last 身后
// 整个过程是 in-place 的，不会 allocate extra space
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums) {
		for nums[last] == nums[finder] {
			if finder == len(nums)-1 {
				return last + 1
			}
			finder++
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

// [0,0,1,1,1,2,2,3,3,4] => [0,1,1,...] => .. => [0, 1, 2, 3, 4]
func removeDuplicates_invariant(nums []int) int {
	ln := len(nums)
	if ln < 2 {
		return ln
	}
	// 循环不变量 nums[0..j] 没有重复元素
	// j 是刚刚赋值完的元素的下标
	j := 0
	for i := 1; i < ln; i++ {
		if nums[i] != nums[j] { // 找到不重复的元素 [0,0,1]
			j++
			nums[j] = nums[i] // [0,1,..] j 指向第一个不重复的元素
		}
	}
	return j + 1
}
