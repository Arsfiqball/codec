package resource

const (
	ErrCodeUnknown = iota
	ErrCodeNotFound
	ErrCodeUnauthorized
	ErrCodeInvalidQuery
	ErrCodeInvalidEntity
	ErrCodeInvalidOpsType
)

type Error struct {
	Original error
	Message  string
	Code     int
}

func (e Error) Error() string {
	return e.Message
}

func NewError(original error, message string, code int) Error {
	return Error{
		Original: original,
		Message:  message,
		Code:     code,
	}
}
