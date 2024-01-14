package encoder

type TermQuery struct {
	name  string
	value any
	boost float64
}

func (tq *TermQuery) Name(name string) {
	tq.name = name
}

func (tq *TermQuery) Value(v any) {
	tq.value = v
}

func (tq *TermQuery) Boost(boost float64) {
	tq.boost = boost
}

func (tq *TermQuery) Json() map[string]any {
	// {"term":{"name":"value"}}
	return map[string]any{
		esTerm: map[string]any{
			tq.name: map[string]any{
				"value": tq.value,
				"boost": tq.boost,
			},
		},
	}
}

const (
	esTerm = "term"
)
