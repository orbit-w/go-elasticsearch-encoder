package encoder

import (
	"fmt"
	"github.com/olivere/elastic/v6"
	"reflect"
	"strings"
)

// MarshalBoolQuery 将结构体解析成elasticsearch BoolQuery 语句
func MarshalBoolQuery(v any) (elastic.Query, error) {
	val := reflect.ValueOf(v)
	query, err := marshal(val, &Field{est: esBool})
	if err != nil {
		return nil, err
	}

	return query, nil
}

func marshal(val reflect.Value, info *Field) (elastic.Query, error) {
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

		bq := elastic.NewBoolQuery()
		if err := parseFields(val, func(v reflect.Value, f *Field) error {
			s, err := marshal(v, f)
			if err != nil {
				return err
			}
			if s != nil {
				bq.Must(s)
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

func marshalTerm(f *Field) (elastic.Query, error) {
	if f.op.omitempty && IsZero(f.value) {
		return nil, nil
	}
	return elastic.NewTermQuery(f.esName, f.value), nil
}

func marshalTerms(f *Field) (elastic.Query, error) {
	switch t := f.value.(type) {
	case string:
		if t == "" {
			return nil, nil
		}
		parts := strings.Split(t, ",")
		if len(parts) == 0 {
			return nil, fmt.Errorf("invalid terms value: %s", f.esName)
		}

		values := make([]any, len(parts))
		for i := range parts {
			values[i] = parts[i]
		}

		return elastic.NewTermsQuery(f.esName, values...), nil
	default:
		return nil, fmt.Errorf("invalid terms value kind: %s", f.esName)
	}
}

func marshalRange(val reflect.Value, info *Field) (elastic.Query, error) {
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

	query := elastic.NewRangeQuery(info.esName)

	if err := parseFields(val, func(v reflect.Value, f *Field) error {
		switch f.est {
		case "gte":
			query.Gte(f.value)
		case "gt":
			query.Gt(f.value)
		case "lte":
			query.Lte(f.value)
		case "lt":
			query.Lt(f.value)
		default:
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return query, nil
}
