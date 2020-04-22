package intersection_of_two_arrays_ii

import (
	"sort"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		Name   string
		Num1   []int
		Num2   []int
		Expect []int
	}{
		{
			"len(Nums1) = 0 && len(Num2) = 0",
			[]int{},
			[]int{},
			[]int{},
		},
		{
			"len(Nums1) = 0 && len(Num2) > 0",
			[]int{},
			[]int{1},
			[]int{},
		},
		{
			"len(Nums1) > 0 && len(Num2) = 0",
			[]int{1},
			[]int{},
			[]int{},
		},
		{
			"len(Nums1) = len(Num2) and Have intersect",
			[]int{1},
			[]int{1},
			[]int{1},
		},
		{
			"len(Nums1) = len(Num2) and Have not intersect",
			[]int{1},
			[]int{2},
			[]int{},
		},
		{
			"len(Nums1) < len(Num2) and Have intersect",
			[]int{1},
			[]int{1, 2},
			[]int{1},
		},
		{
			"len(Nums1) < len(Num2) and Have not intersect",
			[]int{1},
			[]int{2, 3},
			[]int{},
		},
		{
			"len(Nums1) > len(Num2) and Have intersect",
			[]int{1, 2, 2, 3},
			[]int{2, 2, 3},
			[]int{2, 2, 3},
		},
		{
			"len(Nums1) > len(Num2) and Have not intersect",
			[]int{1, 2, 3, 4, 5},
			[]int{6, 7, 8},
			[]int{},
		},
		{
			"Other cases",
			[]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
			[]int{4, 3, 2, 1, 1, 2, 3, 4},
			[]int{1, 2, 3, 4, 1, 2, 3, 4},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			sort.Ints(test.Expect)
			inter := intersect(test.Num1, test.Num2)
			sort.Ints(inter)
			assert.Equal(t, test.Expect, inter)
		})
	}
}

func TestIntersectV2(t *testing.T) {
	cases := []struct {
		Name   string
		Num1   []int
		Num2   []int
		Expect []int
	}{
		{
			"len(Nums1) = 0 && len(Num2) = 0",
			[]int{},
			[]int{},
			[]int{},
		},
		{
			"len(Nums1) = 0 && len(Num2) > 0",
			[]int{},
			[]int{1},
			[]int{},
		},
		{
			"len(Nums1) > 0 && len(Num2) = 0",
			[]int{1},
			[]int{},
			[]int{},
		},
		{
			"len(Nums1) = len(Num2) and Have intersect",
			[]int{1},
			[]int{1},
			[]int{1},
		},
		{
			"len(Nums1) = len(Num2) and Have not intersect",
			[]int{1},
			[]int{2},
			[]int{},
		},
		{
			"len(Nums1) < len(Num2) and Have intersect",
			[]int{1},
			[]int{1, 2},
			[]int{1},
		},
		{
			"len(Nums1) < len(Num2) and Have not intersect",
			[]int{1},
			[]int{2, 3},
			[]int{},
		},
		{
			"len(Nums1) > len(Num2) and Have intersect",
			[]int{1, 2, 2, 3},
			[]int{2, 2, 3},
			[]int{2, 2, 3},
		},
		{
			"len(Nums1) > len(Num2) and Have not intersect",
			[]int{1, 2, 3, 4, 5},
			[]int{6, 7, 8},
			[]int{},
		},
		{
			"Other cases",
			[]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
			[]int{4, 3, 2, 1, 1, 2, 3, 4},
			[]int{1, 2, 3, 4, 1, 2, 3, 4},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			sort.Ints(test.Expect)
			inter := intersectV2(test.Num1, test.Num2)
			sort.Ints(inter)
			assert.Equal(t, test.Expect, inter)
		})
	}
}

func TestIntersectV3(t *testing.T) {
	cases := []struct {
		Name   string
		Num1   []int
		Num2   []int
		Expect []int
	}{
		{
			"len(Nums1) = 0 && len(Num2) = 0",
			[]int{},
			[]int{},
			[]int{},
		},
		{
			"len(Nums1) = 0 && len(Num2) > 0",
			[]int{},
			[]int{1},
			[]int{},
		},
		{
			"len(Nums1) > 0 && len(Num2) = 0",
			[]int{1},
			[]int{},
			[]int{},
		},
		{
			"len(Nums1) = len(Num2) and Have intersect",
			[]int{1},
			[]int{1},
			[]int{1},
		},
		{
			"len(Nums1) = len(Num2) and Have not intersect",
			[]int{1},
			[]int{2},
			[]int{},
		},
		{
			"len(Nums1) < len(Num2) and Have intersect",
			[]int{1},
			[]int{1, 2},
			[]int{1},
		},
		{
			"len(Nums1) < len(Num2) and Have not intersect",
			[]int{1},
			[]int{2, 3},
			[]int{},
		},
		{
			"len(Nums1) > len(Num2) and Have intersect",
			[]int{1, 2, 2, 3},
			[]int{2, 2, 3},
			[]int{2, 2, 3},
		},
		{
			"len(Nums1) > len(Num2) and Have not intersect",
			[]int{1, 2, 3, 4, 5},
			[]int{6, 7, 8},
			[]int{},
		},
		{
			"Other cases",
			[]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
			[]int{4, 3, 2, 1, 1, 2, 3, 4},
			[]int{1, 2, 3, 4, 1, 2, 3, 4},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			sort.Ints(test.Expect)
			inter := intersectV3(test.Num1, test.Num2)
			sort.Ints(inter)
			assert.Equal(t, test.Expect, inter)
		})
	}
}
