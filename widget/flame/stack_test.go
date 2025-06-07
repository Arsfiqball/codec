package flame_test

import (
	"errors"
	"testing"

	"github.com/Arsfiqball/codec/widget/flame"
)

func TestStackFrom(t *testing.T) {
	errSample1 := flame.New("test 1")
	errSample2 := flame.New("test 2")
	errSample3 := flame.New("test 3")

	t.Run("Success extracting stacks", func(t *testing.T) {
		err := errors.New("sentinel error")
		err1 := errSample1.Wrap(err)
		err2 := errSample2.Wrap(err1)
		err3 := errSample3.Wrap(err2)

		stack := flame.StackFrom(err3, 10)

		if stack == "" {
			t.Error("stack is empty")
		}

		// fmt.Println(stack)
	})
}
