package main

import . "leetcode/utils"

// Given the sorted array: [-10,-3,0,5,9]
// One possible answer is [0,-3,9,-10,null,5], which represents
//
//		 0
//		/ \
//	-3   9
//	/   /
//
// -10  5
func sortedArrayToBST(nums []int) *TreeNode {
	ln := len(nums)
	if ln == 0 {
		return nil
	}
	return &TreeNode{Val: nums[ln/2], Left: sortedArrayToBST(nums[:ln/2]), Right: sortedArrayToBST(nums[ln/2+1:])}
}
