# Commands

## readability [path]

Analyze markdown files for readability metrics.

### Arguments

| Argument | Description |
|----------|-------------|
| `path` | File or directory to analyze (required) |

### Flags

#### Output Options

`--format, -f`
:   Output format. Options: `table`, `json`, `markdown`, `summary`, `report`, `diagnostic`. Default: `table`. See [Diagnostic Output](diagnostic-output.md) for linter-style format details.

`--verbose, -v`
:   Show all available metrics in output.

#### Threshold Options

`--check`
:   Enable check mode. Exit with code 1 if any file fails thresholds.

`--max-grade`
:   Maximum Flesch-Kincaid grade level. Overrides config file.

`--max-ari`
:   Maximum ARI score. Overrides config file.

`--max-lines`
:   Maximum lines per file. Set to 0 to disable. Overrides config file.

`--min-admonitions`
:   Minimum MkDocs-style admonitions required. Set to 0 to disable. Overrides config file.

#### Configuration

`--config, -c`
:   Path to configuration file. Default: auto-detect `.readability.yml`.

### Examples

```bash
# Basic analysis
readability docs/

# JSON output
readability --format json docs/

# Diagnostic output (linter-style)
readability --format diagnostic docs/

# Check with custom thresholds
readability --check --max-grade 10 --max-ari 12 docs/

# Verbose output
readability -v docs/

# Use specific config
readability -c .readability.yml docs/
```

### Exit Codes

| Code | Description |
|------|-------------|
| 0 | Success (all files pass or check mode disabled) |
| 1 | Failure (one or more files failed thresholds in check mode) |
