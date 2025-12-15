# Component 1: Schema Creation

## Overview

Create a comprehensive JSON Schema file that describes the structure and validation rules for `.readability.yml` configuration files.

## Technical Approach

### JSON Schema Standard

**Selected**: JSON Schema Draft 2020-12
**Specification**: https://json-schema.org/draft/2020-12/schema

**Rationale**:
- Latest stable JSON Schema specification
- Broad tooling support (VS Code, IntelliJ, yaml-language-server)
- Rich validation features (types, ranges, patterns, conditionals)
- Industry standard for YAML/JSON configuration validation

### Schema File Location

```
schemas/readability-config.schema.json
```

**Rationale**:
- Standard convention used by most projects
- Separate from code to enable independent versioning
- Easy to reference from documentation
- Can be published to CDN/GitHub Pages

### Complete Schema Structure

Based on the Go structs in `pkg/config/config.go:12-33`, the schema will define:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://readability.adaptive-enforcement-lab.com/schemas/config.json",
  "title": "Readability Configuration",
  "description": "Configuration schema for readability markdown analyzer",
  "type": "object",
  "additionalProperties": false,
  "properties": {
    "thresholds": {
      "type": "object",
      "description": "Base readability thresholds applied to all files",
      "additionalProperties": false,
      "properties": {
        "max_grade": {
          "type": "number",
          "description": "Maximum Flesch-Kincaid grade level (12 = high school senior, 16 = college senior)",
          "minimum": 0,
          "maximum": 30,
          "default": 16,
          "examples": [12, 14, 16]
        },
        "max_ari": {
          "type": "number",
          "description": "Maximum Automated Readability Index (similar to grade level)",
          "minimum": 0,
          "maximum": 30,
          "default": 16,
          "examples": [12, 14, 16]
        },
        "max_fog": {
          "type": "number",
          "description": "Maximum Gunning Fog index (years of formal education needed)",
          "minimum": 0,
          "maximum": 30,
          "default": 18,
          "examples": [14, 16, 18]
        },
        "min_ease": {
          "type": "number",
          "description": "Minimum Flesch Reading Ease (0-100 scale, higher = easier). Use negative value to disable.",
          "minimum": -100,
          "maximum": 100,
          "default": 25,
          "examples": [30, 40, 50, -100]
        },
        "max_lines": {
          "type": "integer",
          "description": "Maximum lines of prose per file",
          "minimum": 1,
          "default": 375,
          "examples": [250, 375, 500]
        },
        "min_words": {
          "type": "integer",
          "description": "Minimum words before applying readability formulas (sparse docs are unreliable)",
          "minimum": 0,
          "default": 100,
          "examples": [50, 100, 150]
        },
        "min_admonitions": {
          "type": "integer",
          "description": "Minimum MkDocs-style admonitions required (!!! note, !!! warning). Use -1 to disable.",
          "minimum": -1,
          "default": 1,
          "examples": [0, 1, 2, -1]
        },
        "max_dash_density": {
          "type": "number",
          "description": "Maximum mid-sentence dash pairs per 100 sentences (detects AI-generated slop). Use -1 to disable. 0 = no dashes allowed.",
          "minimum": -1,
          "default": 0,
          "examples": [0, 2, 5, -1]
        }
      }
    },
    "overrides": {
      "type": "array",
      "description": "Path-specific threshold overrides (first match wins)",
      "items": {
        "type": "object",
        "required": ["path"],
        "additionalProperties": false,
        "properties": {
          "path": {
            "type": "string",
            "description": "Path prefix to match (e.g., 'docs/developer-guide/' or 'api/')",
            "minLength": 1,
            "examples": [
              "docs/developer-guide/",
              "docs/user-guide/",
              "api/",
              "README.md"
            ]
          },
          "thresholds": {
            "type": "object",
            "description": "Threshold overrides for this path (inherits unspecified values from base)",
            "additionalProperties": false,
            "properties": {
              "max_grade": {
                "$ref": "#/properties/thresholds/properties/max_grade"
              },
              "max_ari": {
                "$ref": "#/properties/thresholds/properties/max_ari"
              },
              "max_fog": {
                "$ref": "#/properties/thresholds/properties/max_fog"
              },
              "min_ease": {
                "$ref": "#/properties/thresholds/properties/min_ease"
              },
              "max_lines": {
                "$ref": "#/properties/thresholds/properties/max_lines"
              },
              "min_words": {
                "$ref": "#/properties/thresholds/properties/min_words"
              },
              "min_admonitions": {
                "$ref": "#/properties/thresholds/properties/min_admonitions"
              },
              "max_dash_density": {
                "$ref": "#/properties/thresholds/properties/max_dash_density"
              }
            }
          }
        }
      }
    }
  }
}
```

## Viability Analysis

### ✅ High Viability

**Evidence**:
1. **Mature Standard**: JSON Schema Draft 2020-12 is stable and widely adopted
2. **Tooling Ecosystem**: Excellent support across IDEs:
   - VS Code: Built-in via `yaml-language-server`
   - JetBrains IDEs: Native JSON Schema support
   - Vim/Neovim: Via LSP clients + yaml-language-server
   - Emacs: Via lsp-mode
3. **Go Validation Libraries**: Multiple mature options available (see Runtime Validation component)
4. **Industry Adoption**: Used by major projects:
   - ESLint: `.eslintrc` schema
   - GitHub Actions: `workflow.yml` schema
   - MkDocs Material: `mkdocs.yml` schema
   - Kubernetes: All resource schemas

### Schema-to-Code Mapping

The schema directly maps to Go structs with 1:1 correspondence:

| Go Struct Field | YAML Key | JSON Schema Type | Validation |
|-----------------|----------|------------------|------------|
| `Config.Thresholds` | `thresholds` | object | See nested properties |
| `Config.Overrides` | `overrides` | array[object] | Optional, items validated |
| `Thresholds.MaxGrade` | `max_grade` | number | min: 0, max: 30 |
| `Thresholds.MaxARI` | `max_ari` | number | min: 0, max: 30 |
| `Thresholds.MaxFog` | `max_fog` | number | min: 0, max: 30 |
| `Thresholds.MinEase` | `min_ease` | number | min: -100, max: 100 |
| `Thresholds.MaxLines` | `max_lines` | integer | min: 1 |
| `Thresholds.MinWords` | `min_words` | integer | min: 0 |
| `Thresholds.MinAdmonitions` | `min_admonitions` | integer | min: -1 |
| `Thresholds.MaxDashDensity` | `max_dash_density` | number | min: -1 |
| `PathOverride.Path` | `path` | string | minLength: 1, required |
| `PathOverride.Thresholds` | `thresholds` | object | Same as base thresholds |

### Special Handling: Negative Values

The Go code in `pkg/config/config.go:137-164` uses negative values to explicitly disable checks:
- `min_ease: -100` → allow any reading ease score
- `min_admonitions: -1` → disable admonition requirement
- `max_dash_density: -1` → disable dash density check

The schema accommodates this via minimum bounds (e.g., `"minimum": -100` for `min_ease`).

## Alternatives Considered

### Alternative 1: YAML Schema Language

**Approach**: Use YAML-specific schema language (e.g., Kwalify, Rx)

**Pros**:
- Native YAML syntax
- No JSON conversion needed

**Cons**:
- ❌ Poor tooling support (few IDE integrations)
- ❌ Smaller ecosystem
- ❌ Less mature than JSON Schema
- ❌ Not widely adopted

**Verdict**: **Rejected** - JSON Schema is industry standard with better tooling

### Alternative 2: Generate Schema from Go Structs

**Approach**: Use tools like `go-jsonschema` to auto-generate schema from Go struct tags

**Pros**:
- Single source of truth (Go code)
- Automatically stays in sync
- Reduces maintenance burden

**Cons**:
- ⚠️ Generated schemas often lack rich descriptions/examples
- ⚠️ Requires build step
- ⚠️ Less control over schema presentation
- ⚠️ May need manual post-processing

**Verdict**: **Consider for Phase 4** - Start with manual schema, automate later if maintenance burden increases

**Example Tool**: https://github.com/invopop/jsonschema

### Alternative 3: Embedded Schema Comments

**Approach**: Use Go struct tags or comments to embed schema metadata

**Pros**:
- Co-located with code
- No separate file to maintain

**Cons**:
- ❌ Doesn't produce standalone schema file
- ❌ Can't publish to SchemaStore
- ❌ Awkward syntax in Go comments
- ❌ Limited expressiveness

**Verdict**: **Rejected** - Doesn't solve the core problem of providing IDE support

### Alternative 4: Use OpenAPI/Swagger Schema

**Approach**: Repurpose OpenAPI schema definitions for config files

**Pros**:
- Similar to JSON Schema
- Tooling available

**Cons**:
- ❌ Designed for APIs, not config files
- ❌ Extra complexity (endpoints, responses, etc.)
- ❌ YAML language servers expect JSON Schema, not OpenAPI

**Verdict**: **Rejected** - Wrong tool for the job

## Implementation Guidance

### Step 1: Create Schema File

```bash
mkdir -p schemas
touch schemas/readability-config.schema.json
```

### Step 2: Implement Complete Schema

Use the complete schema structure provided above. Key considerations:

1. **Use `$ref`** for DRY - override thresholds reference base definitions
2. **Set `additionalProperties: false`** to catch typos in field names
3. **Provide rich descriptions** - these appear in IDE tooltips
4. **Include examples** - helps users understand valid values
5. **Set realistic bounds** - prevent nonsensical values (e.g., grade level 1000)

### Step 3: Validate Schema Itself

```bash
# Install JSON Schema validator
npm install -g ajv-cli

