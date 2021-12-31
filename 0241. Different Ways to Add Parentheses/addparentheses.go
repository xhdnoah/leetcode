package main

import (
	"strconv"
)

// Input: expression = "2-1-1" Output: [0,2]
// Explanation: ((2-1)-1) = 0 (2-(1-1)) = 2
func diffWaysToCompute(expression string) []int {
	res := []int{}
	for i, r := range expression {
		if r == '+' || r == '-' || r == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, i := range left {
				for _, j := range right {
					if r == '+' {
						res = append(res, i+j)
					} else if r == '-' {
						res = append(res, i-j)
					} else if r == '*' {
						res = append(res, i*j)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		e, _ := strconv.Atoi(expression)
		res = append(res, e)
	}
	return res
}
