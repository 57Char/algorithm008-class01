package valid_sudoku

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestIsValidSudoku(t *testing.T) {
	cases := []struct {
		Name   string
		Board  [][]byte
		Expect bool
	}{
		{
			"Valid Sudoku",
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			true,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, isValidSudoku(test.Board))
			assert.Equal(t, test.Expect, isValidSudokuV2(test.Board))
			assert.Equal(t, test.Expect, isValidSudokuV3(test.Board))
		})
	}
}
