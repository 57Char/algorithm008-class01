## DFS 代码模板
[模板](https://shimo.im/docs/ddgwCccJQKxkrcTq/read)

### 递归
```go
visited = map[*TreeNode]bool{}

func dfs(root, visited) {
    // terminator
    if visited[root] {
    	// already visited 
    	return 
    }

	visited[root] = true 

	// process current node here. 
	...
	for _, child := range root.Children() {
		if !visited[child] {
            dfs(next_node, visited)
        }
    }
}
```

### 非递归
```go
func dfs(root) {
	if root != nil {
		return
	}

	visited, stack := map[*TreeNode]bool{}, []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[node] = true
		process(node)
		nodes := generateRelatedNodes(node)
		stack = append(stack, nodes...)
	}

	// other processing work
}
```

## BFS 代码模板
[模板](https://shimo.im/docs/P8TqKHGKt3ytkYYd/read)

```go
func bfs(graph, start, end) {
	visited := map[*TreeNode]bool{}
	queue := []*TreeNode{}
	queue = append(queue, start)

	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		visited[node] = true
		process(node)
		nodes := generateRelatedNodes(node)
		queue = append(queue, nodes...)
	}

	// other processing work
}
```

## 二分查找代码模板
[模板](https://shimo.im/docs/hjQqRQkGgwd9g36J/read)

```go
func binarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)>>1
		if array[mid] == target {
			// find the target
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
```

## 深度优先搜索和广度优先搜索

### 实战题目
- https://leetcode.com/problems/binary-tree-level-order-traversal/
- https://leetcode.com/problems/minimum-genetic-mutation/
- https://leetcode.com/problems/generate-parentheses/
- https://leetcode.com/problems/find-largest-value-in-each-tree-row/

### 课后作业
- https://leetcode.com/problems/word-ladder/
- https://leetcode.com/problems/word-ladder-ii/
- https://leetcode.com/problems/number-of-islands/
- https://leetcode.com/problems/minesweeper/

## 贪心算法

### 参考链接
动态规划：https://zh.wikipedia.org/wiki/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92

### 实战题目
- https://leetcode.com/problems/coin-change/

### 课后作业
- https://leetcode.com/problems/lemonade-change/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
- https://leetcode.com/problems/assign-cookies/
- https://leetcode.com/problems/walking-robot-simulation/
- https://leetcode.com/problems/jump-game/
- https://leetcode.com/problems/jump-game-ii/

## 二分查找

### 参考链接
Fast InvSqrt() 扩展阅读: https://www.beyond3d.com/content/articles/8/

### 实战题目
- https://leetcode.com/problems/sqrtx/
_ https://leetcode.com/problems/valid-perfect-square/

### 课后作业
- https://leetcode.com/problems/search-in-rotated-sorted-array/
- https://leetcode.com/problems/search-a-2d-matrix/
- https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
- 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方
- 说明：同学们可以将自己的思路、代码写在第 4 周的学习总结中

## 本周作业

### 简单
- https://leetcode.com/problems/lemonade-change/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
- https://leetcode.com/problems/assign-cookies/ 完成1遍
- https://leetcode.com/problems/walking-robot-simulation/
- 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方
- 说明：同学们可以将自己的思路、代码写在第 4 周的学习总结中

### 中等
- https://leetcode.com/problems/word-ladder/
- https://leetcode.com/problems/number-of-islands/
- https://leetcode.com/problems/minesweeper/
- https://leetcode.com/problems/jump-game/ 完成1遍
- https://leetcode.com/problems/search-in-rotated-sorted-array/
- https://leetcode.com/problems/search-a-2d-matrix/
- https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

### 困难
- https://leetcode.com/problems/word-ladder-ii/
- https://leetcode.com/problems/jump-game-ii/ 完成2遍

## 每日一题
- https://leetcode.com/problems/search-in-rotated-sorted-array/
- https://leetcode.com/problems/jump-game-ii/
- https://leetcode.com/problems/sqrtx/
- https://leetcode.com/problems/climbing-stairs/
- https://leetcode.com/problems/binary-tree-level-order-traversal/
- https://leetcode.com/problems/assign-cookies/
- https://leetcode.com/problems/linked-list-cycle-ii/

