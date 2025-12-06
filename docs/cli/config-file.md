# Configuration File

Readability automatically looks for `.readability.yml` in the target directory or git root.

## File Location

The tool searches for the config file in this order:

1. The directory being analyzed
2. Parent directories up to the git root
3. If not found, uses built-in defaults

You can also specify a config file explicitly:

```bash
readability --config /path/to/.readability.yml docs/
```

## Complete Example

```yaml
# .readability.yml
thresholds:
  max_grade: 12        # Maximum Flesch-Kincaid grade level
  max_ari: 12          # Maximum Automated Readability Index
  max_fog: 14          # Maximum Gunning Fog index
  min_ease: 30         # Minimum Flesch Reading Ease (0-100)
  max_lines: 500       # Maximum lines per file
  min_words: 100       # Skip readability check if below this word count
  min_admonitions: 1   # Require at least one MkDocs-style admonition

overrides:
  # API reference docs can be more technical
  - path: docs/api/
    thresholds:
      max_grade: 16
      max_ari: 16
      max_lines: 1000
      min_admonitions: 0  # API docs don't need admonitions

  # Tutorials should be beginner-friendly with callouts
  - path: docs/tutorials/
    thresholds:
      max_grade: 8
      max_ari: 8
      min_ease: 60
      min_admonitions: 2  # Tutorials benefit from more callouts

  # Reference docs with lots of lists/tables break formulas
  - path: docs/reference/
    thresholds:
      max_grade: 50
      max_ari: 50
      min_ease: -100  # Negative value disables this check
```

## Threshold Fields

| Field | Description | Default |
|-------|-------------|---------|
| `max_grade` | Maximum Flesch-Kincaid grade level | `16.0` |
| `max_ari` | Maximum Automated Readability Index | `16.0` |
| `max_fog` | Maximum Gunning Fog index | `18.0` |
| `min_ease` | Minimum Flesch Reading Ease score | `25.0` |
| `max_lines` | Maximum lines per file | `375` |
| `min_words` | Skip readability checks if word count is below this | `100` |
| `min_admonitions` | Minimum MkDocs-style admonitions required | `1` |

### Understanding the Defaults

The defaults target **college senior level** reading comprehension:

| Grade Level | Audience |
|-------------|----------|
| 6-8 | Middle school |
| 8-10 | MIL-STD-38784 (military technical manuals) |
| 10-12 | High school |
| 12-14 | College freshman/sophomore |
| 14-16 | College junior/senior |
| 16+ | Graduate level |

### Disabling Checks

Use extreme values to effectively disable specific checks:

```yaml
thresholds:
  max_grade: 50        # Effectively no grade limit
  min_ease: -100       # Negative values disable ease check
  max_lines: 0         # Zero disables line limit (via CLI only)
  min_admonitions: 0   # Disable admonition requirement
```

## Path Overrides

Use `overrides` to apply different thresholds to specific directories or files.

### Matching Rules

- Paths use **prefix matching**
- First matching override wins (order matters)
- Unmatched files use the base `thresholds`
- Paths are normalized (leading `./` and `../` stripped)

### Override Examples

```yaml
overrides:
  # Match all files under docs/api/
  - path: docs/api/
    thresholds:
      max_grade: 16

  # Match a specific file
  - path: docs/changelog.md
    thresholds:
      max_lines: 2000

  # More specific paths should come first
  - path: docs/guides/advanced/
    thresholds:
      max_grade: 14
  - path: docs/guides/
    thresholds:
      max_grade: 10
```

### Partial Overrides

Override only the thresholds you need; others inherit from the base:

```yaml
thresholds:
  max_grade: 12
  max_ari: 12
  max_lines: 500

overrides:
  - path: docs/api/
    thresholds:
      max_grade: 16  # Only override grade; ari and lines inherit
```

## CLI Overrides

Command-line flags take precedence over config file values:

```bash
# Config says max_grade: 12, but use 10 for this run
readability --max-grade 10 docs/

# Disable line limit for a single run
readability --max-lines 0 docs/

# Use a different config file
readability --config custom-config.yml docs/
```

## GitHub Action Usage

The GitHub Action automatically detects `.readability.yml`:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true  # Fail if thresholds exceeded
```

See [GitHub Action Configuration](../github-action/configuration.md) for more options.
