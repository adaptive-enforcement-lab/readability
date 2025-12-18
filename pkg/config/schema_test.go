package config

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSchemaStructSync verifies the JSON Schema and Go structs stay in sync.
// This test catches drift when developers add fields to structs but forget to update the schema.
func TestSchemaStructSync(t *testing.T) {
	// Load schema
	schemaData, err := os.ReadFile("../../docs/schemas/config.json")
	require.NoError(t, err, "Failed to load schema file - run from pkg/config directory")

	var schema map[string]interface{}
	require.NoError(t, json.Unmarshal(schemaData, &schema))

	// Get schema properties
	props, ok := schema["properties"].(map[string]interface{})
	require.True(t, ok, "Schema should have properties object")

	// Check top-level Config struct fields
	assert.Contains(t, props, "thresholds", "Schema missing 'thresholds' (Config.Thresholds)")
	assert.Contains(t, props, "overrides", "Schema missing 'overrides' (Config.Overrides)")

	// Check all Thresholds struct fields exist in schema
	thresholdsProps := getNestedMap(t, props, "thresholds", "properties")
	thresholdsType := reflect.TypeOf(Thresholds{})

	for i := 0; i < thresholdsType.NumField(); i++ {
		field := thresholdsType.Field(i)
		yamlTag := field.Tag.Get("yaml")

		// Skip unexported or excluded fields
		if yamlTag == "" || yamlTag == "-" {
			continue
		}

		// Extract field name from yaml tag (handle "field,omitempty")
		yamlName := yamlTag
		for j, c := range yamlTag {
			if c == ',' {
				yamlName = yamlTag[:j]
				break
			}
		}

		t.Run("Thresholds."+field.Name, func(t *testing.T) {
			assert.Contains(t, thresholdsProps, yamlName,
				"Schema missing property '%s' (from Go struct field Thresholds.%s with yaml tag '%s')",
				yamlName, field.Name, yamlTag)

			// Verify the field has required schema metadata
			if fieldProps, ok := thresholdsProps[yamlName].(map[string]interface{}); ok {
				assert.Contains(t, fieldProps, "type",
					"Schema field '%s' missing 'type'", yamlName)
				assert.Contains(t, fieldProps, "description",
					"Schema field '%s' missing 'description' (needed for IDE tooltips)", yamlName)
			}
		})
	}

	// Check PathOverride structure
	overridesProps := getNestedMap(t, props, "overrides")
	assert.Equal(t, "array", overridesProps["type"], "overrides should be array type")

	overrideItems := getNestedMap(t, overridesProps, "items", "properties")
	assert.Contains(t, overrideItems, "path", "Override schema missing 'path' property")
	assert.Contains(t, overrideItems, "thresholds", "Override schema missing 'thresholds' property")
}

// TestSchemaMetadata verifies the schema has proper metadata for IDE support.
func TestSchemaMetadata(t *testing.T) {
	schemaData, err := os.ReadFile("../../docs/schemas/config.json")
	require.NoError(t, err)

	var schema map[string]interface{}
	require.NoError(t, json.Unmarshal(schemaData, &schema))

	// Check required top-level metadata
	assert.Equal(t, "https://json-schema.org/draft/2020-12/schema", schema["$schema"],
		"Schema should use JSON Schema Draft 2020-12")

	assert.Contains(t, schema, "$id", "Schema should have $id for referencing")
	assert.Contains(t, schema, "title", "Schema should have title for IDE display")
	assert.Contains(t, schema, "description", "Schema should have description")

	// Check $id uses the published URL
	idStr, ok := schema["$id"].(string)
	require.True(t, ok, "$id should be a string")
	assert.Contains(t, idStr, "readability.adaptive-enforcement-lab.com",
		"$id should reference published schema location")
}

// TestSchemaRequiredFields verifies required fields are marked correctly.
func TestSchemaRequiredFields(t *testing.T) {
	schemaData, err := os.ReadFile("../../docs/schemas/config.json")
	require.NoError(t, err)

	var schema map[string]interface{}
	require.NoError(t, json.Unmarshal(schemaData, &schema))

	props := schema["properties"].(map[string]interface{})

	// PathOverride items should require 'path' field
	overridesProps := getNestedMap(t, props, "overrides", "items")
	required, ok := overridesProps["required"].([]interface{})
	require.True(t, ok, "Override items should have required array")

	hasPath := false
	for _, r := range required {
		if r == "path" {
			hasPath = true
			break
		}
	}
	assert.True(t, hasPath, "Override items should require 'path' field")
}

// getNestedMap is a helper to navigate nested map structures safely.
func getNestedMap(t *testing.T, m map[string]interface{}, keys ...string) map[string]interface{} {
	t.Helper()
	current := m
	for i, key := range keys {
		next, ok := current[key].(map[string]interface{})
		require.True(t, ok, "Failed to navigate to %v at key '%s' (step %d/%d)",
			keys, key, i+1, len(keys))
		current = next
	}
	return current
}
