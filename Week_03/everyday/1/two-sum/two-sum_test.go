package two_sum

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Target int
		Expect []int
	}{
		{
			"not match",
			[]int{2, 3, 7, 9},
			999,
			[]int{},
		},
		{
			"two sum I",
			[]int{2, 3, 7, 9},
			5,
			[]int{0, 1},
		},
		{
			"two sum II",
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			"two sum III",
			[]int{2, 7, 11, 15},
			22,
			[]int{1, 3},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, twoSum(test.Nums, test.Target))
		})
	}
}
