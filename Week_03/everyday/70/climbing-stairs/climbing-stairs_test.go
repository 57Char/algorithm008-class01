package climbing_stairs

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		Name   string
		N      int
		Expect int
	}{
		{
			"N < 3",
			2,
			2,
		},
		{
			"N = 3",
			3,
			3,
		},
		{
			"N > 3",
			5,
			8,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, climbStairs(test.N))
		})
	}
}
