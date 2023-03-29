package action

import "errors"

var (
	ErrUserRequired     = errors.New("user is required")
	ErrResourceNotFound = errors.New("resource not found")
)
