package output

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
)

// Diagnostic writes results in linter/LSP-style diagnostic format.
// Format: file:line:col: severity: message (rule-id)
func Diagnostic(w io.Writer, results []*analyzer.Result) {
	m := mw{w}

	// Sort results by file path for consistent output
	sorted := make([]*analyzer.Result, len(results))
	copy(sorted, results)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].File < sorted[j].File
	})

	for _, r := range sorted {
		path := cleanPath(r.File)
		for _, d := range r.Diagnostics {
			col := d.Column
			if col == 0 {
				col = 1 // Default to column 1 if not specified
			}
			m.printf("%s:%d:%d: %s: %s (%s)\n",
				path,
				d.Line,
				col,
				d.Severity,
				d.Message,
				d.Rule,
			)
		}
	}
}

// DiagnosticSummary writes a summary line after diagnostics.
func DiagnosticSummary(w io.Writer, results []*analyzer.Result) {
	m := mw{w}

	var errors, warnings, infos int
	for _, r := range results {
		for _, d := range r.Diagnostics {
			switch d.Severity {
			case analyzer.SeverityError:
				errors++
			case analyzer.SeverityWarning:
				warnings++
			case analyzer.SeverityInfo:
				infos++
			}
		}
	}

	total := errors + warnings + infos
	if total == 0 {
		m.println("No issues found.")
		return
	}

	parts := make([]string, 0, 3)
	if errors > 0 {
		parts = append(parts, fmt.Sprintf("%d error(s)", errors))
	}
	if warnings > 0 {
		parts = append(parts, fmt.Sprintf("%d warning(s)", warnings))
	}
	if infos > 0 {
		parts = append(parts, fmt.Sprintf("%d info", infos))
	}

	m.printf("\n%d issue(s): %s\n", total, strings.Join(parts, ", "))
}
