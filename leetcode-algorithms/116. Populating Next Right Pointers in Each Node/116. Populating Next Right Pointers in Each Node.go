func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		tmp := queue

		queue = make([]*Node, 0)

		for idx, node := range tmp {
			if idx+1 < len(tmp) {
				node.Next = tmp[idx+1]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}

	return root
}