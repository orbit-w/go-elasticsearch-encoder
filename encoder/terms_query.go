package encoder

type TermsQuery struct {
	name  string
	value []string
}

const (
	esTerms = "terms"
)

func (re *TermsQuery) Name(name string) {
	re.name = name
}

func (re *TermsQuery) Value(values ...string) {
	re.value = append(values, values...)
}

func (re *TermsQuery) Length() int {
	return len(re.value)
}

func (re *TermsQuery) Json() map[string]any {
	// {"terms":{"name":["value1","value2"]}}
	return map[string]any{
		esTerms: map[string]any{
			re.name: re.value,
		},
	}
}
