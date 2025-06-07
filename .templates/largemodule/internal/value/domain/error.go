package domain

import "errors"

var (
	ErrNotFound           = errors.New("domain not found")
	ErrMissingAccumulator = errors.New("missing accumulator")
)
