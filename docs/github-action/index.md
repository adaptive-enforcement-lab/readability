# GitHub Action

Check documentation readability on every pull request. No local setup needed.

## Quick Setup

Add this to `.github/workflows/docs.yml`:

```yaml
name: Check Docs

on:
  pull_request:
    paths:
      - 'docs/**'

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

!!! tip "What This Does"
    Every PR that changes files in `docs/` will run a readability check. If any file fails, the PR shows a red X.

## What You Get

The action adds a summary to your workflow run:

| File | Grade | ARI | Status |
|------|-------|-----|--------|
| docs/index.md | 8.2 | 9.1 | pass |
| docs/api.md | 15.3 | 17.2 | fail |

Click any run to see the full report.

## Features

- Works out of the box
- Shows results in GitHub's job summary
- Reads your `.readability.yml` config
- Fails the build when docs are too complex

## Learn More

- [Configuration](configuration.md) - All inputs and outputs
- [Examples](examples.md) - Common patterns
