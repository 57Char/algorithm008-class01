package n_ary_tree_preorder_traversal

import "container/list"

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	res := []int{}
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for _, child := range root.Children {
		helper(child, res)
	}
}

func preorderV2(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*Node{root}, []int{}
	for len(stack) > 0 {
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- {
			if child := root.Children[i]; child != nil {
				stack = append(stack, child)
			}
		}
	}
	return res
}

func preorderV3(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*Node{root}, []int{}
	for len(stack) > 0 {
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, root.Val)
		stack2 := []*Node{}
		for _, child := range root.Children {
			if child != nil {
				stack2 = append([]*Node{child}, stack2...)
			}
		}
		stack = append(stack, stack2...)
	}
	return res
}

func preorderV4(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := list.New(), []int{}
	stack.PushFront(root)
	for stack.Len() > 0 {
		root = stack.Remove(stack.Front()).(*Node)
		res = append(res, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- {
			if child := root.Children[i]; child != nil {
				stack.PushFront(child)
			}
		}
	}
	return res
}
