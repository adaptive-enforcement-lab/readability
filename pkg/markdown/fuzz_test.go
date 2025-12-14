package markdown

import (
	"testing"
)

// FuzzParse tests the Parse function with arbitrary markdown input.
// This helps discover panics, hangs, or unexpected behavior with malformed input.
func FuzzParse(f *testing.F) {
	// Seed corpus with various markdown patterns
	seeds := []string{
		// Basic content
		"# Heading\n\nSome text with `code`.",
		"",
		"   ",
		"\n\n\n",

		// Code blocks
		"```go\nfunc main() {}\n```",
		"```\ncode block\n```",
		"```python\nprint('hello')\n```",
		"    indented code block",

		// Nested and complex structures
		"# H1\n## H2\n### H3\n#### H4\n##### H5\n###### H6",
		"# Title\n\n```\ncode\n```\n\n## Subtitle\n\nMore text.",

		// Admonitions (MkDocs-style)
		"!!! note\n    Content",
		"!!! warning \"Title\"\n    Warning content.",
		"!!! tip inline\n    Inline tip.",
		"!!! note+\n    Collapsible.",
		"!!!",
		"!!! ",

		// Edge cases
		"```",           // Unclosed code block
		"```\n```\n```", // Multiple fence markers
		"# ",            // Empty heading
		"######",        // Heading without space
		"#######",       // Invalid heading level
		"**bold** _italic_ `code` [link](url)",
		"!!! unknown_type \"Complex \\\"Escaped\\\" Title\"",

		// Unicode and special characters
		"\u0000",
		"\xff\xfe",
		"# \U0001F600 Emoji Heading",
		"```\n\x00\x01\x02\n```",

		// Large repetitive content
		string(make([]byte, 1000)),
		"#" + string(make([]byte, 100)),
	}

	for _, seed := range seeds {
		f.Add([]byte(seed))
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		result, err := Parse(data)

		// We don't expect errors from Parse (it handles malformed input gracefully)
		// but we verify invariants hold
		if err != nil {
			// If an error occurs, it should be a valid error, not a panic
			return
		}

		// Invariant: result should never be nil when err is nil
		if result == nil {
			t.Error("Parse returned nil result without error")
			return
		}

		// Invariant: TotalLines should always be >= 1 (even empty input has 1 line)
		if result.TotalLines < 1 {
			t.Errorf("TotalLines = %d, want >= 1", result.TotalLines)
		}

		// Invariant: CodeLines + EmptyLines <= TotalLines
		if result.CodeLines+result.EmptyLines > result.TotalLines {
			t.Errorf("CodeLines(%d) + EmptyLines(%d) > TotalLines(%d)",
				result.CodeLines, result.EmptyLines, result.TotalLines)
		}

		// Invariant: slices should never be nil
		if result.CodeBlocks == nil {
			t.Error("CodeBlocks is nil, expected empty slice")
		}
		if result.Headings == nil {
			t.Error("Headings is nil, expected empty slice")
		}
		if result.Admonitions == nil {
			t.Error("Admonitions is nil, expected empty slice")
		}

		// Invariant: heading levels should be 1-6
		for i, h := range result.Headings {
			if h.Level < 1 || h.Level > 6 {
				t.Errorf("Heading[%d].Level = %d, want 1-6", i, h.Level)
			}
			if h.Line < 1 {
				t.Errorf("Heading[%d].Line = %d, want >= 1", i, h.Line)
			}
		}

		// Invariant: admonition lines should be positive
		for i, a := range result.Admonitions {
			if a.Line < 1 {
				t.Errorf("Admonition[%d].Line = %d, want >= 1", i, a.Line)
			}
		}
	})
}

// FuzzParseAdmonition tests the parseAdmonition function directly.
// This function parses MkDocs-style admonition syntax.
func FuzzParseAdmonition(f *testing.F) {
	// Seed corpus with various admonition patterns
	seeds := []string{
		// Valid patterns
		"!!! note",
		"!!! warning",
		"!!! tip",
		"!!! info",
		"!!! danger",
		"!!! example",
		"!!! warning \"Custom Title\"",
		"!!! tip inline",
		"!!! note+",
		"!!! example \"This Is A Long Title\"",

		// Edge cases
		"!!!",
		"!!! ",
		"!!!note",
		"!!! note \"",
		"!!! note \"Unclosed",
		"!!! note \"\"",
		"!!! \"no type\"",
		"!!!  multiple  spaces",
		"!!! type \"title\" extra",
		"!!! type \"nested \"quotes\" here\"",

		// Unicode and special characters
		"!!! note \"\U0001F600\"",
		"!!! \u0000",
		"!!! type \"line\nbreak\"",

		// Long inputs
		"!!! " + string(make([]byte, 1000)),
		"!!! type \"" + string(make([]byte, 1000)) + "\"",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, line string) {
		result := parseAdmonition(line)

		// parseAdmonition should never panic
		// It can return nil for invalid input, which is expected

		// Use result to satisfy the linter - the main purpose is panic detection
		_ = result
	})
}
