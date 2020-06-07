# 第13课 | 字典树和并查集

## 1. Trie树的基本实现和特性

### 参考链接
1. [二叉树的层次遍历](https://leetcode.com/problems/binary-tree-level-order-traversal/)
2. [实现 Trie](https://leetcode.com/problems/implement-trie-prefix-tree/)
3. [Tire 树代码模板](https://shimo.im/docs/Pk6vPY3HJ9hKkh33) TODO: 改成Golang版本

## 2. Trie树实战题目解析：单词搜索2

### 实战题目
- https://leetcode.com/problems/implement-trie-prefix-tree/
- https://leetcode.com/problems/word-search-ii/
- 分析单词搜索 2 用 Tire 树方式实现的时间复杂度，请同学们提交在学习总结中。

## 3. 并查集的基本实现、特性和实战题目解析

### 参考链接
- [岛屿数量](https://leetcode.com/problems/number-of-islands/)
- [并查集代码模板](https://shimo.im/docs/ydPCH33xDhK9YwWR)

```go
type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		parent: make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

func (u *UnionFind) Find(x int) int {
	p := u.parent
	for p[x] != x {
		p[x] = p[p[x]]
		x = p[x]
	}
	return x
}

func (u *UnionFind) Union(p, q int) {
	rootP, rootQ := u.Find(p), u.Find(q)
	if rootP == rootQ {
		return
	}
	u.parent[rootP] = rootQ
	u.count--
}

func (u *UnionFind) Count() int {
	return u.count
}
```

# 第14课 | 高级搜索

## 1. 剪枝的实现和特性

## 参考链接
- [DFS 代码模板](https://shimo.im/docs/ddgwCccJQKxkrcTq)
- [BFS 代码模板](https://shimo.im/docs/P8TqKHGKt3ytkYYd)
- [AlphaZero Explained](https://nikcheerla.github.io/deeplearningschool/2018/01/01/AlphaZero-Explained/)
- [棋类复杂度](https://en.wikipedia.org/wiki/Game_complexity)

## 2. 剪枝实战题目解析：数独

### 实战题目
- https://leetcode.com/problems/climbing-stairs/
- https://leetcode.com/problems/generate-parentheses/
- https://leetcode.com/problems/n-queens/
- https://leetcode.com/problems/valid-sudoku/
- https://leetcode.com/problems/sudoku-solver/

## 3. 双向BFS的实现、特性和题解

### 实战题目
- https://leetcode-cn.com/problems/word-ladder/
- https://leetcode-cn.com/problems/minimum-genetic-mutation/

### 课后作业
- 总结双向 BFS 代码模版，请同学们提交在学习总结中。

## 4. 启发式搜索的实现、特性和题解

### 参考链接
- [A* 代码模板](https://shimo.im/docs/CXvjHyWhpQcxXjcw)
- [相似度测量方法](https://dataaspirant.com/2015/04/11/five-most-popular-similarity-measures-implementation-in-python/)
- [二进制矩阵中的最短路径的 A* 解法](https://leetcode.com/problems/shortest-path-in-binary-matrix/discuss/313347/A*-search-in-Python)
- [8 puzzles 解法比较](https://zxi.mytechroad.com/blog/searching/8-puzzles-bidirectional-astar-vs-bidirectional-bfs/)

### 实战题目
- https://leetcode.com/problems/shortest-path-in-binary-matrix/
- https://leetcode.com/problems/sliding-puzzle/
- https://leetcode.com/problems/sudoku-solver/

# 第15课 | 红黑树和AVL树

## AVL树和红黑树的实现和特性

### 参考链接
- [平衡树](https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree)

# 本周作业

## 简单
- https://leetcode.com/problems/climbing-stairs/

## 中等
- https://leetcode.com/problems/implement-trie-prefix-tree/
- https://leetcode.com/problems/friend-circles
- https://leetcode.com/problems/number-of-islands/
- https://leetcode.com/problems/surrounded-regions/
- https://leetcode.com/problems/valid-sudoku/
- https://leetcode.com/problems/generate-parentheses/
- https://leetcode.com/problems/word-ladder/
- https://leetcode.com/problems/minimum-genetic-mutation/

## 困难
- https://leetcode.com/problems/word-search-ii/
- https://leetcode.com/problems/n-queens/
- https://leetcode.com/problems/sudoku-solver/

# 每日一题
- Day50 https://leetcode.com/problems/binary-tree-inorder-traversal/
- Day51 https://leetcode.com/problems/friend-circles/
- Day52 https://leetcode.com/problems/n-queens/
- Day53 https://leetcode.com/problems/surrounded-regions/
- Day54 https://leetcode.com/problems/valid-sudoku/
- Day55 https://leetcode.com/problems/contains-duplicate/
- Day55 https://leetcode.com/problems/intersection-of-two-arrays-ii/
- Day56 https://leetcode.com/problems/word-ladder/

# 其他
- [算法8期期中考试错题详情](https://shimo.im/sheets/3htvPPpgpdRYDYwg/BTaQL/)
- [算法训练营-期中考试题](https://shimo.im/docs/69uDNhxTzjADzqwQ/)  