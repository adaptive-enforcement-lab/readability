# Installation

Pick the installation method that matches how you want to use Readability.

## GitHub Action

No installation needed. Add this step to any workflow file in `.github/workflows/`:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
```

!!! note "Version Pinning"
    Use `@v1` for the latest stable release. This automatically updates to new minor versions while staying on major version 1.

## Docker

Pull the official image from GitHub Container Registry:

```bash
docker pull ghcr.io/adaptive-enforcement-lab/readability:latest
```

Verify it works:

```bash
docker run --rm ghcr.io/adaptive-enforcement-lab/readability:latest --version
```

Analyze your documentation:

```bash
docker run --rm -v "$(pwd):/workspace" \
  ghcr.io/adaptive-enforcement-lab/readability:latest /workspace/docs/
```

!!! tip "Advanced Docker Usage"
    See the [Docker Guide](docker.md) for details on image tags, security verification, CI/CD examples, and volume mounting patterns.

## Pre-commit Hook

Pre-commit hooks run checks before each commit. This catches issues early, on your local machine.

### Step 1: Add the Hook

Create or update `.pre-commit-config.yaml` in your repository root:

```yaml
repos:
  - repo: https://github.com/adaptive-enforcement-lab/readability
    rev: v0.11.0  # Check releases for latest version
    hooks:
      - id: readability-docs
```

### Step 2: Install and Test

```bash
pre-commit install
pre-commit run readability-docs --all-files
```

### Available Hooks

| Hook ID | What It Checks |
|---------|----------------|
| `readability` | All markdown files in the commit |
| `readability-docs` | Only files in the `docs/` folder |

!!! tip "Configuration"
    Create a `.readability.yml` file to customize thresholds. See [Configuration File](../cli/config-file.md) for options.

## CLI Tool

The command-line tool lets you run checks manually or in scripts.

### Option 1: Go Install

If you have Go 1.21 or later:

```bash
go install github.com/adaptive-enforcement-lab/readability/cmd/readability@latest
```

Verify it works:

```bash
readability --version
```

### Option 2: Download Binary

Download a pre-built binary from the [releases page](https://github.com/adaptive-enforcement-lab/readability/releases).

**Available platforms:**

- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

### Option 3: Build from Source

```bash
git clone https://github.com/adaptive-enforcement-lab/readability.git
cd readability
go build -o readability ./cmd/readability
```

## Next Step

Continue to [Quick Start](quick-start.md) to run your first analysis.
