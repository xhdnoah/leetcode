package main

// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1]
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		another := target - num
		if ai, ok := m[another]; ok {
			return []int{ai, i}
		}
		m[num] = i
	}
	return nil
}
