# Readability

Measure how easy your documentation is to read. Catch confusing content before your users do.

## Why Measure Readability?

Good documentation should be easy to understand. But as writers, we often miss when our own writing becomes too complex. Technical jargon creeps in. Sentences grow longer. Before you know it, your docs require a PhD to decode.

Readability analysis gives you objective feedback. It answers questions like:

- Is this page too complex for my audience?
- Which sections need simplification?
- Are my sentences too long?

!!! tip "The Goal"
    Most technical documentation should target a high school reading level. That's not "dumbing down" - it's respecting your reader's time.

## What You Get

This tool analyzes your Markdown files and reports:

| Metric | What It Measures |
|--------|------------------|
| **Grade Level** | School grade needed to understand the text |
| **Reading Ease** | How comfortable the text is to read (0-100 scale) |
| **Word Count** | Total words and reading time estimate |
| **Sentence Length** | Average words per sentence |

## Two Ways to Use It

### GitHub Action

Add readability checks to your pull requests. Catch problems before they merge.

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
```

### Command Line

Run checks locally while you write.

```bash
readability docs/
```

## Next Steps

<div class="grid cards" markdown>

- :material-rocket-launch: **[Getting Started](getting-started/index.md)**

    Install the tool and run your first check

- :material-github: **[GitHub Action](github-action/index.md)**

    Set up automated checks in your CI pipeline

- :material-console: **[CLI Reference](cli/index.md)**

    All command-line options and examples

- :material-chart-bar: **[Understanding Metrics](metrics/index.md)**

    Learn what each score means and how to improve it

</div>
