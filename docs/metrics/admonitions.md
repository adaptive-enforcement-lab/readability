# Admonitions

Readability checks for MkDocs-style admonitions to ensure documentation includes helpful callouts for notes, warnings, tips, and examples.

## What Are Admonitions?

Admonitions are visually distinct callout blocks that highlight important information. They use the `!!!` syntax popularized by MkDocs and its Material theme:

```markdown
!!! note "Optional Title"
    Content indented by 4 spaces.

!!! warning
    This is a warning without a custom title.
```

## Why Check for Admonitions?

Admonitions improve documentation quality by:

- Highlighting important information that readers might otherwise miss
- Breaking up walls of text with visual variety
- Providing contextual cues (warnings, tips, examples)
- Improving scannability for readers seeking specific information

## Supported Types

Readability detects these common admonition types:

| Type | Purpose |
|------|---------|
| `note` | Supplementary information |
| `warning` | Potential pitfalls or breaking changes |
| `tip` | Best practices or shortcuts |
| `example` | Code samples or use cases |
| `info` | Additional context |
| `danger` | Critical warnings |
| `abstract` | Summary or overview |
| `question` | FAQs or discussion points |

Any word following `!!!` is detected as an admonition type.

## Configuration

### Default Threshold

By default, readability requires **at least 1 admonition per file**. This encourages adding at least one helpful callout to each document.

### Config File

Set the minimum in `.readability.yml`:

```yaml
thresholds:
  min_admonitions: 1   # Require at least 1 admonition (default)
  # min_admonitions: 0 # Disable admonition check
  # min_admonitions: 2 # Require at least 2 admonitions
```

### CLI Override

Override via command line:

```bash
# Disable admonition check for this run
readability --min-admonitions 0 docs/

# Require at least 3 admonitions
readability --min-admonitions 3 docs/
```

### Path-Specific Overrides

Different directories may have different requirements:

```yaml
thresholds:
  min_admonitions: 1

overrides:
  # Changelog doesn't need admonitions
  - path: docs/changelog.md
    thresholds:
      min_admonitions: 0

  # Tutorials should have more callouts
  - path: docs/tutorials/
    thresholds:
      min_admonitions: 3
```

## Output

### JSON Output

When using `--format json`, admonition data is included:

```json
{
  "admonitions": {
    "count": 2,
    "types": ["note", "warning"]
  }
}
```

### Check Mode Warning

When `--check` fails due to missing admonitions, you'll see:

```
ADMONITIONS: Files are missing MkDocs-style admonitions (note, warning, tip, etc.).
Admonitions improve documentation by highlighting important information:
- Use !!! note for supplementary information
- Use !!! warning for potential pitfalls or breaking changes
- Use !!! tip for best practices or shortcuts
- Use !!! example for code samples or use cases

Example syntax:
  !!! note "Optional Title"
      Content indented by 4 spaces.

Do NOT add empty or meaningless admonitions. Add value with relevant context.
```

## Best Practices

### Do

- Use `!!! note` for supplementary context that adds value
- Use `!!! warning` for potential pitfalls or breaking changes
- Use `!!! tip` for best practices the reader might not discover on their own
- Use `!!! example` to show real-world usage

### Don't

- Add admonitions just to pass the check
- Use admonitions for routine information that belongs in regular text
- Overuse admonitions (they lose impact if everything is a callout)

## Syntax Variants

Readability detects all these formats:

```markdown
!!! note
    Basic admonition.

!!! warning "Custom Title"
    With a custom title.

!!! tip inline
    Inline modifier (title optional).

!!! note+
    Collapsible variant marker.
```
