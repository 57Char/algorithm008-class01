# 学习笔记

## 树

### 二叉树遍历
- 前序(Pre-order):  根-左-右 
- 中序(In-order):   左-根-右 
- 后序(Post-order): 左-右-根

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