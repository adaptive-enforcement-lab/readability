# Maintaining the Schema

Developer guide for maintaining the JSON Schema for `.readability.yml`.

!!! warning "For Contributors"
    This page is for developers maintaining the schema. For end-user documentation, see [Schema Reference](schema-reference.md).

## Schema Location

The canonical schema file is:

```
docs/schemas/config.json
```

This location is important because:

- MkDocs publishes it to `https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json`
- Runtime validation loads it from this path
- IDE language servers fetch it from the published URL

**Never move or rename this file**, as it would break all existing config files.

## Schema-Struct Synchronization

The schema **must** stay synchronized with Go structs in `pkg/config/`.

### Critical Rule

**When you add, remove, or modify a Go struct field, you MUST update the schema.**

### Affected Structs

- **`Thresholds`** → `properties.thresholds.properties`
- **`PathOverride`** → `properties.overrides.items.properties`
- **`Config`** → Top-level `properties`

### Sync Requirements

For each Go struct field, the schema needs:

1. **Field name**: Must match `yaml:` tag
2. **Type**: Must match Go type (`number`, `integer`, `string`, etc.)
3. **Range**: Must match validation logic
4. **Description**: Required for IDE tooltips
5. **Default**: Should match Go default or zero value
6. **Examples**: At least 2-3 realistic values

## Adding a New Field

### Step 1: Add to Go Struct

Example: Adding `max_complexity` field.

```go
// pkg/config/config.go
type Thresholds struct {
    // ... existing fields ...

    MaxComplexity int `yaml:"max_complexity,omitempty"`
}
```

### Step 2: Update Schema

Add the field to `docs/schemas/config.json`:

```json
{
  "properties": {
    "thresholds": {
      "properties": {
        "max_complexity": {
          "type": "integer",
          "description": "Maximum cyclomatic complexity score per document",
          "minimum": 1,
          "maximum": 100,
          "default": 20,
          "examples": [10, 15, 20]
        }
      }
    }
  }
}
```

### Step 3: Add to Override Schema

If the field should be overridable, add a `$ref`:

```json
{
  "properties": {
    "overrides": {
      "items": {
        "properties": {
          "thresholds": {
            "properties": {
              "max_complexity": {
                "$ref": "#/properties/thresholds/properties/max_complexity"
              }
            }
          }
        }
      }
    }
  }
}
```

### Step 4: Run Tests

The sync test will catch missing fields:

```bash
go test ./pkg/config/ -run TestSchemaStructSync
```

**Expected output**:
```
PASS
```

If the test fails:
```
--- FAIL: TestSchemaStructSync (0.00s)
    --- FAIL: TestSchemaStructSync/Thresholds.MaxComplexity (0.00s)
        schema_test.go:54: Schema missing property 'max_complexity'
```

Fix the schema and re-run tests.

### Step 5: Update Documentation

Add the field to:

- [Schema Reference](schema-reference.md)
- [Configuration File](../cli/config-file.md) examples
- README.md examples (if relevant)

### Step 6: Validate Schema

Run the metaschema validation:

```bash
# Install check-jsonschema
pipx install check-jsonschema

# Validate schema
check-jsonschema --check-metaschema docs/schemas/config.json
```

**Expected output**:
```
ok -- validation done, no errors
```

## Modifying an Existing Field

### Changing Type

**Breaking change**: requires major version bump.

```json
{
  "max_grade": {
    "type": "integer"  // Changed from "number"
  }
}
```

Update:
1. Schema type
2. Go struct type
3. All examples in docs
4. Bump major version

### Changing Range

**May be breaking** if tightening constraints.

```json
{
  "max_grade": {
    "maximum": 50  // Changed from 100
  }
}
```

If lowering maximum or raising minimum:
- Major version bump
- Update docs
- Notify users

If relaxing constraints (higher max, lower min):
- Minor version bump
- Update docs

### Changing Description

**Non-breaking**: patch version acceptable.

```json
{
  "max_grade": {
    "description": "Updated description with more details"
  }
}
```

Update only the description, patch version bump.

## Removing a Field

**Breaking change**: requires major version bump and deprecation period.

### Deprecation Process

1. **Mark as deprecated** (current version):
    ```json
    {
      "max_old_field": {
        "type": "number",
        "description": "DEPRECATED: Use max_new_field instead. Will be removed in v2.0.0",
        "deprecated": true
      }
    }
    ```

2. **Add migration guide** to docs

3. **Wait for next major version**

4. **Remove from schema**:
    ```json
    {
      // max_old_field removed
    }
    ```

5. **Remove from Go struct**

6. **Update all docs**

## Next Steps

- [Schema Testing](schema-testing.md): Testing and deployment
- [Schema Reference](schema-reference.md): Complete schema documentation
- [Validation Guide](validation-guide.md): Common error examples and fixes
- [Validation Workflow](validation-workflow.md): Step-by-step validation process
- [IDE Setup](ide-setup.md): Configure validation in editors
