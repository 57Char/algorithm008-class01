package n_ary_tree_postorder_traversal

import "container/list"

// https://leetcode.com/problems/n-ary-tree-postorder-traversal/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	res := []int{}
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		helper(child, res)
	}
	*res = append(*res, root.Val)
}

// very slow
func postorderV2(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*Node{root}, []int{}
	for len(stack) > 0 {
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append([]int{root.Val}, res...) // slow
		for _, child := range root.Children {
			if child != nil {
				stack = append(stack, child)
			}
		}
	}
	return res
}

// very fast
func postorderV3(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*Node{root}, []int{}
	for len(stack) > 0 {
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, root.Val)
		for _, child := range root.Children {
			if child != nil {
				stack = append(stack, child)
			}
		}
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func postorderV4(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, res := list.New(), []int{}
	stack.PushFront(root)
	for stack.Len() > 0 {
		root = stack.Remove(stack.Front()).(*Node)
		res = append(res, root.Val)
		for _, child := range root.Children {
			if child != nil {
				stack.PushFront(child)
			}
		}
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}