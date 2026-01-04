package config

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// TestGetCompiledSchema_CorruptedJSON tests error handling for corrupted embedded schema
func TestGetCompiledSchema_CorruptedJSON(t *testing.T) {
	// Save original values
	origBytes := embeddedSchemaBytes
	origSchema := compiledSchema
	origError := schemaCompileError

	// Restore after test
	defer func() {
		embeddedSchemaBytes = origBytes
		schemaOnce = sync.Once{}
		compiledSchema = origSchema
		schemaCompileError = origError
	}()

	// Inject corrupted JSON
	embeddedSchemaBytes = []byte(`{invalid json`)
	schemaOnce = sync.Once{}
	compiledSchema = nil
	schemaCompileError = nil

	_, err := getCompiledSchema()
	if err == nil {
		t.Fatal("Expected error for corrupted JSON schema")
	}
	if !strings.Contains(err.Error(), "failed to unmarshal embedded schema") {
		t.Errorf("Expected unmarshal error, got: %v", err)
	}
}

// TestGetCompiledSchema_InvalidSchemaStructure tests error handling for invalid schema structure
func TestGetCompiledSchema_InvalidSchemaStructure(t *testing.T) {
	// Save original values
	origBytes := embeddedSchemaBytes
	origSchema := compiledSchema
	origError := schemaCompileError

	// Restore after test
	defer func() {
		embeddedSchemaBytes = origBytes
		schemaOnce = sync.Once{}
		compiledSchema = origSchema
		schemaCompileError = origError
	}()

	// Inject valid JSON but invalid schema structure
	invalidSchema := map[string]interface{}{
		"$schema": "invalid-schema-version-that-doesnt-exist",
	}
	invalidBytes, _ := json.Marshal(invalidSchema)
	embeddedSchemaBytes = invalidBytes
	schemaOnce = sync.Once{}
	compiledSchema = nil
	schemaCompileError = nil

	_, err := getCompiledSchema()
	if err == nil {
		t.Fatal("Expected error for invalid schema structure")
	}
	// Accept any error from schema loading/compilation
	errMsg := err.Error()
	validError := strings.Contains(errMsg, "failed to add schema resource") ||
		strings.Contains(errMsg, "failing loading") ||
		strings.Contains(errMsg, "invalid file url")
	if !validError {
		t.Errorf("Expected schema loading error, got: %v", err)
	}
}

// TestValidateAgainstSchema_SchemaCompilationFailure tests ValidateAgainstSchema when schema compilation fails
func TestValidateAgainstSchema_SchemaCompilationFailure(t *testing.T) {
	// Save original values
	origBytes := embeddedSchemaBytes
	origSchema := compiledSchema
	origError := schemaCompileError

	// Restore after test
	defer func() {
		embeddedSchemaBytes = origBytes
		schemaOnce = sync.Once{}
		compiledSchema = origSchema
		schemaCompileError = origError
	}()

	// Inject corrupted schema to cause compilation failure
	embeddedSchemaBytes = []byte(`{invalid`)
	schemaOnce = sync.Once{}
	compiledSchema = nil
	schemaCompileError = nil

	testData := map[string]interface{}{
		"thresholds": map[string]interface{}{
			"max_grade": 14.0,
		},
	}

	err := ValidateAgainstSchema(testData)
	if err == nil {
		t.Fatal("Expected error when schema compilation fails")
	}
	if !strings.Contains(err.Error(), "schema validation unavailable") {
		t.Errorf("Expected 'schema validation unavailable' error, got: %v", err)
	}
}

// TestFormatSchemaError_NonValidationError tests formatSchemaError with non-ValidationError
func TestFormatSchemaError_NonValidationError(t *testing.T) {
	err := formatSchemaError(errors.New("not a validation error"))
	if err == nil {
		t.Fatal("Expected error for non-ValidationError input")
	}
	if !strings.Contains(err.Error(), "schema validation failed") {
		t.Errorf("Expected 'schema validation failed' error, got: %v", err)
	}
}

