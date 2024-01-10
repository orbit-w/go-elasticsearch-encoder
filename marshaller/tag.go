package marshaller

import (
	"fmt"
	"strings"
)

const (
	tagOptionOmitempty = "omitempty"
)

type tagOption struct {
	omitempty bool
}

// 顺序很重要, es tag 用',' 分割，
//
//	第一个是query 类型
//	后续元素是可选参数
//		分析器会以找到的第一个非参数字符作为name
//
//	目前的可选参数：
//		omitempty: 基本类型如果为零值,
func parseFieldTag(fieldName, tag string) (est, esName string, err error, op tagOption) {
	parts := strings.Split(tag, ",")
	if len(parts) == 0 {
		err = fmt.Errorf("invalid es tag: %s", tag)
		return
	}
	est = parts[0]
	esName = parseName(parts, fieldName)
	op = parseTagOption(parts)
	return
}

func parseTagOption(parts []string) (op tagOption) {
	if len(parts) <= 1 {
		return
	}

	for i := 1; i < len(parts); i++ {
		switch parts[i] {
		case tagOptionOmitempty:
			op.omitempty = true
		default:

		}
	}
	return
}

func parseName(parts []string, fieldName string) (esName string) {
	if len(parts) >= 2 {
		for i := 1; i < 1; i++ {
			if !isOption(parts[i]) {
				esName = parts[i]
			}
		}
		if esName == "" {
			esName = SnakeCase(fieldName)
		}
	} else {
		esName = SnakeCase(fieldName)
	}
	return
}

func isOption(str string) bool {
	switch str {
	case tagOptionOmitempty:
		return true
	default:
		return false
	}
}
