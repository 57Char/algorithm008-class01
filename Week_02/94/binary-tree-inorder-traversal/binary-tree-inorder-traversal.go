package binary_tree_inorder_traversal

// https://leetcode.com/problems/binary-tree-inorder-traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func inorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}