package ugly_number_ii

// https://leetcode.com/problems/ugly-number-ii/

func nthUglyNumber(n int) int {
	if n <= 1 {
		return 1
	}
	uglys := make([]int, 0, n)
	uglys = append(uglys, 1)
	i2, i3, i5, nth := 0, 0, 0, len(uglys)
	for nth < n {
		u2, u3, u5 := uglys[i2]*2, uglys[i3]*3, uglys[i5]*5
		u := min(min(u2, u3), u5)
		if u == u2 {
			i2++
		}
		if u == u3 {
			i3++
		}
		if u == u5 {
			i5++
		}
		uglys = append(uglys, u)
		nth++
	}
	return uglys[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
