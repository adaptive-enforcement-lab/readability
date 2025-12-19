//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adaptive-enforcement-lab/readability/pkg/config"
	"github.com/invopop/jsonschema"
)

// addExamples adds example values to schema fields
func addExamples(schema *jsonschema.Schema) {
	examples := map[string][]interface{}{
		"max_grade":        {12, 14, 16},
		"max_ari":          {12, 14, 16},
		"max_fog":          {14, 16, 18},
		"min_ease":         {30, 40, 50, -100},
		"max_lines":        {250, 375, 500},
		"min_words":        {50, 100, 150},
		"min_admonitions":  {0, 1, 2, -1},
		"max_dash_density": {0, 2, 5, -1},
		"path":             {"docs/developer-guide/", "docs/user-guide/", "api/", "README.md"},
	}

	// Apply examples to thresholds
	if thresholds, ok := schema.Properties.Get("thresholds"); ok {
		for field, exampleValues := range examples {
			if prop, ok := thresholds.Properties.Get(field); ok {
				prop.Examples = exampleValues
			}
		}
	}

	// Apply examples to override path and thresholds
	if overrides, ok := schema.Properties.Get("overrides"); ok {
		if overrides.Items != nil {
			if pathProp, ok := overrides.Items.Properties.Get("path"); ok {
				pathProp.Examples = examples["path"]
			}
			if thresholds, ok := overrides.Items.Properties.Get("thresholds"); ok {
				for field, exampleValues := range examples {
					if field != "path" {
						if prop, ok := thresholds.Properties.Get(field); ok {
							prop.Examples = exampleValues
						}
					}
				}
			}
		}
	}
}

// removeRequired recursively removes "required" from all schema nodes
func removeRequired(schema *jsonschema.Schema, isRoot bool) {
	if schema == nil {
		return
	}

	// Remove required array
	schema.Required = nil

	// Process properties
	if schema.Properties != nil {
		for pair := schema.Properties.Oldest(); pair != nil; pair = pair.Next() {
			removeRequired(pair.Value, false)
		}
	}

	// Process items (for arrays)
	if schema.Items != nil {
		removeRequired(schema.Items, false)
	}

	// Process additional properties
	if schema.AdditionalProperties != nil {
		removeRequired(schema.AdditionalProperties, false)
	}
}

func main() {
	// Generate schema from Go structs
	reflector := &jsonschema.Reflector{
		AllowAdditionalProperties:  false,
		DoNotReference:             true, // Inline all types instead of using $ref
		RequiredFromJSONSchemaTags: true, // Only mark fields as required if explicitly tagged
	}

	schema := reflector.Reflect(&config.Config{})

	// Post-process schema to remove "required" from all fields except PathOverride.path
	removeRequired(schema, true)

	// Set path as required in overrides
	if overrides, ok := schema.Properties.Get("overrides"); ok {
		if overrides.Items != nil {
			overrides.Items.Required = []string{"path"}
		}
	}

	// Add examples manually (invopop/jsonschema doesn't support examples in tags)
	addExamples(schema)

	// Set metadata
	schema.ID = jsonschema.ID("https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json")
	schema.Title = "Readability Configuration"
	schema.Description = "Configuration schema for readability markdown analyzer"

	// Marshal to JSON with indentation
	schemaBytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal schema: %v\n", err)
		os.Exit(1)
	}

	// Add trailing newline (required by pre-commit end-of-file-fixer)
	schemaBytes = append(schemaBytes, '\n')

	// Write to schema.json in current directory (where config.go lives, for embedding)
	pkgPath := "schema.json"
	if err := os.WriteFile(pkgPath, schemaBytes, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write schema to %s: %v\n", pkgPath, err)
		os.Exit(1)
	}
	fmt.Printf("✓ Generated schema at %s (%d bytes)\n", pkgPath, len(schemaBytes))

	// Copy to canonical docs location (relative to repo root)
	docsPath := filepath.Join("..", "..", "docs", "schemas", "config.json")
	if err := os.MkdirAll(filepath.Dir(docsPath), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(docsPath, schemaBytes, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write schema to %s: %v\n", docsPath, err)
		os.Exit(1)
	}
	fmt.Printf("✓ Copied schema to docs/schemas/config.json\n")
}
