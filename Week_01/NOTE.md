# 学习笔记

## 283. Move Zeroes

https://leetcode.com/problems/move-zeroes

```go
func moveZeroes(nums []int) {
	// 为空，退出
	if len(nums) == 0 {
		return
	}
	// 双指针: i 和 j
	j := 0
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
	j := 0
	for i, val := range nums {
		// 值不为0则处理
		if val != 0 {
            // 下标不重合则设置为0
			if i != j {
				nums[i] = 0
			}
			nums[j] = val
			// 下标前移
			j++
		}
	}
}
```
