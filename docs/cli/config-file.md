# Configuration File

Readability automatically looks for `.content-analyzer.yml` in the target directory or git root.

## File Format

```yaml
thresholds:
  max_grade: 12      # Maximum Flesch-Kincaid grade
  max_ari: 12        # Maximum ARI score
  max_fog: 12        # Maximum Gunning Fog index
  min_ease: 30       # Minimum Flesch Reading Ease
  max_lines: 500     # Maximum lines per file
  min_words: 100     # Minimum words (skip readability check if below)

overrides:
  - pattern: "docs/api/**"
    thresholds:
      max_grade: 14    # Allow higher grade for API docs
      max_lines: 1000
  - pattern: "docs/tutorials/**"
    thresholds:
      max_grade: 8     # Stricter for tutorials
```

## Threshold Fields

| Field | Description |
|-------|-------------|
| `max_grade` | Maximum Flesch-Kincaid grade level |
| `max_ari` | Maximum ARI score |
| `max_fog` | Maximum Gunning Fog index |
| `min_ease` | Minimum Flesch Reading Ease score |
| `max_lines` | Maximum lines per file |
| `min_words` | Skip readability checks if word count is below this |

## Overrides

Use `overrides` to set different thresholds for specific paths:

```yaml
overrides:
  - pattern: "docs/api/**"
    thresholds:
      max_grade: 14
```

Patterns use glob syntax.

## CLI Overrides

CLI flags override config file values:

```bash
# Config says max_grade: 12, but use 10 for this run
readability --max-grade 10 docs/
```
