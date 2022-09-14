## str

更方便操作 string 的函数：
- `Int64ToString(i int64) string`
- `StringToInt64(s string) (int64, error)`
- `IntSliceToString(intSlice []int, sep string) string`
- `Int64SliceToString(int64Slice []int64, sep string) string`
- `StringToIntSlice(str string,sep string) ([]int,error)`
- `StringToInt64Slice(str string,sep string) ([]int64,error)`
- `IsBlank(str string) bool`
- `TruncateString(str string, maxLen int) string`
- `FieldValueJoin[T any](slice []T, sep string, f func(T) string) (string, error)`