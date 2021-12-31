package main

// Input: nums = [-4,-1,0,3,10] Output: [0,1,9,16,100]
// After squaring, the array becomes [16,1,0,9,100]. After sorting, it becomes [0,1,9,16,100]
// 数组本身有序 数组平方的最大值就在数组的两端 不会是中间
// 双指针：i 指向起始位置 j 指向终止位置 定义 res []int, k 指向 res 终止位置
// 如果A[i] * A[i] < A[j] * A[j] 那么result[k--] = A[j] * A[j];
// 如果A[i] * A[i] >= A[j] * A[j] 那么result[k--] = A[i] * A[i];
func sortedSquares(nums []int) []int {
	ln := len(nums)
	i, j, k := 0, ln-1, ln-1
	res := make([]int, ln)
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			res[k] = lm
			i++
		} else {
			res[k] = rm
			j--
		}
		k--
	}
	return res
}
