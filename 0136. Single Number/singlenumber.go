package main

// Input: nums = [2,2,1] Output: 1
// Use only constant extra space: XOR 异或运算
// 任何数与 0 做异或运算，结果仍为原数 a⊕0=a
// 任何数与自身做异或运算，结果为 0: a⊕a=0
// 异或运算满足交换律和结合律
// 因此数组所有元素做异或运算的结构即为只出现一次的数字
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
