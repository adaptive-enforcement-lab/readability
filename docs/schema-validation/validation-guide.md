# Validation Guide

This guide helps you troubleshoot and fix validation errors in `.readability.yml` files.

## Running Validation

### CLI Validation

The quickest way to validate your config:

```bash
readability --validate-config
```

**Success output**:
```
✓ Configuration is valid
```

**Error output**:
```
Configuration validation failed:

  • thresholds.max_grade
    got string, want number
    Suggestion: Remove quotes around numeric values

See https://github.com/adaptive-enforcement-lab/readability/blob/main/docs/cli/config-file.md for configuration reference.
```

### check-jsonschema

For more detailed validation output:

```bash
# Install check-jsonschema
pipx install check-jsonschema

# Validate config
check-jsonschema --schemafile docs/schemas/config.json .readability.yml
```

### IDE Validation

Most IDEs show validation errors inline:

- **Red squiggles**: Mark invalid values
- **Hover tooltip**: Show error message
- **Problems panel**: List all errors

## Common Errors

### Type Errors

#### String Instead of Number

**Error**:
```
got string, want number
```

**Cause**: Numeric values wrapped in quotes.

**Invalid**:
```yaml
thresholds:
  max_grade: "12"  # ❌ String
```

**Valid**:
```yaml
thresholds:
  max_grade: 12    # ✅ Number
```

#### String Instead of Integer

**Error**:
```
got string, want integer
```

**Cause**: Integer values wrapped in quotes.

**Invalid**:
```yaml
thresholds:
  max_lines: "500"  # ❌ String
```

**Valid**:
```yaml
thresholds:
  max_lines: 500    # ✅ Integer
```

#### Wrong Type for Complex Values

**Error**:
```
got string, want object
```

**Cause**: Object field has a simple value instead of nested structure.

**Invalid**:
```yaml
thresholds: "max_grade: 12"  # ❌ String
```

**Valid**:
```yaml
thresholds:                  # ✅ Object
  max_grade: 12
```

### Range Errors

#### Value Too High

**Error**:
```
value 200 is greater than the maximum 100
```

**Cause**: Numeric value exceeds schema maximum.

**Invalid**:
```yaml
thresholds:
  max_grade: 200  # ❌ Exceeds max of 100
```

**Valid**:
```yaml
thresholds:
  max_grade: 16   # ✅ Within 0-100 range
```

#### Value Too Low

**Error**:
```
value -50 is less than the minimum 0
```

**Cause**: Value below schema minimum.

**Invalid**:
```yaml
thresholds:
  max_lines: -10  # ❌ Below minimum of 1
```

**Valid**:
```yaml
thresholds:
  max_lines: 500  # ✅ Minimum is 1
```

!!! tip "Disabling Checks"
    Some fields support negative values to disable checks:
    - `min_ease: -100` disables ease check
    - `min_admonitions: -1` disables admonition requirement
    - `max_dash_density: -1` disables dash check

### Property Errors

#### Unknown Field

**Error**:
```
additional properties 'max_smog' not allowed
```

**Cause**: Field name doesn't exist in schema (typo or removed field).

**Invalid**:
```yaml
thresholds:
  max_grade: 12
  max_smog: 18      # ❌ Field doesn't exist
  max_syllables: 180  # ❌ Field doesn't exist
```

**Valid**:
```yaml
thresholds:
  max_grade: 12
  max_fog: 18       # ✅ Correct field name
  max_lines: 500    # ✅ Correct field name
```

**Tip**: Use IDE autocomplete (Ctrl+Space / Cmd+Space) to see available fields.

#### Missing Required Field

**Error**:
```
missing property 'path'
```

**Cause**: Required field omitted from override.

**Invalid**:
```yaml
overrides:
  - thresholds:     # ❌ Missing 'path'
      max_grade: 16
```

**Valid**:
```yaml
overrides:
  - path: docs/api/  # ✅ Required field present
    thresholds:
      max_grade: 16
```

### Structure Errors

#### Array Expected

**Error**:
```
got object, want array
```

**Cause**: Single object used instead of array.

**Invalid**:
```yaml
overrides:           # ❌ Missing array syntax
  path: docs/api/
  thresholds:
    max_grade: 16
```

**Valid**:
```yaml
overrides:           # ✅ Array with dash
  - path: docs/api/
    thresholds:
      max_grade: 16
```

Note the `-` before `path`. This indicates an array item.

#### Object Expected

**Error**:
```
got array, want object
```

**Cause**: Array used instead of object.

**Invalid**:
```yaml
thresholds:          # ❌ Array instead of object
  - max_grade: 12
  - max_ari: 12
```

**Valid**:
```yaml
thresholds:          # ✅ Object
  max_grade: 12
  max_ari: 12
```

### YAML Syntax Errors

#### Indentation Error

**Error**:
```
yaml: line 5: mapping values are not allowed in this context
```

**Cause**: Incorrect indentation (YAML is whitespace-sensitive).

**Invalid**:
```yaml
thresholds:
max_grade: 12        # ❌ Not indented
```

**Valid**:
```yaml
thresholds:
  max_grade: 12      # ✅ Properly indented (2 spaces)
```

#### Duplicate Keys

**Error**:
```
yaml: line 8: found duplicate key
```

**Cause**: Same field defined twice.

**Invalid**:
```yaml
thresholds:
  max_grade: 12
  max_grade: 16      # ❌ Duplicate
```

**Valid**:
```yaml
thresholds:
  max_grade: 12      # ✅ Single definition
```

For validation workflows and debugging strategies, see [Validation Workflow](validation-workflow.md).

## Next Steps

- [Validation Workflow](validation-workflow.md): Step-by-step validation process, pre-commit hooks, CI validation, and debugging strategies
- [Schema Reference](schema-reference.md): Complete field documentation
- [IDE Setup](ide-setup.md): Configure validation in your editor
- [Configuration File](../cli/config-file.md): Usage examples
