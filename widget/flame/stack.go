package flame

import (
	"fmt"
)

func StackFrom(err error, depth int) string {
	var stack string

	for err != nil && depth > 0 {
		if e, ok := err.(Error); ok {
			stack += fmt.Sprintf("%s\n", e.info)

			if e.caller != "" {
				stack += fmt.Sprintf("  %s\n", e.caller)
			}
		} else {
			stack += fmt.Sprintf("%s\n", err.Error())
		}

		if e, ok := err.(unwrapper); ok {
			err = e.Unwrap()
		} else {
			err = nil
		}

		depth--
	}

	return stack
}
