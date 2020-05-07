package majority_element

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Expect int
	}{
		{
			"One element",
			[]int{1},
			1,
		},
		{
			"Some elements I",
			[]int{1, 1, 1, 1, 1},
			1,
		},
		{
			"Some elements II",
			[]int{1, 2, 2, 1, 2},
			2,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, majorityElement(test.Nums))
			assert.Equal(t, test.Expect, majorityElementV2(test.Nums))
		})
	}
}
