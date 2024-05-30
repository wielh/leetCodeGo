package main

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	answer := postorderTraversal(root.Left)
	answer = append(answer, postorderTraversal(root.Right)...)
	answer = append(answer, root.Val)
	return answer
}
