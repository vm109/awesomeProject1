package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Add(t *testing.T) {
	arr := make([]int, 0)
	stack := &Stack{
		arr,
	}

	stack.Add(1)
	val := stack.Peek()
	assert.Equal(t, val, 1)
	assert.Equal(t, stack.Length(), 1)
}

func TestStack_Get(t *testing.T) {
	stack := Stack{}

	stack.Add(1)
	stack.Add(2)
	stack.Add(3)
	stack.Add(4)
	stack.Add(5)

	assert.Equal(t, stack.Get(), 5)
	assert.Equal(t, stack.Peek(), 4)
	assert.Equal(t, stack.Length(), 4)

	assert.Equal(t, stack.Get(), 4)
	assert.Equal(t, stack.Get(), 3)
	assert.Equal(t, stack.Get(), 2)
	assert.Equal(t, stack.Get(), 1)

	assert.Equal(t, stack.Get(), -1)
}
