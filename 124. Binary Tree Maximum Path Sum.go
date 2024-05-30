package main

import "math"

func halfPathSum(root *TreeNode) (int, int) {
	if root == nil {
		return math.MinInt, math.MinInt
	} else if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}

	leftHalf, leftSum := halfPathSum(root.Left)
	rightHalf, rightSum := halfPathSum(root.Right)
	Sum := max(leftHalf, 0) + root.Val + max(rightHalf, 0)
	Sum = max(Sum, leftSum)
	Sum = max(Sum, rightSum)
	return max(max(leftHalf, rightHalf), 0) + root.Val, Sum
}

func maxPathSum(root *TreeNode) int {
	_, answer := halfPathSum(root)
	return answer
}
