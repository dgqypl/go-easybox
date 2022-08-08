package goroutine

import (
	"fmt"
	"testing"
)

func TestCurGoroutineID(t *testing.T) {
	fmt.Println(CurGoroutineID())
}
