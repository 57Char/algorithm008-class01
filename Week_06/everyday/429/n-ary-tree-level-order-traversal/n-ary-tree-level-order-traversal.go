package n_ary_tree_level_order_traversal

// https://leetcode.com/problems/n-ary-tree-level-order-traversal/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*Node{root}
	for len(queue) > 0 {
		vals := make([]int, len(queue))
		next := []*Node{}
		for i, node := range queue {
			vals[i] = node.Val
			for _, child := range node.Children {
				next = append(next, child)
			}
		}
		res = append(res, vals)
		queue = next
	}
	return res
}
