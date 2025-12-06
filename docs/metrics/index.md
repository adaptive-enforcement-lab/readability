# Metrics

Readability calculates several standard readability metrics to help you understand how accessible your documentation is.

## Overview

| Metric | Range | What It Measures |
|--------|-------|------------------|
| Flesch Reading Ease | 0-100 | Overall readability (higher = easier) |
| Flesch-Kincaid Grade | 0-18+ | US grade level required |
| ARI | 0-20+ | US grade level (character-based) |
| Gunning Fog | 0-20+ | Years of education needed |
| SMOG | 0-20+ | Years of education (based on polysyllables) |
| Coleman-Liau | 0-20+ | US grade level (character-based) |

## Recommended Targets

| Audience | Grade Level | Flesch Ease |
|----------|-------------|-------------|
| General public | 6-8 | 60-70 |
| High school | 9-12 | 50-60 |
| Technical professionals | 12-14 | 30-50 |
| Academic/specialist | 14+ | 0-30 |

## Learn More

- [Flesch Reading Ease](flesch-reading-ease.md) - The most common readability score
- [Grade Level Scores](grade-level.md) - Understanding grade-level metrics
- [Thresholds](thresholds.md) - Setting appropriate limits
- [Admonitions](admonitions.md) - MkDocs-style callout detection
