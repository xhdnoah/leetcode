package main

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	sum := 1 << uint(len(nums)) // 子集总数
	for i := 0; i < sum; i++ {
		stack := []int{}
		tmp := i                              // i 范围 000..000 ~ 111..111
		for j := len(nums) - 1; j >= 0; j-- { // 注意区间是 [0, len(nums) - 1] 代表 i 每一位
			if tmp&1 == 1 {
				stack = append(stack, nums[j]) // 如果当前位为 1 添加当前位对应的元素
			}
			tmp >>= 1 // 通过右移遍历 i 每一位
		}
		res = append(res, stack)
	}
	return res
}
