package Week_02

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0, 3)
	result = append(result, root.Val)
	for _, node := range root.Children {
		result = append(result, preorder(node)...)
	}
	return result
}

type Node struct {
	Val      int
	Children []*Node
}
