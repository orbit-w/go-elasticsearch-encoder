package encoder

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	tagOptionOmitempty = "omitempty"
)

type Field struct {
	name   string
	est    string
	esName string
	value  any //field 值
	op     tagOption
}

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
//		omitempty: 基本类型如果为零值, 则忽略，不会自动生成对应的query子语句
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

func parseFields(val reflect.Value, handle func(v reflect.Value, f *Field) error) error {
	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return nil
		}
		// 获取指针指向的元素（结构体）
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return ErrKind
	}

	// 遍历结构体的所有字段
	for j := 0; j < val.NumField(); j++ {
		// 获取字段的值
		field := val.Field(j)
		// 获取字段的类型
		tf := val.Type().Field(j)
		tag := tf.Tag.Get(esTag)
		if tag == "" || tag == ignoreTag {
			continue
		}
		est, esName, err, op := parseFieldTag(tf.Name, tag)
		if err != nil {
			return err
		}
		if err = handle(field, &Field{
			name:   tf.Name,
			value:  field.Interface(),
			est:    est,
			esName: esName,
			op:     op,
		}); err != nil {
			return err
		}
	}
	return nil
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
