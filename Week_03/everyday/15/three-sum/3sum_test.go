package three_sum

import (
	"fmt"
	"sort"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func Test3Sum(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Expect [][]int
	}{
		{
			"len(Nums) < 3",
			[]int{1, 2},
			[][]int{},
		},
		{
			"len(Nums) = 3 but not match",
			[]int{1, 2, 3},
			[][]int{},
		},
		{
			"len(Nums) = 3 and match",
			[]int{10, -10, 0},
			[][]int{{-10, 0, 10}},
		},
		{
			"3sum I",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			"3sum II",
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			"3sum III",
			[]int{-2, 0, 1, 1, 2},
			[][]int{{-2, 1, 1}, {-2, 0, 2}},
		},
		{
			"3sum IV",
			[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			[][]int{{-4, 2, 2}, {-4, 1, 3}, {-4, 0, 4}, {-4, -2, 6}, {-2, 0, 2}, {-2, -2, 4}},
		},
	}

	transfer := func(nums [][]int) map[string]bool {
		m := map[string]bool{}
		for _, n := range nums {
			sort.Ints(n)
			m[fmt.Sprintf("%v", n)] = true
		}
		return m
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, transfer(test.Expect), transfer(threeSum(test.Nums)))
			assert.Equal(t, transfer(test.Expect), transfer(threeSumV2(test.Nums)))
		})
	}
}
