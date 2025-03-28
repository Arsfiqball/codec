package lord_test

import (
	"context"
	"feature/widget/lord"
	"testing"
)

func plusOne(ctx context.Context, v int) (int, error) {
	return v + 1, nil
}

func double(ctx context.Context, v int) (int, error) {
	return v * 2, nil
}

func TestPipeline(t *testing.T) {
	// Create a pipeline that will increment a number by 1 and then double it.
	pipeline := lord.Pipeline(plusOne, double)

	// Run the pipeline with the initial value of 1.
	v, err := pipeline(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Ensure the result is 4.
	if v != 4 {
		t.Fatalf("unexpected value: %v", v)
	}
}
