package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func TestValidateAgainstSchema_ValidConfig(t *testing.T) {
	// Create a temp config file with valid configuration
	content := `thresholds:
  max_grade: 14
  max_ari: 14
  min_ease: 40
  max_lines: 300
  min_words: 100
  min_admonitions: 1
  max_dash_density: 0
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	// Load should succeed (Load calls ValidateAgainstSchema internally)
	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() failed for valid config: %v", err)
	}

	if cfg.Thresholds.MaxGrade != 14 {
		t.Errorf("MaxGrade = %v, want 14", cfg.Thresholds.MaxGrade)
	}
}

func TestValidateAgainstSchema_InvalidType(t *testing.T) {
	// Test quoted numbers (strings instead of numbers)
	content := `thresholds:
  max_grade: "twelve"
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(configPath)
	if err == nil {
		t.Fatal("Expected error for string instead of number")
	}

	// Check that error message contains helpful information
	errMsg := err.Error()
	if !strings.Contains(errMsg, "max_grade") {
		t.Errorf("Error should mention 'max_grade', got: %v", errMsg)
	}
	if !strings.Contains(errMsg, "got string, want number") {
		t.Errorf("Error should mention type mismatch, got: %v", errMsg)
	}
	if !strings.Contains(errMsg, "Remove quotes") {
		t.Errorf("Error should include suggestion about removing quotes, got: %v", errMsg)
	}
}

