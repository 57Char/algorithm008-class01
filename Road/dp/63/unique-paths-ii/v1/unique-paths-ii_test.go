package v1

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	cases := []struct {
		Name         string
		ObstacleGrid [][]int
		Expect       int
	}{
		{
			"[[1,0]]",
			[][]int{{1, 0}},
			0,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, uniquePathsWithObstacles(test.ObstacleGrid))
		})
	}
}
