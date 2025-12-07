# Readability

Documentation readability analyzer - GitHub Action and CLI tool for measuring content quality metrics.

## Features

- **Flesch Reading Ease** - How easy is your content to read?
- **Grade Level Scores** - Flesch-Kincaid, Gunning Fog, Coleman-Liau, SMOG, ARI
- **Word & Sentence Metrics** - Count, averages, complexity indicators
- **MkDocs Admonitions** - Detect and require `!!! note`, `!!! warning`, etc.
- **Multiple Output Formats** - Table, Markdown, JSON, Summary, Report
- **Threshold Enforcement** - Fail CI when quality drops
- **Job Summary** - Automatic GitHub Actions job summary with formatted report

## Quick Start

### GitHub Action

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    format: markdown
    check: true
    max-grade: 12
```

### CLI

```bash
# Install
go install github.com/adaptive-enforcement-lab/readability/cmd/readability@latest

# Analyze a directory
readability docs/

# Check with thresholds
readability --check --max-grade 12 docs/

# Output as JSON
readability --format json docs/

# Use a config file
readability --config .readability.yml docs/
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
  max_grade: 12
  max_ari: 12
  max_fog: 12
  min_ease: 30
  max_lines: 500
  min_words: 100
  min_admonitions: 1  # Require at least one MkDocs admonition

overrides:
  - path: docs/api/
    thresholds:
      max_grade: 14
      max_lines: 1000
      min_admonitions: -1  # Disable admonition requirement
```

## Action Inputs

| Input | Description | Default |
|-------|-------------|---------|
| `path` | Path to analyze (file or directory) | `docs/` |
| `format` | Output format (table, markdown, json, summary, report) | `markdown` |
| `config` | Path to config file | (auto-detect) |
| `check` | Fail on threshold violations | `false` |
| `max-grade` | Maximum Flesch-Kincaid grade level | (from config) |
| `max-ari` | Maximum ARI score | (from config) |
| `max-lines` | Maximum lines per file | (from config) |
| `summary` | Write formatted report to job summary | `true` |
| `summary-title` | Title for the job summary section | `Documentation Readability Report` |
| `version` | Version of readability to use | `latest` |

## Action Outputs

| Output | Description |
|--------|-------------|
| `report` | Analysis report in JSON format |
| `passed` | Whether all thresholds were met (`true`/`false`) |
| `files-analyzed` | Number of files analyzed |

## CLI Flags

| Flag | Description |
|------|-------------|
| `--format, -f` | Output format: table, json, markdown, summary, report, diagnostic |
| `--verbose, -v` | Show all metrics |
| `--check` | Check against thresholds (exit 1 on failure) |
| `--config, -c` | Path to config file |
| `--max-grade` | Maximum Flesch-Kincaid grade level |
| `--max-ari` | Maximum ARI score |
| `--max-lines` | Maximum lines per file (0 to disable) |
| `--min-admonitions` | Minimum MkDocs-style admonitions required (-1 to disable) |

## License

MIT
