package move_zeroes

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Expect []int
	}{
		{
			"Nil slice",
			nil,
			nil,
		},
		{
			"Empty slice",
			[]int{},
			[]int{},
		},
		{
			"Only 0",
			[]int{0, 0},
			[]int{0, 0},
		},
		{
			"Have 0 and others I",
			[]int{1, 2, 3, 0, 0},
			[]int{1, 2, 3, 0, 0},
		},
		{
			"Have 0 and others II",
			[]int{1, 0, 2, 3, 0},
			[]int{1, 2, 3, 0, 0},
		},
		{
			"Have 0 and others III",
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			moveZeroes(test.Nums)
			assert.Equal(t, test.Expect, test.Nums)
		})
	}
}