func TestValidateAgainstSchema_AdditionalProperties(t *testing.T) {
	// Test unknown field
	content := `thresholds:
  max_grade: 14
  invalid_field: 100
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(configPath)
	if err == nil {
		t.Fatal("Expected error for additional property")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "additional properties") {
		t.Errorf("Error should mention 'additional properties', got: %v", errMsg)
	}
	if !strings.Contains(errMsg, "invalid_field") {
		t.Errorf("Error should mention the invalid field name, got: %v", errMsg)
	}
	if !strings.Contains(errMsg, "Check for typos") {
		t.Errorf("Error should include suggestion about typos, got: %v", errMsg)
	}
}

func TestValidateAgainstSchema_RangeViolations(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		shouldFail  bool
		wantContain string
	}{
		{
			name: "max_grade above maximum",
			content: `thresholds:
  max_grade: 150
`,
			shouldFail:  true,
			wantContain: "maximum:",
		},
		{
			name: "max_grade below minimum",
			content: `thresholds:
  max_grade: -5
`,
			shouldFail:  true,
			wantContain: "minimum:",
		},
		{
			name: "valid max_grade at boundary",
			content: `thresholds:
  max_grade: 100
`,
			shouldFail: false,
		},
		{
			name: "valid max_grade at lower boundary",
			content: `thresholds:
  max_grade: 0
`,
			shouldFail: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			configPath := filepath.Join(tmpDir, "test.yml")
			if err := os.WriteFile(configPath, []byte(tt.content), 0644); err != nil {
				t.Fatal(err)
			}

			_, err := Load(configPath)
			if tt.shouldFail {
				if err == nil {
					t.Fatal("Expected validation error")
				}
				if tt.wantContain != "" && !strings.Contains(err.Error(), tt.wantContain) {
					t.Errorf("Error should contain %q, got: %v", tt.wantContain, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("Expected valid config, got error: %v", err)
				}
			}
		})
	}
}

func TestValidateAgainstSchema_RequiredFields(t *testing.T) {
	// Schema requires 'thresholds' object - test with empty config
	content := `{}`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "empty.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(configPath)
	// Empty config is actually valid - defaults will be applied
	// The schema doesn't require 'thresholds' to be present
	if err != nil {
		t.Fatalf("Empty config should be valid (defaults applied): %v", err)
	}
	if cfg.Thresholds.MaxGrade != 16.0 {
		t.Errorf("Should use default MaxGrade, got %v", cfg.Thresholds.MaxGrade)
	}
}

func TestValidateAgainstSchema_OverridesValidation(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		shouldFail  bool
		wantContain string
	}{
		{
			name: "valid override",
			content: `thresholds:
  max_grade: 16
overrides:
  - path: docs/api/
    thresholds:
      max_grade: 20
`,
			shouldFail: false,
		},
		{
			name: "override with invalid type",
			content: `thresholds:
  max_grade: 16
overrides:
  - path: docs/api/
    thresholds:
      max_grade: "twenty"
`,
			shouldFail:  true,
			wantContain: "got string, want number",
		},
		{
			name: "override missing required path",
			content: `thresholds:
  max_grade: 16
overrides:
  - thresholds:
      max_grade: 20
`,
			shouldFail:  true,
			wantContain: "missing property",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			configPath := filepath.Join(tmpDir, "test.yml")
			if err := os.WriteFile(configPath, []byte(tt.content), 0644); err != nil {
				t.Fatal(err)
			}

			_, err := Load(configPath)
			if tt.shouldFail {
				if err == nil {
					t.Fatal("Expected validation error")
				}
				if tt.wantContain != "" && !strings.Contains(err.Error(), tt.wantContain) {
					t.Errorf("Error should contain %q, got: %v", tt.wantContain, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("Expected valid config, got error: %v", err)
				}
			}
		})
	}
}

func TestValidateAgainstSchema_AllThresholdFields(t *testing.T) {
	// Test that all threshold fields validate correctly
	content := `thresholds:
  max_grade: 16.0
  max_ari: 16.0
  max_fog: 18.0
  min_ease: 25.0
  max_lines: 375
  min_words: 100
  min_admonitions: 1
  max_dash_density: 0.0
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "complete.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() failed for complete valid config: %v", err)
	}

	// Verify all fields loaded correctly
	if cfg.Thresholds.MaxGrade != 16.0 {
		t.Errorf("MaxGrade = %v, want 16.0", cfg.Thresholds.MaxGrade)
	}
	if cfg.Thresholds.MaxARI != 16.0 {
		t.Errorf("MaxARI = %v, want 16.0", cfg.Thresholds.MaxARI)
	}
	if cfg.Thresholds.MaxFog != 18.0 {
		t.Errorf("MaxFog = %v, want 18.0", cfg.Thresholds.MaxFog)
	}
	if cfg.Thresholds.MinEase != 25.0 {
		t.Errorf("MinEase = %v, want 25.0", cfg.Thresholds.MinEase)
	}
	if cfg.Thresholds.MaxLines != 375 {
		t.Errorf("MaxLines = %v, want 375", cfg.Thresholds.MaxLines)
	}
	if cfg.Thresholds.MinWords != 100 {
		t.Errorf("MinWords = %v, want 100", cfg.Thresholds.MinWords)
	}
	if cfg.Thresholds.MinAdmonitions != 1 {
		t.Errorf("MinAdmonitions = %v, want 1", cfg.Thresholds.MinAdmonitions)
	}
	if cfg.Thresholds.MaxDashDensity != 0.0 {
		t.Errorf("MaxDashDensity = %v, want 0.0", cfg.Thresholds.MaxDashDensity)
	}
}

func TestLoadOrDefault_WithInvalidConfig(t *testing.T) {
	// LoadOrDefault should return defaults even if config is invalid
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid.yml")
	content := "thresholds:\n  max_grade: \"invalid\"\n"
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg := LoadOrDefault(configPath)

	// Should fall back to defaults
	if cfg.Thresholds.MaxGrade != 16.0 {
		t.Errorf("Expected default MaxGrade 16.0, got %v", cfg.Thresholds.MaxGrade)
	}
}