# Validate schema is well-formed
ajv compile -s schemas/readability-config.schema.json
```

### Step 4: Test with Sample Configs

Create test files:

```yaml
# test-valid.yml
# yaml-language-server: $schema=../schemas/readability-config.schema.json
thresholds:
  max_grade: 12
  min_ease: 40
```

```yaml
# test-invalid.yml
# yaml-language-server: $schema=../schemas/readability-config.schema.json
thresholds:
  max_grade: "twelve"  # Should error: expected number
  typo_field: 100      # Should error: additionalProperties=false
```

Open these files in VS Code and verify:
- Autocomplete works for all fields
- Invalid values show errors
- Tooltips display descriptions

## Testing Requirements

### Unit Tests

1. **Schema Validity**: Schema itself is valid JSON Schema Draft 2020-12
2. **Example Validation**: All example values in schema are valid
3. **Type Constraints**: Test each field's type requirements
4. **Range Validation**: Test min/max bounds for numeric fields
5. **Required Fields**: Test path overrides require `path` field

### Integration Tests

1. **IDE Validation**: Test schema loads in VS Code, IntelliJ
2. **YAML Language Server**: Verify yaml-language-server recognizes schema
3. **Real Config Files**: Validate repo's `.readability.yml` against schema
4. **Error Messages**: Verify helpful error messages for common mistakes

### Test Coverage Matrix

| Test Case | Expected Behavior |
|-----------|-------------------|
| Valid config with all fields | ✅ Passes validation |
| Valid config with minimal fields | ✅ Passes validation |
| Invalid type (string for number) | ❌ Type error |
| Out-of-range value | ❌ Range error |
| Typo in field name | ❌ Unknown property error |
| Missing required path in override | ❌ Required property error |
| Negative value for disabling check | ✅ Passes validation |

## Success Metrics

- ✅ Schema validates against JSON Schema meta-schema
- ✅ All Go struct fields represented in schema
- ✅ VS Code provides autocomplete for all fields
- ✅ Invalid configs show errors in real-time
- ✅ Tooltips display helpful descriptions
- ✅ All existing `.readability.yml` files validate successfully

## Next Steps

After schema creation:
1. Proceed to [Schema Publishing](02-schema-publishing.md) for distribution
2. Implement [Runtime Validation](04-runtime-validation.md) in Go code
3. Update [YAML Integration](03-yaml-integration.md) in example files
4. Create [Testing Strategy](05-testing-strategy.md) for comprehensive coverage

## References

- [JSON Schema Specification](https://json-schema.org/draft/2020-12/schema)
- [Understanding JSON Schema](https://json-schema.org/understanding-json-schema/)
- [JSON Schema Best Practices](https://json-schema.org/learn/miscellaneous-examples)
- [YAML Language Server](https://github.com/redhat-developer/yaml-language-server)
