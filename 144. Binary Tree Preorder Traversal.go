package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	answer := []int{root.Val}
	answer = append(answer, preorderTraversal(root.Left)...)
	answer = append(answer, preorderTraversal(root.Right)...)
	return answer
}
