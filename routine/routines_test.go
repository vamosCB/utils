package routine

import (
	"fmt"
	"testing"
)

func panicFunc() (interface{}, error) {
	testData := []int{1, 2, 3}
	return testData[len(testData)], nil
}
func handlePanic(err interface{}) {
	fmt.Printf("panic is %+v\n", err)
}
func Test_SaftyRun(t *testing.T) {
	SetPanicHandler(handlePanic)
	SaftyRun(panicFunc)
}
