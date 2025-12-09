# Diagnostic Output

The diagnostic format produces linter-style output that IDEs and CI tools understand.

## The Format

```
file:line:column: severity: message (rule-id)
```

Each issue gets one line. This format works with VS Code, Vim, and most CI systems.

!!! example "Sample Output"
    ```
    docs/api.md:1:1: error: Grade 18.5 exceeds threshold 16.0 (readability/grade-level)
    docs/api.md:1:1: warning: Found 0 admonitions, minimum is 1 (content/admonitions)
    ```

## How to Use It

```bash
readability --format diagnostic docs/
```

For CI pipelines, add `--check` to fail on errors:

```bash
readability --check --format diagnostic docs/
```

## Rule IDs

Each issue has a rule ID you can reference:

| Rule ID | Level | What It Checks |
|---------|-------|----------------|
| `readability/grade-level` | error | Flesch-Kincaid grade |
| `readability/ari` | error | ARI score |
| `readability/gunning-fog` | error | Gunning Fog index |
| `readability/flesch-ease` | error | Reading ease score |
| `structure/max-lines` | error | File length |
| `content/admonitions` | warning | Callout boxes |

## Severity Levels

- **error** - Fails the check, blocks CI
- **warning** - Should fix, but won't block CI
- **info** - Informational only (future use)

## IDE Setup

### VS Code

Create `.vscode/tasks.json`:

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
          "regexp": "^(.+):(\\d+):(\\d+): (error|warning): (.+) \\((.+)\\)$",
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

Run with **Terminal > Run Task > Check Readability**. Issues appear in the Problems panel.

### Vim / Neovim

Add to your config:

```vim
set errorformat+=%f:%l:%c:\ %t%*[^:]:\ %m
```

Then run `:make` with the readability command.

## CI Integration

### GitHub Actions

```yaml
- name: Check docs
  uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    format: diagnostic
    check: true
```

### Generic CI

```bash
readability --check --format diagnostic docs/
```

The command exits with code 1 when errors are found. This fails the CI job.

## Pre-commit Hooks

The pre-commit hook uses diagnostic format by default:

```yaml
repos:
  - repo: https://github.com/adaptive-enforcement-lab/readability
    rev: v0.11.0
    hooks:
      - id: readability-docs
```
