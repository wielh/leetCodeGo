package main

func bstToGst(root *TreeNode) *TreeNode {
	traverse(root, 0)
	return root
}

func traverse(root *TreeNode, currentSum int) int {
	if root == nil {
		return currentSum
	}

	sum := 0
	sum += traverse(root.Right, currentSum)
	sum += root.Val
	root.Val = sum
	sum = traverse(root.Left, sum)
	return sum
}
