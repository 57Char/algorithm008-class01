package binary_tree_preorder_traversal

// https://leetcode.com/problems/binary-tree-preorder-traversal/

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	preorder(root, &res)
	return res
}

func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
}

func preorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append([]int{root.Val}, preorderTraversalV2(root.Left)...), preorderTraversalV2(root.Right)...)
}

func preorderTraversalV3(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil {
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node)
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
		} else {
			node = nil
			for i := len(stack) - 1; i >= 0; i-- {
				if node = stack[i].Right; node != nil {
					stack = stack[:i]
					break
				}
			}
		}
	}
	return res
}
