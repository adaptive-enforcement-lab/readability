package output

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
)

// mw wraps an io.Writer to simplify error handling for fmt functions.
// Write errors to stdout/file are typically unrecoverable, so we ignore them.
type mw struct {
	w io.Writer
}

func (m mw) println(a ...any) {
	_, _ = fmt.Fprintln(m.w, a...)
}

func (m mw) printf(format string, a ...any) {
	_, _ = fmt.Fprintf(m.w, format, a...)
}

// cleanPath strips relative path prefixes for cleaner display.
func cleanPath(path string) string {
	// Strip leading ../ sequences
	for strings.HasPrefix(path, "../") {
		path = path[3:]
	}
	// Strip leading ./
	path = strings.TrimPrefix(path, "./")
	return path
}

// Markdown writes full results as a GitHub-flavored markdown report.
func Markdown(w io.Writer, results []*analyzer.Result) {
	m := mw{w}
	passed, failed, totalWords, totalLines := aggregateCounts(results)

	// Summary table
	m.println("| Metric | Value |")
	m.println("|--------|------:|")
	m.printf("| Files | %d |\n", len(results))
	m.printf("| Passed | %d |\n", passed)
	m.printf("| Failed | %d |\n", failed)
	m.printf("| Words | %d |\n", totalWords)
	m.printf("| Lines | %d |\n", totalLines)
	m.printf("| Reading time | %d min |\n", totalWords/200)
	m.println()

	// Sort by status (failed first), then by file path
	sorted := make([]*analyzer.Result, len(results))
	copy(sorted, results)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Status != sorted[j].Status {
			return sorted[i].Status == "fail"
		}
		return sorted[i].File < sorted[j].File
	})

	// Results table
	m.println("| Status | File | Lines | Read | FK Grade | ARI | Flesch |")
	m.println("|:------:|------|------:|-----:|---------:|----:|-------:|")

	for _, r := range sorted {
		status := "✅"
		if r.Status == "fail" {
			status = "❌"
		}
		readTime := readingTime(r.Structural.Words)
		m.printf("| %s | %s | %d | %s | %.1f | %.1f | %.1f |\n",
			status,
			cleanPath(r.File),
			r.Structural.Lines,
			readTime,
			r.Readability.FleschKincaidGrade,
			r.Readability.ARI,
			r.Readability.FleschReadingEase,
		)
	}
}

// readingTime formats word count as reading time estimate.
// Uses 200 WPM for technical content with ceiling division.
func readingTime(words int) string {
	if words <= 0 {
		return "<1m"
	}
	// Ceiling division: (words + 199) / 200
	minutes := (words + 199) / 200
	return fmt.Sprintf("%dm", minutes)
}

// Summary writes only an aggregate summary in markdown format.
func Summary(w io.Writer, results []*analyzer.Result) {
	m := mw{w}
	passed, failed, totalWords, totalLines := aggregateCounts(results)

	// Overall status
	if failed == 0 {
		m.println("✅ **All files pass readability checks**")
	} else {
		m.printf("❌ **%d file(s) failed**\n", failed)
	}
	m.println()

	// Summary table
	m.println("| Metric | Value |")
	m.println("|--------|------:|")
	m.printf("| Files | %d |\n", len(results))
	m.printf("| Passed | %d |\n", passed)
	m.printf("| Failed | %d |\n", failed)
	m.printf("| Words | %d |\n", totalWords)
	m.printf("| Lines | %d |\n", totalLines)
	m.printf("| Reading time | %d min |\n", totalWords/200)
	m.println()

	// Failed files list if any
	if failed > 0 {
		m.println("### Failed Files")
		m.println()
		m.println("| File | FK Grade | ARI | Issue |")
		m.println("|------|:--------:|:---:|-------|")

		for _, r := range results {
			if r.Status == "fail" {
				issue := identifyIssue(r)
				m.printf("| %s | %.1f | %.1f | %s |\n",
					cleanPath(r.File),
					r.Readability.FleschKincaidGrade,
					r.Readability.ARI,
					issue,
				)
			}
		}
		m.println()
	}

	// Readability distribution
	m.println("### Distribution")
	m.println()
	dist := calculateDistribution(results)
	m.println("| Level | Count | % |")
	m.println("|-------|------:|--:|")
	for _, d := range dist {
		m.printf("| %s | %d | %.0f |\n", d.Label, d.Count, d.Percent)
	}
}

