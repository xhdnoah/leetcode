package main

import "sort"

// Given an array where time[i] denotes the time taken by the ith bus to complete one trip
// Also given an integer totalTrips which denotes the number of trips all buses should make in total
// Return the minimum time required for all buses to complete at least totalTrips trips.
// Input: time [1,2,3] totalTrips 5 Output: 3
// 时刻 t=1 每辆车完成旅途数 [1,0,0] 总数 1；时刻 t=2 每辆车完成 [2,1,0] 总数 3；时刻 t=3 每辆车完成 [3,1,1] 总数 5 因此至少完成总数 5 需要时间 3
// 二分查找：由于时间越多 可以完成的旅途也就越多 因此可以二分
// 答案不可能超过让最快的车跑 totalTrips 所花费的时间 可以将其作为二分的上界
func minimumTime(time []int, totalTrips int) int64 {
	sort.Ints(time)
	return int64(sort.Search(totalTrips*time[0], func(x int) bool {
		total := 0
		for _, t := range time {
			total += x / t
		}
		return total >= totalTrips
	}))
}
