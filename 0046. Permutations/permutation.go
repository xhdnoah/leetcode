package main

// Input: nums = [1,2,3] Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 回溯算法：深度优先遍历、递归和栈 本质都是后进先出
func permute(nums []int) [][]int {
	var (
		ln        = len(nums)
		res       = [][]int{}
		backtrack func(*[]int)
	)
	if ln == 0 {
		return res
	}

	backtrack = func(track *[]int) {
		if len(*track) == ln { // 递归结束条件
			res = append(res, copyFrom(*track))
			return
		}

		for i := 0; i < ln; i++ {
			if contains(track, nums[i]) {
				continue
			}
			*track = append(*track, nums[i])
			backtrack(track)
			*track = (*track)[:len(*track)-1] // 撤销
		}
	}

	backtrack(&[]int{})
	return res
}

// Input: [1,2,3] Output: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
func permute_dynamic(nums []int) [][]int {
	var (
		ln        = len(nums)
		res       = [][]int{}
		backtrack func([]int, int)
	)
	// 将 nums 划分左右两部分 左边表示已填过的数 右边表示待填数字集合
	// 当已经填到第 first 位置 那 nums[:first-1] 为已填集合 [first, ln-1] 为待填集合
	// 尝试用右边集合中的数去填第 first 位置 假设待填数下标 i 那填完后将 i 与 first 交换 动态维护左右集合的一致性
	backtrack = func(output []int, first int) {
		if first == ln {
			res = append(res, copyFrom(output))
			return
		}
		for i := first; i < ln; i++ {
			swap(&output, i, first)
			backtrack(output, first+1)
			swap(&output, i, first)
		}
	}
	backtrack(nums, 0)
	return res
}

// 每一位都有三种选择 1 2 3 每一次做选择 展开一棵决策树 利用约束条件「不能重复选」做剪枝
func permute_visited(nums []int) [][]int {
	var (
		ln        = len(nums)
		res       = [][]int{}
		visited   = map[int]bool{}
		backtrack func([]int)
	)
	// backtrack 函数可以想象成一个在树上游走的指针 不断做选择
	backtrack = func(track []int) {
		if len(track) == ln {
			res = append(res, copyFrom(track))
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue // 剪枝
			}
			// 做选择
			track = append(track, n)
			visited[n] = true
			backtrack(track)
			// 上一轮递归结束 回溯到树的父结点 撤销选择
			track = track[:len(track)-1]
			visited[n] = false
		}
	}

	backtrack([]int{})
	return res
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func copyFrom(src []int) []int {
	tmp := make([]int, len(src))
	copy(tmp, src)
	return tmp
}

func contains(els *[]int, v int) bool {
	for _, s := range *els {
		if v == s {
			return true
		}
	}
	return false
}
