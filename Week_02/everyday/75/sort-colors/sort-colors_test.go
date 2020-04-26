package sort_colors

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Expect []int
	}{
		{
			"Empty Nums",
			[]int{},
			[]int{},
		},
		{
			"len(Nums) = 1",
			[]int{1},
			[]int{1},
		},
		{
			"Colors",
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			"Colors II",
			[]int{2, 1, 0},
			[]int{0, 1, 2},
		},
		{
			"Colors III",
			[]int{1, 0},
			[]int{0, 1},
		},
		{
			"Colors IV",
			[]int{1, 1, 2, 2, 0, 0, 1, 2, 2, 1, 0},
			[]int{0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			sortColors(test.Nums)
			assert.Equal(t, test.Expect, test.Nums)
		})
	}
}
