# What is Readability?

You wrote something. But will people actually understand it?

Readability tools answer that question with numbers. They scan your text and tell you how hard it is to read. This page explains how that works and why you might care.

## The Simple Version

Readability scores measure two things:

1. **How long are your sentences?** Longer sentences are harder to follow.
2. **How complex are your words?** Words with more syllables take more effort.

That's it. The computer counts words, sentences, and syllables. Then it does some math. Out comes a score.

!!! example "A Quick Test"
    Consider these two sentences:

    **Hard:** "The implementation of the authentication mechanism necessitates the utilization of secure protocols."

    **Easy:** "Login requires secure protocols."

    Same meaning. But the second one has shorter words and fewer of them. It scores as easier to read.

## Where This Came From

Readability formulas aren't new. They've been around since the 1940s.

| Year | Formula | Created For |
|------|---------|-------------|
| 1948 | Flesch Reading Ease | US Navy training manuals |
| 1968 | Flesch-Kincaid | Military documents |
| 1969 | Gunning Fog | Newspapers and business |
| 1975 | ARI | Real-time computer analysis |

The US military wanted to make sure soldiers could understand their manuals. Newspapers wanted to reach the widest audience. Publishers wanted to match books to student grade levels.

These same formulas now help developers write clearer documentation.

!!! note "A Strange History"
    The Flesch Reading Ease formula was created by Rudolf Flesch, an Austrian immigrant who became obsessed with clear writing. His 1955 book "Why Johnny Can't Read" became a bestseller and changed how American schools taught reading.

## What the Scores Mean

### Grade Level

A grade level score tells you what US school grade can read the text. Grade 8 means an eighth grader should understand it.

| Score | Who Can Read It |
|-------|-----------------|
| 6 | Sixth graders (age 11-12) |
| 8 | Eighth graders (age 13-14) |
| 10 | High school sophomores |
| 12 | High school seniors |
| 14 | College sophomores |
| 16+ | Graduate students |

Most popular websites aim for grade 7-8. Technical docs usually land around grade 10-14.

### Reading Ease

Reading Ease flips the scale. Higher numbers mean easier reading.

| Score | Difficulty |
|-------|------------|
| 90-100 | Very easy (comics, simple fiction) |
| 60-70 | Standard (newspapers, magazines) |
| 30-50 | Difficult (academic papers, legal docs) |
| 0-30 | Very difficult (scientific journals) |

## Who Uses This?

### Technical Writers

The people who write software documentation, API guides, and user manuals. They use readability scores to catch overly complex explanations before publishing.

### Content Teams

Marketing teams, documentation teams, and content strategists use readability checks to maintain consistent quality across many writers.

### Open Source Projects

Projects that want their docs accessible to non-native English speakers often set readability standards. Simpler English is easier to translate and understand globally.

### Students and Bloggers

Anyone who writes and wants feedback on clarity. If you're writing a blog post or an essay, readability scores tell you if you're losing readers with complex prose.

### Newsrooms

Journalists have used these formulas for decades. The Associated Press stylebook recommends keeping sentences under 20 words. Most newspapers target grade 6-8 reading level.

!!! tip "You Don't Need Users"
    Even if you're just writing for yourself, readability scores help you improve. They're like a spell checker, but for clarity instead of spelling.

## What This Tool Actually Shows You

When you run this tool on a Markdown file, you get a report. Here's what real output looks like:

```
$ readability docs/getting-started.md

┌────────────────────────┬───────┬──────┬───────┬───────┬───────┬──────┐
│ File                   │ Lines │ Read │ Grade │ ARI   │ Ease  │ Stat │
├────────────────────────┼───────┼──────┼───────┼───────┼───────┼──────┤
│ docs/getting-started.md│   45  │ <1m  │  8.2  │  9.1  │ 62.3  │ pass │
└────────────────────────┴───────┴──────┴───────┴───────┴───────┴──────┘
```

Here's what each column means:

| Column | What It Tells You |
|--------|-------------------|
| **Lines** | How long the file is |
| **Read** | How long it takes to read (at 200 words/minute) |
| **Grade** | Flesch-Kincaid grade level |
| **ARI** | Automated Readability Index (another grade metric) |
| **Ease** | Flesch Reading Ease score (higher = easier) |
| **Stat** | Pass or fail based on your thresholds |

If a file fails, the tool tells you why:

```
$ readability --check docs/complex-api.md

docs/complex-api.md:1:1: error: grade level 15.3 exceeds maximum 12 (grade-level)
docs/complex-api.md:1:1: error: reading ease 28.1 below minimum 30 (reading-ease)

GRADE: Files exceed the maximum grade level threshold.
To fix, try:
- Breaking long sentences into shorter ones
- Using simpler words ("use" instead of "utilize")
- Removing unnecessary jargon
```

## Why Put This in CI/CD?

CI/CD pipelines run automated checks on your code. Tests, linting, security scans. Why add writing quality?

**The same reason you lint code.** Linters catch style issues before code review. Readability checks catch confusing docs before they ship.

Without automation:
- Complex docs slip through review
- Quality varies by author
- Nobody notices until users complain

With automation:
- Every PR gets checked
- Standards stay consistent
- Problems are caught early

!!! warning "It's Not Perfect"
    Readability scores can't tell if your writing is accurate, complete, or well-organized. They only measure sentence and word complexity. You still need human review for everything else.

## The Limits of Scores

Readability formulas are blunt instruments. They can't detect:

- **Wrong information** - A factually incorrect sentence can score as "easy"
- **Missing context** - Short sentences without enough explanation score well but confuse readers
- **Good complexity** - Sometimes technical terms are necessary and simplifying would be wrong
- **Tone and voice** - A robotic easy-to-read doc might be less engaging than a slightly complex one

Use scores as one signal, not the only signal. A high grade level is a prompt to review, not an automatic failure.

## Getting Started

Ready to try it? Here's the quickest path:

```bash
# If you have Go installed
go install github.com/adaptive-enforcement-lab/readability/cmd/readability@latest

# Run it on any Markdown file
readability README.md
```

Or see the [Getting Started guide](getting-started/index.md) for more installation options.
