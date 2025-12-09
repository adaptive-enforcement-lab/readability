# Commands

Full reference for the `readability` command.

## Synopsis

```bash
readability [flags] <path>
```

The `path` argument is required. It can be a single file or a directory.

!!! info "Directory Scanning"
    When given a directory, the tool recursively finds all `.md` files.

## Output Flags

### --format, -f

Choose how results are displayed.

| Value | Description |
|-------|-------------|
| `table` | ASCII table (default) |
| `markdown` | Markdown table for GitHub |
| `json` | Machine-readable JSON |
| `summary` | Brief pass/fail summary |
| `report` | Detailed report with distribution |
| `diagnostic` | Linter-style `file:line:col` format |

**Example:**

```bash
readability -f markdown docs/
```

### --verbose, -v

Show all available metrics in the output, not just the defaults.

```bash
readability -v docs/
```

## Check Mode

### --check

Enable check mode. The command exits with code 1 if any file exceeds thresholds.

```bash
readability --check docs/
```

!!! warning "CI Usage"
    Always use `--check` in CI pipelines. Without it, the command exits 0 even when files fail.

## Threshold Flags

Override thresholds from the config file. Useful for testing different limits.

### --max-grade

Maximum Flesch-Kincaid grade level allowed.

```bash
readability --max-grade 12 docs/
```

### --max-ari

Maximum Automated Readability Index score allowed.

```bash
readability --max-ari 12 docs/
```

### --max-lines

Maximum lines per file. Set to 0 to disable this check.

```bash
readability --max-lines 500 docs/
```

### --min-admonitions

Minimum callout boxes (notes, warnings, tips) required per file. Set to 0 to disable.

```bash
readability --min-admonitions 2 docs/
```

## Configuration

### --config, -c

Path to a config file. By default, the tool looks for `.readability.yml` in the target directory or git root.

```bash
readability -c custom-config.yml docs/
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | All files pass (or check mode disabled) |
| 1 | One or more files failed thresholds |

## Examples

```bash
# Basic analysis with table output
readability docs/

# JSON output for scripting
readability -f json docs/ > results.json

# CI check with custom thresholds
readability --check --max-grade 10 --max-ari 10 docs/

# Diagnostic output for IDE integration
readability -f diagnostic docs/

# Verbose output showing all metrics
readability -v docs/
```
