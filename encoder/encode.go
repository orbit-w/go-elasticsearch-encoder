package encoder

import (
	"fmt"
	"reflect"
	"strings"
)

// Marshal 将结构体解析成elasticsearch query 语句
func Marshal(v any) (map[string]any, error) {
	val := reflect.ValueOf(v)
	query, err := marshal(val, &Field{est: esBool})
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"query": query.Json(),
	}, nil
}

func marshal(val reflect.Value, info *Field) (IQuery, error) {
	switch info.est {
	case ignoreTag:
		return nil, nil
	case esBool:
		//必须是指针类型
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				if info.op.omitempty {
					return nil, nil
				}
				return nil, ErrKind
			}
		} else {
			return nil, ErrKind
		}

		bq := NewBool()
		if err := parseFields(val, func(v reflect.Value, f *Field) error {
			s, err := marshal(v, f)
			if err != nil {
				return err
			}
			if s != nil {
				bq.AppendMust(s)
			}
			return nil
		}); err != nil {
			return nil, err
		}
		return bq, nil
	case esRange:
		return marshalRange(val, info)
	case esTerm:
		return marshalTerm(info)
	case esTerms:
		return marshalTerms(info)
	default:
		return nil, nil
	}
}

func marshalTerm(f *Field) (IQuery, error) {
	tq := new(TermQuery)
	if f.op.omitempty && IsZero(f.value) {
		return nil, nil
	}
	tq.Name(f.esName)
	tq.value = f.value
	return tq, nil
}

func marshalTerms(f *Field) (IQuery, error) {
	switch t := f.value.(type) {
	case string:
		if t == "" {
			return nil, nil
		}
		parts := strings.Split(t, ",")
		if len(parts) == 0 {
			return nil, fmt.Errorf("invalid terms value: %s", f.esName)
		}

		query := new(TermsQuery)
		query.Name(f.esName)
		query.Value(parts...)
		if query.Length() == 0 {
			return nil, fmt.Errorf("invalid terms value: %s", f.esName)
		}
		return query, nil
	default:
		return nil, fmt.Errorf("invalid terms value kind: %s", f.esName)
	}
}

func marshalRange(val reflect.Value, info *Field) (IQuery, error) {
	//必须是指针类型
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			if info.op.omitempty {
				return nil, nil
			}
			return nil, ErrKind
		}
	} else {
		return nil, ErrKind
	}

	query := new(RangeTerm)
	query.Name(info.esName)
	if err := parseFields(val, func(v reflect.Value, f *Field) error {
		query.Range(f.est, f.value)
		return nil
	}); err != nil {
		return nil, err
	}

	if query.Length() == 0 {
		return nil, nil
	}
	return query, nil
}
