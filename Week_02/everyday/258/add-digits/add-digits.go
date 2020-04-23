package add_digits

// https://leetcode.com/problems/add-digits/

func addDigits(num int) int {
	return (num-1)%9 + 1
}

func addDigitsV2(num int) int {
	if num < 10 {
		return num
	}

	if r := num % 9; r == 0 {
		return 9
	} else {
		return r
	}
}

func addDigitsV3(num int) int {
	if num < 10 {
		return num
	}

	acc := 0
	for num+acc >= 10 {
		acc += num % 10
		num /= 10
		if num == 0 {
			num = acc
			acc = 0
		}
	}
	return num + acc
}

func addDigitsV4(num int) int {
	if num < 10 {
		return num
	}

	next := 0
	for num > 0 {
		next += num % 10
		num /= 10
	}
	return addDigitsV4(next)
}
