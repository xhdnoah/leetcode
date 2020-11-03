package main

// Input: [1,8,6,2,5,4,8,3,7] Output: 49
// 对撞指针，首尾分别两个指针，每次移动后比较 tmp max 得出当前最大值
// 指针从当前最小可能的位置向中间移动，即若 start 低于 end 则 start++ 反之 end--
func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		var h int
		w := end - start
		if height[start] < height[end] {
			h = height[start]
			start++
		} else {
			h = height[end]
			end--
		}

		tmp := w * h
		if tmp > max {
			max = tmp
		}
	}
	return max
}
