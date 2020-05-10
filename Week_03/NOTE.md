## 递归代码模板
[模板](https://shimo.im/docs/DjqqGCT3xqDYwPyY/read)

```go
func recursion(level, param int) { 
    // terminator 
    if (level > MAX_LEVEL) { 
        // process result 
        return
    }
 
    // process current logic 
    process(level, param)
 
    // drill down 
    recursion(level + 1, newParam)
 
    // restore current status
}
```

## 分治代码模板
[模板](https://shimo.im/docs/3xvghYh3JJPKwdvt/read)

```go
func divideConquer(problem, param1, param2, ...) {
    // recursion terminator
    if problem == nil {
        printResult()
        return
    }
    
    // prepare data 
    data := prepareData(problem) 
    subproblems := splitProblem(problem, data)

    //   conquer subproblems 
    subresult1 := divideConquer(subproblems[0], p1, ...) 
    subresult2 := divideConquer(subproblems[1], p1, ...) 
    subresult3 := divideConquer(subproblems[2], p1, ...) 
    …

    //   process and generate the final result 
    result := processResult(subresult1, subresult2, subresult3, …)
       	
    // revert the current level states
}
```

## 预习题目
- https://leetcode-cn.com/problems/powx-n/
- https://leetcode-cn.com/problems/subsets/
- https://leetcode.com/problems/generate-parentheses/

## 实战题目
- https://leetcode.com/problems/climbing-stairs/
- https://leetcode.com/problems/generate-parentheses/
- https://leetcode.com/problems/invert-binary-tree/description/
- https://leetcode.com/problems/validate-binary-search-tree
- https://leetcode.com/problems/maximum-depth-of-binary-tree
- https://leetcode.com/problems/minimum-depth-of-binary-tree
- https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
- https://leetcode.com/problems/majority-element/description/ （简单、但是高频）
- https://leetcode.com/problems/letter-combinations-of-a-phone-number/
- https://leetcode.com/problems/n-queens/

## 本周作业
|题目|难度|是否完成|第几遍|
|----|----|----|----|
|[236. 二叉树的最近公共祖先](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)|中等|是|1|
|[105. 从前序与中序遍历序列构造二叉树](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)|中等|是|1|
|[77. 组合](https://leetcode.com/problems/combinations/)|中等|否|0|
|[46. 全排列](https://leetcode.com/problems/permutations/)|中等|是|1|
|[47. 全排列 II](https://leetcode.com/problems/permutations-ii/)|中等|否|0|

- https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
- https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
- https://leetcode.com/problems/combinations/
- https://leetcode.com/problems/permutations/
- https://leetcode.com/problems/permutations-ii/

## 每日一题
- https://leetcode.com/problems/two-sum/
- https://leetcode.com/problems/3sum/
- https://leetcode.com/problems/n-queens/
- https://leetcode.com/problems/climbing-stairs/
- https://leetcode.com/problems/maximum-depth-of-binary-tree/
- https://leetcode.com/problems/majority-element/
- https://leetcode.com/problems/number-of-islands/
- https://leetcode.com/problems/ugly-number-ii/
- https://leetcode.com/problems/add-strings/
- https://leetcode.com/problems/minimum-genetic-mutation/
- https://leetcode.com/problems/minesweeper/
- https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
- https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
- https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
