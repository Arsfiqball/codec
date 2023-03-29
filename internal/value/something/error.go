package something

import "errors"

var (
	ErrIDRequired   = errors.New("id is required")
	ErrNameRequired = errors.New("name is required")
)
