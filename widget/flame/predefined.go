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
	badRequest      = New(CodeBadRequest)
	unauthorized    = New(CodeUnauthorized)
	forbidden       = New(CodeForbidden)
	notFound        = New(CodeNotFound)
	conflict        = New(CodeConflict)
	requestTooLarge = New(CodeRequestTooLarge)
	unprocessable   = New(CodeUnprocessable)
	locked          = New(CodeLocked)
	tooManyRequests = New(CodeTooManyRequests)
)

func BadRequest() Error {
	return badRequest.HereBefore(1)
}

func Unauthorized() Error {
	return unauthorized.HereBefore(1)
}

func Forbidden() Error {
	return forbidden.HereBefore(1)
}

func NotFound() Error {
	return notFound.HereBefore(1)
}

func Conflict() Error {
	return conflict.HereBefore(1)
}

func RequestTooLarge() Error {
	return requestTooLarge.HereBefore(1)
}

func Unprocessable() Error {
	return unprocessable.HereBefore(1)
}

func Locked() Error {
	return locked.HereBefore(1)
}

func TooManyRequests() Error {
	return tooManyRequests.HereBefore(1)
}

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
