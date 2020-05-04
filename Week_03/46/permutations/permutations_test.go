package permutations

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Expect [][]int
	}{
		{
			"[1,2,3]",
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			permute := permute(test.Nums)
			assert.Equal(t, test.Expect, permute)
		})
	}
}
