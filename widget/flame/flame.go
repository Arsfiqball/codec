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

// Here is a method that sets the caller field of the Error struct to the file and line number of the caller.
func (e Error) Here() Error {
	var caller string

	_, file, line, ok := runtime.Caller(1)

	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	e.caller = caller

	return e
}

// Wrap method that sets the parent field of the Error struct to the error passed in and returns the Error struct.
// It also sets the caller field of the Error struct to the file and line number of the caller.
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

const CodeUnexpected = "unexpected"

// Unexpected create "unexpected" which function as error forwarding.
// In "Unpack" function, it will be skipped and return the first expected error.
// It also sets the caller field of the Error struct to the file and line number of the caller (useful for debugging).
func Unexpected(err error) Error {
	var caller string

	_, file, line, ok := runtime.Caller(1)

	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	return Error{
		code:   CodeUnexpected,
		info:   fmt.Sprintf("error with code %s", CodeUnexpected),
		data:   Data{},
		caller: caller,
	}
}

type Unpacked struct {
	Code string `json:"code"`
	Info string `json:"info"`
	Data Data   `json:"data"`
}

func Unpack(err error) Unpacked {
	var code string
	var info string
	var data Data

	if fe := firstExpected(err); fe != nil {
		err = fe
	}

	if e, ok := err.(Error); ok {
		code = e.Code()
		info = e.Info()
		data = e.Data()
	} else {
		code = CodeUnexpected
		info = err.Error()
	}

	return Unpacked{
		Code: code,
		Info: info,
		Data: data,
	}
}

func firstExpected(err error) error {
	if e, ok := err.(Error); ok && e.Code() == CodeUnexpected {
		return firstExpected(e.Unwrap())
	}

	return err
}
