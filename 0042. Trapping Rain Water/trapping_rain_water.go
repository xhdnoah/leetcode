package main

// 双指针: 左右指针分别从头尾往中间扫，每次得出较小值，在较小侧进行容量计算(容量取决于短板）移动较小侧指针
// 遍历过程中，记录左右当前最大高度，如果数组中元素比局部最大高度小，则累加 res 值，否则更新最大高度
// 注意 left right 的起点应为 0, len - 1 因为需要比较得出最外围的短板
func trap(height []int) int {
	left, right, leftMax, rightMax, res := 0, len(height)-1, 0, 0, 0
	for left <= right {
		if height[left] < height[right] {
			if height[left] < leftMax {
				res += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				res += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right--
		}
	}
	return res
}
