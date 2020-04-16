package rotate_array

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		K      int
		Expect []int
	}{
		{
			"Nil slice",
			nil,
			0,
			nil,
		},
		{
			"Empty slice",
			[]int{},
			0,
			[]int{},
		},
		{
			"k < n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"k = n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			7,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			"k > n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			8,
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			"k % n = 0",
			[]int{1, 2, 3, 4, 5, 6, 7},
			700,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			rotate(test.Nums, test.K)
			assert.Equal(t, test.Expect, test.Nums)
		})
	}
}

func TestRotateV2(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		K      int
		Expect []int
	}{
		{
			"Nil slice",
			nil,
			0,
			nil,
		},
		{
			"Empty slice",
			[]int{},
			0,
			[]int{},
		},
		{
			"k < n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"k = n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			7,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			"k > n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			8,
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			"k % n = 0",
			[]int{1, 2, 3, 4, 5, 6, 7},
			700,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			rotateV2(test.Nums, test.K)
			assert.Equal(t, test.Expect, test.Nums)
		})
	}
}

func TestRotateV3(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		K      int
		Expect []int
	}{
		{
			"Nil slice",
			nil,
			0,
			nil,
		},
		{
			"Empty slice",
			[]int{},
			0,
			[]int{},
		},
		{
			"k < n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"k = n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			7,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			"k > n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			8,
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			"k % n = 0",
			[]int{1, 2, 3, 4, 5, 6, 7},
			700,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			rotateV3(test.Nums, test.K)
			assert.Equal(t, test.Expect, test.Nums)
		})
	}
}

func TestRotateV4(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		K      int
		Expect []int
	}{
		{
			"Nil slice",
			nil,
			0,
			nil,
		},
		{
			"Empty slice",
			[]int{},
			0,
			[]int{},
		},
		{
			"k < n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"k = n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			7,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			"k > n",
			[]int{1, 2, 3, 4, 5, 6, 7},
			8,
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			"k % n = 0",
			[]int{1, 2, 3, 4, 5, 6, 7},
			700,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			rotateV4(test.Nums, test.K)
			assert.Equal(t, test.Expect, test.Nums)
		})
	}
}