package access

import (
	"embed"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type Access struct {
	Name    string
	RootDir string
}

func (a Access) Filepath(p ...string) string {
	return filepath.Join(append([]string{a.RootDir}, p...)...)
}

type Generator func() error
type Matcher func(acc Access) (bool, Generator)

func (a Access) Generate() error {
	if a.Name == "" {
		return errors.New("access name cannot be empty")
	}

	if strings.TrimSpace(a.Name) != a.Name {
		return errors.New("access name cannot begin or end with whitespace")
	}

	// Check for leading number or number after slash
	parts := strings.SplitSeq(a.Name, "/")
	for part := range parts {
		if part == "" {
			continue
		}
		if len(part) > 0 && (part[0] >= '0' && part[0] <= '9') {
			return errors.New("access name cannot have leading number or number after slash")
		}
	}

	// Only allow alphanumeric, slashes, and spaces
	for _, char := range a.Name {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '/' && char != ' ' {
			return errors.New("access name can only contain alphanumeric characters, spaces, and slashes")
		}
	}

	if a.RootDir == "" {
		rootDir, err := os.Getwd()
		if err != nil {
			return errors.New("failed to get current working directory")
		}

		a.RootDir = rootDir
	}

	matchers := []Matcher{
		isPackagePattern,
		isFeaturePattern,
		isModelPattern,
		isDomainPattern,
	}

	for _, matcher := range matchers {
		if matched, generator := matcher(a); matched {
			return generator()
		}
	}

	return errors.New("invalid access name pattern")
}

func isPackagePattern(acc Access) (bool, Generator) {
	name := acc.Name

	splited := strings.Split(name, "/")
	if len(splited) == 2 && splited[0] != "" && splited[1] != "" && splited[1] != "internal" {
		return true, func() error {
			// Logic to summon a package
			return nil
		}
	}

	return false, nil
}

func makeGoPackageName(s string) string {
	if len(s) == 0 {
		return s
	}

	// Replace spaces with empty string and convert to lowercase
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	// Remove any underscores
	s = strings.ReplaceAll(s, "_", "")

	// Ensure the first character is a letter
	if len(s) > 0 && (s[0] < 'a' || s[0] > 'z') && (s[0] < 'A' || s[0] > 'Z') {
		s = "pkg" + s // Prefix with "pkg" if it doesn't start with a letter
	}

	return s
}

func makePascalCaseOf(s string) string {
	if len(s) == 0 {
		return s
	}

	// Split the string by spaces, underscores, and hyphens
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '_' || r == '-'
	})

	// Capitalize each part
	for i, part := range parts {
		if len(part) > 0 {
			r := []rune(part)
			r[0] = unicode.ToUpper(r[0])
			parts[i] = string(r)
		}
	}

	// Join without separators for PascalCase
	result := strings.Join(parts, "")

	// Ensure the first character is a letter for valid Go identifier
	if len(result) > 0 && !unicode.IsLetter([]rune(result)[0]) {
		result = "Type" + result
	}

	return result
}

func makeReceiverVar(s string) string {
	if len(s) == 0 {
		return s
	}

	// Take only the first character and convert to lowercase
	firstChar := []rune(s)[0]
	return string(unicode.ToLower(firstChar))
}

func dirAlreadyExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false // If stat fails, assume directory does not exist
	}

	return info.IsDir() // Check if the path is a directory
}

