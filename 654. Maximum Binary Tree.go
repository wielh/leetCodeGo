package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := -1
	max := -1
	for index, num := range nums {
		if num > max {
			maxIndex = index
			max = num
		}
	}

	leftTree := constructMaximumBinaryTree(nums[:maxIndex])
	rightTree := constructMaximumBinaryTree(nums[maxIndex+1:])
	return &TreeNode{Val: max, Left: leftTree, Right: rightTree}
}
