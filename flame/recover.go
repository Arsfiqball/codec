package flame

import (
	"fmt"
	"runtime"
)

const CodePanic = "panic"

func RecoverAs(out *error, depth int) {
	const skip = 2
	if r := recover(); r != nil {
		var err error

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
				code:   CodePanic,
				info:   fmt.Sprintf("stack %d: %s", i-skip, name),
				caller: fmt.Sprintf("%s:%d", file, line),
				parent: err,
			}

			err = childErr
		}

		*out = err
	}
}