func isFeaturePattern(acc Access) (bool, Generator) {
	name := acc.Name

	splited := strings.Split(name, "/")
	if len(splited) == 1 && splited[0] != "" {
		return true, func() error {
			featurePkg := makeGoPackageName(splited[0])
			featureType := makePascalCaseOf(splited[0])

			type dataT struct {
				FeaturePkg  string
				FeatureType string
			}

			data := dataT{
				FeaturePkg:  featurePkg,
				FeatureType: featureType,
			}

			if yes := dirAlreadyExists(acc.Filepath(featurePkg)); yes {
				return fmt.Errorf("feature package %s already exists", featurePkg)
			}

			if err := renderToFile("main.go.tmpl", data, acc.Filepath(featurePkg, "main.go")); err != nil {
				return fmt.Errorf("failed to generate main.go for feature: %w", err)
			}

			if err := renderToFile("wire.go.tmpl", data, acc.Filepath(featurePkg, "wire.go")); err != nil {
				return fmt.Errorf("failed to generate wire.go for feature: %w", err)
			}

			if err := renderToFile("context.md.tmpl", data, acc.Filepath(featurePkg, "context.md")); err != nil {
				return fmt.Errorf("failed to generate context.md for feature: %w", err)
			}

			// Create base model
			genModel := Access{
				Name:    featurePkg + "/internal/model/" + featurePkg,
				RootDir: acc.RootDir,
			}

			if err := genModel.Generate(); err != nil {
				return fmt.Errorf("failed to generate model for feature: %w", err)
			}

			genDomain := Access{
				Name:    featurePkg + "/internal/" + featurePkg,
				RootDir: acc.RootDir,
			}

			if err := genDomain.Generate(); err != nil {
				return fmt.Errorf("failed to generate domain for feature: %w", err)
			}

			return nil
		}
	}

	return false, nil
}

func makeModelFilenameNoExt(s string) string {
	if len(s) == 0 {
		return s
	}

	// Replace spaces with empty string and convert to lowercase
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	// Remove any underscores
	s = strings.ReplaceAll(s, "_", "")

	// Ensure the first character is a letter
	if len(s) > 0 && (s[0] < 'a' || s[0] > 'z') && (s[0] < 'A' || s[0] > 'Z') {
		s = "tbl_" + s // Prefix with "tbl_" if it doesn't start with a letter
	}

	return s
}

