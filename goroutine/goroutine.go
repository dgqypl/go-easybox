package goroutine

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

var goroutineSpace = []byte("goroutine ")

// CurGoroutineID 摘自："跟踪函数调用链，理解代码更直观"一文（https://time.geekbang.org/column/article/471138）
func CurGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}
