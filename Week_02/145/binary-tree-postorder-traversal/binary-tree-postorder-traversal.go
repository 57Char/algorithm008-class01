package binary_tree_postorder_traversal

// https://leetcode.com/problems/binary-tree-postorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, res)
	postorder(root.Right, res)
	*res = append(*res, root.Val)
}

func postorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(postorderTraversalV2(root.Left), postorderTraversalV2(root.Right)...), root.Val)
}