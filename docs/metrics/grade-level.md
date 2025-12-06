# Grade Level Scores

Grade level metrics estimate the US school grade level required to understand the text.

## Metrics Comparison

| Metric | Focus | Best For |
|--------|-------|----------|
| Flesch-Kincaid | Sentence length, syllables | General use |
| ARI | Characters per word | Technical content |
| Gunning Fog | Complex words | Business writing |
| SMOG | Polysyllables | Healthcare/education |
| Coleman-Liau | Character counts | Academic assessment |

## Flesch-Kincaid Grade Level

The most common grade-level metric.

**Formula:**
```
0.39 × (words / sentences) + 11.8 × (syllables / words) - 15.59
```

## ARI (Automated Readability Index)

Uses characters instead of syllables, making it more consistent for technical content.

**Formula:**
```
4.71 × (characters / words) + 0.5 × (words / sentences) - 21.43
```

## Gunning Fog Index

Focuses on "complex words" (3+ syllables, excluding common suffixes).

**Formula:**
```
0.4 × ((words / sentences) + 100 × (complex words / words))
```

## Which to Use?

- **Technical documentation**: Use ARI or Flesch-Kincaid
- **Business content**: Use Gunning Fog
- **Healthcare/education**: Use SMOG
- **General purpose**: Use Flesch-Kincaid as primary
