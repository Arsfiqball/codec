package lord_test

import (
	"errors"
	"feature/widget/lord"
	"testing"
)

func TestWithResult(t *testing.T) {
	// Create a collection of numbers.
	collection := lord.Range(10, func(i int) int {
		return i
	})

	// Add 10 to each number in the collection.
	collection = collection.Map(func(v int) int {
		return v + 10
	})

	// Remove all numbers that are not divisible by 2.
	collection = collection.Filter(func(v int) bool {
		return v%2 == 0
	})

	// Ensure the collection contains 5 numbers.
	if len(collection) != 5 {
		t.Fatalf("unexpected length: %v", len(collection))
	}

	// Ensure the collection contains the numbers 10, 12, 14, 16, and 18.
	for i, v := range collection {
		if v != i*2+10 {
			t.Fatalf("unexpected value at index %v: %v", i, v)
		}
	}

	// Ensure the collection can be converted to a result.
	result := lord.Convert(collection, lord.WithResult)

	// Change number 14 to error.
	result = result.Map(func(r lord.Result[int]) lord.Result[int] {
		if r.Value == 14 {
			r.Err = errors.New("error")
		}

		return r
	})

	// Get the first error in the result.
	v, ok := result.Find(func(r lord.Result[int]) bool { return r.Err != nil })
	if !ok {
		t.Fatalf("error not found")
	}

	// Ensure the error is on number 14.
	if v.Value != 14 {
		t.Fatalf("unexpected value: %v", v.Value)
	}
}
