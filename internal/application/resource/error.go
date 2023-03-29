package resource

import "errors"

var (
	ErrUserRequired     = errors.New("user is required")
	ErrUserNotAdmin     = errors.New("user is not admin")
	ErrUserNotAuthor    = errors.New("user is not author")
	ErrResourceNotFound = errors.New("resource not found")
	ErrInvalidQuery     = errors.New("invalid query")
)
