package main

import . "leetcode/utils"

// Input: [3,1,2,4] Output: 17
// Explanation: Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
// Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.  Sum is 17.
// Since the answer may be large, return the answer modulo 10^9 + 7.
// 看到这种 mod1e9+7 的题目，首先要想到的就是 dp 最终的优化解即是利用 DP + 单调栈
// 单调栈维护数组中的值逐渐递增的对应下标序列。定义 dp[i+1] 代表以 A[i] 结尾的子区间内最小值总和
// 状态转移方程为 dp[i+1] = dp[prev+1] + (i-prev)*A[i]
// 以 [3, 1, 2, 4, 3] 为例，当 i = 4, 所有以 A[4] 为结尾的子区间有
// [3]
// [4, 3]
// [2, 4, 3]
// [1, 2, 4, 3]
// [3, 1, 2, 4, 3]
// 前两个子区间 [3] and [4, 3] 最小是 A[4]: 3
// 后 3 个子区间都包含 2 == A[2] 是比 3 小的前一个数
// dp[4+1] = dp[2+1] + (4-2)*3
// 因为子区间的增长就是从左到右所以是 i-prev 个
func sumSubarrayMins(A []int) int {
	stack, dp, res := []int{}, make([]int, len(A)+1), 0
	stack = append(stack, -1) // 单调增长栈的作用是得到比当前元素小的前一个数

	for i := 0; i < len(A); i++ {
		for stack[len(stack)-1] != -1 && A[i] <= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1] // 栈顶元素出栈直到满足单调递增
		}
		// 此时保证 A[prev] 是比 A[i] 小的前一个数
		// (i-prev)*A[i] 代表没出现 prev 前有 i-prev 个区间内 A[i] 最小
		prev := stack[len(stack)-1]
		dp[i+1] = (dp[prev+1] + (i-prev)*A[i]) % Mod
		stack = append(stack, i)
		res += dp[i+1]
		res %= Mod
	}
	return res
}

// stack 维护逐渐递增的下标序列
// [-1]
// [-1 0]
// [-1 1] 3 > 1 则排除下标 0
// [-1 1 2]
// [-1 1 2 3]

// 暴力解法
func subSubarrayMins_violent(A []int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		stack := []int{} // 只包含当前元素和遍历过程中最小元素
		stack = append(stack, A[i])
		for j := i; j < len(A); j++ {
			if stack[len(stack)-1] >= A[j] {
				stack = stack[:len(stack)-1]
				stack = append(stack, A[j])
			}
			res += stack[len(stack)-1] // 在遍历过程中累加
		}
	}
	return res % Mod
}

func main() {
	sumSubarrayMins([]int{3, 1, 2, 4, 3})
}
