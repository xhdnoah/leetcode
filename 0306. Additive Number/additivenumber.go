package main

import (
	"strconv"
	"strings"
)

// Input: "112358" Output: true
// The additive sequence: 1, 1, 2, 3, 5, 8
// 1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
// Input: "199100199" Output: true
// The additive sequence: 1, 99, 100, 199
// 1 + 99 = 100, 99 + 100 = 199
// 对给出的字符串判断是否为斐波那契数列形式
// 每次判断需累加两个数字，故在 DFS 遍历过程中需维护两个数的边界
// firstEnd secondEnd 相加之和数的起始位置为 secondEnd+1
// 每次移动后需判断后面的字符串是否以两个数字之和为开头
func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	// 第一个数字不可能超过字符串的一半
	for firstEnd := 0; firstEnd < len(num)/2; firstEnd++ {
		// first second 可以为 0 但不能为以 0 开头的数字如 03
		if num[0] == '0' && firstEnd > 0 {
			break
		}
		first, _ := strconv.Atoi(num[:firstEnd+1])
		// 剪枝条件：第一个数字和第二个数字不能超过之和数字的长度
		for secondEnd := firstEnd + 1; max(firstEnd, secondEnd-firstEnd) <= len(num)-secondEnd; secondEnd++ {
			if num[firstEnd+1] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			second, _ := strconv.Atoi(num[firstEnd+1 : secondEnd+1])
			if recursiveCheck(num, first, second, secondEnd+1) {
				return true
			}
		}
	}
	return false
}

// Propagate for rest of the string
// x1 x2: 两个要相加的数字; left: secondEnd 后一位
func recursiveCheck(num string, x1 int, x2 int, left int) bool {
	if left == len(num) {
		return true
	}
	if strings.HasPrefix(num[left:], strconv.Itoa(x1+x2)) {
		return recursiveCheck(num, x2, x1+x2, left+len(strconv.Itoa(x1+x2)))
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// isAdditiveNumber("199100199")
	isAdditiveNumber("198019823962")
}

// 关键是找到前两个正确的数字之后 firstEnd secondEnd 不再更新
// recursiveCheck 递归调用过程
// 199100199 1 9 2 不符合 判断 return false; secondEnd++ 进入下一次遍历
// 199100199 1 99 3 符合 x1 x2 右移并继续递归判断
// 199100199 99 100 6
// 199100199 100 199 9
