package powx_n

// https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n < 0 {
		x, n = 1/x, -n
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := pow(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}

func myPowV2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	half := myPow(x, n/2)
	res := myPow(x, n%2)
	return res * half * half
}

// TODO：牛顿迭代法