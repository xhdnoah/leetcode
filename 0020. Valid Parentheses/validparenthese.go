package main

// Input: "([)]" Output: false Input: "{[]}" Output: true
// 遇到左括号就进栈 push，遇到右括号并且栈顶为与之对应的左括号，就把栈顶元素出栈
// 最后看栈里面还有没有其他元素，如果为空，即匹配
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v) // 左括号入栈
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1] // 遇到右括号且有匹配的左括号将对应左括号出栈
		} else {
			return false
		}
	}
	return len(stack) == 0
}
