package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	answer := &TreeNode{}
	answer.Val = preorder[0]
	spiltIndex := 0
	for inorder[spiltIndex] != preorder[0] {
		spiltIndex++
	}
	answer.Left = buildTree(preorder[1:spiltIndex+1], inorder[:spiltIndex])
	answer.Right = buildTree(preorder[spiltIndex+1:], inorder[spiltIndex+1:])
	return answer
}
