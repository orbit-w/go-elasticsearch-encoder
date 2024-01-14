package encoder

type Query struct {
	BoolQuery *BoolQuery `json:"query"`
	Sort      map[string]any
}

// BoolQuery represents a bool query.
type BoolQuery struct {
	Must    []IQuery `json:"must,omitempty"`
	Filter  []IQuery `json:"filter,omitempty"`
	Should  []IQuery `json:"should,omitempty"`
	MustNot []IQuery `json:"must_not,omitempty"`
}

func NewBool() *BoolQuery {
	return new(BoolQuery)
}

func (b *BoolQuery) Json() map[string]any {
	c := make(map[string]any)
	if v := b.parse(b.Must); v != nil {
		c["must"] = v
	}

	if v := b.parse(b.Filter); v != nil {
		c["filter"] = v
	}

	if v := b.parse(b.Should); v != nil {
		c["should"] = v
	}

	if v := b.parse(b.MustNot); v != nil {
		c["must_not"] = v
	}
	return map[string]any{
		esBool: c,
	}
}

func (b *BoolQuery) parse(queries []IQuery) []map[string]any {
	if len(queries) == 0 {
		return nil
	}
	v := make([]map[string]any, 0, len(queries))
	for i := range queries {
		query := queries[i]
		v = append(v, query.Json())
	}
	return v
}

func (b *BoolQuery) AppendMust(v IQuery) {
	if b.Must == nil {
		b.Must = make([]IQuery, 0, 1<<3)
	}
	if v != nil {
		b.Must = append(b.Must, v)
	}
}

func (b *BoolQuery) AppendShould(v IQuery) {
	if b.Should == nil {
		b.Should = make([]IQuery, 0, 1<<3)
	}
	if v != nil {
		b.Should = append(b.Should, v)
	}
}

func (b *BoolQuery) AppendFilter(v IQuery) {
	if b.Filter == nil {
		b.Filter = make([]IQuery, 0, 1<<3)
	}
	if v != nil {
		b.Filter = append(b.Filter, v)
	}
}
