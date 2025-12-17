# Configuration File

Store your thresholds in a `.readability.yml` file instead of passing flags every time.

## Quick Start

Create `.readability.yml` in your repository root:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade: 12
  max_ari: 12
```

That's it. The tool finds it automatically.

!!! tip "Where to Put It"
    Place the config file in your repository root. The tool searches from the target directory up to the git root.

## All Options

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade: 12       # Flesch-Kincaid grade level
  max_ari: 12         # Automated Readability Index
  max_fog: 14         # Gunning Fog index
  min_ease: 40        # Flesch Reading Ease (0-100)
  max_lines: 400      # Lines per file
  min_words: 100      # Skip check if fewer words
  min_admonitions: 1  # Required callout boxes
  max_dash_density: 0 # Mid-sentence dash pairs per 100 sentences
```

## What Each Threshold Means

| Option | What It Controls | Default |
|--------|------------------|---------|
| `max_grade` | School grade needed to read | 16 |
| `max_ari` | Similar to grade, different formula | 16 |
| `max_fog` | Complexity from long words | 18 |
| `min_ease` | Comfort level (higher = easier) | 25 |
| `max_lines` | File length limit | 375 |
| `min_words` | Skip short files | 100 |
| `min_admonitions` | Notes, tips, warnings needed | 1 |
| `max_dash_density` | Mid-sentence dashes per 100 sentences (prevents AI slop) | 0 |

!!! info "Grade Level Scale"
    A grade of 12 means "high school senior" level. Most technical docs should target grades 10-14.

## Different Rules for Different Folders

Use `overrides` to apply stricter or looser rules to specific paths:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade: 12

overrides:
  # API docs can be more technical
  - path: docs/api/
    thresholds:
      max_grade: 16
      min_admonitions: 0

  # Tutorials should be simple
  - path: docs/tutorials/
    thresholds:
      max_grade: 8
```

### How Path Matching Works

- Paths match from the start (prefix matching)
- First matching rule wins
- Put specific paths before general ones
- Unmatched files use the base thresholds

**Example order:**

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
overrides:
  # Specific path first
  - path: docs/api/advanced/
    thresholds:
      max_grade: 18

  # General path second
  - path: docs/api/
    thresholds:
      max_grade: 16
```

## Disabling Checks

Set extreme values to skip specific checks:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade: 100        # No grade limit
  min_ease: -100        # No ease requirement
  max_lines: 0          # No line limit (CLI only)
  min_admonitions: 0    # No admonition requirement
  max_dash_density: -1  # No dash density check
```

## Command Line Overrides

Flags override config file values for a single run:

```bash
# Use grade 10 instead of config value
readability --max-grade 10 docs/

# Use a different config file
readability --config strict.yml docs/
```

## With GitHub Actions

The action finds `.readability.yml` automatically:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
```

See [GitHub Action Configuration](../github-action/configuration.md) for action-specific options.
