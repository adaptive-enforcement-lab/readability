# Flesch Reading Ease

This score tells you how comfortable your text is to read. Higher numbers mean easier reading.

## The Scale

| Score | How It Reads | Who Can Follow |
|-------|--------------|----------------|
| 90-100 | Very easy | 5th graders |
| 80-90 | Easy | 6th graders |
| 70-80 | Fairly easy | 7th graders |
| 60-70 | Normal | 8th-9th graders |
| 50-60 | Fairly hard | 10th-12th graders |
| 30-50 | Hard | College students |
| 0-30 | Very hard | Grad students |

!!! info "For Tech Docs"
    Most technical docs land between 30-50. That's fine for expert readers. Aim higher (50-60+) for tutorials or user guides.

## What Affects the Score

The formula looks at:

- **Sentence length** - Longer sentences lower the score
- **Word complexity** - More syllables per word lower the score

## Targets by Content Type

| Content | Target Score |
|---------|--------------|
| API reference | 30-50 |
| User guides | 50-60 |
| Tutorials | 60-70 |
| Marketing | 70+ |

## How to Improve

1. **Shorten sentences** - Break long ones apart
2. **Use simple words** - "use" beats "utilize"
3. **Go active** - "The API sends data" not "Data is sent"
4. **Cut filler** - Remove words that don't add meaning

!!! example "Before and After"
    **Before (score ~40):**
    "The implementation of the authentication mechanism necessitates the utilization of secure token-based verification protocols."

    **After (score ~70):**
    "Login uses secure tokens to verify users."
