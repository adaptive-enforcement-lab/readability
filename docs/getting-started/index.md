# Getting Started

This guide will help you get up and running with Readability in just a few minutes.

## Choose Your Method

Readability can be used in two ways:

1. **GitHub Action** - Integrate into your CI/CD pipeline for automated checks
2. **CLI Tool** - Run locally during development or in scripts

## Quick Start

### GitHub Action (Recommended)

Add to your workflow:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
```

### CLI Installation

```bash
go install github.com/adaptive-enforcement-lab/readability/cmd/readability@latest
```

Then run:

```bash
readability docs/
```

## What's Next?

- [Installation](installation.md) - Detailed installation options
- [Quick Start](quick-start.md) - Your first readability analysis
