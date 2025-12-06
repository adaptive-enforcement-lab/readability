# Flesch Reading Ease

The Flesch Reading Ease score is one of the most widely used readability metrics. Higher scores indicate easier-to-read content.

## Score Interpretation

| Score | Difficulty | Typical Audience |
|-------|------------|------------------|
| 90-100 | Very Easy | 5th grade |
| 80-90 | Easy | 6th grade |
| 70-80 | Fairly Easy | 7th grade |
| 60-70 | Standard | 8th-9th grade |
| 50-60 | Fairly Difficult | 10th-12th grade |
| 30-50 | Difficult | College |
| 0-30 | Very Difficult | College graduate |

## Formula

```
206.835 - 1.015 × (words / sentences) - 84.6 × (syllables / words)
```

The formula penalizes:

- Long sentences (more words per sentence)
- Complex words (more syllables per word)

## Recommendations

For technical documentation:

- **API Reference**: 30-50 is acceptable
- **User Guides**: Aim for 50-60
- **Tutorials**: Target 60-70

## Improving Your Score

1. **Shorter sentences** - Break up complex sentences
2. **Simpler words** - Use "use" instead of "utilize"
3. **Active voice** - "The API returns data" vs "Data is returned by the API"
4. **Remove filler** - Cut unnecessary words
