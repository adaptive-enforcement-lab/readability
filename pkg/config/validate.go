package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

const schemaURL = "https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json"

//go:embed schema.json
var embeddedSchemaBytes []byte

var (
	compiledSchema     *jsonschema.Schema
	schemaCompileError error
	schemaOnce         sync.Once
)

// getCompiledSchema loads and compiles the embedded schema on first use.
//
// Note: Error handling in this function covers defensive "should never happen" cases
// since the schema is generated from Go structs and embedded at compile time.
// These error paths are not easily testable without mocking the embed directive.
func getCompiledSchema() (*jsonschema.Schema, error) {
	schemaOnce.Do(func() {
		// Unmarshal embedded schema JSON
		var schemaData interface{}
		if err := json.Unmarshal(embeddedSchemaBytes, &schemaData); err != nil {
			// This would only happen if the embedded schema is malformed JSON,
			// which should be caught by go:generate and CI validation
			schemaCompileError = fmt.Errorf("failed to unmarshal embedded schema: %w", err)
			return
		}

		// Compile schema
		compiler := jsonschema.NewCompiler()
		if err := compiler.AddResource(schemaURL, schemaData); err != nil {
			// This would only happen if the schema structure is invalid
			schemaCompileError = fmt.Errorf("failed to add schema resource: %w", err)
			return
		}

		compiledSchema, schemaCompileError = compiler.Compile(schemaURL)
	})

	return compiledSchema, schemaCompileError
}

// ValidateAgainstSchema validates the parsed YAML data against the JSON Schema.
// It returns a formatted error with helpful suggestions if validation fails.
func ValidateAgainstSchema(data interface{}) error {
	schema, err := getCompiledSchema()
	if err != nil {
		return fmt.Errorf("schema validation unavailable: %w", err)
	}

	if err := schema.Validate(data); err != nil {
		return formatSchemaError(err)
	}
	return nil
}

// formatSchemaError converts a JSON Schema validation error into a user-friendly
// error message with YAML paths and helpful suggestions.
func formatSchemaError(err error) error {
	validationErr, ok := err.(*jsonschema.ValidationError)
	if !ok {
		return fmt.Errorf("schema validation failed: %w", err)
	}

	var buf strings.Builder
	buf.WriteString("Configuration validation failed:\n\n")

	// Get detailed errors (includes all nested errors)
	detailedErrors := flattenValidationErrors(validationErr)

	for _, e := range detailedErrors {
		// Convert instance location slice to YAML path
		yamlPath := instanceLocationToYAMLPath(e.InstanceLocation)
		errMsg := e.Error()

		buf.WriteString(fmt.Sprintf("  • %s\n", yamlPath))
		buf.WriteString(fmt.Sprintf("    %s\n", errMsg))

		// Add suggestion if possible
		if suggestion := getSuggestion(errMsg); suggestion != "" {
			buf.WriteString(fmt.Sprintf("    Suggestion: %s\n", suggestion))
		}
		buf.WriteString("\n")
	}

	buf.WriteString("See https://github.com/adaptive-enforcement-lab/readability/blob/main/docs/cli/config-file.md for configuration reference.\n")

	return fmt.Errorf("%s", buf.String())
}

// flattenValidationErrors extracts leaf validation errors (actual problems)
func flattenValidationErrors(err *jsonschema.ValidationError) []*jsonschema.ValidationError {
	var result []*jsonschema.ValidationError

	// If this error has causes, recurse into them (parent errors are just containers)
	if len(err.Causes) > 0 {
		for _, cause := range err.Causes {
			result = append(result, flattenValidationErrors(cause)...)
		}
	} else {
		// Leaf error - this is an actual validation problem
		result = append(result, err)
	}

	return result
}

// instanceLocationToYAMLPath converts instance location slice to YAML dot notation path
// Example: ["thresholds", "max_grade"] → thresholds.max_grade
func instanceLocationToYAMLPath(location []string) string {
	if len(location) == 0 {
		return "(root)"
	}
	return strings.Join(location, ".")
}

// getSuggestion provides helpful suggestions based on the error message
func getSuggestion(errMsg string) string {
	msg := strings.ToLower(errMsg)

	switch {
	case strings.Contains(msg, "got string, want number"):
		return "Remove quotes around numeric values"
	case strings.Contains(msg, "got string, want integer"):
		return "Remove quotes around numeric values"
	case strings.Contains(msg, "additional properties") && strings.Contains(msg, "not allowed"):
		return "Check for typos in field names - unknown properties are not allowed"
	case strings.Contains(msg, "missing property") || strings.Contains(msg, "missing properties"):
		return "This field is required and cannot be omitted"
	case strings.Contains(msg, "got") && strings.Contains(msg, "want object"):
		return "This should be a YAML object with nested fields"
	case strings.Contains(msg, "got") && strings.Contains(msg, "want array"):
		return "This should be a YAML array (list)"
	case strings.Contains(msg, "maximum:"):
		return "Value exceeds the allowed maximum"
	case strings.Contains(msg, "minimum:"):
		return "Value is below the allowed minimum"
	default:
		return ""
	}
}
