package sort

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Expect []int
	}{
		{
			"MergeSort",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			mergeSort(test.Nums, 0, len(test.Nums)-1)
			assert.Equal(t, test.Expect, test.Nums)
		})
	}
}

