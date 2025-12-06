# Readability

Documentation readability analyzer - GitHub Action and CLI tool for measuring content quality metrics.

## Features

- **Flesch Reading Ease** - How easy is your content to read?
- **Grade Level Scores** - Flesch-Kincaid, Gunning Fog, Coleman-Liau, SMOG, ARI
- **Word & Sentence Metrics** - Count, averages, complexity indicators
- **Multiple Output Formats** - Table, Markdown, JSON
- **Threshold Enforcement** - Fail CI when quality drops

## Quick Start

### GitHub Action

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    format: markdown
    check: true
    threshold-flesch: 30
```

### CLI

```bash
# Install
go install github.com/adaptive-enforcement-lab/readability/cmd/readability@latest

# Analyze a directory
readability --recursive docs/

# Check with thresholds
readability --check --threshold-flesch 30 docs/

# Output as JSON
readability --format json docs/
```

## Metrics

| Metric | Range | Interpretation |
|--------|-------|----------------|
| Flesch Reading Ease | 0-100 | Higher = easier (60-70 is standard) |
| Flesch-Kincaid Grade | 0-18+ | US grade level needed to understand |
| Gunning Fog Index | 0-20+ | Years of education needed |
| SMOG Index | 0-20+ | Years of education needed |
| Coleman-Liau Index | 0-20+ | US grade level |
| ARI | 0-20+ | US grade level |

## Configuration

Create `.readability.yml` in your repo:

```yaml
thresholds:
  flesch: 30
  grade: 12
  words: 3000

ignore:
  - "docs/api/**"
  - "CHANGELOG.md"
```

## Action Inputs

| Input | Description | Default |
|-------|-------------|---------|
| `path` | Path to analyze | `docs/` |
| `format` | Output format (table, markdown, json) | `markdown` |
| `check` | Fail on threshold violations | `false` |
| `threshold-flesch` | Minimum Flesch score | `30` |
| `threshold-grade` | Maximum grade level | `12` |
| `threshold-words` | Maximum words per file | `3000` |
| `recursive` | Analyze subdirectories | `true` |

## License

MIT
