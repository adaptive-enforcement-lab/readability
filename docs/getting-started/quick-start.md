# Quick Start

## Your First Analysis

Run readability on a directory:

```bash
readability docs/
```

Output:

```
┌─────────────────────────┬───────┬───────────┬─────────┐
│ File                    │ Grade │ Flesch    │ Status  │
├─────────────────────────┼───────┼───────────┼─────────┤
│ docs/index.md           │ 8.2   │ 62.5      │ pass    │
│ docs/getting-started.md │ 10.1  │ 55.3      │ pass    │
│ docs/api-reference.md   │ 14.5  │ 38.2      │ fail    │
└─────────────────────────┴───────┴───────────┴─────────┘
```

## Output Formats

### Markdown

```bash
readability --format markdown docs/
```

### JSON

```bash
readability --format json docs/
```

### Summary

```bash
readability --format summary docs/
```

## Check Mode

Fail if thresholds are exceeded:

```bash
readability --check --max-grade 12 docs/
```

Exit code 1 if any file exceeds grade level 12.

## Configuration File

Create `.content-analyzer.yml`:

```yaml
thresholds:
  max_grade: 12
  max_ari: 12
  max_lines: 500
```

Then run without flags:

```bash
readability docs/
```
