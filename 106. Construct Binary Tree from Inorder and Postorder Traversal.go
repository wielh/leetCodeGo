package main

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	answer := &TreeNode{}
	l := len(postorder)
	answer.Val = postorder[l-1]
	spiltIndex := 0
	for inorder[spiltIndex] != postorder[l-1] {
		spiltIndex++
	}
	answer.Left = buildTree(inorder[:spiltIndex], postorder[:spiltIndex])
	answer.Right = buildTree(inorder[spiltIndex+1:], postorder[spiltIndex:l-1])
	return answer
}
