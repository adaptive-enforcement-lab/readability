# Readability

[![CI](https://github.com/adaptive-enforcement-lab/readability/actions/workflows/ci.yml/badge.svg)](https://github.com/adaptive-enforcement-lab/readability/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/adaptive-enforcement-lab/readability/graph/badge.svg)](https://codecov.io/gh/adaptive-enforcement-lab/readability)
[![Go Report Card](https://goreportcard.com/badge/github.com/adaptive-enforcement-lab/readability)](https://goreportcard.com/report/github.com/adaptive-enforcement-lab/readability)
[![Go Reference](https://pkg.go.dev/badge/github.com/adaptive-enforcement-lab/readability.svg)](https://pkg.go.dev/github.com/adaptive-enforcement-lab/readability)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/11610/badge)](https://www.bestpractices.dev/projects/11610)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/adaptive-enforcement-lab/readability/badge)](https://scorecard.dev/viewer/?uri=github.com/adaptive-enforcement-lab/readability)

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

### Docker

```bash
# Pull the image
docker pull ghcr.io/adaptive-enforcement-lab/readability:latest

# Analyze local docs
docker run --rm -v "$(pwd):/workspace" ghcr.io/adaptive-enforcement-lab/readability:latest /workspace/docs

# With thresholds
docker run --rm -v "$(pwd):/workspace" ghcr.io/adaptive-enforcement-lab/readability:latest \
  --check --max-grade 12 /workspace/docs

# Verify image signature
cosign verify ghcr.io/adaptive-enforcement-lab/readability:latest \
  --certificate-identity-regexp 'https://github.com/adaptive-enforcement-lab/readability/.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
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
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade: 12       # Maximum Flesch-Kincaid grade level
  max_ari: 12         # Maximum Automated Readability Index
  max_fog: 12         # Maximum Gunning Fog index
  min_ease: 30        # Minimum Flesch Reading Ease (0-100 scale)
  max_lines: 500      # Maximum lines of prose per file
  min_words: 100      # Skip files with fewer words (formulas unreliable)
  min_admonitions: 1  # Require at least one MkDocs admonition

overrides:
  - path: docs/api/
    thresholds:
      max_grade: 14           # Allow more complexity for API docs
      max_lines: 1000         # API docs can be longer
      min_admonitions: -1     # Disable admonition requirement
```

Your editor will provide autocomplete and validation as you type. See the [Configuration Guide](https://readability.adaptive-enforcement-lab.com/latest/cli/config-file/#ide-support) for IDE setup.

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
