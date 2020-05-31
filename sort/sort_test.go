package sort

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_quickSort(t *testing.T) {
	testData := []int{3, 1, 5, 8, 6, 9}
	QuickSort(testData, 0, 5)
	assert.Equal(t, true, reflect.DeepEqual([]int{1, 3, 5, 6, 8, 9}, testData))
}
