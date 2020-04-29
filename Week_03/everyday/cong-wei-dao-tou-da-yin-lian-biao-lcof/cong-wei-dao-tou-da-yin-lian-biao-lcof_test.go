package cong_wei_dao_tou_da_yin_lian_biao_lcof

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestReversePrint(t *testing.T) {
	cases := []struct {
		Name   string
		Head   *ListNode
		Expect []int
	}{
		{
			"Empty Head",
			nil,
			[]int{},
		},
		{
			"Only one node",
			&ListNode{Val: 1},
			[]int{1},
		},
		{
			"Many nodes",
			NewListNode(1, 2, 3, 4),
			[]int{4, 3, 2, 1},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, reversePrint(test.Head))
			assert.Equal(t, test.Expect, reversePrintV2(test.Head))
		})
	}
}

func NewListNode(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	v := vals[0]

	l := &ListNode{
		Val: v,
	}

	if len(vals) > 1 {
		l.Next = NewListNode(vals[1:]...)
	}

	return l
}
