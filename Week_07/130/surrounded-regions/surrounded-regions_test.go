package surrounded_regions

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		Name   string
		Board  [][]byte
		Expect [][]byte
	}{
		{
			"Board",
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			solve(test.Board)
			assert.Equal(t, test.Expect, test.Board)
		})
	}
}
