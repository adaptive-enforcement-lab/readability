# Readability

Measure how easy your writing is to read. Get a score. Improve.

## New Here?

If you've never heard of "readability scores" before, start with these:

<div class="grid cards" markdown>

- :material-book-open-variant: **[What is Readability?](introduction.md)**

    How computers measure writing difficulty. The history, the math, and why it matters.

- :material-account-question: **[Who Is This For?](use-cases.md)**

    Find your situation: blogger, tech writer, docs team, student, or just curious.

</div>

## The Short Version

Readability formulas count words, sentences, and syllables. They output a grade level: the US school grade that can understand the text. A score of 8 means an eighth grader should follow it.

This tool runs those formulas on your Markdown files and tells you:

| Metric | What It Tells You |
|--------|-------------------|
| **Grade Level** | What school grade can read this |
| **Reading Ease** | How comfortable it is (0-100, higher = easier) |
| **Reading Time** | How long it takes at 200 words/minute |

!!! example "What the Output Looks Like"
    ```
    $ readability docs/

    ┌─────────────────┬───────┬──────┬───────┬───────┬───────┬──────┐
    │ File            │ Lines │ Read │ Grade │ ARI   │ Ease  │ Stat │
    ├─────────────────┼───────┼──────┼───────┼───────┼───────┼──────┤
    │ docs/index.md   │   42  │ <1m  │  7.8  │  8.9  │ 65.2  │ pass │
    │ docs/setup.md   │   89  │  2m  │ 11.2  │ 12.4  │ 48.1  │ pass │
    │ docs/api.md     │  156  │  4m  │ 14.8  │ 16.1  │ 29.3  │ fail │
    └─────────────────┴───────┴──────┴───────┴───────┴───────┴──────┘
    ```

## Two Ways to Run It

### In CI (GitHub Action)

Check every pull request automatically:

```yaml
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
```

PRs with overly complex docs fail the check. Writers fix issues before merge.

### Locally (Command Line)

Check files while you write:

```bash
readability docs/getting-started.md
```

See scores instantly. Revise and re-run until you're happy.

## Next Steps

<div class="grid cards" markdown>

- :material-rocket-launch: **[Getting Started](getting-started/index.md)**

    Install and run your first check in under 5 minutes

- :material-github: **[GitHub Action](github-action/index.md)**

    Automate checks in your CI pipeline

- :material-console: **[CLI Reference](cli/index.md)**

    All flags, options, and output formats

- :material-chart-bar: **[Understanding Metrics](metrics/index.md)**

    What each score means and how to improve it

</div>
