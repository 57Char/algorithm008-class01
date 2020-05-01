package number_of_islands

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	cases := []struct {
		Name   string
		Grid   [][]byte
		Expect int
	}{
		{
			"row is empty",
			nil,
			0,
		},
		{
			"col is empty",
			[][]byte{},
			0,
		},
		{
			"0 island",
			[][]byte{{}},
			0,
		},
		{
			"0 island II",
			[][]byte{{'0'}},
			0,
		},
		{
			"0 island III",
			[][]byte{{'0', '0'}, {'0', '0'}},
			0,
		},
		{
			"1 island",
			[][]byte{{'1'}},
			1,
		},
		{
			"1 island II",
			[][]byte{{'1', '0'}},
			1,
		},
		{
			"1 island III",
			[][]byte{{'1', '0'}, {'1', '0'}},
			1,
		},
		{
			"3 islands",
			[][]byte{{'1', '0', '0'}, {'0', '1', '0'}, {'0', '0', '1'}},
			3,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, numIslands(test.Grid))
		})
	}
}
