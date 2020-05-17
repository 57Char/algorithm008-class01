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