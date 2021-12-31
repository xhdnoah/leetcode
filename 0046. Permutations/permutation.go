package main

// Input: nums = [1,2,3] Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res, track := [][]int{}, make([]int, 0)
	backtrack(nums, &track, &res)
	return res
}

func backtrack(nums []int, track *[]int, res *[][]int) {
	if len(*track) == len(nums) {
		tmp := make([]int, len(*track))
		copy(tmp, *track)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}
		*track = append(*track, nums[i])
		backtrack(nums, track, res)
		*track = (*track)[:len(*track)-1]
	}
}

func contains(els *[]int, v int) bool {
	for _, s := range *els {
		if v == s {
			return true
		}
	}
	return false
}
