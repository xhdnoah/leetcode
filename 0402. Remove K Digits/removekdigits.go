package main

import "strings"

// Input: num = "1432219", k = 3 Output: "1219"
// Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
// 贪心算法在每一步选择中都采取在当前状态下最好或最优的选择，适用于最优子结构问题，即局部最优解能决定全局最优解：旅行推销员问题
// 若要使得剩下的数字最小，需要保证靠前的数字尽可能小。贪心策略：对数字序列从左往右找到第一个位置 i 使得 num[i] < num[i-1] 并删去 num[i-1]
// 删去一个字符后，剩下 n−1 长度的序列形成新的子问题，可以继续使用同样的策略，直至删除 k 次，暴力的实现复杂度最差会达到 O(nk)（考虑整个序列单调不降）
// 考虑从左往右增量的构造最后的答案。我们可以用一个单调栈维护当前的答案序列，栈中的元素代表截止到当前位置，删除不超过 k 次个数字后，所能得到的最小整数
// 对于每个数字，如果该数字小于栈顶元素，我们就不断地弹出栈顶元素，直到 1、栈为空 2、新的栈顶元素不大于当前数字 3、我们已经删除了 k 位数字
func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		ls, digit := len(stack), num[i]
		for k > 0 && ls > 0 && digit < stack[ls-1] {
			stack = stack[:ls-1] // 遍历数字小于栈顶元素，弹出栈顶
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k] // 删除最后几个数字
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
