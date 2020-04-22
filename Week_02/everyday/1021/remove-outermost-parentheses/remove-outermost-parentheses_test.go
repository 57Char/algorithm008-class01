package remove_outermost_parentheses

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestRemoveOuterParentheses(t *testing.T) {
	cases := []struct {
		Name   string
		S      string
		Expect string
	}{
		{
			"Empty -> Empty",
			"",
			"",
		},
		{
			"() -> Empty",
			"()",
			"",
		},
		{
			"()() -> Empty",
			"()()",
			"",
		},
		{
			"(()())(()) -> ()()()",
			"(()())(())",
			"()()()",
		},
		{
			"(()())(())(()(())) -> ()()()()(())",
			"(()())(())(()(()))",
			"()()()()(())",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, removeOuterParentheses(test.S))
		})
	}
}

func TestRemoveOuterParenthesesV2(t *testing.T) {
	cases := []struct {
		Name   string
		S      string
		Expect string
	}{
		{
			"Empty -> Empty",
			"",
			"",
		},
		{
			"() -> Empty",
			"()",
			"",
		},
		{
			"()() -> Empty",
			"()()",
			"",
		},
		{
			"(()())(()) -> ()()()",
			"(()())(())",
			"()()()",
		},
		{
			"(()())(())(()(())) -> ()()()()(())",
			"(()())(())(()(()))",
			"()()()()(())",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, removeOuterParenthesesV2(test.S))
		})
	}
}

func TestRemoveOuterParenthesesV3(t *testing.T) {
	cases := []struct {
		Name   string
		S      string
		Expect string
	}{
		{
			"Empty -> Empty",
			"",
			"",
		},
		{
			"() -> Empty",
			"()",
			"",
		},
		{
			"()() -> Empty",
			"()()",
			"",
		},
		{
			"(()())(()) -> ()()()",
			"(()())(())",
			"()()()",
		},
		{
			"(()())(())(()(())) -> ()()()()(())",
			"(()())(())(()(()))",
			"()()()()(())",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, removeOuterParenthesesV3(test.S))
		})
	}
}

func TestRemoveOuterParenthesesV4(t *testing.T) {
	cases := []struct {
		Name   string
		S      string
		Expect string
	}{
		{
			"Empty -> Empty",
			"",
			"",
		},
		{
			"() -> Empty",
			"()",
			"",
		},
		{
			"()() -> Empty",
			"()()",
			"",
		},
		{
			"(()())(()) -> ()()()",
			"(()())(())",
			"()()()",
		},
		{
			"(()())(())(()(())) -> ()()()()(())",
			"(()())(())(()(()))",
			"()()()()(())",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, removeOuterParenthesesV4(test.S))
		})
	}
}
