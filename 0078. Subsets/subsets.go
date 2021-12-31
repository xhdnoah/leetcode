package main

// 数学归纳法：依次遍历在前一次的基础上 append nums[i] 进行迭代
// [] => [1] => [2],[1 2] => [3],[1 3],[2 3],[1 2 3]
func subsets(nums []int) [][]int {
	res := make([][]int, 1) // 注意 length 此时 res 为 [[]]
	for _, n := range nums {
		for _, origin := range res {
			clone := make([]int, len(origin), len(origin)+1)
			copy(clone, origin)
			clone = append(clone, n)
			res = append(res, clone)
		}
	}
	return res
}

// DFS 回溯算法
func subsets_dfs(nums []int) [][]int {
	track := []int{} // 记录走过路径
	res := [][]int{}
	backtrack(nums, 0, track, &res)
	return res
}

func backtrack(nums []int, start int, track []int, res *[][]int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i]) // 做选择
		backtrack(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func subsets_bitwise(nums []int) [][]int {
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
