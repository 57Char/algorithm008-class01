package zero_matrix_lcci

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	cases := []struct {
		Name   string
		matrix [][]int
		Expect [][]int
	}{
		{
			"0 row and 0 col",
			[][]int{},
			[][]int{},
		},
		{
			"1 row and 0 col",
			[][]int{
				{},
			},
			[][]int{
				{},
			},
		},
		{
			"1 row and 1 col",
			[][]int{
				{1},
			},
			[][]int{
				{1},
			},
		},
		{
			"1 row and 1 col II",
			[][]int{
				{0},
			},
			[][]int{
				{0},
			},
		},
		{
			"N row and M col",
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			"N row and M col II",
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			setZeroes(test.matrix)
			assert.Equal(t, test.Expect, test.matrix)
		})
	}
}

func TestSetZeroesV2(t *testing.T) {
	cases := []struct {
		Name   string
		matrix [][]int
		Expect [][]int
	}{
		{
			"0 row and 0 col",
			[][]int{},
			[][]int{},
		},
		{
			"1 row and 0 col",
			[][]int{
				{},
			},
			[][]int{
				{},
			},
		},
		{
			"1 row and 1 col",
			[][]int{
				{1},
			},
			[][]int{
				{1},
			},
		},
		{
			"1 row and 1 col II",
			[][]int{
				{0},
			},
			[][]int{
				{0},
			},
		},
		{
			"N row and M col",
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			"N row and M col II",
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			setZeroesV2(test.matrix)
			assert.Equal(t, test.Expect, test.matrix)
		})
	}
}
