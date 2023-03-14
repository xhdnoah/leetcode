package main

import (
	. "leetcode/utils"
	"strconv"
	"strings"
)

// Input: "11235813" Output: [1,1,2,3,5,8,13]; Input: "1101111" Output: [110, 1, 111]
// 0 <= F[i] <= 2^31 - 1 (整数符合 32 位有符号整数类型); F.length >= 3
func splitIntoFibonacci(S string) []int {
	if len(S) < 3 {
		return []int{}
	}
	res, isComplete := []int{}, false
	for firstEnd := 0; firstEnd < len(S)/2; firstEnd++ {
		if S[0] == '0' && firstEnd > 0 {
			break
		}
		first, _ := strconv.Atoi(S[:firstEnd+1])
		if first >= 1<<31 { // 剪枝条件：每个数字要求小于 2^31 - 1
			break
		}
		for secondEnd := firstEnd + 1; Max(firstEnd, secondEnd-firstEnd) <= len(S)-secondEnd; secondEnd++ {
			if S[firstEnd+1] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			second, _ := strconv.Atoi(S[firstEnd+1 : secondEnd+1])
			if second >= 1<<31 {
				break
			}
			findRecursiveCheck(S, first, second, secondEnd+1, &res, &isComplete)
		}
	}
	return res
}

func findRecursiveCheck(S string, x1, x2 int, left int, res *[]int, isComplete *bool) {
	if x1 >= 1<<31 || x2 >= 1<<31 {
		return
	}
	if left == len(S) {
		if !*isComplete {
			*isComplete = true
			*res = append(*res, x1)
			*res = append(*res, x2)
		}
		return
	}
	if strings.HasPrefix(S[left:], strconv.Itoa(x1+x2)) && !*isComplete {
		*res = append(*res, x1)
		findRecursiveCheck(S, x2, x1+x2, left+len(strconv.Itoa(x1+x2)), res, isComplete)
		return
	}
	if len(*res) > 0 && !*isComplete {
		*res = (*res)[:len(*res)-1]
	}
	return
}
