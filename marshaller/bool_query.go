package marshaller

type Query struct {
	Bool *BoolQuery `json:"bool"`
}

func NewQuery() *Query {
	return &Query{
		Bool: new(BoolQuery),
	}
}

func (q *Query) AppendMust(v ...any) {
	q.Bool.AppendMust(v...)
}

func (q *Query) AppendShould(v ...any) {
	q.Bool.AppendShould(v...)
}

func (q *Query) AppendFilter(v ...any) {
	q.Bool.AppendFilter(v...)
}

// BoolQuery represents a bool query.
type BoolQuery struct {
	Must    []any `json:"must,omitempty"`
	Filter  []any `json:"filter,omitempty"`
	Should  []any `json:"should,omitempty"`
	MustNot []any `json:"must_not,omitempty"`
}

func (bq *BoolQuery) AppendMust(list ...any) {
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

func (bq *BoolQuery) AppendShould(list ...any) {
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

func (bq *BoolQuery) AppendFilter(list ...any) {
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
