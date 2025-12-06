# Thresholds

Setting appropriate thresholds helps enforce documentation quality without being overly restrictive.

## Recommended Thresholds

### Technical Documentation

```yaml
thresholds:
  max_grade: 12
  max_ari: 14
  min_ease: 30
  max_lines: 500
```

### User Guides

```yaml
thresholds:
  max_grade: 10
  max_ari: 12
  min_ease: 50
  max_lines: 300
```

### API Reference

```yaml
thresholds:
  max_grade: 14
  max_ari: 16
  min_ease: 20
  max_lines: 1000
```

## Content-Specific Overrides

Different content types need different thresholds:

```yaml
thresholds:
  max_grade: 10

overrides:
  - pattern: "docs/api/**"
    thresholds:
      max_grade: 14

  - pattern: "docs/tutorials/**"
    thresholds:
      max_grade: 8
```

## When to Skip Checks

Set `min_words` to skip readability checks for very short documents:

```yaml
thresholds:
  min_words: 100
```

Documents with fewer than 100 words won't fail readability checks (the formulas are unreliable with sparse content).

## Line Limits

Use `max_lines` to encourage document splitting:

```yaml
thresholds:
  max_lines: 500
```

Long documents should typically be split into focused sections.
