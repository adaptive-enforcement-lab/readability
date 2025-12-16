# Getting Started

Get Readability running in under five minutes. Choose the method that fits your workflow.

## Choose Your Installation Method

<div class="grid cards" markdown>

- :material-docker: **Docker**

    ---

    Use Docker for CI/CD pipelines. Works in GitHub Actions, GitLab CI, and Jenkins.

    No setup needed. Just pull and run. Every environment gets the same results.

    [:octicons-arrow-right-24: Docker Guide](docker.md)

- :material-download: **Binary Download**

    ---

    Download a single file. Run it on macOS, Linux, or Windows.

    Fast and simple. No dependencies. Works offline.

    [:octicons-arrow-right-24: Installation Guide](installation.md#download-binary)

- :material-language-go: **Go Install**

    ---

    Install from source using Go. Best for contributors and developers.

    Get the latest code. Easy to modify and rebuild.

    [:octicons-arrow-right-24: Installation Guide](installation.md#go-install)

- :material-github: **GitHub Action**

    ---

    Add one step to your workflow file. No installation needed.

    Checks run on every pull request. Blocks merges when quality drops.

    [:octicons-arrow-right-24: GitHub Action Guide](../github-action/index.md)

</div>

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
