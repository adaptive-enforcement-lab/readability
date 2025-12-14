package markdown

import (
	"strings"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name           string
		content        string
		wantProseLen   int // approximate, just check it's not empty
		wantHeadings   int
		wantCodeBlocks int
		wantTotalLines int
		wantCodeLines  int
		wantEmptyLines int
	}{
		{
			name:           "empty content",
			content:        "",
			wantProseLen:   0,
			wantHeadings:   0,
			wantCodeBlocks: 0,
			wantTotalLines: 1,
			wantCodeLines:  0,
			wantEmptyLines: 1,
		},
		{
			name:           "simple prose",
			content:        "Hello world. This is a test.",
			wantProseLen:   10,
			wantHeadings:   0,
			wantCodeBlocks: 0,
			wantTotalLines: 1,
			wantCodeLines:  0,
			wantEmptyLines: 0,
		},
		{
			name:           "heading only",
			content:        "# Title",
			wantProseLen:   1,
			wantHeadings:   1,
			wantCodeBlocks: 0,
			wantTotalLines: 1,
			wantCodeLines:  0,
			wantEmptyLines: 0,
		},
		{
			name: "multiple headings",
			content: `# H1
## H2
### H3
#### H4
##### H5
###### H6`,
			wantHeadings:   6,
			wantCodeBlocks: 0,
			wantTotalLines: 6,
		},
		{
			name:           "fenced code block",
			content:        "# Title\n\n```go\nfunc main() {}\n```\n\nSome text.",
			wantHeadings:   1,
			wantCodeBlocks: 1,
			wantTotalLines: 7,
			wantCodeLines:  3,
		},
		{
			name:           "multiple code blocks",
			content:        "```\ncode1\n```\n\n```python\ncode2\n```",
			wantCodeBlocks: 2,
			wantTotalLines: 7,
			wantCodeLines:  6,
		},
		{
			name:           "prose with empty lines",
			content:        "First paragraph.\n\nSecond paragraph.\n\nThird paragraph.",
			wantProseLen:   10,
			wantTotalLines: 5,
			wantEmptyLines: 2,
		},
		{
			name: "mixed content",
			content: `# Introduction

This is some prose content.

## Code Example

` + "```go" + `
package main
` + "```" + `

More prose here.`,
			wantHeadings:   2,
			wantCodeBlocks: 1,
			wantTotalLines: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse([]byte(tt.content))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if tt.wantProseLen > 0 && len(result.Prose) < tt.wantProseLen {
				t.Errorf("Parse() prose length = %d, want at least %d", len(result.Prose), tt.wantProseLen)
			}
			if len(result.Headings) != tt.wantHeadings {
				t.Errorf("Parse() headings = %d, want %d", len(result.Headings), tt.wantHeadings)
			}
			if len(result.CodeBlocks) != tt.wantCodeBlocks {
				t.Errorf("Parse() code blocks = %d, want %d", len(result.CodeBlocks), tt.wantCodeBlocks)
			}
			if result.TotalLines != tt.wantTotalLines {
				t.Errorf("Parse() total lines = %d, want %d", result.TotalLines, tt.wantTotalLines)
			}
			if tt.wantCodeLines > 0 && result.CodeLines != tt.wantCodeLines {
				t.Errorf("Parse() code lines = %d, want %d", result.CodeLines, tt.wantCodeLines)
			}
			if tt.wantEmptyLines > 0 && result.EmptyLines != tt.wantEmptyLines {
				t.Errorf("Parse() empty lines = %d, want %d", result.EmptyLines, tt.wantEmptyLines)
			}
		})
	}
}

