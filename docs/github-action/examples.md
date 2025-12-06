# Examples

## Report Only (No Failure)

Generate a report without failing the build:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    format: markdown
```

## Strict Enforcement

Fail if any document exceeds thresholds:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
    max-grade: 10
```

## JSON Output for Processing

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  id: readability
  with:
    path: docs/
    format: json

- name: Process results
  run: |
    echo "${{ steps.readability.outputs.report }}" | jq '.[] | select(.status == "fail")'
```

## Multiple Paths

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/api/

- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/guides/
    max-grade: 8  # Stricter for user guides
```

## With Configuration File

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    config: .readability.yml
    check: true
```
