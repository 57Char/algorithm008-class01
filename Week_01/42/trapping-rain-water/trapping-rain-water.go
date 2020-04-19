package trapping_rain_water

// https://leetcode.com/problems/trapping-rain-water/

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	res, left, right, maxLeft, maxRight := 0, 0, n-1, 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}

	return res
}

func trapV2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	max, maxIndex := 0, 0
	for i, val := range height {
		if max < val {
			max = val
			maxIndex = i
		}
	}

	res, left, right, maxLeft, maxRight := 0, 0, n-1, 0, 0
	for ; left < maxIndex; left++ {
		if height[left] >= maxLeft {
			maxLeft = height[left]
		} else {
			res += maxLeft - height[left]
		}
	}
	for ; right > maxIndex; right-- {
		if height[right] >= maxRight {
			maxRight = height[right]
		} else {
			res += maxRight - height[right]
		}
	}

	return res
}
