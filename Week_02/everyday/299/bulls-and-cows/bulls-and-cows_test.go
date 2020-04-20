package bulls_and_cows

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestGetHint(t *testing.T) {
	cases := []struct {
		Name   string
		Secret string
		Guess  string
		Expect string
	}{
		{
			"1A3B",
			"1807",
			"7810",
			"1A3B",
		},
		{
			"0A0B I",
			"1111",
			"2222",
			"0A0B",
		},
		{
			"0A0B II",
			"",
			"",
			"0A0B",
		},
		{
			"0A4B",
			"1122",
			"2211",
			"0A4B",
		},
		{
			"4A0B",
			"1234",
			"1234",
			"4A0B",
		},
		{
			"3A0B",
			"1234",
			"1235",
			"3A0B",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, getHint(test.Secret, test.Guess))
		})
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, getHintV2(test.Secret, test.Guess))
		})
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, getHintV3(test.Secret, test.Guess))
		})
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, getHintV4(test.Secret, test.Guess))
		})
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, getHintV5(test.Secret, test.Guess))
		})
	}
}
