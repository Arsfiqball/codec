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
		out := func() (err error) {
			defer flame.RecoverAs(&err, 10)

			someProxyFunc()
			return nil
		}()

		stack := flame.StackFrom(out, 10)

		if stack == "" {
			t.Error("stack is empty")
		}

		// fmt.Println(stack)
	})

	t.Run("No panic should return nil error", func(t *testing.T) {
		out := func() (err error) {
			defer flame.RecoverAs(&err, 10)
			// No panic here
			return nil
		}()

		if out != nil {
			t.Errorf("Expected nil error but got %v", out)
		}
	})

	t.Run("Recover direct panic", func(t *testing.T) {
		out := func() (err error) {
			defer flame.RecoverAs(&err, 10)
			panic("direct panic")
		}()

		if out == nil {
			t.Error("Expected error but got nil")
		}

		stack := flame.StackFrom(out, 10)
		if stack == "" {
			t.Error("stack is empty")
		}
	})

	t.Run("Recover with deeper call stack", func(t *testing.T) {
		deepFunc3 := func() {
			panic("deep panic")
		}

		deepFunc2 := func() {
			deepFunc3()
		}

		deepFunc1 := func() {
			deepFunc2()
		}

		out := func() (err error) {
			defer flame.RecoverAs(&err, 10)
			deepFunc1()
			return nil
		}()

		if out == nil {
			t.Error("Expected error but got nil")
		}

		stack := flame.StackFrom(out, 10)
		if stack == "" {
			t.Error("stack is empty")
		}
	})
}