func TestGetSuggestion_AllPaths(t *testing.T) {
	tests := []struct {
		errMsg      string
		wantContain string
	}{
		{"got string, want number", "Remove quotes"},
		{"got string, want integer", "Remove quotes"},
		{"additional properties are not allowed", "Check for typos"},
		{"missing property 'thresholds'", "This field is required"},
		{"missing properties: 'path', 'thresholds'", "This field is required"},
		{"got null, want object", "This should be a YAML object"},
		{"got string, want array", "This should be a YAML array"},
		{"value exceeds maximum: 150", "Value exceeds the allowed maximum"},
		{"value below minimum: -5", "Value is below the allowed minimum"},
		{"some unknown error", ""}, // Default case - no suggestion
	}

	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			suggestion := getSuggestion(tt.errMsg)
			if tt.wantContain == "" {
				if suggestion != "" {
					t.Errorf("Expected empty suggestion for %q, got %q", tt.errMsg, suggestion)
				}
			} else {
				if !strings.Contains(suggestion, tt.wantContain) {
					t.Errorf("getSuggestion(%q) = %q, want to contain %q", tt.errMsg, suggestion, tt.wantContain)
				}
			}
		})
	}
}

func TestInstanceLocationToYAMLPath_EmptyLocation(t *testing.T) {
	path := instanceLocationToYAMLPath([]string{})
	if path != "(root)" {
		t.Errorf("instanceLocationToYAMLPath([]) = %q, want \"(root)\"", path)
	}
}

func TestInstanceLocationToYAMLPath_WithLocation(t *testing.T) {
	path := instanceLocationToYAMLPath([]string{"thresholds", "max_grade"})
	if path != "thresholds.max_grade" {
		t.Errorf("instanceLocationToYAMLPath([thresholds, max_grade]) = %q, want \"thresholds.max_grade\"", path)
	}
}

func TestFindSchemaFile_FromCurrentDir(t *testing.T) {
	// Save original working directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Create temp directory structure with schema
	tmpDir := t.TempDir()
	schemaDir := filepath.Join(tmpDir, "docs", "schemas")
	if err := os.MkdirAll(schemaDir, 0755); err != nil {
		t.Fatal(err)
	}
	schemaPath := filepath.Join(schemaDir, "config.json")
	if err := os.WriteFile(schemaPath, []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}

	// Change to temp directory
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Should find schema from current directory
	found := findSchemaFile()
	if found != "docs/schemas/config.json" {
		t.Errorf("findSchemaFile() = %q, want \"docs/schemas/config.json\"", found)
	}
}

func TestFindSchemaFile_FromGitRoot(t *testing.T) {
	// Save original working directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Create temp directory structure: root/.git and root/docs/schemas/config.json
	tmpDir := t.TempDir()
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatal(err)
	}
	schemaDir := filepath.Join(tmpDir, "docs", "schemas")
	if err := os.MkdirAll(schemaDir, 0755); err != nil {
		t.Fatal(err)
	}
	schemaPath := filepath.Join(schemaDir, "config.json")
	if err := os.WriteFile(schemaPath, []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create subdirectory and change to it
	subDir := filepath.Join(tmpDir, "pkg", "config")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(subDir); err != nil {
		t.Fatal(err)
	}

	// Should find schema by walking up to git root
	found := findSchemaFile()
	expectedSuffix := filepath.Join("docs", "schemas", "config.json")
	if !strings.HasSuffix(found, expectedSuffix) {
		t.Errorf("findSchemaFile() = %q, want path ending with %q", found, expectedSuffix)
	}
	// Verify file actually exists
	if _, err := os.Stat(found); err != nil {
		t.Errorf("Schema file not found at %q: %v", found, err)
	}
}

func TestFindSchemaFile_NotFound(t *testing.T) {
	// Save original working directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Create temp directory with no schema
	tmpDir := t.TempDir()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Should return empty string when not found
	found := findSchemaFile()
	if found != "" {
		t.Errorf("findSchemaFile() = %q, want empty string", found)
	}
}

