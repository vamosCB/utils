package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Empty(t *testing.T) {
	list := &SingleLinkedList{}
	assert.Equal(t, true, list.Empty())
}

func Test_Length(t *testing.T) {
	list := &SingleLinkedList{
		head: &Node{
			Data: 1,
			Next: &Node{
				Data: 2,
			},
		},
	}
	assert.Equal(t, 2, list.Length())
}

func Test_Append(t *testing.T) {
	list := &SingleLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	t.Log("aaaaa")
	assert.Equal(t, 4, list.Length())
}
