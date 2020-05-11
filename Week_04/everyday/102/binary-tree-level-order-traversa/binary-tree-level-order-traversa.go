package binary_tree_level_order_traversa

// https://leetcode.com/problems/binary-tree-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var level []*TreeNode
		var curr []int
		for _, node := range queue {
			curr = append(curr, node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		res = append(res, curr)
		queue = level
	}
	return res
}
