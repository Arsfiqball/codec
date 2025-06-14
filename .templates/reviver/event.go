package reviver

import (
	"context"
)

type Event interface {
	SomethingHappened(context.Context, SomethingHappened) error
}
