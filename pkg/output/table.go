package output

import (
	"io"
	"strings"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
)

// Table writes results in human-readable table format.
func Table(w io.Writer, results []*analyzer.Result, verbose bool) {
	m := mw{w}
	for _, r := range results {
		writeFileResult(m, r, verbose)
		m.println()
	}

	if len(results) > 1 {
		writeSummary(m, results)
	}
}

func writeFileResult(m mw, r *analyzer.Result, verbose bool) {
	// File path
	m.printf("%s\n", r.File)

	// Basic metrics
	m.printf("  Lines: %d | Words: %d | Reading time: %d min\n",
		r.Structural.Lines,
		r.Structural.Words,
		r.Structural.ReadingTimeMinutes,
	)

	// Headings
	m.printf("  Headers: H1=%d H2=%d H3=%d H4=%d\n",
		r.Headings.H1,
		r.Headings.H2,
		r.Headings.H3,
		r.Headings.H4,
	)

	// Readability scores
	ease := readabilityLabel(r.Readability.FleschReadingEase)
	m.printf("  Readability: FK=%.1f ARI=%.1f Flesch=%.1f (%s)\n",
		r.Readability.FleschKincaidGrade,
		r.Readability.ARI,
		r.Readability.FleschReadingEase,
		ease,
	)

	// Composition
	m.printf("  Code: %.0f%% | Prose: %d lines\n",
		r.Composition.CodeBlockRatio*100,
		r.Composition.ProseLines,
	)

	// Status
	status := strings.ToUpper(r.Status)
	m.printf("  Status: %s\n", status)

	if verbose {
		m.printf("  Additional metrics:\n")
		m.printf("    Coleman-Liau: %.1f\n", r.Readability.ColemanLiau)
		m.printf("    Gunning Fog: %.1f\n", r.Readability.GunningFog)
		m.printf("    SMOG: %.1f\n", r.Readability.SMOG)
		m.printf("    Sentences: %d\n", r.Structural.Sentences)
		m.printf("    Characters: %d\n", r.Structural.Characters)
	}
}

func writeSummary(m mw, results []*analyzer.Result) {
	m.println("---")
	m.printf("Summary: %d files analyzed\n", len(results))

	passed := 0
	failed := 0
	totalWords := 0
	totalLines := 0

	for _, r := range results {
		if r.Status == "pass" {
			passed++
		} else {
			failed++
		}
		totalWords += r.Structural.Words
		totalLines += r.Structural.Lines
	}

	m.printf("  Passed: %d | Failed: %d\n", passed, failed)
	m.printf("  Total: %d words, %d lines\n", totalWords, totalLines)
}

// readabilityLabel converts Flesch Reading Ease score to human label.
func readabilityLabel(score float64) string {
	switch {
	case score >= 90:
		return "Very Easy"
	case score >= 80:
		return "Easy"
	case score >= 70:
		return "Fairly Easy"
	case score >= 60:
		return "Standard"
	case score >= 50:
		return "Fairly Difficult"
	case score >= 30:
		return "Difficult"
	default:
		return "Very Difficult"
	}
}
