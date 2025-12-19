package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestValidateConfig_ValidConfig tests explicit validation with ValidateConfig
func TestValidateConfig_ValidConfig(t *testing.T) {
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

	err := ValidateConfig(configPath)
	if err != nil {
		t.Fatalf("ValidateConfig() failed for valid config: %v", err)
	}
}

// TestValidateConfig_FileNotFound tests error when file doesn't exist
func TestValidateConfig_FileNotFound(t *testing.T) {
	err := ValidateConfig("/nonexistent/path/config.yml")
	if err == nil {
		t.Fatal("Expected error for non-existent file")
	}
	if !strings.Contains(err.Error(), "no such file") {
		t.Errorf("Expected 'no such file' error, got: %v", err)
	}
}

// TestValidateConfig_MalformedYAML tests error when YAML is malformed
func TestValidateConfig_MalformedYAML(t *testing.T) {
	content := `thresholds:
  max_grade: [this is not valid yaml syntax
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	err := ValidateConfig(configPath)
	if err == nil {
		t.Fatal("Expected error for malformed YAML")
	}
	if !strings.Contains(err.Error(), "invalid YAML") {
		t.Errorf("Expected 'invalid YAML' error, got: %v", err)
	}
}

// TestValidateConfig_InvalidType tests validation errors are caught
func TestValidateConfig_InvalidType(t *testing.T) {
	content := `thresholds:
  max_grade: "twelve"
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	err := ValidateConfig(configPath)
	if err == nil {
		t.Fatal("Expected error for string instead of number")
	}

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

// TestLoad_ValidYAML tests that Load() validates and succeeds with valid YAML
func TestLoad_ValidYAML(t *testing.T) {
	content := `thresholds:
  max_grade: 14
  max_ari: 14
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	if cfg.Thresholds.MaxGrade != 14 {
		t.Errorf("MaxGrade = %v, want 14", cfg.Thresholds.MaxGrade)
	}
}

// TestLoad_InvalidType tests that Load() validates and rejects type errors
func TestLoad_InvalidType(t *testing.T) {
	content := `thresholds:
  max_grade: "sixteen"
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(configPath)
	if err == nil {
		t.Fatal("Expected Load() to fail for invalid config")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "Configuration validation failed") {
		t.Errorf("Expected validation error, got: %v", err)
	}
	if !strings.Contains(errMsg, "max_grade") {
		t.Errorf("Error should mention 'max_grade', got: %v", err)
	}
}

// TestValidateAgainstSchema_Success tests successful schema validation
func TestValidateAgainstSchema_Success(t *testing.T) {
	validData := map[string]interface{}{
		"thresholds": map[string]interface{}{
			"max_grade": 14.0,
			"max_ari":   14.0,
		},
	}

	err := ValidateAgainstSchema(validData)
	if err != nil {
		t.Errorf("ValidateAgainstSchema() failed for valid data: %v", err)
	}

	// Call again to exercise the schema caching (sync.Once) path
	err = ValidateAgainstSchema(validData)
	if err != nil {
		t.Errorf("ValidateAgainstSchema() failed on second call: %v", err)
	}
}

// TestValidateAgainstSchema_InvalidData tests schema validation failure
func TestValidateAgainstSchema_InvalidData(t *testing.T) {
	invalidData := map[string]interface{}{
		"thresholds": map[string]interface{}{
			"max_grade":     "not a number",
			"unknown_field": 123,
		},
	}

	err := ValidateAgainstSchema(invalidData)
	if err == nil {
		t.Error("Expected validation error for invalid data")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "Configuration validation failed") {
		t.Errorf("Expected validation error message, got: %v", err)
	}
}

// TestInstanceLocationToYAMLPath tests path conversion
func TestInstanceLocationToYAMLPath(t *testing.T) {
	tests := []struct {
		location []string
		want     string
	}{
		{[]string{}, "(root)"},
		{[]string{"thresholds"}, "thresholds"},
		{[]string{"thresholds", "max_grade"}, "thresholds.max_grade"},
		{[]string{"overrides", "0", "path"}, "overrides.0.path"},
	}

	for _, tt := range tests {
		got := instanceLocationToYAMLPath(tt.location)
		if got != tt.want {
			t.Errorf("instanceLocationToYAMLPath(%v) = %q, want %q", tt.location, got, tt.want)
		}
	}
}

// TestGetSuggestion_AllPaths tests error message suggestions
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
		{"some other error", ""}, // No suggestion for unknown errors
	}

	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			got := getSuggestion(tt.errMsg)
			if tt.wantContain == "" {
				if got != "" {
					t.Errorf("getSuggestion(%q) = %q, want empty string", tt.errMsg, got)
				}
			} else {
				if !strings.Contains(got, tt.wantContain) {
					t.Errorf("getSuggestion(%q) = %q, want to contain %q", tt.errMsg, got, tt.wantContain)
				}
			}
		})
	}
}
