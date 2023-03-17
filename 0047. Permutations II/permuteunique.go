package main

import "sort"

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums) // 排序保证重复数紧密分布
	var (
		ln        = len(nums)
		visited   = make([]bool, ln)
		track     = []int{}
		backtrack func(int)
	)
	backtrack = func(idx int) {
		if idx == ln {
			ans = append(ans, append([]int(nil), track...))
			return
		}
		for i := range nums {
			// 去重关键：除对已访问集合的判断外
			// 还保证了对于重复数的集合 一定是从左往右逐个填入的 不存在左边的重复数还没取到的而再次填入的可能
			if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) {
				continue
			}
			track = append(track, nums[i])
			visited[i] = true
			backtrack(idx + 1)
			visited[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return
}
