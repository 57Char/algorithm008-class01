package plus_one

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Expect []int
	}{
		{
			"Plus 1 I",
			[]int{0, 0},
			[]int{0, 1},
		},
		{
			"Plus 1 II",
			[]int{1, 2, 3, 0, 0},
			[]int{1, 2, 3, 0, 1},
		},
		{
			"Last 9",
			[]int{1, 9},
			[]int{2, 0},
		},
		{
			"Last 99",
			[]int{1, 9, 9},
			[]int{2, 0, 0},
		},
		{
			"All 999..",
			[]int{9, 9, 9},
			[]int{1, 0, 0, 0},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, plusOne(test.Nums))
		})
	}
}
