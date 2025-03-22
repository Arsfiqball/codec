package flame_test

import (
	"errors"
	"feature/widget/flame"
	"testing"
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
}
