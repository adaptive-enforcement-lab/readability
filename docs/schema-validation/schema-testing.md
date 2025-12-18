# Schema Testing and Deployment

This guide covers testing and deploying JSON Schema changes.

!!! warning "Schema Development Guide"
    This page covers testing and deployment of schema changes. For information on adding, modifying, or removing schema fields, see [Maintaining the Schema](maintaining-schema.md).

## Validation Testing

Multiple test suites verify schema correctness.

### Schema Completeness Test

The test `pkg/config/schema_test.go:TestSchemaStructSync` verifies two things:

- All Go struct fields exist in schema
- Each field has `type` and `description`

Run the test:

```bash
go test ./pkg/config/ -run TestSchemaStructSync -v
```

### Schema Metaschema Test

This test validates the schema against JSON Schema Draft 2020-12 specification.

Run the test:

```bash
check-jsonschema --check-metaschema docs/schemas/config.json
```

### Runtime Validation Test

The test `pkg/config/validate_test.go` covers runtime validation. It tests three scenarios:

- Valid configs (should pass)
- Invalid configs (should fail with helpful errors)
- Edge cases (boundary values, missing fields, etc.)

Run the test:

```bash
go test ./pkg/config/ -run TestValidate -v
```

### Integration Test

This test ensures that config files in examples validate. Run it:

```bash
find . -name ".readability.yml" -exec check-jsonschema --schemafile docs/schemas/config.json {} \;
```

## Pre-commit Hooks

Pre-commit hooks run automatically before commits. They validate two things:

1. **Schema metaschema**:
    ```bash
    check-jsonschema --check-metaschema docs/schemas/config.json
    ```

2. **Config file**:
    ```bash
    check-jsonschema --schemafile docs/schemas/config.json .readability.yml
    ```

Set up pre-commit hooks:

```bash
pip install pre-commit
pre-commit install
```

Run hooks manually:

```bash
pre-commit run validate-json-schema --all-files
pre-commit run validate-readability-config --all-files
```

## CI Validation

The CI workflow includes a `validate-schema` job. It runs these steps:

```yaml
- name: Validate schema against JSON Schema Draft 2020-12
  run: check-jsonschema --check-metaschema docs/schemas/config.json

- name: Validate .readability.yml against schema
  run: check-jsonschema --schemafile docs/schemas/config.json .readability.yml
```

This validation runs on every PR.

## Performance Benchmarks

The file `pkg/config/validate_bench_test.go` measures validation performance.

Run benchmarks:

```bash
go test -bench=. -benchmem ./pkg/config/
```

These are the expected results:

- Schema compilation: < 5ns (cached via `sync.Once`)
- Valid config validation: < 10µs
- Invalid config validation: < 20µs
- Error formatting: < 20µs

Investigate if benchmarks regress by more than 2x.

## Schema Publishing

MkDocs publishes the schema to this URL:

```
https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
```

### How It Works

The publishing process has four steps:

1. **MkDocs build** includes `docs/schemas/` directory
2. **mike** deploys to gh-pages branch
3. **GitHub Pages** serves at `readability.adaptive-enforcement-lab.com`
4. `/latest/` points to current stable version

### Versioning

Schema URLs support two versioning styles:

- **Latest**: `/latest/schemas/config.json` (recommended)
- **Specific version**: `/v1.11.0/schemas/config.json`

Users should reference `/latest/` to get automatic updates.

### Testing Published Schema

Verify the schema is accessible after deployment:

```bash
curl -I https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
```

You should see this response:

```
HTTP/2 200
content-type: application/json
```

## Common Pitfalls

Watch out for these common mistakes.

### Forgetting to Update Overrides

When you add a field to `thresholds`, add a `$ref` in `overrides.items.properties.thresholds`:

```json
{
  "properties": {
    "thresholds": {
      "properties": {
        "new_field": { /* definition */ }
      }
    },
    "overrides": {
      "items": {
        "properties": {
          "thresholds": {
            "properties": {
              "new_field": {
                "$ref": "#/properties/thresholds/properties/new_field"
              }
            }
          }
        }
      }
    }
  }
}
```

### Missing Description

Every field needs a description for IDE tooltips:

```json
{
  "max_grade": {
    "type": "number",
    "description": "Maximum Flesch-Kincaid grade level"  // Required!
  }
}
```

### Wrong Type Mapping

Go types map to specific JSON Schema types:

| Go Type | JSON Schema Type |
|---------|------------------|
| `int`, `int64`, `uint` | `"integer"` |
| `float64` | `"number"` |
| `string` | `"string"` |
| `bool` | `"boolean"` |
| `[]T` | `"array"` with `items` |
| `struct` | `"object"` with `properties` |

Here is a wrong example:

```json
{
  "max_lines": {
    "type": "number"  // Wrong! Should be "integer"
  }
}
```

Here is the correct version:

```json
{
  "max_lines": {
    "type": "integer"
  }
}
```

### Forgetting Examples

Always provide 2-3 realistic examples:

```json
{
  "max_grade": {
    "type": "number",
    "examples": [12, 14, 16]  // Helps users understand typical values
  }
}
```

### Not Setting `additionalProperties: false`

This setting prevents typos from passing validation:

```json
{
  "properties": {
    "thresholds": {
      "type": "object",
      "additionalProperties": false,  // Reject unknown fields
      "properties": { /* ... */ }
    }
  }
}
```

## Review Checklist

Review this checklist before submitting a schema change:

- [ ] Schema matches Go struct fields (yaml tags)
- [ ] All fields have `type` and `description`
- [ ] Range constraints match validation logic
- [ ] Examples provided for each field
- [ ] Overrides section includes `$ref` for new fields
- [ ] `additionalProperties: false` set on objects
- [ ] `TestSchemaStructSync` passes
- [ ] Metaschema validation passes
- [ ] Runtime validation tests pass
- [ ] Benchmarks don't regress
- [ ] Documentation updated
- [ ] Pre-commit hooks pass
- [ ] CI validates successfully

## Getting Help

If you need help with a schema change, try these resources:

1. Check [JSON Schema documentation](https://json-schema.org/)
2. Look at existing field definitions for examples
3. Ask in PR review
4. Run tests early and often

## Next Steps

- [Maintaining the Schema](maintaining-schema.md): Adding, modifying, or removing schema fields
- [Schema Reference](schema-reference.md): Complete schema documentation
- [Validation Guide](validation-guide.md): Common error examples and fixes
- [Validation Workflow](validation-workflow.md): Step-by-step validation process
- [IDE Setup](ide-setup.md): Configure validation in editors
