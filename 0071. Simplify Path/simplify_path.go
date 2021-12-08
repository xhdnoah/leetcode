package main

import "strings"

// Input: "/home/" Output: "/home"; Input: "/../" Output: "/"
// Input: "/home//foo/" Output: "/home/foo"; Input: "/a/./b/../../c/" Output: "/c"
func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		cur := strings.TrimSpace(arr[i])
		if cur == ".." { // 遇到 .. 出栈
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && len(cur) > 0 { // 非 "." 入栈
			stack = append(stack, arr[i])
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	return "/" + strings.Join(stack, "/")
}
