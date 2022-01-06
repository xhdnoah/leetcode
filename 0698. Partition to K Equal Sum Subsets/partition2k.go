package main

// Input: nums = [4,3,2,3,5,2,1], k = 4 Output: true
// Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.
// 回溯算法 backtrack
func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	used := make([]bool, len(nums))
	target := sum / k
	// k 号桶初始时什么都没装 从 nums[0] 开始做选择
	return backtrack(k, 0, nums, 0, used, target)
}

// 目前 k 号桶里已装入数字之和为 bucket；used 标志某一个元素是否已经被装到桶中；target 是每个桶需要达成的目标和
func backtrack(k, bucket int, nums []int, start int, used []bool, target int) bool {
	if k == 0 { // base case 所有桶都装满了
		return true
	}
	if bucket == target {
		// 装满当前桶，递归穷举下一个桶的选择，从 nums[0] 开始
		return backtrack(k-1, 0, nums, 0, used, target)
	}
	// 从 start 开始向后探查有效的 nums[i] 装入当前桶
	for i := start; i < len(nums); i++ {
		if used[i] {
			continue // 剪枝: nums[i] 已被装入别的桶
		}
		if nums[i]+bucket > target {
			continue // 当前桶装不下 nums[i]
		}
		used[i] = true    // 做选择
		bucket += nums[i] // 将 nums[i] 装入当前桶
		if backtrack(k, bucket, nums, i+1, used, target) {
			return true
		}
		used[i] = false // 撤销选择
		bucket -= nums[i]
	}
	return false
}
