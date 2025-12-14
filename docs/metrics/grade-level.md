# Grade Level Scores

These scores estimate what school grade can read your text. A score of 8 means an 8th grader should understand it.

## The Four Metrics

| Metric | What It Measures | Best For |
|--------|------------------|----------|
| **Flesch-Kincaid** | Words per sentence, syllables | General docs |
| **ARI** | Characters per word | Tech content |
| **Gunning Fog** | Long words (3+ syllables) | Business writing |
| **Coleman-Liau** | Character counts | Academic papers |

!!! tip "Which One?"
    Start with Flesch-Kincaid. It works well for most content. Add ARI if you write technical docs with code terms.

## Flesch-Kincaid Grade Level

The most common metric. It looks at sentence length and syllable count.

**What affects it:**

- Longer sentences raise the score
- Words with more syllables raise the score

**Target:** Grade 8-12 for most docs.

## ARI (Automated Readability Index)

Uses character count instead of syllables. This makes it better for technical content. Code terms often have many characters but few syllables.

**What affects it:**

- More characters per word raise the score
- Longer sentences raise the score

**Target:** Grade 10-14 for technical docs.

!!! info "Why ARI for Tech Docs?"
    Words like "config" and "param" are short in syllables but long in characters. ARI catches this better than other metrics.

## Gunning Fog Index

Counts "complex words": those with three or more syllables. Common endings like "-ing" and "-ed" don't count.

**What affects it:**

- More complex words raise the score
- Longer sentences raise the score

**Target:** Grade 12-14 for business content.

## Picking the Right Metric

| Your Content | Primary Metric | Secondary |
|--------------|----------------|-----------|
| User guides | Flesch-Kincaid | - |
| API docs | ARI | Flesch-Kincaid |
| Business docs | Gunning Fog | Flesch-Kincaid |
| Tutorials | Flesch-Kincaid | ARI |

!!! note "Multiple Metrics"
    You can check more than one. Use `max_grade` for Flesch-Kincaid and `max_ari` for ARI in your config.
