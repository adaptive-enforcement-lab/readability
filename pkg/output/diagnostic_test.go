package output

import (
	"bytes"
	"strings"
	"testing"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
)

func TestDiagnostic_Format(t *testing.T) {
	results := []*analyzer.Result{
		{
			File: "docs/example.md",
			Diagnostics: []analyzer.Diagnostic{
				{
					Line:     1,
					Column:   1,
					Severity: analyzer.SeverityError,
					Rule:     "readability/grade-level",
					Message:  "Flesch-Kincaid grade 18.5 exceeds threshold 16.0",
				},
				{
					Line:     45,
					Column:   0, // Should default to 1
					Severity: analyzer.SeverityWarning,
					Rule:     "content/admonitions",
					Message:  "Found 0 admonitions, minimum required is 1",
				},
			},
		},
		{
			File: "docs/api/reference.md",
			Diagnostics: []analyzer.Diagnostic{
				{
					Line:     1,
					Column:   1,
					Severity: analyzer.SeverityError,
					Rule:     "structure/max-lines",
					Message:  "450 lines exceeds threshold 375",
				},
			},
		},
	}

	var buf bytes.Buffer
	Diagnostic(&buf, results)

	output := buf.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Should have 3 diagnostics total
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d: %q", len(lines), output)
	}

	// Check format: file:line:col: severity: message (rule-id)
	expectedPatterns := []struct {
		file     string
		line     string
		col      string
		severity string
		rule     string
	}{
		{"docs/api/reference.md", "1", "1", "error", "structure/max-lines"},
		{"docs/example.md", "1", "1", "error", "readability/grade-level"},
		{"docs/example.md", "45", "1", "warning", "content/admonitions"}, // col defaults to 1
	}

	for i, expected := range expectedPatterns {
		if i >= len(lines) {
			t.Errorf("Missing expected line %d", i)
			continue
		}
		line := lines[i]

		if !strings.HasPrefix(line, expected.file+":") {
			t.Errorf("Line %d: expected file %q, got %q", i, expected.file, line)
		}
		if !strings.Contains(line, ":"+expected.line+":"+expected.col+":") {
			t.Errorf("Line %d: expected line:col %s:%s, got %q", i, expected.line, expected.col, line)
		}
		if !strings.Contains(line, expected.severity+":") {
			t.Errorf("Line %d: expected severity %q, got %q", i, expected.severity, line)
		}
		if !strings.Contains(line, "("+expected.rule+")") {
			t.Errorf("Line %d: expected rule %q, got %q", i, expected.rule, line)
		}
	}
}

func TestDiagnostic_SortsByFile(t *testing.T) {
	results := []*analyzer.Result{
		{
			File: "z-file.md",
			Diagnostics: []analyzer.Diagnostic{
				{Line: 1, Severity: analyzer.SeverityError, Rule: "test", Message: "z-file issue"},
			},
		},
		{
			File: "a-file.md",
			Diagnostics: []analyzer.Diagnostic{
				{Line: 1, Severity: analyzer.SeverityError, Rule: "test", Message: "a-file issue"},
			},
		},
	}

	var buf bytes.Buffer
	Diagnostic(&buf, results)

	output := buf.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")

	if len(lines) != 2 {
		t.Fatalf("Expected 2 lines, got %d", len(lines))
	}

	// a-file should come before z-file
	if !strings.HasPrefix(lines[0], "a-file.md:") {
		t.Errorf("Expected a-file.md first, got %q", lines[0])
	}
	if !strings.HasPrefix(lines[1], "z-file.md:") {
		t.Errorf("Expected z-file.md second, got %q", lines[1])
	}
}

func TestDiagnostic_EmptyResults(t *testing.T) {
	var buf bytes.Buffer
	Diagnostic(&buf, []*analyzer.Result{})

	if buf.Len() != 0 {
		t.Errorf("Expected empty output for empty results, got %q", buf.String())
	}
}

func TestDiagnostic_NoDiagnostics(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:        "docs/good-file.md",
			Diagnostics: []analyzer.Diagnostic{}, // No issues
		},
	}

	var buf bytes.Buffer
	Diagnostic(&buf, results)

	if buf.Len() != 0 {
		t.Errorf("Expected empty output for file with no diagnostics, got %q", buf.String())
	}
}

func TestDiagnosticSummary_Counts(t *testing.T) {
	results := []*analyzer.Result{
		{
			File: "test.md",
			Diagnostics: []analyzer.Diagnostic{
				{Severity: analyzer.SeverityError, Rule: "test", Message: "error 1"},
				{Severity: analyzer.SeverityError, Rule: "test", Message: "error 2"},
				{Severity: analyzer.SeverityWarning, Rule: "test", Message: "warning 1"},
				{Severity: analyzer.SeverityInfo, Rule: "test", Message: "info 1"},
			},
		},
	}

	var buf bytes.Buffer
	DiagnosticSummary(&buf, results)

	output := buf.String()

	if !strings.Contains(output, "4 issue(s)") {
		t.Errorf("Expected total count '4 issue(s)', got %q", output)
	}
	if !strings.Contains(output, "2 error(s)") {
		t.Errorf("Expected '2 error(s)', got %q", output)
	}
	if !strings.Contains(output, "1 warning(s)") {
		t.Errorf("Expected '1 warning(s)', got %q", output)
	}
	if !strings.Contains(output, "1 info") {
		t.Errorf("Expected '1 info', got %q", output)
	}
}

func TestDiagnosticSummary_NoIssues(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:        "test.md",
			Diagnostics: []analyzer.Diagnostic{},
		},
	}

	var buf bytes.Buffer
	DiagnosticSummary(&buf, results)

	output := buf.String()
	if !strings.Contains(output, "No issues found") {
		t.Errorf("Expected 'No issues found', got %q", output)
	}
}

func TestDiagnostic_PathCleaning(t *testing.T) {
	results := []*analyzer.Result{
		{
			File: "./docs/example.md",
			Diagnostics: []analyzer.Diagnostic{
				{Line: 1, Severity: analyzer.SeverityError, Rule: "test", Message: "test"},
			},
		},
	}

	var buf bytes.Buffer
	Diagnostic(&buf, results)

	output := buf.String()

	// Path should be cleaned (no ./ prefix)
	if strings.Contains(output, "./docs/") {
		t.Errorf("Path should be cleaned, got %q", output)
	}
	if !strings.HasPrefix(output, "docs/example.md:") {
		t.Errorf("Expected clean path 'docs/example.md', got %q", output)
	}
}
