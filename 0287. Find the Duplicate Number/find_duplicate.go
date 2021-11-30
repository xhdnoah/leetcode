package main

// Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive)
// Input: [1,3,4,2,2] Output: 2
// 将这些数字想象成链表中的结点，数组数字代表下一结点元素在数组中的下标，找重复的数字就是找链表中成环的那个点
// 由于题目保证了有重复数字，一定会成环，可用快慢指针的方法，快指针一次走 2 步，慢指针一次走 1 步
// 相交以后，快指针从头开始，每次走一步，再次遇见的时候就是成环的交点处，也即是重复数字所在的地方
// 第一次相交是进入环中任意位置，之后 walker 指针从起点出发与在环中的 slow 相遇即为成环交点
// Linked List: 1 -> nums[1]: 3 -> nums[3]: 2 -> nums[2]: 4 -> nums[4]: 2 -> nums[2]: 4 -> ... 成环
// slow: 1 -> 3 -> 2 -> 4 -> 2 -> 4 -> 2 fast: 3 -> 4 -> 4 -> 4 walker: 1 -> 3 -> 2
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[slow]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	walker := 0
	for walker != slow {
		walker = nums[walker]
		slow = nums[slow]
	}
	return walker
}

// 二分搜索 [1, n] 以内的数 + 1 个重复数构成 [n+1]int
// 如果过半数在平均值及以下，说明重复数 [low, mid] 否则在 [mid + 1, high]
func findDuplicate1(nums []int) int {
	low, high := 1, len(nums)
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, n := range nums {
			if n <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
