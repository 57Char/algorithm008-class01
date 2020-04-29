package ti_huan_kong_ge_lcof

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestReplaceSpace(t *testing.T) {
	cases := []struct {
		Name   string
		S      string
		Expect string
	}{
		{
			"S is empty",
			"",
			"",
		},
		{
			"S without spaces",
			"abc",
			"abc",
		},
		{
			"S with spaces I",
			"   ",
			"%20%20%20",
		},
		{
			"S with spaces II",
			"We are happy.",
			"We%20are%20happy.",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, replaceSpace(test.S))
			assert.Equal(t, test.Expect, replaceSpaceV2(test.S))
		})
	}
}
