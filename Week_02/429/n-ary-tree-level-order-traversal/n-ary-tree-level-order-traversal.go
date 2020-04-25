package n_ary_tree_level_order_traversal

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	level := []*Node{root}
	for len(level) > 0 {
		vals := []int{}
		children := []*Node{}
		for _, root := range level {
			vals = append(vals, root.Val)
			for _, child := range root.Children {
				children = append(children, child)
			}
		}
		res = append(res, vals)
		level = children
	}
	return res
}

// bfs
func levelOrderV2(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	bfs := func(root *Node) {
		queue := []*Node{root}
		for len(queue) > 0 {
			vals := []int{}
			next := []*Node{}
			for _, root := range queue {
				vals = append(vals, root.Val)
				for _, node := range root.Children {
					next = append(next, node)
				}
			}
			res = append(res, vals)
			queue = next
		}
	}
	bfs(root)
	return res
}

// dfs
func levelOrderV3(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	dfs(root, 0, &res)
	return res
}

func dfs(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	for _, node := range root.Children {
		dfs(node, level+1, res)
	}
}
