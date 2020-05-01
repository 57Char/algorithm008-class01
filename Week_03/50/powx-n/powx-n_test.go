package powx_n

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMyPow(t *testing.T) {
	cases := []struct {
		Name   string
		X      float64
		N      int
		Expect float64
	}{
		{
			"pow(0,0)=1",
			0,
			0,
			1,
		},
		{
			"pow(0,1)=0",
			0,
			1,
			0,
		},
		{
			"pow(10,0)=1",
			10,
			0,
			1,
		},
		{
			"pow(10,1)=10",
			10,
			1,
			10,
		},
		{
			"pow(2,10)=1024",
			2,
			10,
			1024,
		},
		{
			"pow(2,9)=512",
			2,
			9,
			512,
		},
		{
			"pow(2,-1)=0.5",
			2,
			-1,
			0.5,
		},
		{
			"pow(2,-2)=0.25",
			2,
			-2,
			0.25,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expect, myPow(test.X, test.N))
			assert.Equal(t, test.Expect, myPowV2(test.X, test.N))
		})
	}
}
