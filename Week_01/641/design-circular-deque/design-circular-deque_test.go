package design_circular_deque

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMyCircularDeque(t *testing.T) {
	circularDeque := Constructor(3)
	assert.Equal(t, true, circularDeque.InsertLast(1))   // return true
	assert.Equal(t, true, circularDeque.InsertLast(2))   // return true
	assert.Equal(t, true, circularDeque.InsertFront(3))  // return true
	assert.Equal(t, false, circularDeque.InsertFront(4)) // return false, the queue is full
	assert.Equal(t, 2, circularDeque.GetRear())          // return 2
	assert.Equal(t, true, circularDeque.IsFull())        // return true
	assert.Equal(t, true, circularDeque.DeleteLast())    // return true
	assert.Equal(t, true, circularDeque.InsertFront(4))  // return true
	assert.Equal(t, 4, circularDeque.GetFront())         // return 4
}
