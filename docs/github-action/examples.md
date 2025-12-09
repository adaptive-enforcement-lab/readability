# Examples

Common patterns for using the GitHub Action.

## Report Only

Show results without failing the build:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
```

!!! tip "When to Use This"
    Good for getting started. See the reports, then add `check: true` once you've fixed issues.

## Block Bad Docs

Fail the PR if docs are too complex:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
```

## Custom Limits

Set your own grade level:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
    max-grade: 10
    max-ari: 10
```

## Process Results in Script

Get JSON output and use it:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  id: readability
  with:
    path: docs/
    format: json

- name: Show failed files
  run: |
    echo '${{ steps.readability.outputs.report }}' \
      | jq '.[] | select(.status == "fail")'
```

## Check Multiple Folders

Run separate checks with different rules:

```yaml
# User guides: strict
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/guides/
    check: true
    max-grade: 8

# API docs: relaxed
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/api/
    check: true
    max-grade: 14
```

## With Config File

Use a `.readability.yml` for settings:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
```

The action finds the config file automatically.

## Only on Doc Changes

Run only when docs change:

```yaml
name: Check Docs

on:
  pull_request:
    paths:
      - 'docs/**'
      - '.readability.yml'

jobs:
  readability:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: adaptive-enforcement-lab/readability@v1
        with:
          path: docs/
          check: true
```
