package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. Create a root node whose value is the maximum value in nums.
// 2. Recursively build the left subtree on the subarray prefix to the left of the maximum value.
// 3. Recursively build the right subtree on the subarray suffix to the right of the maximum value.
// Input: nums = [3,2,1,6,0,5] Output: [6,3,5,null,2,0,null,null,1]
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}

	index, maxVal := -1, 0
	for i := lo; i <= hi; i++ {
		if maxVal <= nums[i] {
			index = i
			maxVal = nums[i]
		}
	}

	root := &TreeNode{Val: maxVal}
	root.Left = build(nums, lo, index-1)
	root.Right = build(nums, index+1, hi)

	return root
}

func main() {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}
