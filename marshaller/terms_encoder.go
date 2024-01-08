package marshaller

import (
	"fmt"
	"strings"
)

type TermsEncoder struct {
}

const (
	esTerms = "terms"
)

var termsEncoder = new(TermsEncoder)

func init() {
	reg(esTerms, termsEncoder.Encode, termsEncoder.Sql)
}

func (re *TermsEncoder) Encode(query any, cur, field *Field) error {
	switch field.est {
	case esEnumLogic:
	default:
		m := query.(map[string]map[string][]any)
		str := field.value.(string)
		parts := strings.Split(str, ",")
		if len(parts) == 0 {
			return fmt.Errorf("invalid terms valud: %s", cur.esName)
		}

		for i := range parts {
			p := parts[i]
			m[esTerms][cur.esName] = append(m[esTerms][cur.esName], p)
		}
	}
	return nil
}

func (re *TermsEncoder) Sql(field *Field) any {
	return map[string]map[string][]any{
		esTerms: {
			field.esName: []any{},
		},
	}
}
