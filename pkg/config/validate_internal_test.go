package config

import (
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"testing"
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