// Report writes a standalone markdown report suitable for job summaries.
// This is the recommended format for CI integration.
func Report(w io.Writer, results []*analyzer.Result) {
	m := mw{w}
	passed, failed, totalWords, totalLines := aggregateCounts(results)

	m.println("## Documentation Readability")
	m.println()

	// Status badge
	if failed == 0 {
		m.println("✅ All files pass")
	} else {
		m.printf("❌ %d/%d failed\n", failed, len(results))
	}
	m.println()

	// Summary table
	m.println("| Metric | Value |")
	m.println("|--------|------:|")
	m.printf("| Files | %d |\n", len(results))
	m.printf("| Passed | %d |\n", passed)
	m.printf("| Failed | %d |\n", failed)
	m.printf("| Words | %d |\n", totalWords)
	m.printf("| Lines | %d |\n", totalLines)
	m.printf("| Reading time | ~%d min |\n", totalWords/200)
	m.println()

	// Only show failed files in report (keep it concise)
	if failed > 0 {
		m.println("### Files Requiring Attention")
		m.println()
		m.println("| File | FK | ARI | Issue |")
		m.println("|------|---:|----:|-------|")

		for _, r := range results {
			if r.Status == "fail" {
				issue := identifyIssue(r)
				m.printf("| %s | %.1f | %.1f | %s |\n",
					cleanPath(r.File),
					r.Readability.FleschKincaidGrade,
					r.Readability.ARI,
					issue,
				)
			}
		}
		m.println()
	}

	// Distribution (collapsed for brevity)
	m.println("<details>")
	m.println("<summary>Readability Distribution</summary>")
	m.println()
	dist := calculateDistribution(results)
	m.println("| Level | Count | % |")
	m.println("|-------|------:|--:|")
	for _, d := range dist {
		m.printf("| %s | %d | %.0f |\n", d.Label, d.Count, d.Percent)
	}
	m.println("</details>")
}

func aggregateCounts(results []*analyzer.Result) (passed, failed, totalWords, totalLines int) {
	for _, r := range results {
		if r.Status == "pass" {
			passed++
		} else {
			failed++
		}
		totalWords += r.Structural.Words
		totalLines += r.Structural.Lines
	}
	return
}

func identifyIssue(r *analyzer.Result) string {
	issues := []string{}

	if r.Readability.FleschKincaidGrade > 14 {
		issues = append(issues, "Grade level too high")
	}
	if r.Readability.ARI > 14 {
		issues = append(issues, "ARI too high")
	}
	if r.Readability.FleschReadingEase < 30 {
		issues = append(issues, "Reading ease too low")
	}
	if r.Structural.Lines > 375 {
		issues = append(issues, "Too many lines")
	}

	if len(issues) == 0 {
		return "Threshold exceeded"
	}
	return issues[0]
}

type distribution struct {
	Label   string
	Count   int
	Percent float64
}

func calculateDistribution(results []*analyzer.Result) []distribution {
	counts := map[string]int{
		"Very Easy (90+)":          0,
		"Easy (80-89)":             0,
		"Fairly Easy (70-79)":      0,
		"Standard (60-69)":         0,
		"Fairly Difficult (50-59)": 0,
		"Difficult (30-49)":        0,
		"Very Difficult (<30)":     0,
	}

	for _, r := range results {
		score := r.Readability.FleschReadingEase
		switch {
		case score >= 90:
			counts["Very Easy (90+)"]++
		case score >= 80:
			counts["Easy (80-89)"]++
		case score >= 70:
			counts["Fairly Easy (70-79)"]++
		case score >= 60:
			counts["Standard (60-69)"]++
		case score >= 50:
			counts["Fairly Difficult (50-59)"]++
		case score >= 30:
			counts["Difficult (30-49)"]++
		default:
			counts["Very Difficult (<30)"]++
		}
	}

	total := float64(len(results))
	order := []string{
		"Very Easy (90+)",
		"Easy (80-89)",
		"Fairly Easy (70-79)",
		"Standard (60-69)",
		"Fairly Difficult (50-59)",
		"Difficult (30-49)",
		"Very Difficult (<30)",
	}

	dist := make([]distribution, 0, len(order))
	for _, label := range order {
		count := counts[label]
		if count > 0 {
			dist = append(dist, distribution{
				Label:   label,
				Count:   count,
				Percent: float64(count) / total * 100,
			})
		}
	}
	return dist
}
