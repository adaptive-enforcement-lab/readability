package markdown

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	extast "github.com/yuin/goldmark/extension/ast"
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
	// Strip frontmatter and admonition blocks before parsing to exclude them from prose
	cleanedContent := stripFrontmatter(content)
	cleanedContent = stripAdmonitions(cleanedContent)

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM), // Enable GitHub Flavored Markdown (includes tables)
	)
	reader := text.NewReader(cleanedContent)
	doc := md.Parser().Parse(reader)

	result := &ParseResult{
		CodeBlocks:  make([]string, 0),
		Headings:    make([]Heading, 0),
		Admonitions: make([]Admonition, 0),
	}

	prose := extractAST(doc, cleanedContent, result)
	// Normalize whitespace: collapse multiple spaces to single space
	prose = strings.Join(strings.Fields(prose), " ")
	result.Prose = strings.TrimSpace(prose)

	countLines(content, result)

	return result, nil
}

// extractAST walks the AST and extracts headings, code blocks, and prose.
func extractAST(doc ast.Node, content []byte, result *ParseResult) string {
	var proseBuilder strings.Builder

	_ = ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch n := node.(type) {
		case *ast.Heading:
			result.Headings = append(result.Headings, extractHeading(n, content))
		case *ast.FencedCodeBlock:
			result.CodeBlocks = append(result.CodeBlocks, extractCodeBlock(n, content))
		case *ast.CodeBlock:
			result.CodeBlocks = append(result.CodeBlocks, extractCodeBlock(n, content))
		case *ast.Text:
			extractText(n, content, &proseBuilder)
		case *ast.String:
			extractString(n, &proseBuilder)
		}

		return ast.WalkContinue, nil
	})

	return proseBuilder.String()
}

// extractHeading extracts a heading from an AST node.
func extractHeading(n *ast.Heading, content []byte) Heading {
	line := 1
	if n.Lines().Len() > 0 {
		line = bytes.Count(content[:n.Lines().At(0).Start], []byte("\n")) + 1
	}
	return Heading{
		Line:  line,
		Level: n.Level,
		Text:  extractHeadingText(n, content),
	}
}

// codeBlocker is an interface for nodes that have line segments.
type codeBlocker interface {
	Lines() *text.Segments
}

// extractCodeBlock extracts code content from a code block node.
func extractCodeBlock(n codeBlocker, content []byte) string {
	var codeContent bytes.Buffer
	lines := n.Lines()
	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)
		codeContent.Write(line.Value(content))
	}
	return codeContent.String()
}

// extractText extracts text content if not inside a code block, table, or list.
func extractText(n *ast.Text, content []byte, builder *strings.Builder) {
	parent := n.Parent()
	if isInsideCodeBlock(parent) || isInsideTable(parent) || isInsideList(parent) {
		return
	}
	builder.Write(n.Segment.Value(content))
	builder.WriteString(" ")
}

// extractString extracts string content if not inside a code block, table, or list.
func extractString(n *ast.String, builder *strings.Builder) {
	parent := n.Parent()
	if isInsideCodeBlock(parent) || isInsideTable(parent) || isInsideList(parent) {
		return
	}
	builder.Write(n.Value)
	builder.WriteString(" ")
}

// stripFrontmatter removes YAML (---) or TOML (+++) frontmatter from content.
// Frontmatter is metadata at the start of a file enclosed in delimiters.
func stripFrontmatter(content []byte) []byte {
	lines := bytes.Split(content, []byte("\n"))

	// Check if file starts with frontmatter delimiter
	firstLine := bytes.TrimSpace(lines[0])
	if !bytes.Equal(firstLine, []byte("---")) && !bytes.Equal(firstLine, []byte("+++")) {
		return content // No frontmatter
	}

	delimiter := firstLine

	// Find closing delimiter (must match opening)
	for i := 1; i < len(lines); i++ {
		if bytes.Equal(bytes.TrimSpace(lines[i]), delimiter) {
			// Found closing delimiter, return everything after it
			if i+1 < len(lines) {
				return bytes.Join(lines[i+1:], []byte("\n"))
			}
			return []byte{}
		}
	}

	// No closing delimiter found, return original content
	return content
}

// stripAdmonitions removes MkDocs-style admonition blocks from content.
// Admonitions are lines starting with !!! followed by indented content.
func stripAdmonitions(content []byte) []byte {
	lines := bytes.Split(content, []byte("\n"))
	var result [][]byte
	i := 0

	for i < len(lines) {
		line := lines[i]
		trimmed := bytes.TrimSpace(line)

		// Check if this is an admonition start
		if bytes.HasPrefix(trimmed, []byte("!!!")) {
			// Skip the !!! line
			i++
			// Skip all following indented lines (admonition content)
			for i < len(lines) {
				nextLine := lines[i]
				// If line is indented (starts with spaces/tabs), it's admonition content
				if len(nextLine) > 0 && (nextLine[0] == ' ' || nextLine[0] == '\t') {
					i++
					continue
				}
				// If line is empty, could be part of admonition or end of it
				if len(bytes.TrimSpace(nextLine)) == 0 {
					i++
					continue
				}
				// Non-indented, non-empty line - end of admonition
				break
			}
			continue
		}

		result = append(result, line)
		i++
	}

	return bytes.Join(result, []byte("\n"))
}

// countLines counts total, code, and empty lines, and detects admonitions.
func countLines(content []byte, result *ParseResult) {
	lines := bytes.Split(content, []byte("\n"))
	result.TotalLines = len(lines)

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
			continue
		}

		if len(trimmed) == 0 {
			result.EmptyLines++
			continue
		}

		// Detect MkDocs-style admonitions: !!! type or !!! type "title"
		if bytes.HasPrefix(trimmed, []byte("!!!")) {
			if adm := parseAdmonition(string(trimmed)); adm != nil {
				adm.Line = lineNum + 1 // 1-based line numbers
				result.Admonitions = append(result.Admonitions, *adm)
			}
		}
	}
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

// isInsideTable checks if a node is inside a table.
func isInsideTable(node ast.Node) bool {
	for node != nil {
		switch node.(type) {
		case *extast.Table, *extast.TableRow, *extast.TableCell, *extast.TableHeader:
			return true
		}
		node = node.Parent()
	}
	return false
}

// isInsideList checks if a node is inside a list item.
func isInsideList(node ast.Node) bool {
	for node != nil {
		switch node.(type) {
		case *ast.List, *ast.ListItem:
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
