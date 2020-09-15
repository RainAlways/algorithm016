package Week_02

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	currentNodes := []*Node{root}
	nextNodes := make([]*Node, 0, 10)
	result := make([][]int, 0, 3)

	for len(currentNodes) != 0 {
		levelValues := make([]int, 0, len(currentNodes))
		nextNodes = nextNodes[0:0]
		for _, node := range currentNodes {
			if node == nil {
				continue
			}
			levelValues = append(levelValues, node.Val)
			nextNodes = append(nextNodes, node.Children...)
		}

		currentNodes = currentNodes[0:0]
		currentNodes = append(currentNodes, nextNodes...)
		result = append(result, levelValues)
	}

	return result
}
