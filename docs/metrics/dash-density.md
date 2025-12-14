# Dash Density

Dash density measures mid-sentence dash patterns that often indicate AI-generated content.

## What It Detects

The tool counts three patterns:

1. **Space-hyphen-space** (`-`) - Example: "The system - which processes data - runs fast"
2. **Em-dash with spaces** (`—`) - Example: "Documentation — especially technical — needs clarity"
3. **Em-dash without spaces** (`—`) - Example: "This feature—unlike others—offers benefits"

!!! note "Why This Matters"
    AI-generated content frequently overuses mid-sentence dashes as a stylistic crutch. Human writers use them occasionally, but AI tends to insert them formulaically when adding parenthetical information.

## How It's Calculated

Dash density is the number of mid-sentence dash occurrences per 100 sentences.

**Example:**

```markdown
The system - which processes data - runs efficiently.
Documentation — especially technical docs — requires clarity.
This feature—unlike others—offers benefits.
```

This has 6 dashes in 3 sentences = **200 dashes per 100 sentences**.

## Default Threshold

By default, **no mid-sentence dashes are allowed** (`max_dash_density: 0`).

This strict default helps prevent AI slop patterns from entering your documentation.

**In your config:**

```yaml
thresholds:
  max_dash_density: 0  # Default - no dashes allowed
```

**From the command line:**

```bash
# Allow some dashes (5 per 100 sentences)
readability --max-dash-density 5 docs/

# Disable the check entirely
readability --max-dash-density -1 docs/
```

## How to Fix Violations

When you see a dash density error, rewrite the sentence using these techniques:

### Use Commas

**Before:**
```markdown
The system - which processes data quickly - runs efficiently.
```

**After:**
```markdown
The system, which processes data quickly, runs efficiently.
```

### Split Into Separate Sentences

**Before:**
```markdown
Documentation — especially technical documentation — requires clarity.
```

**After:**
```markdown
Documentation requires clarity. Technical documentation especially needs clear writing.
```

### Restructure to Avoid Parenthetical Constructions

**Before:**
```markdown
This feature — unlike traditional approaches — offers significant benefits.
```

**After:**
```markdown
Unlike traditional approaches, this feature offers significant benefits.
```

Or:

```markdown
This feature offers significant benefits compared to traditional approaches.
```

## Different Rules for Different Folders

You might want to allow dashes in some content:

```yaml
thresholds:
  max_dash_density: 0  # Strict by default

overrides:
  # Academic content might need some dashes
  - path: docs/research/
    thresholds:
      max_dash_density: 10

  # Disable check for legacy docs
  - path: docs/legacy/
    thresholds:
      max_dash_density: -1
```

## Do's and Don'ts

**Do:**

- Use commas for parenthetical information
- Split complex sentences into simpler ones
- Restructure to eliminate unnecessary asides

**Don't:**

- Use mid-sentence dashes as a replacement for proper sentence structure
- Insert parenthetical clauses that break reading flow
- Rely on dashes to connect loosely related ideas

## What Gets Excluded

The dash density check only examines **prose content**. The following are automatically excluded:

!!! success "Excluded Content"
    - **Frontmatter**: YAML (`---`) and TOML (`+++`) metadata at file start
    - **Code blocks**: All fenced code blocks (` ```yaml `) and inline code
    - **Tables**: All table content including headers and cells
    - **Lists**: Unordered (`- Item`) and ordered list markers and content
    - **Admonitions**: MkDocs-style callouts (`!!! note`) and their content

**Example:**

```markdown
---
title: Document - with dash
author: John - Doe
---

# This dash is DETECTED ❌
The system - which processes data - runs fast.

# These dashes are IGNORED ✓
- List item with - dash in content
| Column - 1 | Column - 2 |
!!! note "Title - with dash"
    Content with - dashes here
```yaml
code: value - with dash
```

Only the prose line ("The system...") counts toward dash density. Frontmatter, lists, tables, admonitions, and code are all excluded from analysis.

## AI Slop Indicators

High dash density correlates with:

- Formulaic sentence structure
- Excessive parenthetical information
- Reduced reading flow
- Potential AI-generated content

By keeping dash density at or near zero, you ensure your documentation maintains human-quality writing patterns.
