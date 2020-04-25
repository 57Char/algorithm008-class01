# 学习笔记

## 树

### 二叉树遍历

Name | Rule
---------------|-----------------
前序(Pre-order) |Root->Left->Right
中序(In-order)  |Left->Root->Right
后序(Post-order)|Left->Right->Root

### Go模板

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序
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

// 中序
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

// 后序
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
```

### Go模板之二

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append([]int{root.Val}, preorderTraversal(root.Left)...), preorderTraversal(root.Right)...)
}

// 中序
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}

// 后序
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(postorderTraversal(root.Left), postorderTraversal(root.Right)...), root.Val)
}

```

## BFS&DFS

- BFS: 广度优先
- DFS：深度优先

### 题目：429. N叉树的层序遍历

[429. N叉树的层序遍历](https://leetcode.com/problems/n-ary-tree-level-order-traversal/)

```go
type Node struct {
	Val      int
	Children []*Node
}

// bfs
func levelOrder(root *Node) [][]int {
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
func levelOrder(root *Node) [][]int {
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
```
