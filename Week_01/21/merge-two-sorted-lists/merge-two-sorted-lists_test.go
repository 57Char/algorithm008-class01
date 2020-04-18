package merge_two_sorted_lists

import (
	"fmt"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		Name   string
		L1     *ListNode
		L2     *ListNode
		Expect string
	}{
		{
			"Empty L1",
			nil,
			&ListNode{},
			(&ListNode{}).String(),
		},
		{
			"Empty L2",
			&ListNode{},
			nil,
			(&ListNode{}).String(),
		},
		{
			"Empty L1 and Empty L2",
			nil,
			nil,
			"",
		},
		{
			"Merge L1 and L2 I",
			NewListNode(1, 2),
			NewListNode(1, 2),
			NewListNode(1, 1, 2, 2).String(),
		},
		{
			"Merge L1 and L2 II",
			NewListNode(1, 2),
			NewListNode(1, 2, 3),
			NewListNode(1, 1, 2, 2, 3).String(),
		},
		{
			"Merge L1 and L2 III",
			NewListNode(1, 2, 3, 4),
			NewListNode(2, 3, 5, 6),
			NewListNode(1, 2, 2, 3, 3, 4, 5, 6).String(),
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			merged := mergeTwoLists(test.L1, test.L2)
			if merged != nil {
				assert.Equal(t, test.Expect, merged.String())
			}
		})
	}
}

func TestMergeTwoListsV2(t *testing.T) {
	cases := []struct {
		Name   string
		L1     *ListNode
		L2     *ListNode
		Expect string
	}{
		{
			"Empty L1",
			nil,
			&ListNode{},
			(&ListNode{}).String(),
		},
		{
			"Empty L2",
			&ListNode{},
			nil,
			(&ListNode{}).String(),
		},
		{
			"Empty L1 and Empty L2",
			nil,
			nil,
			"",
		},
		{
			"Merge L1 and L2 I",
			NewListNode(1, 2),
			NewListNode(1, 2),
			NewListNode(1, 1, 2, 2).String(),
		},
		{
			"Merge L1 and L2 II",
			NewListNode(1, 2),
			NewListNode(1, 2, 3),
			NewListNode(1, 1, 2, 2, 3).String(),
		},
		{
			"Merge L1 and L2 III",
			NewListNode(1, 2, 3, 4),
			NewListNode(2, 3, 5, 6),
			NewListNode(1, 2, 2, 3, 3, 4, 5, 6).String(),
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			merged := mergeTwoListsV2(test.L1, test.L2)
			if merged != nil {
				assert.Equal(t, test.Expect, merged.String())
			}
		})
	}
}

func (l *ListNode) String() string {
	vals := []int{l.Val}

	ln := l
	for {
		ln = ln.Next
		if ln == nil {
			break
		}
		vals = append(vals, ln.Val)
	}

	return fmt.Sprintf("%v", vals)
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
