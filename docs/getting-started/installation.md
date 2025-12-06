# Installation

## GitHub Action

The GitHub Action requires no installation - just add it to your workflow:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
```

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
