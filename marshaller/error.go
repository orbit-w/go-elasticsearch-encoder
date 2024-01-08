package marshaller

import "errors"

var (
	ErrNil    = errors.New("is nil")
	ErrSyntax = errors.New("syntax error")
	ErrKind   = errors.New("kind error")
)
