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

func (re *TermEncoder) Encode(query any, cur, field *Field) error {
	switch field.est {
	case esEnumLogic:
	case "value":
		m := query.(map[string]map[string]map[string]any)
		m[esTerm][field.esName] = map[string]any{
			"value": field.value,
			"boost": 1.0,
		}
	}
	return nil
}

func (re *TermEncoder) Sql(field *Field) any {
	return map[string]map[string]map[string]any{
		esTerm: map[string]map[string]any{},
	}
}
