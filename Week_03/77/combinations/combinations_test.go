package combinations

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	cases := []struct {
		Name   string
		N      int
		K      int
		Expect [][]int
	}{
		{
			"n = 4 and k = 2",
			4,
			2,
			[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, combine(test.N, test.K))
			assert.Equal(t, test.Expect, combineV2(test.N, test.K))
		})
	}
}

//=== RUN   TestCombine/n_=_4_and_k_=_2
//Before: v=1, nums=[1 2 3 4], track=[]
//Before: v=2, nums=[2 3 4], track=[1]
//Result: res=[[1 2]]
//After: v=2, nums=[2 3 4], track=[1]
//Before: v=3, nums=[2 3 4], track=[1]
//Result: res=[[1 2] [1 3]]
//After: v=3, nums=[2 3 4], track=[1]
//Before: v=4, nums=[2 3 4], track=[1]
//Result: res=[[1 2] [1 3] [1 4]]
//After: v=4, nums=[2 3 4], track=[1]
//After: v=1, nums=[1 2 3 4], track=[]
//Before: v=2, nums=[1 2 3 4], track=[]
//Before: v=3, nums=[3 4], track=[2]
//Result: res=[[1 2] [1 3] [1 4] [2 3]]
//After: v=3, nums=[3 4], track=[2]
//Before: v=4, nums=[3 4], track=[2]
//Result: res=[[1 2] [1 3] [1 4] [2 3] [2 4]]
//After: v=4, nums=[3 4], track=[2]
//After: v=2, nums=[1 2 3 4], track=[]
//Before: v=3, nums=[1 2 3 4], track=[]
//Before: v=4, nums=[4], track=[3]
//Result: res=[[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
//After: v=4, nums=[4], track=[3]
//After: v=3, nums=[1 2 3 4], track=[]
//Before: v=4, nums=[1 2 3 4], track=[]
//After: v=4, nums=[1 2 3 4], track=[]
//--- PASS: TestCombine/n_=_4_and_k_=_2 (0.00s)