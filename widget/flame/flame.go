package flame

import (
	"fmt"
	"runtime"
)

type Error struct {
	code   string
	info   string
	data   Data
	caller string
	parent error
}

var _ error = Error{}

func New(code string) Error {
	var caller string

	_, file, line, ok := runtime.Caller(1)

	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	return Error{
		code:   code,
		info:   fmt.Sprintf("error with code %s", code),
		data:   Data{},
		caller: caller,
	}
}

func (e Error) Code() string {
	return e.code
}

func (e Error) Caller() string {
	return e.caller
}

func (e Error) Wrap(err error) Error {
	var caller string

	_, file, line, ok := runtime.Caller(1)

	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	e.parent = err
	e.caller = caller

	return e
}

func (e Error) WithInfo(message string) Error {
	e.info = message

	return e
}

func (e Error) Info() string {
	return e.info
}

func (e Error) WithData(data Data) Error {
	e.data = data

	return e
}

func (e Error) Data() Data {
	return e.data
}

func (e Error) Error() string {
	return e.info
}

func (e Error) Unwrap() error {
	return e.parent
}

func (e Error) Is(target error) bool {
	if target == nil {
		return false
	}

	err, ok := target.(Error)

	if ok && e.code == err.code {
		return true
	}

	return false
}
