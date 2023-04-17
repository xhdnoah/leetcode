package main

// 整数数组 nums 按升序排列在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行旋转
// 使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 在旋转数组中搜索一个给定的目标值
// Input: nums = [4,5,6,7,0,1,2], target = 0 Output: 4 算法时间复杂度必须是 O(log n)
// 在常规二分查找时查看当前 mid 为分割位置分割出来的两个部分 [l, mid] 和 [mid + 1, r] 哪个部分是有序的
// 并根据有序的那个部分确定我们该如何改变二分查找的上下界，根据有序的那部分判断出 target 在不在这个部分
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // [0, mid] 有序
			if nums[0] <= target && target < nums[mid] { // target 在 [lo, mid-1] 区间内
				hi = mid - 1
			} else { // target 在 [mid+1, hi] 区间内
				lo = mid + 1
			}
		} else { // [mid+1, hi] 有序
			if nums[mid] < target && target <= nums[hi] { // target 在 [mid+1, hi] 区间内
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func main() {
	search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
}
