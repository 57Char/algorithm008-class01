## 动态规划
### 定义
- Wiki 定义: https://en.wikipedia.org/wiki/Dynamic_programming
- “Simplifying a complicated problem by breaking it down into simpler sub-problems”
(in a recursive manner)
- Divide & Conquer + Optimal substructure 分治 + 最优子结构

### 关键点
- 动态规划 和 递归或者分治 没有根本上的区别(关键看有无最优的子结构)
- 共性: 找到重复子问题
- 差异性: 最优子结构、中途可以淘汰次优解

- 最优子结构: opt[n] = best_of(opt[n-1], opt[n-2], ...)
- 储存中间状态: opt[i]
- 递推公式(美其名曰: 状态转移方程或者 DP 方程)
 - Fib:    opt[i] = opt[n-1] + opt[n-2]
 - 二维路径: opt[i,j] = opt[i+1][j] + opt[i][j+1] (且判断a[i,j]是否空地)

### 小结
- 打破自己的思维惯性，形成机器思维
- 理解复杂逻辑的关键
- 也是职业进阶的要点要领

## 参考链接
- [不同路径](https://leetcode.com/problems/unique-paths/)
- [不同路径2](https://leetcode.com/problems/unique-paths-ii/)
- [最长公共子序列](https://leetcode.com/problems/longest-common-subsequence/)
- [MIT 动态规划课程最短路径算法](https://www.bilibili.com/video/av53233912?from=search&seid=2847395688604491997)

## 实战题目
- https://leetcode.com/problems/climbing-stairs/
- https://leetcode.com/problems/triangle/
- https://leetcode.com/problems/triangle/discuss/38735/Python-easy-to-understand-solutions-(top-down-bottom-up)
- https://leetcode.com/problems/maximum-subarray/
- https://leetcode.com/problems/maximum-product-subarray/
- https://leetcode.com/problems/coin-change/

- https://leetcode.com/problems/house-robber/
- https://leetcode.com/problems/house-robber-ii/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
- [一个方法团灭 6 道股票问题](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/)

## 高级 DP 实战题目
- https://leetcode.com/problems/perfect-squares/
- https://leetcode.com/problems/edit-distance/ （重点）
- https://leetcode.com/problems/jump-game/
- https://leetcode.com/problems/jump-game-ii/
- https://leetcode.com/problems/unique-paths/
- https://leetcode.com/problems/unique-paths-ii/
- https://leetcode.com/problems/unique-paths-iii/
- https://leetcode.com/problems/coin-change/
- https://leetcode.com/problems/coin-change-2/

## 本周作业
### 中等
- https://leetcode.com/problems/minimum-path-sum/
- https://leetcode.com/problems/decode-ways
- https://leetcode.com/problems/maximal-square/
- https://leetcode.com/problems/task-scheduler/
- https://leetcode.com/problems/palindromic-substrings/
### 困难
- https://leetcode.com/problems/longest-valid-parentheses/
- https://leetcode.com/problems/edit-distance/
- https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/
- https://leetcode.com/problems/frog-jump/
- https://leetcode.com/problems/split-array-largest-sum
- https://leetcode.com/problems/student-attendance-record-ii/
- https://leetcode.com/problems/minimum-window-substring/
- https://leetcode.com/problems/burst-balloons/

## 每日一题
- https://leetcode.com/problems/4sum/
- https://leetcode.com/problems/permutations-ii/
- https://leetcode.com/problems/n-queens/
- https://leetcode.com/problems/search-a-2d-matrix/
- https://leetcode.com/problems/combinations/
- https://leetcode.com/problems/triangle/
- https://leetcode.com/problems/word-ladder/
- https://leetcode.com/problems/majority-element/
- https://leetcode.com/problems/house-robber/
- https://leetcode.com/problems/maximal-square/
- https://leetcode.com/problems/coin-change/
- https://leetcode.com/problems/n-ary-tree-level-order-traversal/
- https://leetcode.com/problems/number-of-segments-in-a-string/
- https://leetcode.com/problems/lemonade-change/
- https://leetcode.com/problems/task-scheduler/
- https://leetcode.com/problems/burst-balloons/ 未做

## 参考资料 
- https://leetcode-cn.com/problems/house-robber/solution/dong-tai-gui-hua-jie-ti-si-bu-zou-xiang-jie-cjavap/