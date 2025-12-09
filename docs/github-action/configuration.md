# Configuration

All inputs and outputs for the GitHub Action.

## Inputs

| Input | What It Does | Default |
|-------|--------------|---------|
| `path` | Folder or file to check | `docs/` |
| `format` | Output format | `markdown` |
| `config` | Config file path | Auto-detect |
| `check` | Fail on violations | `false` |
| `max-grade` | Grade limit | From config |
| `max-ari` | ARI limit | From config |
| `max-lines` | Line limit | From config |
| `summary` | Show job summary | `true` |
| `summary-title` | Summary heading | `Documentation Readability Report` |

!!! note "Config File"
    The action finds `.readability.yml` in your repo root automatically. You don't need to set the `config` input unless using a different file.

## Outputs

Use outputs in later steps by adding `id` to the action:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  id: readability
  with:
    path: docs/

- run: echo "Passed: ${{ steps.readability.outputs.passed }}"
```

| Output | What It Contains |
|--------|------------------|
| `report` | Full results as JSON |
| `passed` | `true` or `false` |
| `files-analyzed` | Number of files checked |

## Full Example

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    format: markdown
    check: true
    max-grade: 12
    max-ari: 14
    max-lines: 500
    summary: true
    summary-title: Docs Check
```

## Job Summary

The action writes a report to your workflow's summary page. This shows:

| Column | Meaning |
|--------|---------|
| **File** | Which file |
| **Lines** | File length |
| **Read** | Reading time |
| **Grade** | [Grade level](../metrics/grade-level.md) score |
| **ARI** | [ARI](../metrics/grade-level.md#ari_automated_readability_index) score |
| **Ease** | [Reading ease](../metrics/flesch-reading-ease.md) score |
| **Status** | Pass or fail |

### Turn Off Summary

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    summary: false
```

### Custom Title

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    summary-title: "Docs Quality"
```