// TestInstanceLocationToYAMLPath tests YAML path conversion
func TestInstanceLocationToYAMLPath(t *testing.T) {
	tests := []struct {
		name string
		loc  []string
		want string
	}{
		{"empty", []string{}, "(root)"},
		{"single", []string{"thresholds"}, "thresholds"},
		{"nested", []string{"overrides", "0", "thresholds", "max_grade"}, "overrides.0.thresholds.max_grade"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := instanceLocationToYAMLPath(tt.loc)
			if got != tt.want {
				t.Errorf("instanceLocationToYAMLPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetSuggestion tests all suggestion patterns
func TestGetSuggestion(t *testing.T) {
	tests := []struct {
		errMsg string
		want   string
	}{
		{"got string, want number", "Remove quotes around numeric values"},
		{"got string, want integer", "Remove quotes around numeric values"},
		{"additional properties not allowed", "Check for typos in field names - unknown properties are not allowed"},
		{"missing property", "This field is required and cannot be omitted"},
		{"missing properties", "This field is required and cannot be omitted"},
		{"got array, want object", "This should be a YAML object with nested fields"},
		{"got string, want array", "This should be a YAML array (list)"},
		{"maximum: 100", "Value exceeds the allowed maximum"},
		{"minimum: 0", "Value is below the allowed minimum"},
		{"unknown error", ""},
	}

	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			got := getSuggestion(tt.errMsg)
			if got != tt.want {
				t.Errorf("getSuggestion(%q) = %q, want %q", tt.errMsg, got, tt.want)
			}
		})
	}
}

// TestValidateConfig tests the public ValidateConfig function
func TestValidateConfig(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := tmpDir + "/valid.yml"
		content := `thresholds:
  max_grade: 16
`
		if err := writeTestFile(configPath, content); err != nil {
			t.Fatal(err)
		}

		err := ValidateConfig(configPath)
		if err != nil {
			t.Errorf("ValidateConfig() unexpected error: %v", err)
		}
	})

	t.Run("file not found", func(t *testing.T) {
		err := ValidateConfig("/nonexistent/file.yml")
		if err == nil {
			t.Error("Expected error for non-existent file")
		}
	})

	t.Run("invalid YAML", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := tmpDir + "/invalid.yml"
		content := `{invalid yaml`
		if err := writeTestFile(configPath, content); err != nil {
			t.Fatal(err)
		}

		err := ValidateConfig(configPath)
		if err == nil {
			t.Error("Expected error for invalid YAML")
		}
	})

	t.Run("invalid schema", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := tmpDir + "/schema-invalid.yml"
		content := `thresholds:
  max_grade: "not a number"
`
		if err := writeTestFile(configPath, content); err != nil {
			t.Fatal(err)
		}

		err := ValidateConfig(configPath)
		if err == nil {
			t.Error("Expected schema validation error")
		}
	})
}

// TestFlattenValidationErrors tests nested error flattening
func TestFlattenValidationErrors(t *testing.T) {
	// Create nested validation errors
	leaf1 := &jsonschema.ValidationError{
		InstanceLocation: []string{"field1"},
	}
	leaf2 := &jsonschema.ValidationError{
		InstanceLocation: []string{"field2"},
	}
	parent := &jsonschema.ValidationError{
		InstanceLocation: []string{"root"},
		Causes:           []*jsonschema.ValidationError{leaf1, leaf2},
	}

	result := flattenValidationErrors(parent)
	if len(result) != 2 {
		t.Errorf("Expected 2 flattened errors, got %d", len(result))
	}

	// Verify only leaf errors are returned (not parent)
	for _, err := range result {
		if err == parent {
			t.Error("Parent error should not be in flattened results")
		}
	}
}

// Helper function to write test files
func writeTestFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
