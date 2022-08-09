package numeric

import (
	"math/rand"
	"time"
)

func RandomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func RandomInt64(max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max)
}

// Int64SliceSubtraction 示例：src=[1,2,3,4], minuend=[1,3]，调用完函数返回结果：[2,4]
func Int64SliceSubtraction(src []int64, minuend []int64) []int64 {
	var minuendMap = make(map[int64]struct{})
	for _, v := range minuend {
		minuendMap[v] = struct{}{}
	}

	var result []int64
	for _, v := range src {
		if _, ok := minuendMap[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}
