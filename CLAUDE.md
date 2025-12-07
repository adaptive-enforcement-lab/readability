# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Test Commands

```bash
# Build
go build ./...

# Run tests
go test ./...

# Run a single test
go test ./pkg/config/... -run TestThresholdsForPath

# Install locally
go install ./cmd/readability

# Run CLI
go run ./cmd/readability docs/
go run ./cmd/readability --check --format json docs/
```

## Architecture

This is a Go CLI tool and GitHub Action that analyzes markdown documentation for readability metrics.

### Package Structure

- **cmd/readability** - CLI entry point using Cobra. Handles flag parsing, config loading, and output formatting.
- **pkg/analyzer** - Core analysis engine. `Analyzer.Analyze()` processes markdown and computes all metrics. Uses `textstats` library for readability formulas (Flesch-Kincaid, ARI, Gunning Fog, etc.).
- **pkg/config** - YAML config loading. Supports path-based threshold overrides via `ThresholdsForPath()`. Config file is `.readability.yml`.
- **pkg/markdown** - Goldmark-based parser. Extracts prose (excluding code blocks), headings, and MkDocs-style admonitions (`!!! note`, `!!! warning`, etc.).
- **pkg/output** - Output formatters (table, JSON, markdown, summary, report).

### Data Flow

1. CLI loads config (auto-detects `.readability.yml` or uses `--config`)
2. `Analyzer.AnalyzeFile()` reads markdown, calls `markdown.Parse()` to extract prose
3. `textstats` computes readability scores on the extracted prose
4. `checkStatus()` applies path-specific thresholds from config
5. Output formatter renders results

### GitHub Action

`action.yml` is a composite action that:
1. Downloads the pre-built binary from GitHub releases
2. Runs analysis with `--format json` internally
3. Converts to markdown for job summary
4. Sets outputs: `report`, `passed`, `files-analyzed`

## Key Implementation Details

### Path Override Matching

`ThresholdsForPath()` matches both relative and absolute paths. Override paths like `docs/developer-guide/` will match:
- `docs/developer-guide/test.md` (relative)
- `/home/runner/work/repo/docs/developer-guide/test.md` (absolute, CI environment)

### Threshold Merging

Zero values in override thresholds are treated as "not specified" and inherit from base. To explicitly disable a check in an override, use negative values:
- `min_ease: -100` to allow any reading ease score
- `min_admonitions: -1` to disable admonition requirement

### Skipping Readability Checks

Files with fewer than `min_words` (default 100) skip readability formula checks, as these formulas are unreliable for sparse prose. Line limits and admonition checks still apply.

## Release Process

Uses release-please for automated releases. Conventional commits trigger version bumps:
- `fix:` → patch version
- `feat:` → minor version
