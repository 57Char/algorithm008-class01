package minimum_genetic_mutation

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMinMutation(t *testing.T) {
	cases := []struct {
		Name   string
		Start  string
		End    string
		Bank   []string
		Expect int
	}{
		{
			"Bank is empty",
			"AAAACCCC",
			"AACCGGTA",
			[]string{},
			-1,
		},
		{
			"AAAACCCC -> AAAACCCC",
			"AAAACCCC",
			"AAAACCCC",
			[]string{"AAAACCCC"},
			0,
		},
		{
			"AACCGGTT -> AACCGGTA",
			"AACCGGTT",
			"AACCGGTA",
			[]string{"AACCGGTA"},
			1,
		},
		{
			"AACCGGTT -> AAACGGTA",
			"AACCGGTT",
			"AAACGGTA",
			[]string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			2,
		},
		{
			"AACCCCCC -> AAAAACCC",
			"AAAAACCC",
			"AACCCCCC",
			[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			3,
		},
		{
			"AAAACCCC -> CCCCCCCC",
			"AAAACCCC",
			"CCCCCCCC",
			[]string{"AAAACCCA", "AAACCCCA", "AACCCCCA", "AACCCCCC", "ACCCCCCC", "CCCCCCCC", "AAACCCCC", "AACCCCCC"},
			4,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, minMutation(test.Start, test.End, test.Bank))
			assert.Equal(t, test.Expect, minMutationV2(test.Start, test.End, test.Bank))
			assert.Equal(t, test.Expect, minMutationV3(test.Start, test.End, test.Bank))
			assert.Equal(t, test.Expect, minMutationV4(test.Start, test.End, test.Bank))
		})
	}
}
