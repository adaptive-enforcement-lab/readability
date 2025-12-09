# Who Is This For?

Different people use readability tools for different reasons. Find your situation below.

## "I Write Blog Posts"

You want feedback on your writing. Are your posts too dense? Will readers bail halfway through?

**How to use this:**

```bash
readability my-post.md
```

Look at the Grade column. If it's above 10, your post might lose casual readers. Try:

- Breaking long sentences in half
- Swapping complex words for simple ones
- Reading sentences aloud (if you run out of breath, it's too long)

!!! tip "No Installation Needed"
    Paste your text into any online Flesch-Kincaid calculator to get a quick score. Use this tool when you want to check files regularly.

## "I Write Technical Documentation"

You maintain docs for software, APIs, or tools. You want to make sure your explanations are clear.

**How to use this:**

Set up a config file with your standards:

```yaml
# .readability.yml
thresholds:
  max_grade: 12      # High school senior level
  max_ari: 14
  min_ease: 30
```

Run checks locally:

```bash
readability --check docs/
```

Or add it to your CI pipeline so every PR gets checked automatically.

!!! note "Technical Terms Are Okay"
    Don't dumb down necessary jargon. "API endpoint" is fine in API docs. The goal is clear explanations around the jargon, not removing it.

## "I Manage a Docs Team"

Multiple writers contribute to your documentation. Quality varies. You want consistent standards.

**How to use this:**

1. Set thresholds that match your audience
2. Add the GitHub Action to your docs repository
3. PRs that exceed thresholds fail the check

This creates a forcing function. Writers know the standards before they submit. Reviewers don't have to be the "readability police."

```yaml
# In your GitHub workflow
- uses: adaptive-enforcement-lab/readability@v1
  with:
    path: docs/
    check: true
    max-grade: 10  # Stricter for user-facing docs
```

## "I'm a Student Learning to Write"

You're working on essays, reports, or applications. You want to improve your writing clarity.

**How to use this:**

Check your draft:

```bash
readability essay.md
```

Compare before and after revisions. Did your edits improve the score? Worse scores after editing might mean you made things more complex, not better.

!!! example "Revision Practice"
    Take any paragraph scoring above grade 12. Rewrite it to hit grade 8 without losing meaning. This exercise builds the habit of clear writing.

## "I Run an Open Source Project"

Contributors come from everywhere. Non-native English speakers read your docs. You want maximum accessibility.

**How to use this:**

Set stricter thresholds than commercial projects:

```yaml
thresholds:
  max_grade: 10     # More accessible globally
  min_ease: 40      # Easier to translate
  max_lines: 300    # Shorter, focused pages
```

Simpler English is easier to:
- Translate with automated tools
- Understand for non-native speakers
- Maintain by rotating contributors

## "I'm Curious What This Measures"

You don't have a specific use case. You just want to understand readability scores.

**Start here:**

Read the [Introduction](introduction.md) for the history and theory. Then run the tool on any text you have lying around:

```bash
readability README.md
```

See what scores your existing writing gets. That context makes the numbers meaningful.

## "I Need to Convince My Team"

You think your team should care about readability. They're skeptical.

**Arguments that work:**

1. **Support ticket reduction** - Clearer docs mean fewer "how do I...?" questions
2. **Onboarding speed** - New users get productive faster with accessible docs
3. **Global reach** - Simpler English works better for international audiences
4. **Consistent quality** - Automated checks beat inconsistent human review

**Arguments that don't work:**

- "The score says it's bad" - Scores are signals, not verdicts
- "Hemingway wrote at grade 4" - Your API docs aren't literature
- "It's industry best practice" - Show your own data instead

!!! warning "Don't Weaponize Scores"
    Readability tools help writers improve. They shouldn't shame or gatekeep. Use them to start conversations, not end them.

## Not Sure Where to Start?

Try this:

1. Run `readability README.md` on any project
2. Look at the Grade score
3. Pick one sentence with complex words and simplify it
4. Run again and see if the score improved

That's it. One file, one sentence, one improvement. Scale up from there.
