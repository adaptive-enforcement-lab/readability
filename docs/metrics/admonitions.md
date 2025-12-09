# Admonitions

Admonitions are callout boxes. They highlight tips, warnings, and notes in your docs.

## What They Look Like

In MkDocs, you write them like this:

```markdown
!!! tip "Pro Tip"
    Use short sentences. They're easier to read.

!!! warning
    This breaks in version 2.0.
```

!!! note "Why Check for These?"
    Admonitions break up walls of text. They make docs easier to scan. This tool checks that you use at least one per file.

## Common Types

| Type | When to Use |
|------|-------------|
| `note` | Extra info that's good to know |
| `tip` | Best practices and shortcuts |
| `warning` | Things that might cause problems |
| `danger` | Critical issues to avoid |
| `example` | Code samples and use cases |
| `info` | Background context |

## Setting the Minimum

By default, each file needs at least one admonition.

**In your config:**

```yaml
thresholds:
  min_admonitions: 1  # Default
```

**From the command line:**

```bash
# Turn off the check
readability --min-admonitions 0 docs/

# Require at least 3
readability --min-admonitions 3 docs/
```

## Different Rules for Different Folders

Some content needs more callouts than others:

```yaml
thresholds:
  min_admonitions: 1

overrides:
  # Changelogs don't need callouts
  - path: docs/changelog.md
    thresholds:
      min_admonitions: 0

  # Tutorials need more
  - path: docs/tutorials/
    thresholds:
      min_admonitions: 3
```

## Do's and Don'ts

**Do:**

- Add `!!! tip` for shortcuts readers might miss
- Add `!!! warning` for common mistakes
- Add `!!! example` to show real usage

**Don't:**

- Add empty admonitions just to pass the check
- Put basic info in callouts (use regular text)
- Overuse them (too many loses impact)

!!! warning "Quality Over Quantity"
    The check counts admonitions. It can't judge if they're helpful. Don't game the metric with empty boxes.

## Syntax Guide

All of these work:

```markdown
!!! note
    Basic callout.

!!! warning "Watch Out"
    With a custom title.

!!! tip inline
    Flows with text.
```

The tool detects any word after `!!!` as a valid type.
