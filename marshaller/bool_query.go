package marshaller

type Query struct {
	BoolQuery *BoolQuery `json:"query"`
}

type BoolQuery struct {
	Bool *Bool `json:"bool"`
}

func NewBoolQuery() *BoolQuery {
	return &BoolQuery{
		Bool: new(Bool),
	}
}

func (q *BoolQuery) AppendMust(v ...any) {
	q.Bool.AppendMust(v...)
}

func (q *BoolQuery) AppendShould(v ...any) {
	q.Bool.AppendShould(v...)
}

func (q *BoolQuery) AppendFilter(v ...any) {
	q.Bool.AppendFilter(v...)
}

// Bool represents a bool query.
type Bool struct {
	Must    []any `json:"must,omitempty"`
	Filter  []any `json:"filter,omitempty"`
	Should  []any `json:"should,omitempty"`
	MustNot []any `json:"must_not,omitempty"`
}

func (bq *Bool) AppendMust(list ...any) {
	if bq.Must == nil {
		bq.Must = make([]any, 0, 1<<3)
	}
	for i := range list {
		v := list[i]
		if v != nil {
			bq.Must = append(bq.Must, v)
		}
	}
}

func (bq *Bool) AppendShould(list ...any) {
	if bq.Should == nil {
		bq.Should = make([]any, 0, 1<<3)
	}
	for i := range list {
		v := list[i]
		if v != nil {
			bq.Should = append(bq.Should, v)
		}
	}
}

func (bq *Bool) AppendFilter(list ...any) {
	if bq.Filter == nil {
		bq.Filter = make([]any, 0, 1<<3)
	}
	for i := range list {
		v := list[i]
		if v != nil {
			bq.Filter = append(bq.Filter, v)
		}
	}
}
