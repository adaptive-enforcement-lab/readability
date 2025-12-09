# Getting Started

Get Readability running in under five minutes. Choose the method that fits your workflow.

## Choose Your Path

| Method | Best For | Setup Time |
|--------|----------|------------|
| **GitHub Action** | Automated CI/CD checks | 2 minutes |
| **Pre-commit Hook** | Local checks before commits | 3 minutes |
| **CLI Tool** | Manual analysis, scripting | 2 minutes |

!!! tip "New to Readability?"
    Start with the **GitHub Action**. It requires no local installation and catches issues automatically on every pull request.

## Quick Preview

Here's what you'll get - a clear report showing which files need attention:

```
docs/api-reference.md:1:1: error: Grade 14.5 exceeds threshold 12.0
docs/getting-started.md:1:1: pass
docs/index.md:1:1: pass

3 files analyzed: 2 passed, 1 failed
```

## Next Steps

1. **[Installation](installation.md)** - Set up your preferred method
2. **[Quick Start](quick-start.md)** - Run your first analysis
