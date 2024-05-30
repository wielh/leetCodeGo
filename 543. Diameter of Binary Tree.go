package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(root *TreeNode, diameter int) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftDepth, leftDiameter := depth(root.Left, diameter)
	rightDepth, rightDiameter := depth(root.Right, diameter)
	localDiameter := leftDepth + rightDepth
	diameter = max(diameter, leftDiameter)
	diameter = max(diameter, localDiameter)
	diameter = max(diameter, rightDiameter)
	return max(leftDepth, rightDepth) + 1, diameter
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := depth(root, 0)
	return diameter
}