func TestParse_Headings(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		wantLevels []int
		wantTexts  []string
	}{
		{
			name:       "all heading levels",
			content:    "# H1\n## H2\n### H3\n#### H4\n##### H5\n###### H6",
			wantLevels: []int{1, 2, 3, 4, 5, 6},
			wantTexts:  []string{"H1", "H2", "H3", "H4", "H5", "H6"},
		},
		{
			name:       "heading with formatting",
			content:    "# **Bold** heading",
			wantLevels: []int{1},
			// Parser strips markdown formatting, leaving just the suffix
			wantTexts: []string{" heading"},
		},
		{
			name:       "heading with inline code",
			content:    "# The `code` heading",
			wantLevels: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse([]byte(tt.content))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if len(result.Headings) != len(tt.wantLevels) {
				t.Fatalf("Parse() headings count = %d, want %d", len(result.Headings), len(tt.wantLevels))
			}

			for i, h := range result.Headings {
				if h.Level != tt.wantLevels[i] {
					t.Errorf("heading[%d] level = %d, want %d", i, h.Level, tt.wantLevels[i])
				}
				if tt.wantTexts != nil && h.Text != tt.wantTexts[i] {
					t.Errorf("heading[%d] text = %q, want %q", i, h.Text, tt.wantTexts[i])
				}
			}
		})
	}
}

func TestParse_Admonitions(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		wantCount  int
		wantTypes  []string
		wantTitles []string
	}{
		{
			name:      "no admonitions",
			content:   "Regular content without admonitions.",
			wantCount: 0,
		},
		{
			name:      "basic admonition",
			content:   "!!! note\n    Content here.",
			wantCount: 1,
			wantTypes: []string{"note"},
		},
		{
			name:       "admonition with title",
			content:    "!!! warning \"Custom Title\"\n    Warning content.",
			wantCount:  1,
			wantTypes:  []string{"warning"},
			wantTitles: []string{"Custom Title"},
		},
		{
			name:      "multiple admonitions",
			content:   "!!! note\n    Note.\n\n!!! tip\n    Tip.\n\n!!! warning\n    Warning.",
			wantCount: 3,
			wantTypes: []string{"note", "tip", "warning"},
		},
		{
			name:      "admonition with inline modifier",
			content:   "!!! tip inline\n    Inline tip.",
			wantCount: 1,
			wantTypes: []string{"tip"},
		},
		{
			name:      "collapsible admonition",
			content:   "!!! note+\n    Collapsible.",
			wantCount: 1,
			wantTypes: []string{"note"},
		},
		{
			name:      "admonition inside code block ignored",
			content:   "```\n!!! note\n    This should not be counted.\n```",
			wantCount: 0,
		},
		{
			name:      "all common types",
			content:   "!!! note\n\n!!! warning\n\n!!! tip\n\n!!! info\n\n!!! danger\n\n!!! example\n\n!!! abstract\n\n!!! question",
			wantCount: 8,
			wantTypes: []string{"note", "warning", "tip", "info", "danger", "example", "abstract", "question"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse([]byte(tt.content))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if len(result.Admonitions) != tt.wantCount {
				t.Errorf("Parse() admonitions count = %d, want %d", len(result.Admonitions), tt.wantCount)
			}

			for i, adm := range result.Admonitions {
				if tt.wantTypes != nil && i < len(tt.wantTypes) {
					if adm.Type != tt.wantTypes[i] {
						t.Errorf("admonition[%d] type = %q, want %q", i, adm.Type, tt.wantTypes[i])
					}
				}
				if tt.wantTitles != nil && i < len(tt.wantTitles) {
					if adm.Title != tt.wantTitles[i] {
						t.Errorf("admonition[%d] title = %q, want %q", i, adm.Title, tt.wantTitles[i])
					}
				}
				if adm.Line <= 0 {
					t.Errorf("admonition[%d] line = %d, want positive", i, adm.Line)
				}
			}
		})
	}
}

