package main

// Input: n = 4, k = 2
// Output: [[2,4], [3,4], [2,3], [1,2], [1,3], [1,4]]
func combine(n int, k int) [][]int {
	res := [][]int{}
	if n <= 0 || k <= 0 {
		return res
	}
	track := []int{}
	backtrack(n, k, 1, track, &res)
	return res
}

func backtrack(n, k, i int, track []int, res *[][]int) {
	l := len(track)
	if l == k {
		tmp := make([]int, l)
		copy(tmp, track)
		*res = append(*res, tmp)
	}
	for j := i; j <= n; j++ {
		track = append(track, j)
		backtrack(n, k, j+1, track, res)
		track = track[:l]
	}
}
