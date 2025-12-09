# Thresholds

Thresholds are your limits. Set them too strict and everything fails. Too loose and bad docs slip through.

## Starting Points

Pick based on who reads your docs.

### For Technical Docs

```yaml
thresholds:
  max_grade: 12
  max_ari: 14
  min_ease: 30
  max_lines: 500
```

!!! tip "This Is the Default"
    These values work for most developer docs. Start here and adjust as needed.

### For User Guides

```yaml
thresholds:
  max_grade: 10
  max_ari: 12
  min_ease: 50
  max_lines: 300
```

Stricter limits because non-technical readers need simpler text.

### For API Reference

```yaml
thresholds:
  max_grade: 14
  max_ari: 16
  min_ease: 20
  max_lines: 1000
```

Looser limits because expert readers can handle dense content.

## Different Rules Per Folder

Not all docs are alike. Use overrides:

```yaml
thresholds:
  max_grade: 10

overrides:
  # API docs can be harder
  - path: docs/api/
    thresholds:
      max_grade: 14

  # Tutorials must be easy
  - path: docs/tutorials/
    thresholds:
      max_grade: 8
```

!!! note "Path Matching"
    Paths are relative to your repo root. End folder paths with `/` to match all files inside.

## Skip Short Files

Very short docs give bad scores. The formulas need enough text to work well.

```yaml
thresholds:
  min_words: 100
```

Files under 100 words won't fail checks.

## Limit File Length

Long files are hard to navigate. Set a line limit:

```yaml
thresholds:
  max_lines: 500
```

!!! warning "Split Long Docs"
    If a file hits this limit, break it into smaller focused pages. One topic per page works best.

## All Available Thresholds

| Setting | What It Limits | Default |
|---------|----------------|---------|
| `max_grade` | Flesch-Kincaid grade level | 12 |
| `max_ari` | Automated Readability Index | 14 |
| `min_ease` | Flesch Reading Ease (higher = easier) | 30 |
| `max_fog` | Gunning Fog Index | 16 |
| `max_lines` | Total lines in file | 500 |
| `min_words` | Minimum words to check | 50 |
| `min_admonitions` | Required callout boxes | 1 |