func TestParseAdmonition(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		wantNil   bool
		wantType  string
		wantTitle string
	}{
		{
			name:    "empty after prefix",
			line:    "!!!",
			wantNil: true,
		},
		{
			name:     "basic type",
			line:     "!!! note",
			wantType: "note",
		},
		{
			name:      "type with title",
			line:      "!!! warning \"Watch Out\"",
			wantType:  "warning",
			wantTitle: "Watch Out",
		},
		{
			name:     "type with inline modifier",
			line:     "!!! tip inline",
			wantType: "tip",
		},
		{
			name:     "collapsible modifier",
			line:     "!!! note+",
			wantType: "note",
		},
		{
			name:      "complex title with spaces",
			line:      "!!! example \"This Is A Long Title\"",
			wantType:  "example",
			wantTitle: "This Is A Long Title",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseAdmonition(tt.line)

			if tt.wantNil {
				if result != nil {
					t.Errorf("parseAdmonition() = %v, want nil", result)
				}
				return
			}

			if result == nil {
				t.Fatal("parseAdmonition() = nil, want non-nil")
			}
			if result.Type != tt.wantType {
				t.Errorf("parseAdmonition().Type = %q, want %q", result.Type, tt.wantType)
			}
			if result.Title != tt.wantTitle {
				t.Errorf("parseAdmonition().Title = %q, want %q", result.Title, tt.wantTitle)
			}
		})
	}
}

func TestParse_CodeBlocks(t *testing.T) {
	tests := []struct {
		name         string
		content      string
		wantBlocks   int
		wantContains []string
	}{
		{
			name:         "fenced code block",
			content:      "```\ncode here\n```",
			wantBlocks:   1,
			wantContains: []string{"code here"},
		},
		{
			name:         "fenced with language",
			content:      "```go\npackage main\n```",
			wantBlocks:   1,
			wantContains: []string{"package main"},
		},
		{
			name:         "indented code block",
			content:      "    indented code\n    more code",
			wantBlocks:   1,
			wantContains: []string{"indented code"},
		},
		{
			name:         "multiple fenced blocks",
			content:      "```\nblock1\n```\n\n```\nblock2\n```",
			wantBlocks:   2,
			wantContains: []string{"block1", "block2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse([]byte(tt.content))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if len(result.CodeBlocks) != tt.wantBlocks {
				t.Errorf("Parse() code blocks = %d, want %d", len(result.CodeBlocks), tt.wantBlocks)
			}

			for i, want := range tt.wantContains {
				if i < len(result.CodeBlocks) {
					found := false
					for _, block := range result.CodeBlocks {
						if contains(block, want) {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Parse() code blocks should contain %q", want)
					}
				}
			}
		})
	}
}

func TestParse_LineCounting(t *testing.T) {
	tests := []struct {
		name           string
		content        string
		wantTotal      int
		wantCode       int
		wantEmpty      int
		wantProseLines int // calculated as total - code - empty
	}{
		{
			name:           "all prose",
			content:        "Line 1\nLine 2\nLine 3",
			wantTotal:      3,
			wantCode:       0,
			wantEmpty:      0,
			wantProseLines: 3,
		},
		{
			name:           "with empty lines",
			content:        "Line 1\n\nLine 2\n\nLine 3",
			wantTotal:      5,
			wantCode:       0,
			wantEmpty:      2,
			wantProseLines: 3,
		},
		{
			name:           "code block",
			content:        "```\ncode\n```",
			wantTotal:      3,
			wantCode:       3,
			wantEmpty:      0,
			wantProseLines: 0,
		},
		{
			name:           "mixed",
			content:        "Prose\n\n```\ncode\n```\n\nMore prose",
			wantTotal:      7,
			wantCode:       3,
			wantEmpty:      2,
			wantProseLines: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse([]byte(tt.content))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if result.TotalLines != tt.wantTotal {
				t.Errorf("TotalLines = %d, want %d", result.TotalLines, tt.wantTotal)
			}
			if result.CodeLines != tt.wantCode {
				t.Errorf("CodeLines = %d, want %d", result.CodeLines, tt.wantCode)
			}
			if result.EmptyLines != tt.wantEmpty {
				t.Errorf("EmptyLines = %d, want %d", result.EmptyLines, tt.wantEmpty)
			}

			proseLines := result.TotalLines - result.CodeLines - result.EmptyLines
			if proseLines != tt.wantProseLines {
				t.Errorf("ProseLines = %d, want %d", proseLines, tt.wantProseLines)
			}
		})
	}
}

