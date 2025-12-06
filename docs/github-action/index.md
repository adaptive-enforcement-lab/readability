# GitHub Action

Integrate Readability into your CI/CD pipeline to automatically check documentation quality on every pull request.

## Basic Usage

```yaml
name: Documentation Quality

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

## Features

- Automatic Go installation and build
- Multiple output formats
- Threshold enforcement
- Configuration file support
- Automatic job summary with clickable metric links

## Next Steps

- [Configuration](configuration.md) - All available inputs and outputs
- [Examples](examples.md) - Common workflow patterns
