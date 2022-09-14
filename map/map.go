package _map

func PutAll(src map[any]any, dest map[any]any) {
	for k, v := range src {
		dest[k] = v
	}
}

// SliceToMap 对自定义「结构体」/「结构体指针」切片，转换成 key 为切片元素的某个属性，value 为切片元素的 map
//
// 比如对于如下结构体：
//
// type User struct {
// 	Id int64
// 	Name string
// }
//
func SliceToMap[T any, k comparable](slice []T, f func(T) k) map[k]T {
	var result = make(map[k]T)

	if len(slice) == 0 {
		return result
	}

	for _, v := range slice {
		result[f(v)] = v
	}

	return result
}
