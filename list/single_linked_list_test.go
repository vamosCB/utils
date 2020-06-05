package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Empty(t *testing.T) {
	list := NewSingleLinkedList()
	assert.Equal(t, true, list.Empty())
}

func Test_Len(t *testing.T) {
	list := NewSingleLinkedList()
	list.Append(1)
	list.Append(2)
	assert.Equal(t, 2, list.Len())
}

func Test_Append(t *testing.T) {
	list := &SingleLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	assert.Equal(t, 4, list.Length)
}