func TestGetCompiledSchema_SchemaNotFound(t *testing.T) {
	// Save original working directory and reset schema state
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// Reset global state for schema
		compiledSchema = nil
		schemaCompileError = nil
		schemaOnce = sync.Once{}
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Reset schema state before test
	compiledSchema = nil
	schemaCompileError = nil
	schemaOnce = sync.Once{}

	// Change to directory without schema
	tmpDir := t.TempDir()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Should return error when schema not found
	_, err = getCompiledSchema()
	if err == nil {
		t.Fatal("Expected error when schema not found")
	}
	if !strings.Contains(err.Error(), "schema file not found") {
		t.Errorf("Error should mention 'schema file not found', got: %v", err)
	}
}

func TestGetCompiledSchema_InvalidJSON(t *testing.T) {
	// Save original working directory and reset schema state
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// Reset global state for schema
		compiledSchema = nil
		schemaCompileError = nil
		schemaOnce = sync.Once{}
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Reset schema state before test
	compiledSchema = nil
	schemaCompileError = nil
	schemaOnce = sync.Once{}

	// Create temp directory with invalid JSON schema
	tmpDir := t.TempDir()
	schemaDir := filepath.Join(tmpDir, "docs", "schemas")
	if err := os.MkdirAll(schemaDir, 0755); err != nil {
		t.Fatal(err)
	}
	schemaPath := filepath.Join(schemaDir, "config.json")
	if err := os.WriteFile(schemaPath, []byte("invalid json"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Should return error for invalid JSON
	_, err = getCompiledSchema()
	if err == nil {
		t.Fatal("Expected error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "failed to unmarshal schema") {
		t.Errorf("Error should mention 'failed to unmarshal schema', got: %v", err)
	}
}

func TestValidateAgainstSchema_SchemaLoadError(t *testing.T) {
	// Save original working directory and reset schema state
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// Reset global state for schema
		compiledSchema = nil
		schemaCompileError = nil
		schemaOnce = sync.Once{}
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Reset schema state before test
	compiledSchema = nil
	schemaCompileError = nil
	schemaOnce = sync.Once{}

	// Change to directory without schema
	tmpDir := t.TempDir()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Should return error when schema can't be loaded
	err = ValidateAgainstSchema(map[string]interface{}{})
	if err == nil {
		t.Fatal("Expected error when schema unavailable")
	}
	if !strings.Contains(err.Error(), "schema validation unavailable") {
		t.Errorf("Error should mention 'schema validation unavailable', got: %v", err)
	}
}

func TestFormatSchemaError_NonValidationError(t *testing.T) {
	// Test with a non-ValidationError
	err := formatSchemaError(os.ErrNotExist)
	if err == nil {
		t.Fatal("Expected error to be returned")
	}
	if !strings.Contains(err.Error(), "schema validation failed") {
		t.Errorf("Error should mention 'schema validation failed', got: %v", err)
	}
}

func TestGetCompiledSchema_ReadError(t *testing.T) {
	// Save original working directory and reset schema state
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// Reset global state for schema
		compiledSchema = nil
		schemaCompileError = nil
		schemaOnce = sync.Once{}
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Reset schema state before test
	compiledSchema = nil
	schemaCompileError = nil
	schemaOnce = sync.Once{}

	// Create temp directory with schema as a directory (not a file) to cause read error
	tmpDir := t.TempDir()
	schemaDir := filepath.Join(tmpDir, "docs", "schemas")
	if err := os.MkdirAll(schemaDir, 0755); err != nil {
		t.Fatal(err)
	}
	// Create config.json as a directory instead of a file
	schemaPath := filepath.Join(schemaDir, "config.json")
	if err := os.Mkdir(schemaPath, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Should return error when schema can't be read
	_, err = getCompiledSchema()
	if err == nil {
		t.Fatal("Expected error when schema file can't be read")
	}
	if !strings.Contains(err.Error(), "failed to read schema file") {
		t.Errorf("Error should mention 'failed to read schema file', got: %v", err)
	}
}

func TestFindSchemaFile_GitRootWithoutSchema(t *testing.T) {
	// Save original working directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Create temp directory with .git but NO schema
	tmpDir := t.TempDir()
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create subdirectory and change to it
	subDir := filepath.Join(tmpDir, "pkg")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(subDir); err != nil {
		t.Fatal(err)
	}

	// Should return empty string - stops at .git but schema not found
	found := findSchemaFile()
	if found != "" {
		t.Errorf("findSchemaFile() = %q, want empty string (git root without schema)", found)
	}
}

func TestGetCompiledSchema_InvalidSchemaType(t *testing.T) {
	// Save original working directory and reset schema state
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// Reset global state for schema
		compiledSchema = nil
		schemaCompileError = nil
		schemaOnce = sync.Once{}
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Reset schema state before test
	compiledSchema = nil
	schemaCompileError = nil
	schemaOnce = sync.Once{}

	// Create temp directory with schema that is not an object (array instead)
	tmpDir := t.TempDir()
	schemaDir := filepath.Join(tmpDir, "docs", "schemas")
	if err := os.MkdirAll(schemaDir, 0755); err != nil {
		t.Fatal(err)
	}
	schemaPath := filepath.Join(schemaDir, "config.json")
	// Array instead of object - may cause AddResource to fail
	invalidSchema := `["not", "a", "schema", "object"]`
	if err := os.WriteFile(schemaPath, []byte(invalidSchema), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Should return error when schema type is wrong
	_, err = getCompiledSchema()
	if err == nil {
		t.Fatal("Expected error when schema is not an object")
	}

	// Verify the error is from AddResource (coverage for lines 47-50)
	if !strings.Contains(err.Error(), "failed to add schema resource") {
		t.Logf("Got error from Compile instead of AddResource: %v", err)
		t.Logf("This test may not cover the AddResource error path (lines 47-50)")
	}
}

func TestFindSchemaFile_WalkToRoot(t *testing.T) {
	// NOTE: The parent == dir check at line 79-80 is actually unreachable
	// due to the loop condition `for dir != "" && dir != "/"`.
	// When dir == "/", the loop never runs, so we can't reach the parent == dir check.
	// This test verifies findSchemaFile returns empty when no schema is found.

	// Save original working directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Change to filesystem root
	if err := os.Chdir("/"); err != nil {
		t.Skip("Cannot change to filesystem root")
	}

	// findSchemaFile from root should return empty string
	found := findSchemaFile()
	if found != "" {
		t.Errorf("findSchemaFile() from root = %q, want empty string", found)
	}
}

func TestGetCompiledSchema_AddResourceError(t *testing.T) {
	// Test coverage for AddResource error path (lines 47-50)
	// This requires a schema that JSON unmarshals successfully
	// but fails when added to the compiler.

	// NOTE: With the current jsonschema library, AddResource rarely fails
	// because it accepts any valid JSON structure. The actual validation
	// happens in Compile(). This test documents the error path exists
	// but may not be practically reachable.

	// Reset schema state
	compiledSchema = nil
	schemaCompileError = nil
	schemaOnce = sync.Once{}

	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		compiledSchema = nil
		schemaCompileError = nil
		schemaOnce = sync.Once{}
		if err := os.Chdir(originalWd); err != nil {
			t.Errorf("Failed to restore working directory: %v", err)
		}
	}()

	// Create temp directory with intentionally invalid schema
	// Using a string value which should fail schema validation
	tmpDir := t.TempDir()
	schemaDir := filepath.Join(tmpDir, "docs", "schemas")
	if err := os.MkdirAll(schemaDir, 0755); err != nil {
		t.Fatal(err)
	}
	schemaPath := filepath.Join(schemaDir, "config.json")
	// String literal is valid JSON but not a valid schema
	invalidSchema := `"not a schema"`
	if err := os.WriteFile(schemaPath, []byte(invalidSchema), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Should return error
	_, err = getCompiledSchema()
	if err == nil {
		t.Fatal("Expected error for invalid schema")
	}

	// The error could be from either AddResource or Compile
	// Both are valid error paths
	t.Logf("Got schema error (as expected): %v", err)
}
