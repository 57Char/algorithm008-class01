package jump_game

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Expect bool
	}{
		{
			"Empty",
			[]int{},
			false,
		},
		{
			"No need to jump",
			[]int{1},
			true,
		},
		{
			"Can jump",
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			"Can't jump",
			[]int{3, 2, 1, 0, 4},
			false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, canJump(test.Nums))
			assert.Equal(t, test.Expect, canJumpV2(test.Nums))
		})
	}
}
