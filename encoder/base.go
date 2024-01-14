package encoder

const (
	esTag     = "es"
	ignoreTag = "-"

	esBool = "bool"
)

type IQuery interface {
	Json() map[string]any
}
