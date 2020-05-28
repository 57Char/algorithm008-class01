package combinations

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	cases := []struct {
		Name   string
		N      int
		K      int
		Expect [][]int
	}{
		{
			"n = 4 and k = 2",
			4,
			2,
			[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			"n = 3 and k = 3",
			3,
			3,
			[][]int{{1, 2, 3}},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, combine(test.N, test.K))
			assert.Equal(t, test.Expect, combineV2(test.N, test.K))
		})
	}
}
