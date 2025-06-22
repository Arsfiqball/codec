package access_test

import (
	"os"
	"path"
	"testing"

	"github.com/Arsfiqball/codec/access"
)

func TestGenerate(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	rootDir := path.Join(cwd, ".outputs")

	// Clean up the generated files
	if err := os.RemoveAll(rootDir); err != nil {
		t.Fatalf("failed to clean up generated files: %v", err)
	}

	// Test that it only allows alphanumeric characters and slashes
	a := access.Access{
		Name:    "some invalid name!",
		RootDir: rootDir,
	}

	if err := a.Generate(); err == nil {
		t.Fatal("expected an error for invalid name, got nil")
	}

	type scenarioT struct {
		name string
	}

	scenarios := []scenarioT{
		{name: "book"},
		{name: "book/internal/model/chapter"},
		{name: "book/internal/category"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			a := access.Access{
				Name:    scenario.name,
				RootDir: rootDir,
			}

			if err := a.Generate(); err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}