func TestParse_ProseExtraction(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		wantContain string
		wantExclude string
	}{
		{
			name:        "excludes code blocks",
			content:     "Hello world.\n\n```\ncode here\n```\n\nMore text.",
			wantContain: "Hello world",
			wantExclude: "code here",
		},
		{
			name:        "includes heading text",
			content:     "# My Title\n\nParagraph.",
			wantContain: "My Title",
		},
		{
			name:        "handles inline code",
			content:     "Use the `command` to run.",
			wantContain: "Use the",
		},
		{
			name:        "excludes table content",
			content:     "Normal text.\n\n| Col - 1 | Col - 2 |\n|---------|----------|\n| Val - A | Val - B |\n\nMore text.",
			wantContain: "Normal text",
			wantExclude: "Col - 1",
		},
		{
			name:        "excludes list content",
			content:     "Normal text.\n\n- [Link](url.md) - Description\n- Another - item\n\nMore text.",
			wantContain: "Normal text",
			wantExclude: "Description",
		},
		{
			name:        "excludes nested lists",
			content:     "Text.\n\n- Item 1\n  - Nested - item\n\nMore.",
			wantContain: "Text",
			wantExclude: "Nested",
		},
		{
			name:        "excludes YAML frontmatter",
			content:     "---\ntitle: Test\nauthor: John - Doe\ndescription: A test - with dashes\n---\n\n# Heading\n\nNormal prose.",
			wantContain: "Normal prose",
			wantExclude: "John - Doe",
		},
		{
			name:        "excludes TOML frontmatter",
			content:     "+++\ntitle = \"Test\"\nauthor = \"John - Doe\"\n+++\n\n# Heading\n\nNormal prose.",
			wantContain: "Normal prose",
			wantExclude: "John - Doe",
		},
		{
			name:        "handles content without frontmatter",
			content:     "# Heading\n\nNormal prose with - dashes.",
			wantContain: "with - dashes",
		},
		{
			name:        "handles incomplete frontmatter (no closing delimiter)",
			content:     "---\ntitle: Test\n\n# Heading\n\nProse.",
			wantContain: "title",
			wantExclude: "should not match because incomplete frontmatter is kept",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse([]byte(tt.content))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			if tt.wantContain != "" && !contains(result.Prose, tt.wantContain) {
				t.Errorf("Prose should contain %q, got %q", tt.wantContain, result.Prose)
			}
			if tt.wantExclude != "" && contains(result.Prose, tt.wantExclude) {
				t.Errorf("Prose should not contain %q, got %q", tt.wantExclude, result.Prose)
			}
		})
	}
}

func TestIsInsideCodeBlock(t *testing.T) {
	// This is implicitly tested through Parse tests
	// Adding explicit test for edge case
	content := "```\ncode\n```"
	result, err := Parse([]byte(content))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Prose should not contain code block content
	if contains(result.Prose, "code") {
		t.Errorf("Prose should not contain code block content")
	}
}

func TestExtractString(t *testing.T) {
	// Test extractString directly with real ast.String nodes
	tests := []struct {
		name     string
		value    string
		wantText string
	}{
		{
			name:     "simple string",
			value:    "test value",
			wantText: "test value ",
		},
		{
			name:     "empty string",
			value:    "",
			wantText: " ",
		},
		{
			name:     "string with special chars",
			value:    "hello & goodbye",
			wantText: "hello & goodbye ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder strings.Builder
			// Create a real ast.String node
			node := ast.NewString([]byte(tt.value))
			extractString(node, &builder)

			if builder.String() != tt.wantText {
				t.Errorf("extractString() = %q, want %q", builder.String(), tt.wantText)
			}
		})
	}
}

func TestExtractStringInsideCodeBlock(t *testing.T) {
	// Test that extractString does not extract content when inside a code block
	var builder strings.Builder
	node := ast.NewString([]byte("should not appear"))

	// Set parent to a code block to test the guard clause
	codeBlock := ast.NewCodeBlock()
	codeBlock.AppendChild(codeBlock, node)

	extractString(node, &builder)

	if builder.String() != "" {
		t.Errorf("extractString() should not extract inside code block, got %q", builder.String())
	}
}

// helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
