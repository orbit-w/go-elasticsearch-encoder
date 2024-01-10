package marshaller

const (
	esTag     = "es"
	ignoreTag = "-"

	esBool = "bool"
)

type IQuery interface {
	Invalid() bool
}
