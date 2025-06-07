package flame_test

import (
	"errors"
	"testing"

	"github.com/Arsfiqball/codec/flame"
)

func TestError(t *testing.T) {
	errSample1 := flame.New("test 1")
	errSample2 := flame.New("test 2")

	t.Run("Success on notice scenario", func(t *testing.T) {
		err := errors.New("sentinel error")
		err1 := errSample1.Wrap(err).WithInfo("info 1").WithData(flame.Data{"key1": "value"})
		err2 := errSample2.Wrap(err1).WithInfo("info 2").WithData(flame.Data{"key2": "value"})
		err3 := flame.Unexpected(err2)
		unpacked := flame.Unpack(err3)

		if unpacked.Code != "test 2" {
			t.Error("code is not test 2")
		}

		if unpacked.Info != "info 2" {
			t.Error("info is not info 2")
		}

		if unpacked.Data["key1"] != "value" {
			t.Error("data key1 is not value")
		}

		if unpacked.Data["key2"] != "value" {
			t.Error("data key2 is not value")
		}
	})

	t.Run("Success on validation scenario", func(t *testing.T) {
		errEntityLayer := flame.
			New("entity validation").
			WithInfo("validation error on entity layer").
			WithData(flame.Data{
				"key1": errors.New("error 1").Error(),
				"key2": errors.New("error 2").Error(),
			}).
			Here().
			NilOnEmptyData()

		errServiceLayer := flame.
			New("service validation").
			WithInfo("validation error on service layer").
			Wrap(errEntityLayer)

		errUnexpected := flame.Unexpected(errServiceLayer)
		unpacked := flame.Unpack(errUnexpected)

		// fmt.Println("Code:", unpacked.Code)
		// fmt.Println("Info:", unpacked.Info)
		// fmt.Println("Data:", unpacked.Data)

		if unpacked.Code != "service validation" {
			t.Error("code is not service validation")
		}

		if unpacked.Info != "validation error on service layer" {
			t.Error("info is not validation error on service layer")
		}

		if unpacked.Data["key1"] != "error 1" {
			t.Error("data key1 is not error 1")
		}

		if unpacked.Data["key2"] != "error 2" {
			t.Error("data key2 is not error 2")
		}
	})

	t.Run("Success on notice scenario with nil data", func(t *testing.T) {
		nilErr := flame.New("nil error").WithInfo("nil info").NilOnEmptyData()
		unpacked := flame.Unpack(nilErr)

		// fmt.Println("IsEmpty:", unpacked.IsEmpty())

		if !unpacked.IsEmpty() {
			t.Error("is not empty")
		}
	})

	t.Run("Success flame.Unexpected(nil)", func(t *testing.T) {
		err := flame.Unexpected(nil)
		unpacked := flame.Unpack(err)

		// fmt.Println("Code:", unpacked.Code)
		// fmt.Println("Info:", unpacked.Info)
		// fmt.Println("Data:", unpacked.Data)

		if unpacked.Code != "unexpected" {
			t.Error("code is not unexpected")
		}

		if unpacked.Info != "error with code unexpected" {
			t.Error("info is not error with code unexpected")
		}

		if !unpacked.Data.IsEmpty() {
			t.Error("data is not nil")
		}
	})

	t.Run("Success casting to error", func(t *testing.T) {
		err := errors.New("sentinel error")
		err1 := errSample1.Wrap(err).WithInfo("info 1").WithData(flame.Data{"key1": "value"})
		err2 := errSample2.Wrap(err1).WithInfo("info 2").WithData(flame.Data{"key2": "value"})
		err3 := flame.Unexpected(err2)

		casted, ok := errSample1.From(err3)

		// fmt.Println("Casted Code:", casted.Code())
		// fmt.Println("Casted Info:", casted.Info())
		// fmt.Println("Casted Data:", casted.Data())
		// fmt.Println("OK:", ok)

		if !ok {
			t.Error("is not ok")
		}

		if casted.Code() != "test 1" {
			t.Error("code is not test 1")
		}

		if casted.Info() != "info 1" {
			t.Error("info is not info 1")
		}

		if casted.Data()["key1"] != "value" {
			t.Error("data key1 is not value")
		}

		if _, ok := casted.Data()["key2"]; ok {
			t.Error("data key2 is not nil")
		}
	})
}
