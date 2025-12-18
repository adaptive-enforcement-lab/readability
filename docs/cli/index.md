# CLI Reference

Analyze Markdown files from the command line.

## Install

```bash
go install github.com/adaptive-enforcement-lab/readability/cmd/readability@latest
```

!!! note "Other Options"
    See [Installation](../getting-started/installation.md) for binaries and building from source.

## Basic Usage

```bash
# Check a file
readability README.md

# Check a folder
readability docs/

# Fail if issues found (for CI)
readability --check docs/
```

## Options

| Flag | Short | Purpose |
|------|-------|---------|
| `--format` | `-f` | Output format |
| `--check` | | Exit 1 on failures |
| `--config` | `-c` | Config file path |
| `--verbose` | `-v` | Show all metrics |

## Output Formats

| Format | Best For |
|--------|----------|
| `table` | Reading in terminal |
| `markdown` | GitHub summaries |
| `json` | Scripts |
| `diagnostic` | IDEs and linters |

## Threshold Flags

Set limits from the command line:

```bash
readability --max-grade 12 --max-ari 12 docs/
```

| Flag | Controls |
|------|----------|
| `--max-grade` | Grade level limit |
| `--max-ari` | ARI score limit |
| `--max-lines` | File length limit |
| `--min-admonitions` | Required callouts |

## More Info

- [Commands](commands.md) - All flags
- [Config File](../configuration/index.md) - Save settings
- [Diagnostic Output](diagnostic-output.md) - IDE setup
