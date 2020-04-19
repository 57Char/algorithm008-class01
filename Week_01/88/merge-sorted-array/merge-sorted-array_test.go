package merge_sorted_array

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		Name   string
		Num1   []int
		M      int
		Num2   []int
		N      int
		Expect []int
	}{
		{
			"M == 0",
			[]int{0, 0},
			0,
			[]int{1, 2},
			2,
			[]int{1, 2},
		},
		{
			"N == 0",
			[]int{1, 2},
			2,
			[]int{},
			0,
			[]int{1, 2},
		},
		{
			"M > N",
			[]int{1, 2, 3, 0, 0},
			3,
			[]int{1, 3},
			2,
			[]int{1, 1, 2, 3, 3},
		},
		{
			"M = N",
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{1, 3, 4},
			3,
			[]int{1, 1, 2, 3, 3, 4},
		},
		{
			"M < N",
			[]int{1, 3, 0, 0, 0},
			2,
			[]int{1, 2, 4},
			3,
			[]int{1, 1, 2, 3, 4},
		},
		{
			"len(Num1) > (M + N)",
			[]int{1, 3, 0, 0, 0, 0},
			2,
			[]int{1, 2, 4},
			3,
			[]int{1, 1, 2, 3, 4, 0},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			merge(test.Num1, test.M, test.Num2, test.N)
			assert.Equal(t, test.Expect, test.Num1)
		})
	}
}

func TestMergeV2(t *testing.T) {
	cases := []struct {
		Name   string
		Num1   []int
		M      int
		Num2   []int
		N      int
		Expect []int
	}{
		{
			"M == 0",
			[]int{0, 0},
			0,
			[]int{1, 2},
			2,
			[]int{1, 2},
		},
		{
			"N == 0",
			[]int{1, 2},
			2,
			[]int{},
			0,
			[]int{1, 2},
		},
		{
			"M > N",
			[]int{1, 2, 3, 0, 0},
			3,
			[]int{1, 3},
			2,
			[]int{1, 1, 2, 3, 3},
		},
		{
			"M = N",
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{1, 3, 4},
			3,
			[]int{1, 1, 2, 3, 3, 4},
		},
		{
			"M < N",
			[]int{1, 3, 0, 0, 0},
			2,
			[]int{1, 2, 4},
			3,
			[]int{1, 1, 2, 3, 4},
		},
		{
			"len(Num1) > (M + N)",
			[]int{1, 3, 0, 0, 0, 0},
			2,
			[]int{1, 2, 4},
			3,
			[]int{1, 1, 2, 3, 4, 0},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			mergeV2(test.Num1, test.M, test.Num2, test.N)
			assert.Equal(t, test.Expect, test.Num1)
		})
	}
}
