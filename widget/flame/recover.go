package flame

import (
	"fmt"
	"runtime"
)

const PanicCode = "panic"

var Panic = New(PanicCode)

func RecoverAs(out *Error, depth int) {
	if out == nil {
		return
	}

	const skip = 2
	if r := recover(); r != nil {
		err := *out // Copy the original flame error
		err.info = fmt.Sprintf("%v", r)

		for i := skip; i < depth; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}

			name := "unknown function"

			fn := runtime.FuncForPC(pc)
			if fn != nil {
				name = fn.Name()
			}

			childErr := Error{
				code:   PanicCode,
				info:   fmt.Sprintf("stack %d: %s", i-skip, name),
				caller: fmt.Sprintf("%s:%d", file, line),
				parent: err,
			}

			err = childErr
		}

		*out = err
	}
}
