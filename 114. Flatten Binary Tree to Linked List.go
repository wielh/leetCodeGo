package main

func flatten(root *TreeNode) {
	flattenWithTail(root)
}

func flattenWithTail(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		return root
	} else if root.Left == nil {
		rightTail := flattenWithTail(root.Right)
		return rightTail
	} else if root.Right == nil {
		leftTail := flattenWithTail(root.Left)
		root.Right = root.Left
		root.Left = nil
		return leftTail
	} else {
		rightTail := flattenWithTail(root.Right)
		leftTail := flattenWithTail(root.Left)
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
		return rightTail
	}
}
