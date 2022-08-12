package str

import (
	"strconv"
	"strings"
)

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func IntSliceToString(intSlice []int, sep string) string {
	return numericSliceToString(intSlice, sep, strconv.Itoa)
}

func Int64SliceToString(int64Slice []int64, sep string) string {
	return numericSliceToString(int64Slice, sep, Int64ToString)
}

type numeric interface {
	int | int64
}

func numericSliceToString[T numeric](arr []T, sep string, f func(T) string) string {
	stringSlice := make([]string, len(arr))
	for i, v := range arr {
		stringSlice[i] = f(v)
	}

	return strings.Join(stringSlice, sep)
}

func StringToIntSlice(str string, sep string) ([]int, error) {
	return stringToNumericSlice(str, sep, strconv.Atoi)
}

func StringToInt64Slice(str string, sep string) ([]int64, error) {
	return stringToNumericSlice(str, sep, StringToInt64)
}

func stringToNumericSlice[T numeric](str string, sep string, f func(string) (T, error)) ([]T, error) {
	var target []T

	if strings.Trim(str, BlankString) == "" {
		return target, nil
	}

	strSlice := strings.Split(str, sep)
	for _, v := range strSlice {
		val, err := f(v)
		if err != nil {
			return nil, err
		}
		target = append(target, val)
	}

	return target, nil
}

const BlankString = " "

func IsBlank(str string) bool {
	return strings.Trim(str, BlankString) == ""
}

// TruncateString 截断字符串，支持中文字符截断
func TruncateString(str string, maxLen int) string {
	runeStr := []rune(str)

	if len(runeStr) <= maxLen {
		return str
	}
	return string(runeStr[:maxLen])
}
