package sqrtx

// https://leetcode.com/problems/sqrtx/

func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) >> 1
	}
	return r
}

func mySqrtV2(x int) int {
	if x <= 1 {
		return x
	}
	low, high := 0, x
	var mid, sqr int
	for {
		mid = low + (high-low)>>1
		if mid == low {
			return mid
		}
		sqr = mid * mid
		if sqr == x {
			return mid
		}
		if sqr > x {
			high = mid
		} else {
			low = mid
		}
	}
}