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

type Terms map[string]map[string][]any

func (t Terms) Invalid() bool {
	if v, ok := t[esTerms]; !ok || v == nil {
		return true
	}
	return false
}

func (re *TermsEncoder) Encode(query any, cur, field *Field) error {
	switch field.est {
	case esTerms:
		switch t := field.value.(type) {
		case string:
			if t == "" {
				return nil
			}
			parts := strings.Split(t, ",")
			if len(parts) == 0 {
				return fmt.Errorf("invalid terms valud: %s", cur.esName)
			}

			m := query.(Terms)
			if v := m[esTerms]; v == nil {
				m[esTerms] = map[string][]any{}
			}
			for i := range parts {
				p := parts[i]
				m[esTerms][field.esName] = append(m[esTerms][field.esName], p)
			}
		}
	}
	return nil
}

func (re *TermsEncoder) Sql(_ *Field) any {
	query := Terms{}
	query[esTerms] = nil
	return query
}
