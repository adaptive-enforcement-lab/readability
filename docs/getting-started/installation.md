# Installation

## GitHub Action

The GitHub Action requires no installation - just add it to your workflow:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
```

## Pre-commit Hook

Add to your `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: https://github.com/adaptive-enforcement-lab/readability
    rev: 0.6.0  # Use latest release
    hooks:
      - id: readability
        # Optionally check only docs/ directory:
        # id: readability-docs
```

Then install and run:

```bash
pre-commit install
pre-commit run readability --all-files
```

### Available Hooks

| Hook ID | Description |
|---------|-------------|
| `readability` | Check all markdown files passed by pre-commit |
| `readability-docs` | Check only `docs/` directory (ignores filenames) |

### Configuration

Create `.readability.yml` in your repository root to configure thresholds. See [Configuration File](../cli/config-file.md) for details.

## CLI Tool

### Go Install (Recommended)

If you have Go installed:

```bash
go install github.com/adaptive-enforcement-lab/readability/cmd/readability@latest
```

### Binary Download

Download pre-built binaries from the [releases page](https://github.com/adaptive-enforcement-lab/readability/releases).

Available platforms:

- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

### From Source

```bash
git clone https://github.com/adaptive-enforcement-lab/readability.git
cd readability
go build -o readability ./cmd/readability
```
