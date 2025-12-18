package config

import (
	"testing"

	"gopkg.in/yaml.v3"
)

// BenchmarkSchemaCompilation measures schema loading and compilation performance.
// This should be fast on first call and instant on subsequent calls (cached).
func BenchmarkSchemaCompilation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := getCompiledSchema()
		if err != nil {
			b.Fatalf("Failed to compile schema: %v", err)
		}
	}
}

// BenchmarkValidateValidConfig measures validation performance for a valid configuration.
func BenchmarkValidateValidConfig(b *testing.B) {
	validYAML := `
thresholds:
  max_grade: 16
  max_ari: 16
  max_fog: 18
  min_ease: 25
  max_lines: 1000
  min_words: 100
  min_admonitions: 1
  max_dash_density: 2
overrides:
  - path: docs/reference/
    thresholds:
      max_fog: 20
`
	var data interface{}
	if err := yaml.Unmarshal([]byte(validYAML), &data); err != nil {
		b.Fatalf("Failed to unmarshal test YAML: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ValidateAgainstSchema(data); err != nil {
			b.Fatalf("Validation failed: %v", err)
		}
	}
}

// BenchmarkValidateInvalidConfig measures validation performance when detecting errors.
func BenchmarkValidateInvalidConfig(b *testing.B) {
	invalidYAML := `
thresholds:
  max_grade: 200  # exceeds maximum of 100
  max_ari: "sixteen"  # should be number
  invalid_field: true  # unknown property
`
	var data interface{}
	if err := yaml.Unmarshal([]byte(invalidYAML), &data); err != nil {
		b.Fatalf("Failed to unmarshal test YAML: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ValidateAgainstSchema(data)
		// We expect this to fail, so we don't check the error
	}
}

// BenchmarkValidateMinimalConfig measures validation of a minimal valid config.
func BenchmarkValidateMinimalConfig(b *testing.B) {
	minimalYAML := `
thresholds:
  max_grade: 12
`
	var data interface{}
	if err := yaml.Unmarshal([]byte(minimalYAML), &data); err != nil {
		b.Fatalf("Failed to unmarshal test YAML: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ValidateAgainstSchema(data); err != nil {
			b.Fatalf("Validation failed: %v", err)
		}
	}
}

// BenchmarkValidateComplexConfig measures validation of a config with many overrides.
func BenchmarkValidateComplexConfig(b *testing.B) {
	complexYAML := `
thresholds:
  max_grade: 16
  max_ari: 16
  max_fog: 18
  min_ease: 25
  max_lines: 1000
  min_words: 100
  min_admonitions: 1
  max_dash_density: 2
overrides:
  - path: docs/api/
    thresholds:
      max_grade: 18
      max_fog: 20
  - path: docs/tutorials/
    thresholds:
      min_ease: 40
      max_grade: 12
  - path: docs/reference/
    thresholds:
      max_grade: 20
      max_lines: 2000
  - path: docs/contributing/
    thresholds:
      max_grade: 14
      min_admonitions: 2
`
	var data interface{}
	if err := yaml.Unmarshal([]byte(complexYAML), &data); err != nil {
		b.Fatalf("Failed to unmarshal test YAML: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ValidateAgainstSchema(data); err != nil {
			b.Fatalf("Validation failed: %v", err)
		}
	}
}

// BenchmarkErrorFormatting measures the performance of error message formatting.
func BenchmarkErrorFormatting(b *testing.B) {
	// Create invalid config that will trigger multiple errors
	invalidYAML := `
thresholds:
  max_grade: 200
  max_ari: "invalid"
  unknown_field: true
overrides:
  - thresholds:
      max_grade: 10
`
	var data interface{}
	if err := yaml.Unmarshal([]byte(invalidYAML), &data); err != nil {
		b.Fatalf("Failed to unmarshal test YAML: %v", err)
	}

	// Get the validation error once
	validationErr := ValidateAgainstSchema(data)
	if validationErr == nil {
		b.Fatal("Expected validation error, got nil")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Benchmark just the formatting by triggering validation
		_ = ValidateAgainstSchema(data)
	}
}