func makeModelTableName(s string) string {
	if len(s) == 0 {
		return s
	}

	// Replace spaces with empty string and convert to lowercase
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	// Remove any underscores
	s = strings.ReplaceAll(s, "_", "")

	// Special plural cases
	specialPlurals := map[string]string{
		"person":   "people",
		"man":      "men",
		"woman":    "women",
		"child":    "children",
		"foot":     "feet",
		"tooth":    "teeth",
		"goose":    "geese",
		"mouse":    "mice",
		"ox":       "oxen",
		"leaf":     "leaves",
		"wolf":     "wolves",
		"knife":    "knives",
		"life":     "lives",
		"wife":     "wives",
		"shelf":    "shelves",
		"datum":    "data",
		"medium":   "media",
		"analysis": "analyses",
		"crisis":   "crises",
		"basis":    "bases",
	}

	// Check if word is already plural
	if strings.HasSuffix(s, "ies") ||
		strings.HasSuffix(s, "es") ||
		(strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss") && len(s) > 1) ||
		strings.HasSuffix(s, "ves") ||
		strings.HasSuffix(s, "children") ||
		strings.HasSuffix(s, "people") {
		// Already plural, leave as is
	} else if plural, exists := specialPlurals[s]; exists {
		s = plural
	} else if strings.HasSuffix(s, "status") || strings.HasSuffix(s, "species") ||
		strings.HasSuffix(s, "series") || strings.HasSuffix(s, "equipment") ||
		strings.HasSuffix(s, "information") || strings.HasSuffix(s, "rice") ||
		strings.HasSuffix(s, "money") || strings.HasSuffix(s, "deer") ||
		strings.HasSuffix(s, "fish") || strings.HasSuffix(s, "sheep") {
		// These words are the same in singular and plural
	} else if strings.HasSuffix(s, "y") && len(s) > 1 &&
		strings.IndexByte("aeiou", s[len(s)-2]) == -1 {
		// Words ending in consonant + y: change y to ies
		s = s[:len(s)-1] + "ies"
	} else if strings.HasSuffix(s, "s") ||
		strings.HasSuffix(s, "ch") ||
		strings.HasSuffix(s, "sh") ||
		strings.HasSuffix(s, "x") ||
		strings.HasSuffix(s, "z") {
		// Add "es" for words ending with s, ch, sh, x, z
		s = s + "es"
	} else if strings.HasSuffix(s, "f") {
		// Words ending in f: change f to ves
		s = s[:len(s)-1] + "ves"
	} else if strings.HasSuffix(s, "fe") {
		// Words ending in fe: change fe to ves
		s = s[:len(s)-2] + "ves"
	} else if !strings.HasSuffix(s, "s") {
		// Default case: add s
		s = s + "s"
	}

	// Ensure the first character is a letter
	if len(s) > 0 && (s[0] < 'a' || s[0] > 'z') && (s[0] < 'A' || s[0] > 'Z') {
		s = "tbl_" + s // Prefix with "tbl_" if it doesn't start with a letter
	}

	return s
}

func makeModelTableSingular(s string) string {
	if len(s) == 0 {
		return s
	}

	// Replace spaces with empty string and convert to lowercase
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	// Remove any underscores
	s = strings.ReplaceAll(s, "_", "")

	// Special singular forms
	specialSingulars := map[string]string{
		"people":   "person",
		"men":      "man",
		"women":    "woman",
		"children": "child",
		"feet":     "foot",
		"teeth":    "tooth",
		"geese":    "goose",
		"mice":     "mouse",
		"oxen":     "ox",
		"leaves":   "leaf",
		"wolves":   "wolf",
		"knives":   "knife",
		"lives":    "life",
		"wives":    "wife",
		"shelves":  "shelf",
		"data":     "datum",
		"media":    "medium",
		"analyses": "analysis",
		"crises":   "crisis",
		"bases":    "basis",
	}

	// Check if word is already singular or has special form
	if singular, exists := specialSingulars[s]; exists {
		s = singular
	} else if strings.HasSuffix(s, "ies") && len(s) > 3 {
		// Words ending in "ies" -> "y"
		s = s[:len(s)-3] + "y"
	} else if strings.HasSuffix(s, "ves") && len(s) > 3 {
		// Words ending in "ves" -> "f" or "fe"
		s = s[:len(s)-3] + "f"
	} else if strings.HasSuffix(s, "es") && len(s) > 2 {
		// Words ending with "es"
		if strings.HasSuffix(s[:len(s)-2], "sh") ||
			strings.HasSuffix(s[:len(s)-2], "ch") ||
			strings.HasSuffix(s[:len(s)-2], "x") ||
			strings.HasSuffix(s[:len(s)-2], "s") ||
			strings.HasSuffix(s[:len(s)-2], "z") {
			s = s[:len(s)-2]
		} else {
			// Generic "es" -> remove "s"
			s = s[:len(s)-1]
		}
	} else if strings.HasSuffix(s, "s") && len(s) > 1 && !strings.HasSuffix(s, "ss") {
		// Simply remove trailing "s" for regular plurals
		s = s[:len(s)-1]
	}

	// Ensure the first character is a letter
	if len(s) > 0 && (s[0] < 'a' || s[0] > 'z') && (s[0] < 'A' || s[0] > 'Z') {
		s = "tbl_" + s // Prefix with "tbl_" if it doesn't start with a letter
	}

	return s
}

func isModelPattern(acc Access) (bool, Generator) {
	name := acc.Name

	splited := strings.Split(name, "/")
	if len(splited) == 4 && splited[1] == "internal" && splited[2] == "model" {
		return true, func() error {
			featurePkg := makeGoPackageName(splited[0])
			modelFilenameNoExt := makeModelFilenameNoExt(splited[3])
			modelTableName := makeModelTableName(modelFilenameNoExt)
			modelTableSingular := makeModelTableSingular(modelTableName)
			modelType := makePascalCaseOf(splited[3])
			modelReceiver := makeReceiverVar(modelType)

			type dataT struct {
				ModelType          string
				ModelTable         string
				ModelTableSingular string
				ModelReceiver      string
			}

			data := dataT{
				ModelType:          modelType,
				ModelTable:         modelTableName,
				ModelTableSingular: modelTableSingular,
				ModelReceiver:      modelReceiver,
			}

			if err := renderToFile("internal/model/model.go.tmpl", data, acc.Filepath(featurePkg, "internal", "model", modelFilenameNoExt+".go")); err != nil {
				return fmt.Errorf("failed to generate model.go for feature: %w", err)
			}

			return nil
		}
	}

	return false, nil
}

func isDomainPattern(acc Access) (bool, Generator) {
	name := acc.Name

	splited := strings.Split(name, "/")
	if len(splited) == 3 && splited[1] == "internal" && splited[2] != "model" {
		return true, func() error {
			featurePkg := makeGoPackageName(splited[0])
			domainPkg := makeGoPackageName(splited[2])
			domainType := makePascalCaseOf(splited[2])
			domainReceiver := makeReceiverVar(domainType)
			domainFilename := makeModelFilenameNoExt(splited[2])

			type dataT struct {
				FeaturePkg     string
				DomainPkg      string
				DomainType     string
				DomainReceiver string
			}

			data := dataT{
				FeaturePkg:     featurePkg,
				DomainPkg:      domainPkg,
				DomainType:     domainType,
				DomainReceiver: domainReceiver,
			}

			if err := renderToFile("internal/domain/entity.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, domainFilename+".go")); err != nil {
				return fmt.Errorf("failed to generate entity.go for domain: %w", err)
			}

			if err := renderToFile("internal/domain/patch.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, "patch.go")); err != nil {
				return fmt.Errorf("failed to generate patch.go for domain: %w", err)
			}

			if err := renderToFile("internal/domain/field.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, "field.go")); err != nil {
				return fmt.Errorf("failed to generate field.go for domain: %w", err)
			}

			if err := renderToFile("internal/domain/query.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, "query.go")); err != nil {
				return fmt.Errorf("failed to generate query.go for domain: %w", err)
			}

			if err := renderToFile("internal/domain/repo.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, "repo.go")); err != nil {
				return fmt.Errorf("failed to generate repo.go for domain: %w", err)
			}

			if err := renderToFile("internal/domain/repo_gorm.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, "repo_gorm.go")); err != nil {
				return fmt.Errorf("failed to generate repo_gorm.go for domain: %w", err)
			}

			if err := renderToFile("internal/domain/event.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, "event.go")); err != nil {
				return fmt.Errorf("failed to generate event.go for domain: %w", err)
			}

			if err := renderToFile("internal/domain/event_watermill.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, "event_watermill.go")); err != nil {
				return fmt.Errorf("failed to generate event_watermill.go for domain: %w", err)
			}

			if err := renderToFile("internal/domain/service.go.tmpl", data, acc.Filepath(featurePkg, "internal", domainPkg, "service.go")); err != nil {
				return fmt.Errorf("failed to generate service.go for domain: %w", err)
			}

			return nil
		}
	}

	return false, nil
}

func ensureDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}

	return nil
}

//go:embed .templates/*
var templatesFS embed.FS

func renderToFile(templatePath string, data interface{}, filePath string) error {
	// Read the template file
	temp, err := templatesFS.ReadFile(".templates/" + templatePath)
	if err != nil {
		return errors.New("failed to read template file")
	}

	if err := ensureDirExists(filepath.Dir(filePath)); err != nil {
		return errors.New("failed to ensure directory exists")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	// Execute the template with the provided data
	tmpl, err := template.New(filepath.Base(filePath)).Parse(string(temp))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// If the file is a Go file, format it
	if isGoExtension(filePath) {
		if err := formatGoFile(filePath); err != nil {
			return fmt.Errorf("failed to format Go file: %w", err)
		}
	}

	return nil
}

func isGoExtension(fileName string) bool {
	return strings.HasSuffix(fileName, ".go")
}

func formatGoFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("failed to read Go file for formatting")
	}

	p, err := format.Source(data)
	if err != nil {
		return errors.New("failed to format Go file")
	}

	if err := os.WriteFile(filePath, p, 0644); err != nil {
		return errors.New("failed to write formatted Go file")
	}

	return nil
}
