# Diagnostic Output Format

The diagnostic output format provides linter-style output compatible with IDEs, editors, and CI tooling.

## Format

```
file:line:column: severity: message (rule-id)
```

Each issue is reported on a single line with:

| Field | Description |
|-------|-------------|
| `file` | Relative path to the file |
| `line` | Line number (1-based) |
| `column` | Column number (1-based, defaults to 1) |
| `severity` | Issue severity: `error`, `warning`, or `info` |
| `message` | Human-readable description of the issue |
| `rule-id` | Machine-readable rule identifier |

## Usage

```bash
readability --format diagnostic docs/
```

With check mode for CI:

```bash
readability --check --format diagnostic docs/
```

## Example Output

```
docs/api/reference.md:1:1: error: Flesch-Kincaid grade 18.5 exceeds threshold 16.0 (readability/grade-level)
docs/api/reference.md:1:1: error: ARI 20.4 exceeds threshold 16.0 (readability/ari)
docs/api/reference.md:1:1: warning: Found 0 admonitions, minimum required is 1 (content/admonitions)
docs/getting-started.md:1:1: error: 450 lines exceeds threshold 375 (structure/max-lines)

4 issue(s): 3 error(s), 1 warning(s)
```

## Rule IDs

| Rule ID | Severity | Description |
|---------|----------|-------------|
| `readability/grade-level` | error | Flesch-Kincaid grade level exceeds threshold |
| `readability/ari` | error | Automated Readability Index exceeds threshold |
| `readability/gunning-fog` | error | Gunning Fog index exceeds threshold |
| `readability/flesch-ease` | error | Flesch Reading Ease below threshold |
| `structure/max-lines` | error | File exceeds maximum line count |
| `content/admonitions` | warning | File has fewer admonitions than required |

## Severity Levels

- **error**: Threshold violations that cause check mode to fail
- **warning**: Issues that should be addressed but don't block CI
- **info**: Informational messages (reserved for future use)

## IDE Integration

The diagnostic format is designed to work with tools that parse compiler-style output.

### VS Code

Use the "Problems" panel with a task that runs readability:

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "Check Readability",
      "type": "shell",
      "command": "readability --format diagnostic docs/",
      "problemMatcher": {
        "owner": "readability",
        "fileLocation": ["relative", "${workspaceFolder}"],
        "pattern": {
          "regexp": "^(.+):(\\d+):(\\d+): (error|warning|info): (.+) \\((.+)\\)$",
          "file": 1,
          "line": 2,
          "column": 3,
          "severity": 4,
          "message": 5,
          "code": 6
        }
      }
    }
  ]
}
```

### Vim/Neovim

Add to your `errorformat`:

```vim
set errorformat+=%f:%l:%c:\ %t%*[^:]:\ %m
```

### Pre-commit

The diagnostic format is the default for pre-commit hooks:

```yaml
repos:
  - repo: https://github.com/adaptive-enforcement-lab/readability
    rev: v0.9.0
    hooks:
      - id: readability-docs
```

## CI Integration

### GitHub Actions

```yaml
- name: Check documentation readability
  uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    format: diagnostic
    check: true
```

### Generic CI

```bash
readability --check --format diagnostic docs/ || exit 1
```

The exit code is non-zero when any file fails thresholds, making it suitable for CI gates.
