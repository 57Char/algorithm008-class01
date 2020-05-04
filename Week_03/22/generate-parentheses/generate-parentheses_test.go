package generate_parentheses

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	cases := []struct {
		Name   string
		N      int
		Expect []string
	}{
		{
			"N < 0",
			-1,
			[]string{},
		},
		{
			"N = 0",
			0,
			[]string{},
		},
		{
			"N = 1",
			1,
			[]string{"()"},
		},
		{
			"N = 3",
			3,
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, generateParenthesis(test.N))
		})
	}
}
