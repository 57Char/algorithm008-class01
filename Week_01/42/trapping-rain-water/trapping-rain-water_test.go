package trapping_rain_water

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 0, trap([]int{0}))
	assert.Equal(t, 0, trap([]int{0, 1, 0}))
	assert.Equal(t, 1, trap([]int{0, 1, 0, 1}))
	assert.Equal(t, 0, trap([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, trap([]int{1, 2, 3, 4, 5, 4, 3, 2, 1}))
	assert.Equal(t, 10, trap([]int{5, 1, 2, 3, 4, 5}))
	assert.Equal(t, 0, trap([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, 10, trap([]int{5, 4, 3, 2, 1, 5}))
	assert.Equal(t, 20, trap([]int{5, 1, 2, 3, 4, 5, 4, 3, 2, 1, 5}))
}

func TestTrapV2(t *testing.T) {
	assert.Equal(t, 6, trapV2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 0, trapV2([]int{0}))
	assert.Equal(t, 0, trapV2([]int{0, 1, 0}))
	assert.Equal(t, 1, trapV2([]int{0, 1, 0, 1}))
	assert.Equal(t, 0, trapV2([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, trapV2([]int{1, 2, 3, 4, 5, 4, 3, 2, 1}))
	assert.Equal(t, 10, trapV2([]int{5, 1, 2, 3, 4, 5}))
	assert.Equal(t, 0, trapV2([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, 10, trapV2([]int{5, 4, 3, 2, 1, 5}))
	assert.Equal(t, 20, trapV2([]int{5, 1, 2, 3, 4, 5, 4, 3, 2, 1, 5}))
}
