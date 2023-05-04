package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := map[int][]int{}        // 上级对下属的映射
	for i, m := range manager { // 建树
		g[m] = append(g[m], i)
	}
	var dfs func(int) int
	dfs = func(cur int) (res int) {
		for _, n := range g[cur] {
			res1 := dfs(n)
			if res1 > res {
				res = res1
			}
		}
		return informTime[cur] + res
	}
	return dfs(headID)
}
