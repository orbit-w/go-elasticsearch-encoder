package marshaller

type RangeEncoder struct {
}

const (
	esRange = "range"
)

var rangeEncoder = new(RangeEncoder)

func init() {
	reg(esRange, rangeEncoder.Encode, rangeEncoder.Sql)
}

func (re *RangeEncoder) Encode(query any, cur, field *Field) error {
	switch field.est {
	case esEnumLogic:
	default:
		m := query.(map[string]map[string]map[string]any)
		m[esRange][cur.esName][field.est] = field.value
	}
	return nil
}

func (re *RangeEncoder) Sql(field *Field) any {
	return map[string]map[string]map[string]any{
		esRange: map[string]map[string]any{
			field.esName: map[string]any{},
		},
	}
}
