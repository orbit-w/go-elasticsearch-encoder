package marshaller

type TermEncoder struct {
}

const (
	esTerm = "term"
)

var termEncoder = new(TermEncoder)

func init() {
	reg(esTerm, termEncoder.Encode, termEncoder.Sql)
}

type Term map[string]map[string]map[string]any

func (t Term) Invalid() bool {
	term, ok := t[esTerm]
	if !ok || term == nil {
		return true
	}
	return false
}

func (re *TermEncoder) Encode(query any, _, field *Field) error {
	switch field.est {
	case esTerm:
		if field.op.omitempty && IsZero(field.value) {
			return nil
		}

		m := query.(Term)
		if v := m[esTerm]; v == nil {
			m[esTerm] = map[string]map[string]any{}
		}
		m[esTerm][field.esName] = map[string]any{
			"value": field.value,
			"boost": 1.0,
		}
	}
	return nil
}

func (re *TermEncoder) Sql(_ *Field) any {
	query := Term{}
	query[esTerm] = nil
	return query
}
