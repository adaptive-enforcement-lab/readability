# Schema Overrides and Validation

This page covers path-specific overrides, validation rules, and examples. You can customize thresholds for different file paths in your documentation.

!!! note "Field Definitions"
    For detailed information about individual threshold fields (max_grade, min_ease, etc.), see the [Schema Reference](schema-reference.md).

## Overrides Array

The `overrides` array lets you set different thresholds for specific file paths. When analyzing a file, the first matching override applies its thresholds.

### Override Object Structure

Each override has:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `path` | `string` | ✅ Yes | Path prefix to match (e.g., `docs/api/`, `README.md`) |
| `thresholds` | `object` | ❌ No | Threshold overrides (inherits unspecified values from base) |

**Example**:
```yaml
overrides:
  - path: docs/api/
    thresholds:
      max_grade: 18      # Override only grade
      min_admonitions: 0 # Override admonitions
      # Other thresholds inherit from base
```

### Path Matching Rules

Overrides match file paths using simple prefix matching. Here are the key rules to understand:

1. **Prefix Matching**: The system checks if a file path starts with the override path.
    - `docs/api/` matches `docs/api/rest.md` and `docs/api/auth/tokens.md`
    - Does NOT match `tutorials/api/intro.md`

2. **First Match Wins**: The system stops at the first matching override. Order matters.
    - Put specific paths before general paths

3. **Absolute and Relative**: Both path types work.
    - Matches `docs/api/rest.md` (relative path)
    - Also matches `/home/user/repo/docs/api/rest.md` (absolute path)

4. **Exact File Match**: You can target individual files.
    - `README.md` matches only `README.md`
    - `docs/README.md` matches only `docs/README.md`

**Example Order**:
```yaml
overrides:
  # Specific path first
  - path: docs/api/advanced/
    thresholds:
      max_grade: 20

  # General path second
  - path: docs/api/
    thresholds:
      max_grade: 16

  # If these were reversed, docs/api/advanced/ would match docs/api/ and never reach the specific rule
```

### Threshold Inheritance

When you use an override, the system merges it with your base thresholds. This means you only need to specify the values you want to change.

Override thresholds work this way:

- **Specified fields**: The override value replaces the base value
- **Unspecified fields**: The base threshold value is used
- **Zero values**: These are treated as "not specified" and inherit from base

**Example**:
```yaml
thresholds:
  max_grade: 12
  max_ari: 12
  max_fog: 14
  min_ease: 40

overrides:
  - path: docs/api/
    thresholds:
      max_grade: 18  # Override grade
      # max_ari inherits 12 from base
      # max_fog inherits 14 from base
      # min_ease inherits 40 from base
```

### Disabling Checks in Overrides

You can turn off specific checks for certain file paths. This is useful when some documentation has different requirements.

```yaml
overrides:
  - path: docs/reference/
    thresholds:
      max_grade: 100        # Effectively unlimited
      min_ease: -100        # Disable ease check
      min_admonitions: -1   # Disable admonition requirement
      max_dash_density: -1  # Disable dash check
```

## Complete Example

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
# Base thresholds for all files
thresholds:
  max_grade: 12
  max_ari: 12
  max_fog: 14
  min_ease: 40
  max_lines: 500
  min_words: 100
  min_admonitions: 1
  max_dash_density: 0

# Path-specific overrides
overrides:
  # API reference docs can be more complex
  - path: docs/api/
    thresholds:
      max_grade: 18
      max_fog: 20
      max_lines: 1000
      min_admonitions: 0

  # Tutorials should be simple
  - path: docs/tutorials/
    thresholds:
      max_grade: 10
      min_ease: 50

  # Advanced topics can be complex
  - path: docs/advanced/
    thresholds:
      max_grade: 16
      max_fog: 18

  # README has different standards
  - path: README.md
    thresholds:
      max_grade: 10
      max_lines: 250
      min_admonitions: 0
```

## Validation Rules

The schema automatically checks your configuration for errors. This helps catch mistakes before you commit your changes.

### Type Validation

The schema ensures each field has the correct data type. This prevents common configuration errors.

- Numbers must be numbers (not strings)
- Integers must be whole numbers
- Arrays must be arrays
- Objects must be objects

**Invalid**:
```yaml
thresholds:
  max_grade: "12"  # ❌ String, should be number
```

**Valid**:
```yaml
thresholds:
  max_grade: 12    # ✅ Number
```

### Range Validation

Every number field has limits to keep values in a reasonable range. The schema checks these automatically.

**Invalid**:
```yaml
thresholds:
  max_grade: 200   # ❌ Exceeds maximum of 100
  min_ease: 150    # ❌ Exceeds maximum of 100
```

**Valid**:
```yaml
thresholds:
  max_grade: 16    # ✅ Within 0-100 range
  min_ease: 40     # ✅ Within -100-100 range
```

### Required Fields

Each override needs a path to know which files it applies to. The `path` field is required. All other fields are optional.

**Invalid**:
```yaml
overrides:
  - thresholds:     # ❌ Missing required 'path'
      max_grade: 16
```

**Valid**:
```yaml
overrides:
  - path: docs/api/  # ✅ Required field present
    thresholds:
      max_grade: 16
```

### Additional Properties

The schema rejects unknown field names. This catches typos and prevents invalid configuration.

**Invalid**:
```yaml
thresholds:
  max_grade: 12
  unknown_field: true  # ❌ Not defined in schema
```

**Valid**:
```yaml
thresholds:
  max_grade: 12
  max_ari: 12      # ✅ Defined in schema
```

## Schema Evolution

The schema uses semantic versioning to manage changes. This helps you understand what kind of update you are getting.

- **Patch updates**: Bug fixes and documentation improvements. No breaking changes.
- **Minor updates**: New optional fields. Your existing configs will still work.
- **Major updates**: Breaking changes like field renames, removals, or type changes.

We recommend using `/latest/` in your schema reference. This gives you the most recent version automatically:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
```

## Next Steps

- [Schema Reference](schema-reference.md): Detailed field definitions and thresholds
- [IDE Setup](ide-setup.md): Configure your editor
- [Validation Guide](validation-guide.md): Common error examples and fixes
- [Validation Workflow](validation-workflow.md): Step-by-step validation process
- [Configuration File](../index.md): Examples and usage
