package main

// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
// Output: [1,2,2,3,5,6]
// 为了不大量移动元素，就要从 2个数组长度之和的最后一个位置开始，依次选取两个数组中大的数
// 从第一个数组的尾巴开始往头放，只要循环一次以后，就生成了合并以后的数组
func merge(num1 []int, m int, num2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- { // 注意终止条件
		if num1[m-1] <= num2[n-1] {
			num1[p-1] = num2[n-1]
			n--
		} else {
			num1[p-1] = num1[m-1]
			m--
		}
	}
	// 此时剩余 n 个 num2 元素都是比 num1 小的
	for ; n > 0; n-- {
		num1[n-1] = num2[n-1]
	}
}
