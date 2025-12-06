# Configuration

## Inputs

| Input | Description | Default |
|-------|-------------|---------|
| `path` | Path to analyze (file or directory) | `docs/` |
| `format` | Output format (table, markdown, json, summary, report) | `markdown` |
| `config` | Path to config file | (auto-detect) |
| `check` | Fail on threshold violations | `false` |
| `max-grade` | Maximum Flesch-Kincaid grade level | (from config) |
| `max-ari` | Maximum ARI score | (from config) |
| `max-lines` | Maximum lines per file | (from config) |

## Outputs

| Output | Description |
|--------|-------------|
| `report` | Analysis report in specified format |
| `passed` | Whether all thresholds were met (true/false) |
| `files-analyzed` | Number of files analyzed |

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
