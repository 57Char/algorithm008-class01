package add_strings

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestAddStrings(t *testing.T) {
	cases := []struct {
		Name   string
		Nums1  string
		Nums2  string
		Expect string
	}{
		{
			"Nums1 is empty and Nums2 is empty",
			"",
			"",
			"",
		},
		{
			"Nums1 is empty but Nums2 is not empty",
			"",
			"222",
			"222",
		},
		{
			"Nums1 is not empty but Nums2 is empty",
			"111",
			"",
			"111",
		},
		{
			"0 + 0 = 0",
			"0",
			"0",
			"0",
		},
		{
			"1 + 9 = 10",
			"1",
			"9",
			"10",
		},
		{
			"99 + 1 = 100",
			"99",
			"1",
			"100",
		},
		{
			"998 + 1 = 999",
			"998",
			"1",
			"999",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, addStrings(test.Nums1, test.Nums2))
			assert.Equal(t, test.Expect, addStringsV2(test.Nums1, test.Nums2))
		})
	}
}
