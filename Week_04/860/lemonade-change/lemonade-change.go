package lemonade_change

// https://leetcode.com/problems/lemonade-change/

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			if five >= 1 {
				five--
				ten++
			} else {
				return false
			}
		case 20:
			if five >= 1 && ten >= 1 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func lemonadeChangeV2(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five >= 1 {
				five--
				ten++
			} else {
				return false
			}
		} else {
			if five >= 1 && ten >= 1 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func lemonadeChangeV3(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			five--
			ten++
		} else if ten > 0 {
			five--
			ten--
		} else {
			five -= 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}