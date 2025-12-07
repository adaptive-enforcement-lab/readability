package markdown

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// ParseResult contains extracted content from a markdown file.
type ParseResult struct {
	Prose       string
	CodeBlocks  []string
	Headings    []Heading
	Admonitions []Admonition
	TotalLines  int
	CodeLines   int
	EmptyLines  int
}

// Admonition represents a MkDocs-style admonition block.
type Admonition struct {
	Line  int    // Line number (1-based)
	Type  string // note, warning, tip, etc.
	Title string // optional custom title
}

// Heading represents a markdown heading.
type Heading struct {
	Line  int // Line number (1-based)
	Level int
	Text  string
}

// Parse extracts prose content, code blocks, and headings from markdown.
func Parse(content []byte) (*ParseResult, error) {
	md := goldmark.New()
	reader := text.NewReader(content)
	doc := md.Parser().Parse(reader)

	result := &ParseResult{
		CodeBlocks:  make([]string, 0),
		Headings:    make([]Heading, 0),
		Admonitions: make([]Admonition, 0),
	}

	var proseBuilder strings.Builder

	// Walk the AST to extract content
	_ = ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch n := node.(type) {
		case *ast.Heading:
			// Get line number from the heading's first line segment
			line := 1
			if n.Lines().Len() > 0 {
				line = bytes.Count(content[:n.Lines().At(0).Start], []byte("\n")) + 1
			}
			heading := Heading{
				Line:  line,
				Level: n.Level,
				Text:  extractHeadingText(n, content),
			}
			result.Headings = append(result.Headings, heading)

		case *ast.FencedCodeBlock:
			var codeContent bytes.Buffer
			lines := n.Lines()
			for i := 0; i < lines.Len(); i++ {
				line := lines.At(i)
				codeContent.Write(line.Value(content))
			}
			result.CodeBlocks = append(result.CodeBlocks, codeContent.String())

		case *ast.CodeBlock:
			var codeContent bytes.Buffer
			lines := n.Lines()
			for i := 0; i < lines.Len(); i++ {
				line := lines.At(i)
				codeContent.Write(line.Value(content))
			}
			result.CodeBlocks = append(result.CodeBlocks, codeContent.String())

		case *ast.Text:
			// Only extract text that's not inside code blocks
			parent := n.Parent()
			if !isInsideCodeBlock(parent) {
				proseBuilder.Write(n.Segment.Value(content))
				if n.HardLineBreak() || n.SoftLineBreak() {
					proseBuilder.WriteString(" ")
				} else {
					proseBuilder.WriteString(" ")
				}
			}

		case *ast.String:
			parent := n.Parent()
			if !isInsideCodeBlock(parent) {
				proseBuilder.Write(n.Value)
				proseBuilder.WriteString(" ")
			}
		}

		return ast.WalkContinue, nil
	})

	result.Prose = strings.TrimSpace(proseBuilder.String())

	// Count lines
	lines := bytes.Split(content, []byte("\n"))
	result.TotalLines = len(lines)

	// Count code lines, empty lines, and detect admonitions
	inCodeBlock := false
	for lineNum, line := range lines {
		trimmed := bytes.TrimSpace(line)
		if bytes.HasPrefix(trimmed, []byte("```")) {
			inCodeBlock = !inCodeBlock
			result.CodeLines++
			continue
		}
		if inCodeBlock {
			result.CodeLines++
		} else if len(trimmed) == 0 {
			result.EmptyLines++
		}

		// Detect MkDocs-style admonitions: !!! type or !!! type "title"
		if !inCodeBlock && bytes.HasPrefix(trimmed, []byte("!!!")) {
			admonition := parseAdmonition(string(trimmed))
			if admonition != nil {
				admonition.Line = lineNum + 1 // 1-based line numbers
				result.Admonitions = append(result.Admonitions, *admonition)
			}
		}
	}

	return result, nil
}

// parseAdmonition parses a MkDocs-style admonition line.
// Formats: !!! note, !!! warning "Custom Title", !!! tip inline
func parseAdmonition(line string) *Admonition {
	// Remove the !!! prefix
	line = strings.TrimPrefix(line, "!!!")
	line = strings.TrimSpace(line)

	if line == "" {
		return nil
	}

	adm := &Admonition{}

	// Check for quoted title: !!! type "title"
	if idx := strings.Index(line, "\""); idx != -1 {
		adm.Type = strings.TrimSpace(line[:idx])
		rest := line[idx+1:]
		if endIdx := strings.Index(rest, "\""); endIdx != -1 {
			adm.Title = rest[:endIdx]
		}
	} else {
		// No title, just type (may include modifiers like "inline")
		parts := strings.Fields(line)
		if len(parts) > 0 {
			adm.Type = parts[0]
		}
	}

	// Normalize type (remove "inline" modifier if present)
	adm.Type = strings.TrimSuffix(adm.Type, "+")

	return adm
}

// isInsideCodeBlock checks if a node is inside a code block.
func isInsideCodeBlock(node ast.Node) bool {
	for node != nil {
		switch node.(type) {
		case *ast.FencedCodeBlock, *ast.CodeBlock, *ast.CodeSpan:
			return true
		}
		node = node.Parent()
	}
	return false
}

// extractHeadingText extracts the text content from a heading node.
// This replaces the deprecated n.Text() method.
func extractHeadingText(n *ast.Heading, source []byte) string {
	var buf bytes.Buffer
	for child := n.FirstChild(); child != nil; child = child.NextSibling() {
		if text, ok := child.(*ast.Text); ok {
			buf.Write(text.Segment.Value(source))
		}
	}
	return buf.String()
}
