# Quick Start

Run your first readability analysis and understand the results.

## Your First Analysis

Point the tool at your documentation folder:

=== "Docker"

    ```bash
    docker run --rm -v "$(pwd):/workspace" \
      ghcr.io/adaptive-enforcement-lab/readability:latest /workspace/docs/
    ```

=== "Local Binary"

    ```bash
    readability docs/
    ```

You'll see a table like this:

```
┌─────────────────────────┬───────┬─────────┬─────────┐
│ File                    │ Grade │ Flesch  │ Status  │
├─────────────────────────┼───────┼─────────┼─────────┤
│ docs/index.md           │ 8.2   │ 62.5    │ pass    │
│ docs/getting-started.md │ 10.1  │ 55.3    │ pass    │
│ docs/api-reference.md   │ 14.5  │ 38.2    │ fail    │
└─────────────────────────┴───────┴─────────┴─────────┘
```

!!! info "Understanding the Scores"
    - **Grade**: School grade level needed (lower is simpler)
    - **Flesch**: Reading ease score (higher is easier)
    - **Status**: Pass/fail based on your thresholds

## Output Formats

Choose the format that fits your needs:

| Format | Command | Best For |
|--------|---------|----------|
| Table | `readability docs/` | Human reading |
| Markdown | `readability -f markdown docs/` | GitHub summaries |
| JSON | `readability -f json docs/` | Scripts and automation |
| Diagnostic | `readability -f diagnostic docs/` | IDE integration |

## Check Mode

Add `--check` to fail when thresholds are exceeded. This is useful for CI pipelines.

=== "Docker"

    ```bash
    docker run --rm -v "$(pwd):/workspace" \
      ghcr.io/adaptive-enforcement-lab/readability:latest \
      --check /workspace/docs/
    ```

=== "Local Binary"

    ```bash
    readability --check docs/
    ```

The command exits with code 1 if any file fails. Use this in your CI to block PRs with readability issues.

## Custom Thresholds

Override defaults from the command line:

=== "Docker"

    ```bash
    docker run --rm -v "$(pwd):/workspace" \
      ghcr.io/adaptive-enforcement-lab/readability:latest \
      --check --max-grade 12 --max-ari 12 /workspace/docs/
    ```

=== "Local Binary"

    ```bash
    readability --check --max-grade 12 --max-ari 12 docs/
    ```

Or create a `.readability.yml` file for persistent settings:

```yaml
thresholds:
  max_grade: 12
  max_ari: 12
  max_lines: 400
  min_admonitions: 1
```

!!! tip "Config File Location"
    Place `.readability.yml` in your repository root. The tool finds it automatically.

    When using Docker, the config file is automatically detected when you mount your workspace with `-v "$(pwd):/workspace"`.

## What's Next?

- **[CLI Reference](../cli/index.md)** - All command options
- **[Understanding Metrics](../metrics/index.md)** - What each score means
- **[GitHub Action](../github-action/index.md)** - Automate your checks
