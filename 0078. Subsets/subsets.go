package main

// 依次遍历在前一次的基础上 append nums[i]
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
