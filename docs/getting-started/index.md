# Getting Started

Get Readability running in under five minutes. Choose the method that fits your workflow.

## Choose Your Installation Method

Pick the method that works best for you.

### 🐳 Docker

Use Docker for CI/CD pipelines. Works in GitHub Actions, GitLab CI, and Jenkins.

No setup needed. Just pull and run. Every environment gets the same results.

[Docker Guide](docker.md)

### 📦 Binary Download

Download a single file. Run it on macOS, Linux, or Windows.

Fast and simple. No dependencies. Works offline.

[Installation Guide](installation.md#download-binary)

### 🔧 Go Install

Install from source using Go. Best for contributors and developers.

Get the latest code. Easy to modify and rebuild.

[Installation Guide](installation.md#go-install)

### 🚀 GitHub Action

Add one step to your workflow file. No installation needed.

Checks run on every pull request. Blocks merges when quality drops.

[GitHub Action Guide](../github-action/index.md)

!!! tip "New to Readability?"
    Start with **Docker** for CI/CD workflows or **Binary Download** for local use. Both methods get you running in under 2 minutes.

## Quick Preview

Here's what you'll get: a clear report showing which files need attention.

```
docs/api-reference.md:1:1: error: Grade 14.5 exceeds threshold 12.0
docs/getting-started.md:1:1: pass
docs/index.md:1:1: pass

3 files analyzed: 2 passed, 1 failed
```

## Next Steps

1. **[Installation](installation.md)** - Set up your preferred method
2. **[Quick Start](quick-start.md)** - Run your first analysis
