package flame_test

import (
	"testing"

	"github.com/Arsfiqball/codec/flame"
)

func funcThatPanics() {
	panic("test")
}

type someProxyInterface interface {
	SomeProxyMethod()
}

type someProxyStruct struct{}

func newSomeProxyInterface() someProxyInterface {
	return &someProxyStruct{}
}

func (s *someProxyStruct) SomeProxyMethod() {
	funcThatPanics()
}

func someProxyFunc() {
	newSomeProxyInterface().SomeProxyMethod()
}

func TestRecover(t *testing.T) {
	t.Run("Success recover as flame.Panic", func(t *testing.T) {
		err := flame.Panic

		func() {
			defer flame.RecoverAs(&err, 10)

			someProxyFunc()
		}()

		stack := flame.StackFrom(err, 10)

		if stack == "" {
			t.Error("stack is empty")
		}

		// fmt.Println(stack)
	})
}
