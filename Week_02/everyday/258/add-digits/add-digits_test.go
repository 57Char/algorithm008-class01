package add_digits

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestAddDigits(t *testing.T) {
	cases := []struct {
		Name   string
		Num    int
		Expect int
	}{
		{
			"[< 10] 9 -> 9",
			9,
			9,
		},
		{
			"[= 10] 10 -> 1 + 0 -> 1",
			10,
			1,
		},
		{
			"[> 10] 11 -> 1 + 1 -> 2",
			11,
			2,
		},
		{
			"[add N times] 38 -> 3 + 8 -> 11 -> 1 + 1 -> 2",
			38,
			2,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, addDigits(test.Num))
		})
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, addDigitsV2(test.Num))
		})
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, addDigitsV3(test.Num))
		})
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, addDigitsV4(test.Num))
		})
	}
}
