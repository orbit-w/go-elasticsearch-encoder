package encoder

import (
	"reflect"
	"unicode"
)

// SnakeCase 将驼峰式写法转换成下划线写法
func SnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// IsZeroR 用反射零值判断，在性能敏感的代码中最好用 IsZero
func IsZeroR(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

// IsZero 零值判断
func IsZero(value any) bool {
	switch v := value.(type) {
	case string:
		if v == "" {
			return true
		}
	case int:
		return v == 0
	case uint:
		return v == uint(0)
	case int64:
		return v == int64(0)
	case uint64:
		return v == uint64(0)
	case uint32:
		return v == uint32(0)
	case int32:
		return v == int32(0)
	case uint8:
		return v == uint8(0)
	case int8:
		return v == int8(0)
	case uint16:
		return v == uint16(0)
	case int16:
		return v == int16(0)
	case float32:
		return v == float32(0)
	case float64:
		return v == float64(0)
	case bool:
		return !v
	}
	return false
}
