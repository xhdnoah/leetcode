package main

import "math"

// Input: word1 = "intention", word2 = "execution" Output: 5
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')
func minDistance(word1 string, word2 string) int {
	memo := make(map[[2]int]int, 0)
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		tuple := [2]int{i, j}
		if memo[tuple] > 0 {
			return memo[tuple]
		}
		if word1[i] == word2[j] {
			memo[tuple] = dp(i-1, j-1)
		} else {
			memo[tuple] = min(
				dp(i, j-1)+1,
				dp(i-1, j)+1,
				dp(i-1, j-1)+1,
			)
		}
		return memo[tuple]
	}
	return dp(len(word1)-1, len(word2)-1)
}

func min(nums ...int) int {
	m := math.MaxInt32
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}
