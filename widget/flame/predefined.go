package flame

const (
	CodeBadRequest      = "bad request"
	CodeUnauthorized    = "unauthorized"
	CodeForbidden       = "forbidden"
	CodeNotFound        = "not found"
	CodeConflict        = "conflict"
	CodeRequestTooLarge = "request too large"
	CodeUnprocessable   = "unprocessable"
	CodeLocked          = "locked"
	CodeTooManyRequests = "too many requests"
)

var (
	BadRequest      = New(CodeBadRequest)
	Unauthorized    = New(CodeUnauthorized)
	Forbidden       = New(CodeForbidden)
	NotFound        = New(CodeNotFound)
	Conflict        = New(CodeConflict)
	RequestTooLarge = New(CodeRequestTooLarge)
	Unprocessable   = New(CodeUnprocessable)
	Locked          = New(CodeLocked)
	TooManyRequests = New(CodeTooManyRequests)
)

type HttpUnpacked struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func HttpUnpack(err error) HttpUnpacked {
	unpacked := Unpack(err)

	return HttpUnpacked{
		Code:    httpCodeOf(unpacked.Code),
		Message: unpacked.Info,
		Data:    unpacked.Data,
	}
}

func httpCodeOf(code string) int {
	switch code {
	case CodeBadRequest:
		return 400
	case CodeUnauthorized:
		return 401
	case CodeForbidden:
		return 403
	case CodeNotFound:
		return 404
	case CodeConflict:
		return 409
	case CodeRequestTooLarge:
		return 413
	case CodeUnprocessable:
		return 422
	case CodeLocked:
		return 423
	case CodeTooManyRequests:
		return 429
	default:
		return 500
	}
}
