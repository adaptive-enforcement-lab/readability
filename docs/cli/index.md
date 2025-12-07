# CLI Reference

The Readability CLI provides full access to all analysis features from the command line.

## Installation

```bash
go install github.com/adaptive-enforcement-lab/readability/cmd/readability@latest
```

## Basic Usage

```bash
# Analyze a single file
readability docs/index.md

# Analyze a directory (recursive)
readability docs/

# Output as JSON
readability --format json docs/
```

## Quick Reference

| Flag | Short | Description |
|------|-------|-------------|
| `--format` | `-f` | Output format: table, json, markdown, summary, report, diagnostic |
| `--verbose` | `-v` | Show all metrics |
| `--check` | | Check against thresholds (exit 1 on failure) |
| `--config` | `-c` | Path to config file |
| `--max-grade` | | Maximum Flesch-Kincaid grade level |
| `--max-ari` | | Maximum ARI score |
| `--max-lines` | | Maximum lines per file |
| `--min-admonitions` | | Minimum MkDocs-style admonitions required |

## Next Steps

- [Commands](commands.md) - Detailed command reference
- [Configuration File](config-file.md) - Config file format
- [Diagnostic Output](diagnostic-output.md) - Linter-style output for IDE/CI integration
