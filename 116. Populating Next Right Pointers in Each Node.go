package main

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	currentNodes := []*Node{root}
	nextNodes := []*Node{}
	for {
		for i := 1; i < len(currentNodes); i++ {
			currentNodes[i-1].Next = currentNodes[i]
		}

		for _, node := range currentNodes {
			if node.Left == nil {
				return root
			}
			nextNodes = append(nextNodes, node.Left)
			nextNodes = append(nextNodes, node.Right)
		}

		currentNodes = nextNodes
		nextNodes = []*Node{}
	}
}
