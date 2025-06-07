package fiberhandler

import (
	"errors"

	"github.com/Arsfiqball/codec/internal/application/resource"

	"github.com/gofiber/fiber/v2"
)

const (
	codeSuccess           = "SUCCESS"
	codeErrBadRequest     = "ERR_BAD_REQUEST"
	codeErrInvalidEntity  = "ERR_INVALID_ENTITY"
	codeErrInvalidOpsType = "ERR_INVALID_OPS_TYPE"
	codeErrInvalidQuery   = "ERR_INVALID_QUERY"
	codeErrNotFound       = "ERR_NOT_FOUND"
	codeErrUnauthorized   = "ERR_UNAUTHORIZED"
	codeErrUnknown        = "ERR_UNKNOWN"
)

var codeForResource = map[int]string{
	resource.ErrCodeInvalidEntity:  codeErrInvalidEntity,
	resource.ErrCodeInvalidOpsType: codeErrInvalidOpsType,
	resource.ErrCodeInvalidQuery:   codeErrInvalidQuery,
	resource.ErrCodeNotFound:       codeErrNotFound,
	resource.ErrCodeUnauthorized:   codeErrUnauthorized,
	resource.ErrCodeUnknown:        codeErrUnknown,
}

var httpStatusForCode = map[string]int{
	codeSuccess:           fiber.StatusOK,
	codeErrInvalidEntity:  fiber.StatusUnprocessableEntity,
	codeErrInvalidOpsType: fiber.StatusBadRequest,
	codeErrInvalidQuery:   fiber.StatusBadRequest,
	codeErrNotFound:       fiber.StatusNotFound,
	codeErrUnauthorized:   fiber.StatusUnauthorized,
	codeErrUnknown:        fiber.StatusInternalServerError,
}

type FormatResponse struct {
	TraceID   string      `json:"traceId"`
	RequestID string      `json:"requestId"`
	Code      string      `json:"code"`
	Info      string      `json:"info"`
	Data      interface{} `json:"data,omitempty"`
}

func FormatSuccess(c *fiber.Ctx, data interface{}) error {
	return c.Status(httpStatusForCode[codeSuccess]).JSON(FormatResponse{
		TraceID:   c.Locals("traceId").(string),
		RequestID: c.Locals("requestId").(string),
		Code:      codeSuccess,
		Info:      "success",
		Data:      data,
	})
}

func FormatError(c *fiber.Ctx, err error) error {
	var (
		resourceErr *resource.Error
		code        = codeErrUnknown
		status      = httpStatusForCode[codeErrUnknown]
		traceId     = c.Locals("traceId").(string)
		requestId   = c.Locals("requestId").(string)
	)

	if errors.As(err, &resourceErr) {
		code = codeForResource[resourceErr.Code]
		status = httpStatusForCode[code]

		return c.Status(status).JSON(FormatResponse{
			TraceID:   traceId,
			RequestID: requestId,
			Code:      code,
			Info:      resourceErr.Message,
		})
	}

	return c.Status(status).JSON(FormatResponse{
		TraceID:   traceId,
		RequestID: requestId,
		Code:      code,
		Info:      "unknown error occurred",
	})
}

func FormatErrorWithCode(c *fiber.Ctx, err error, code string) error {
	return c.Status(httpStatusForCode[code]).JSON(FormatResponse{
		TraceID:   c.Locals("traceId").(string),
		RequestID: c.Locals("requestId").(string),
		Code:      code,
		Info:      err.Error(),
	})
}
