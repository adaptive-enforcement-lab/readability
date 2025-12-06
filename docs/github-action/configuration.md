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
| `summary` | Write formatted report to job summary | `true` |
| `summary-title` | Title for the job summary section | `Documentation Readability Report` |

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
    summary: true
    summary-title: Documentation Readability Report
```

## Job Summary

The action automatically writes a formatted report to the GitHub Actions job summary. This is enabled by default (`summary: true`).

### Summary Table Columns

| Column | Description |
|--------|-------------|
| **File** | Path to the analyzed file |
| **Lines** | Number of lines in the file |
| **Read** | Estimated reading time (e.g., `<1m`, `2m`, `5m`) |
| **Grade** | [Flesch-Kincaid Grade Level](../metrics/grade-level.md#flesch-kincaid-grade-level) |
| **ARI** | [Automated Readability Index](../metrics/grade-level.md#ari-automated-readability-index) |
| **Fog** | [Gunning Fog Index](../metrics/grade-level.md#gunning-fog-index) |
| **Ease** | [Flesch Reading Ease](../metrics/flesch-reading-ease.md) score |
| **Status** | `pass` or `fail` based on configured thresholds |

### Disabling the Summary

To disable the job summary:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    summary: false
```

### Custom Title

Change the summary section title:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    summary-title: "Docs Quality Check"
```
