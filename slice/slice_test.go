package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveTargetFromIntSlice(t *testing.T) {
	src := []int{1, 2, 3, 4}
	result := RemoveTargetFromIntSlice(src, 2)
	assert.Equal(t, 2, len(result))
}
