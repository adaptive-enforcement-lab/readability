# Schema Reference

Complete reference for the `.readability.yml` JSON Schema.

## Schema Metadata

- **$schema**: `https://json-schema.org/draft/2020-12/schema`
- **$id**: `https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json`
- **Version**: Draft 2020-12 (latest JSON Schema specification)

## Top-Level Structure

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:   # Base thresholds (object, optional)
  # ... threshold properties

overrides:    # Path-specific overrides (array, optional)
  - path: docs/api/
    thresholds:
      # ... override thresholds
```

## Thresholds Object

The `thresholds` object defines base readability requirements applied to all files (unless overridden).

!!! tip "Override Capabilities"
    All threshold fields can be overridden on a per-path basis. See [Schema Overrides and Validation](schema-overrides.md) for details on path-specific customization.

### max_grade

Maximum Flesch-Kincaid grade level.

| Property | Value |
|----------|-------|
| **Type** | `number` |
| **Range** | 0 to 100 |
| **Default** | 16 |
| **Examples** | `12`, `14`, `16` |

**Description**: School grade level needed to understand the content. Grade 12 = high school senior, 16 = college senior.

**Common Values**:
- `8-10`: Simple, accessible content (blog posts, tutorials)
- `10-12`: General technical documentation
- `12-14`: Standard technical content
- `14-16`: Advanced or API documentation
- `16-18`: Complex technical content

**Example**:
```yaml
thresholds:
  max_grade: 12  # High school level maximum
```

### max_ari

Maximum Automated Readability Index.

| Property | Value |
|----------|-------|
| **Type** | `number` |
| **Range** | 0 to 100 |
| **Default** | 16 |
| **Examples** | `12`, `14`, `16` |

**Description**: Similar to Flesch-Kincaid grade level but uses a different formula. Measures grade level needed to understand the text.

**Example**:
```yaml
thresholds:
  max_ari: 12
```

### max_fog

Maximum Gunning Fog index.

| Property | Value |
|----------|-------|
| **Type** | `number` |
| **Range** | 0 to 100 |
| **Default** | 18 |
| **Examples** | `14`, `16`, `18` |

**Description**: Years of formal education needed to understand the text on first reading. Emphasizes complex words (3+ syllables).

**Common Values**:
- `< 12`: Easily understood by the general public
- `12-14`: High school level
- `14-16`: College level
- `16-18`: College graduate level
- `> 18`: Extremely difficult

**Example**:
```yaml
thresholds:
  max_fog: 14  # Require high school readability
```

### min_ease

Minimum Flesch Reading Ease score.

| Property | Value |
|----------|-------|
| **Type** | `number` |
| **Range** | -100 to 100 |
| **Default** | 25 |
| **Examples** | `30`, `40`, `50`, `-100` |

**Description**: Readability score from 0-100, where higher = easier to read. Use `-100` to disable the check.

**Score Interpretation**:
- `90-100`: Very easy (5th grade)
- `80-90`: Easy (6th grade)
- `70-80`: Fairly easy (7th grade)
- `60-70`: Standard (8th-9th grade)
- `50-60`: Fairly difficult (10th-12th grade)
- `30-50`: Difficult (college level)
- `0-30`: Very difficult (college graduate)
- `< 0`: Extremely difficult

**Example**:
```yaml
thresholds:
  min_ease: 40   # Require fairly readable content
  # OR
  min_ease: -100 # Disable ease check
```

### max_lines

Maximum lines of prose per file.

| Property | Value |
|----------|-------|
| **Type** | `integer` |
| **Range** | 1 to 10000 |
| **Default** | 375 |
| **Examples** | `250`, `375`, `500` |

**Description**: Maximum number of prose lines (excluding code blocks, headings, and blank lines) allowed in a single markdown file.

**Rationale**: Long files are harder to navigate and maintain. Breaking content into smaller files improves discoverability.

**Example**:
```yaml
thresholds:
  max_lines: 500  # Allow up to 500 lines
```

### min_words

Minimum words before applying readability formulas.

| Property | Value |
|----------|-------|
| **Type** | `integer` |
| **Range** | 0 to 10000 |
| **Default** | 100 |
| **Examples** | `50`, `100`, `150` |

**Description**: Files with fewer words skip readability formula checks (Flesch-Kincaid, ARI, Fog, Ease) because these formulas are unreliable for sparse content. Line limits and admonition checks still apply.

**Example**:
```yaml
thresholds:
  min_words: 100  # Skip formulas for files < 100 words
```

### min_admonitions

Minimum MkDocs-style admonitions required.

| Property | Value |
|----------|-------|
| **Type** | `integer` |
| **Range** | -1 to 100 |
| **Default** | 1 |
| **Examples** | `0`, `1`, `2`, `-1` |

**Description**: Minimum number of MkDocs admonitions (`!!! note`, `!!! warning`, etc.) required in the file. Use `-1` to disable, `0` to allow files without admonitions but not require them, or positive numbers to enforce.

**Admonition Types**: `note`, `abstract`, `info`, `tip`, `success`, `question`, `warning`, `failure`, `danger`, `bug`, `example`, `quote`

**Rationale**: Admonitions highlight important information and improve scannability.

**Example**:
```yaml
thresholds:
  min_admonitions: 1   # Require at least one callout
  # OR
  min_admonitions: -1  # Disable check
```

### max_dash_density

Maximum mid-sentence dash pairs per 100 sentences.

| Property | Value |
|----------|-------|
| **Type** | `number` |
| **Range** | -1 to 500 |
| **Default** | 0 |
| **Examples** | `0`, `2`, `5`, `-1` |

**Description**: Detects excessive mid-sentence em-dashes (`—`), a common pattern in AI-generated "slop" content. Value represents maximum dash pairs per 100 sentences. Use `-1` to disable, `0` to disallow dashes entirely.

**What's Counted**: Dash pairs within a single sentence (for example, multiple dashes used for parenthetical information). Dashes at sentence start/end are ignored.

**Rationale**: Human writers rarely use mid-sentence dashes; AI models overuse them.

**Example**:
```yaml
thresholds:
  max_dash_density: 0   # No mid-sentence dashes allowed
  # OR
  max_dash_density: 2   # Allow up to 2 dash pairs per 100 sentences
  # OR
  max_dash_density: -1  # Disable check
```

For path-specific threshold overrides and validation rules, see [Schema Overrides and Validation](schema-overrides.md).

## Next Steps

- [Schema Overrides and Validation](schema-overrides.md): Path-specific overrides, examples, and validation rules
- [IDE Setup](ide-setup.md): Configure your editor
- [Validation Guide](validation-guide.md): Common error examples and fixes
- [Validation Workflow](validation-workflow.md): Step-by-step validation process
- [Configuration File](../index.md): Examples and usage
