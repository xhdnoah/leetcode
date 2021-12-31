package main

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4] Output: [9,4]
// 相比数组 unordered_set 用在数值大小没有限制的场景
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0)
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		if _, ok := set[v]; !ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
