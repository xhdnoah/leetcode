package main

// Input: temperatures = [73,74,75,71,69,72,76,73] Output: [1,1,4,2,1,1,0,0]
// ans[i] 是指在第 i 天之后，才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替
// 维护一个存储下标的单调栈，从栈底到栈顶的下标对应的温度列表中的温度依次递减。如果一个下标在单调栈里，则表示尚未找到下一次温度更高的下标。
func dailyTemperatures_stack(temperatures []int) []int {
	lt := len(temperatures)
	ans := make([]int, lt)
	stack := []int{}
	for i := 0; i < lt; i++ {
		temp := temperatures[i]
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
