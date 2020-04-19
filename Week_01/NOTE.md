# 学习笔记

## 双指针

本周的题目很大一部分都可以用双指针解决，我认为很多题目纯粹是为了创造"双指针"适用条件。相对的，实际工作中却极少出现能用上"双指针"场景。

基本思路：
- 两个指针，两头往中间推，相交则结束
- 两个指针，一个快，一个慢，快指针走完则结束

## 283. Move Zeroes

https://leetcode.com/problems/move-zeroes

```go
func moveZeroes(nums []int) {
	// 为空，退出
	if len(nums) == 0 {
		return
	}
	// 双指针: i 和 j
	j := 0 // 非0下标
	for i, val := range nums {
		// 值不为0则处理
		if val != 0 {
			// 值交换
			nums[i], nums[j] = 0, val
			// 注意交换顺序，下面的代码会有问题
			// nums[j], nums[i] = val, 0
			// 下标前移
			j++
		}
	}
}
```

```go
func moveZeroes(nums []int) {
	// 为空，退出
	if len(nums) == 0 {
		return
	}
	// 双指针: i 和 j
	j := 0 // 非0下标
	for i, val := range nums {
		// 值不为0则处理
		if val != 0 {
			// 值交换，避免上一种解法的问题，也可以nums[j], nums[i] = val, nums[j]
			nums[i], nums[j] = nums[j], val
			// 下标前移
			j++
		}
	}
}
```

```go
func moveZeroes(nums []int) {
	// 为空，退出
	if len(nums) == 0 {
		return
	}
	// 双指针: i 和 j
	j := 0 // 非0下标
	for i, val := range nums {
		// 值不为0则处理
		if val != 0 {
			// 下标不重合则设置为0
			if i != j {
				nums[i] = 0
			}
			// 交换值
			nums[j] = val
			// 下标前移
			j++
		}
	}
}
```

## 66. Plus One

https://leetcode.com/problems/plus-one/

```go
func plusOne(digits []int) []int {
	n := len(digits)
	// 题目中说数组不为空，所以未考虑n==0的情况
	// 从最后一个下标开始处理
	for i := n - 1; i >= 0; i-- {
		// 不进位
		if digits[i] < 9 {
			// 加1，退出
			digits[i]++
			return digits
		}
		// 进位 9 + 1 -> 0
		digits[i] = 0
	}

	// 全是9，进位溢出，新建数组，第1位为1，其他为0
	digits = make([]int, n+1)
	digits[0] = 1

	return digits
}

func plusOne(digits []int) []int {
	n := len(digits)
	// 题目中说数组不为空，所以未考虑n==0的情况
	// 从最后一个下标开始处理
	for i := n - 1; i >= 0; i-- {
		// 先加1
		digits[i]++
		// 求余
		digits[i] = digits[i] % 10
		// 判断是否进位，不进位则直接退出
		if digits[i] != 0 {
			return digits
		}
	}

	// 全是9，进位溢出，新建数组，第1位为1，其他为0
	digits = make([]int, n+1)
	digits[0] = 1

	return digits
}
```