package number_of_1_bits

// https://leetcode.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	var count uint32
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return int(count)
}
