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
