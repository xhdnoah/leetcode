package main

import . "leetcode/utils"

// Input: [[1,3],[2,6],[8,10],[15,18]] Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6]
// 先按照区间起点进行排序。然后从区间起点小的开始扫描，依次合并每个有重叠的区间。
func merge(intervals [][]int) [][]int {
	li := len(intervals)
	if li == 0 {
		return intervals
	}
	quickSort(intervals, 0, li-1)
	res := [][]int{intervals[0]}
	for cur, next := 0, 1; next < li; next++ {
		if res[cur][1] < intervals[next][0] {
			cur++
			res = append(res, intervals[next])
		} else {
			res[cur][1] = Max(intervals[next][1], res[cur][1])
		}
	}
	return res
}

// [1, 6, 3, 7, 5]
func partitionSort(a [][]int, lo, hi int) int {
	pStart, pEnd := a[hi][0], a[hi][1]
	i := lo - 1 // 游标 i
	for j := lo; j < hi; j++ {
		jStart, jEnd := a[j][0], a[j][1]
		// 发现小于 pivot 的区间就交换到 i 处
		if (jStart < pStart) || (jStart == pStart && jEnd < pEnd) {
			i++ // 先自增再交换保证 i 及其左侧的区间小于 pivot
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func quickSort(a [][]int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	merge([][]int{{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2}})
}
