package encoder

const (
	esTag     = "es"
	ignoreTag = "-"

	esBool = "bool"

	esRange = "range"
	esTerm  = "term"
	esTerms = "terms"
)

type IQuery interface {
	Json() map[string]any
}
