package marshaller

import (
	"reflect"
)

type Field struct {
	name   string
	est    string
	esName string
	value  any //field 值
	op     tagOption
}

// MarshalQueryBool 将结构体解析成elasticsearch query bool 语句
func MarshalQueryBool(v any) (*BoolQuery, error) {
	val := reflect.ValueOf(v)
	return marshalBool(val)
}

func marshalBool(val reflect.Value) (*BoolQuery, error) {
	if err := kindErr(val); err != nil {
		return nil, err
	}

	var (
		sql = NewBoolQuery()
	)
	if err := parseFields(val, func(v reflect.Value, f *Field) error {
		switch f.est {
		case esBool:
			if s, err := marshalBool(v); err != nil && s != nil {
				sql.AppendMust(s)
			}
		case esRange:
			results, err := marshalSub(v, f)
			if err != nil {
				return err
			}
			sql.AppendMust(results...)
		case esTerm, esTerms:
			re, err := marshalTerm(f)
			if err != nil {
				return err
			}
			if !re.Invalid() {
				sql.AppendMust(re)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return sql, nil
}

func marshalSub(val reflect.Value, cur *Field) ([]any, error) {
	var (
		result = make([]any, 0, 1<<3)
		sql    any
		parser func(field *Field) error
	)
	fact, _ := getFactory(cur.est)
	parser, sql = fact.Create(cur)
	if err := parseFields(val, func(v reflect.Value, f *Field) error {
		switch f.est {
		case esBool:
			s, err := marshalBool(v)
			if err != nil {
				return err
			}
			if s != nil {
				result = append(result, s)
			}
		case esRange:
			list, err := marshalSub(v, f)
			if err != nil {
				return err
			}
			result = append(result, list...)
		case esTerm, esTerms:
			re, err := marshalTerm(f)
			if err != nil {
				return err
			}
			if !re.Invalid() {
				result = append(result, re)
			}
		default:
			if err := parser(f); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if sql != nil {
		result = append(result, sql)
	}
	return result, nil
}

func marshalTerm(cur *Field) (IQuery, error) {
	fact, _ := getFactory(cur.est)
	parser, sql := fact.Create(cur)

	if err := parser(cur); err != nil {
		return nil, err
	}
	return sql.(IQuery), nil
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

		if err := kindErrField(field); err != nil {
			continue
		}

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

		if err = kindErrTag(est, field); err != nil {
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

func kindErr(val reflect.Value) error {
	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return ErrNil
		}
		return nil
	case reflect.Struct:
		return nil
	default:
		return ErrKind
	}
}

func kindErrField(val reflect.Value) error {
	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return ErrNil
		}
		return nil
	default:
		return nil
	}
}

func kindErrTag(est string, val reflect.Value) error {
	switch est {
	case esRange:
		switch val.Kind() {
		case reflect.Ptr:
			return nil
		default:
			return ErrKind
		}
	default:
		switch val.Kind() {
		case reflect.Ptr, reflect.Struct:
			return ErrKind
		default:
			return nil
		}
	}
}
