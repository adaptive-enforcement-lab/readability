package analyzer

import (
	"strings"
	"testing"
)

// FuzzStripFrontmatter tests the stripFrontmatter function with arbitrary input.
// This function removes YAML frontmatter delimited by --- markers.
func FuzzStripFrontmatter(f *testing.F) {
	// Seed corpus with various frontmatter patterns
	seeds := []string{
		// Valid frontmatter
		"---\ntitle: Test\n---\nContent here.",
		"---\nkey: value\n---\n\n# Heading",
		"---\nmulti:\n  - item1\n  - item2\n---\nBody",

		// No frontmatter
		"Regular content without frontmatter.",
		"# Heading\n\nParagraph.",
		"",

		// Edge cases
		"---",           // Only opening marker
		"---\n---",      // Empty frontmatter
		"---\n---\n",    // Empty frontmatter with trailing newline
		"---\nno close", // Unclosed frontmatter
		"--- ",          // Marker with trailing space
		" ---",          // Leading space (not valid frontmatter)
		"---\n---\n---", // Multiple markers
		"---\n---\n---\ncontent",
		"----", // Four dashes (not frontmatter)
		"--",   // Two dashes (not frontmatter)
		"---\ncontent\n---\nmore---content",

		// Unicode and special characters
		"---\ntitle: \U0001F600\n---\nContent",
		"---\n\x00\x01\n---\nAfter",
		"---\nkey: \"value with ---\"\n---\nBody",

		// Large inputs
		"---\n" + strings.Repeat("key: value\n", 100) + "---\nContent",
		"---\n---\n" + strings.Repeat("x", 10000),
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, content string) {
		result := stripFrontmatter(content)

		// stripFrontmatter should never panic

		// Invariant: result should never be longer than input
		// (we're only removing content, never adding)
		if len(result) > len(content) {
			t.Errorf("stripFrontmatter result (%d bytes) longer than input (%d bytes)",
				len(result), len(content))
		}

		// Invariant: if input doesn't start with ---, output equals input
		if !strings.HasPrefix(content, "---") {
			if result != content {
				t.Error("stripFrontmatter modified content that doesn't start with ---")
			}
		}

		// Invariant: result should not contain the frontmatter if it was valid
		// (This is a weaker check - just verify no panic)
	})
}

// FuzzCountSentences tests the countSentences function with arbitrary input.
// This function counts sentences based on punctuation marks.
func FuzzCountSentences(f *testing.F) {
	// Seed corpus with various sentence patterns
	seeds := []string{
		// Normal sentences
		"This is a sentence.",
		"One. Two. Three.",
		"Hello! How are you? Fine.",
		"No punctuation here",

		// Edge cases
		"",
		".",
		"...",
		"?!.",
		"   ",
		"\n\n\n",

		// False positives for sentence detection
		"Dr. Smith went to the store.",
		"The price is $9.99 today.",
		"Visit example.com for more.",
		"Version 1.2.3 is released.",
		"The U.S.A. is large.",
		"3.14159 is pi.",

		// Unicode
		"\U0001F600 Emoji. More text.",
		"Chinese\u3002 Japanese\u3002", // Ideographic full stop
		"Arabic\u061F Question?",       // Arabic question mark

		// Special characters
		"\x00.\x01?\x02!",
		"Tab\t.\tMore.",
		"Newline.\nMore text.",

		// Long inputs
		strings.Repeat("Sentence. ", 100),
		strings.Repeat("x", 10000) + ".",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, text string) {
		count := countSentences(text)

		// countSentences should never panic

		// Invariant: count should always be >= 0
		if count < 0 {
			t.Errorf("countSentences returned %d, want >= 0", count)
		}

		// Invariant: for non-empty text, count should be >= 1
		// (the function returns 1 if no punctuation found but text exists)
		if len(text) > 0 && count < 1 {
			t.Errorf("countSentences returned %d for non-empty input, want >= 1", count)
		}

		// Invariant: count should not exceed the number of sentence-ending chars + 1
		maxPossible := strings.Count(text, ".") + strings.Count(text, "!") + strings.Count(text, "?")
		if maxPossible == 0 && len(text) > 0 {
			maxPossible = 1 // Function returns 1 for text without punctuation
		}
		if count > maxPossible {
			t.Errorf("countSentences returned %d, but max possible is %d", count, maxPossible)
		}
	})
}

// FuzzCountWords tests the countWords function with arbitrary input.
func FuzzCountWords(f *testing.F) {
	// Seed corpus
	seeds := []string{
		"one two three",
		"",
		"   ",
		"word",
		"multiple   spaces   between",
		"\ttabs\tand\nnewlines",
		"hyphenated-word counts-as-one",
		"\u00A0non-breaking space",
		strings.Repeat("word ", 1000),
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, text string) {
		count := countWords(text)

		// countWords should never panic

		// Invariant: count should always be >= 0
		if count < 0 {
			t.Errorf("countWords returned %d, want >= 0", count)
		}

		// Invariant: empty or whitespace-only input should return 0
		if strings.TrimSpace(text) == "" && count != 0 {
			t.Errorf("countWords returned %d for whitespace-only input, want 0", count)
		}
	})
}

// FuzzCalculateReadingTime tests the calculateReadingTime function.
func FuzzCalculateReadingTime(f *testing.F) {
	// Seed with various word counts
	f.Add(0)
	f.Add(1)
	f.Add(100)
	f.Add(199)
	f.Add(200)
	f.Add(201)
	f.Add(1000)
	f.Add(-1)
	f.Add(-100)
	f.Add(1<<30 - 1)  // Large positive
	f.Add(-(1 << 30)) // Large negative

	f.Fuzz(func(t *testing.T, words int) {
		result := calculateReadingTime(words)

		// calculateReadingTime should never panic

		// Invariant: result should always be >= 0
		if result < 0 {
			t.Errorf("calculateReadingTime(%d) = %d, want >= 0", words, result)
		}

		// Invariant: for words <= 0, result should be 0
		if words <= 0 && result != 0 {
			t.Errorf("calculateReadingTime(%d) = %d, want 0", words, result)
		}

		// Invariant: for words > 0, result should be >= 1
		if words > 0 && result < 1 {
			t.Errorf("calculateReadingTime(%d) = %d, want >= 1", words, result)
		}
	})
}

// FuzzCalculateRatio tests the calculateRatio function.
func FuzzCalculateRatio(f *testing.F) {
	// Seed with various part/total combinations
	f.Add(0, 0)
	f.Add(0, 100)
	f.Add(50, 100)
	f.Add(100, 100)
	f.Add(100, 50) // Part > total
	f.Add(-1, 100)
	f.Add(100, -1)
	f.Add(-1, -1)
	f.Add(1<<30, 1<<30)

	f.Fuzz(func(t *testing.T, part, total int) {
		result := calculateRatio(part, total)

		// calculateRatio should never panic

		// Invariant: for total == 0, result should be 0
		if total == 0 && result != 0.0 {
			t.Errorf("calculateRatio(%d, 0) = %f, want 0.0", part, result)
		}

		// Invariant: result should not be NaN or Inf (for valid inputs)
		if total != 0 {
			if result != result { // NaN check
				t.Errorf("calculateRatio(%d, %d) returned NaN", part, total)
			}
		}
	})
}
