# Readability

Documentation readability analyzer - GitHub Action and CLI tool for measuring content quality metrics.

## Overview

Readability analyzes your documentation and provides actionable metrics to help you write clearer, more accessible content. Whether you're maintaining technical documentation, writing blog posts, or creating user guides, Readability helps ensure your content is easy to understand.

## Features

- **Flesch Reading Ease** - How easy is your content to read?
- **Grade Level Scores** - Flesch-Kincaid, Gunning Fog, Coleman-Liau, SMOG, ARI
- **Word & Sentence Metrics** - Count, averages, complexity indicators
- **Multiple Output Formats** - Table, Markdown, JSON, Summary, Report
- **Threshold Enforcement** - Fail CI when quality drops

## Quick Example

### GitHub Action

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    format: markdown
    check: true
    max-grade: 12
```

### CLI

```bash
# Analyze a directory
readability docs/

# Check with thresholds
readability --check --max-grade 12 docs/
```

## Next Steps

<div class="grid cards" markdown>

- :material-rocket-launch: **[Getting Started](getting-started/index.md)**

    Install and run your first analysis in minutes

- :material-github: **[GitHub Action](github-action/index.md)**

    Integrate readability checks into your CI/CD pipeline

- :material-console: **[CLI Reference](cli/index.md)**

    Full command-line interface documentation

- :material-chart-bar: **[Metrics](metrics/index.md)**

    Understand the readability scores and what they mean

</div>
