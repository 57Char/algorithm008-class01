package maximum_depth_of_binary_tree

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	level := []*TreeNode{root}
	for len(level) > 0 {
		depth++
		q := []*TreeNode{}
		for _, node := range level {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level = q
	}
	return depth
}

func maxDepthV3(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	level++
	return max(helper(root.Left, level), helper(root.Right, level))
}
