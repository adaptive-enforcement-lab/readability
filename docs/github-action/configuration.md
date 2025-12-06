# Configuration

## Inputs

| Input | Description | Default |
|-------|-------------|---------|
| `path` | Path to analyze (file or directory) | `docs/` |
| `format` | Output format (table, markdown, json, summary, report) | `markdown` |
| `config` | Path to config file | (auto-detect `.readability.yml`) |
| `check` | Fail on threshold violations | `false` |
| `max-grade` | Maximum Flesch-Kincaid grade level | (from config) |
| `max-ari` | Maximum ARI score | (from config) |
| `max-lines` | Maximum lines per file | (from config) |

## Outputs

| Output | Description |
|--------|-------------|
| `report` | Analysis report in JSON format |
| `passed` | Whether all thresholds were met (`true`/`false`) |
| `files-analyzed` | Number of files analyzed |

## Configuration File Auto-Detection

The action automatically detects `.readability.yml` in your repository root. You don't need to specify the `config` input unless using a different filename or location.

```yaml
# .readability.yml is auto-detected
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
```

## Using Outputs

Access the action outputs in subsequent steps:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  id: readability
  with:
    path: docs/
    check: false  # Don't fail, just report

- name: Check results
  run: |
    echo "Files analyzed: ${{ steps.readability.outputs.files-analyzed }}"
    echo "All passed: ${{ steps.readability.outputs.passed }}"

- name: Fail if readability issues
  if: steps.readability.outputs.passed == 'false'
  run: exit 1
```

## Example with All Options

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    format: markdown
    config: .readability.yml
    check: true
    max-grade: 12
    max-ari: 14
    max-lines: 500
```
