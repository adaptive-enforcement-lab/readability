package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
)

// --- JSON Tests ---

func TestJSON_EmptyResults(t *testing.T) {
	var buf bytes.Buffer
	err := JSON(&buf, []*analyzer.Result{})
	if err != nil {
		t.Fatalf("JSON() error = %v", err)
	}

	// Should produce empty array
	output := strings.TrimSpace(buf.String())
	if output != "[]" {
		t.Errorf("Expected [], got %q", output)
	}
}

func TestJSON_SingleResult(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "test.md",
			Status: "pass",
			Structural: analyzer.Structural{
				Lines:              50,
				Words:              200,
				Sentences:          10,
				Characters:         1000,
				ReadingTimeMinutes: 1,
			},
			Readability: analyzer.Readability{
				FleschKincaidGrade: 8.5,
				FleschReadingEase:  65.0,
				ARI:                9.0,
			},
		},
	}

	var buf bytes.Buffer
	err := JSON(&buf, results)
	if err != nil {
		t.Fatalf("JSON() error = %v", err)
	}

	// Parse output to verify it's valid JSON
	var parsed []*analyzer.Result
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("JSON output not valid: %v", err)
	}

	if len(parsed) != 1 {
		t.Errorf("Expected 1 result, got %d", len(parsed))
	}
	if parsed[0].File != "test.md" {
		t.Errorf("Expected file 'test.md', got %q", parsed[0].File)
	}
}

func TestJSON_MultipleResults(t *testing.T) {
	results := []*analyzer.Result{
		{File: "doc1.md", Status: "pass"},
		{File: "doc2.md", Status: "fail"},
	}

	var buf bytes.Buffer
	err := JSON(&buf, results)
	if err != nil {
		t.Fatalf("JSON() error = %v", err)
	}

	var parsed []*analyzer.Result
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("JSON output not valid: %v", err)
	}

	if len(parsed) != 2 {
		t.Errorf("Expected 2 results, got %d", len(parsed))
	}
}

// --- Table Tests ---

func TestTable_SingleResult(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "test.md",
			Status: "pass",
			Structural: analyzer.Structural{
				Lines:              50,
				Words:              200,
				ReadingTimeMinutes: 1,
			},
			Headings: analyzer.Headings{H1: 1, H2: 3, H3: 2},
			Readability: analyzer.Readability{
				FleschKincaidGrade: 8.5,
				FleschReadingEase:  65.0,
				ARI:                9.0,
			},
			Composition: analyzer.Composition{
				CodeBlockRatio: 0.2,
				ProseLines:     40,
			},
		},
	}

	var buf bytes.Buffer
	Table(&buf, results, false)

	output := buf.String()

	// Check file name is present
	if !strings.Contains(output, "test.md") {
		t.Errorf("Expected file name in output, got %q", output)
	}
	// Check structural metrics
	if !strings.Contains(output, "Lines: 50") {
		t.Errorf("Expected 'Lines: 50' in output")
	}
	// Check status
	if !strings.Contains(output, "PASS") {
		t.Errorf("Expected 'PASS' in output")
	}
}

func TestTable_Verbose(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "test.md",
			Status: "pass",
			Readability: analyzer.Readability{
				ColemanLiau: 10.5,
				GunningFog:  12.0,
				SMOG:        11.0,
			},
			Structural: analyzer.Structural{
				Sentences:  15,
				Characters: 500,
			},
		},
	}

	var buf bytes.Buffer
	Table(&buf, results, true)

	output := buf.String()

	// Verbose should include additional metrics
	if !strings.Contains(output, "Coleman-Liau") {
		t.Errorf("Verbose should include Coleman-Liau")
	}
	if !strings.Contains(output, "Gunning Fog") {
		t.Errorf("Verbose should include Gunning Fog")
	}
	if !strings.Contains(output, "SMOG") {
		t.Errorf("Verbose should include SMOG")
	}
}

func TestTable_MultipleSummary(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "doc1.md",
			Status: "pass",
			Structural: analyzer.Structural{
				Lines: 50,
				Words: 200,
			},
		},
		{
			File:   "doc2.md",
			Status: "fail",
			Structural: analyzer.Structural{
				Lines: 100,
				Words: 400,
			},
		},
	}

	var buf bytes.Buffer
	Table(&buf, results, false)

	output := buf.String()

	// Multiple files should have summary
	if !strings.Contains(output, "Summary") {
		t.Errorf("Multiple files should show summary")
	}
	if !strings.Contains(output, "Passed: 1") {
		t.Errorf("Summary should show passed count")
	}
	if !strings.Contains(output, "Failed: 1") {
		t.Errorf("Summary should show failed count")
	}
	if !strings.Contains(output, "600 words") {
		t.Errorf("Summary should show total words")
	}
}

func TestReadabilityLabel(t *testing.T) {
	tests := []struct {
		score float64
		want  string
	}{
		{95.0, "Very Easy"},
		{90.0, "Very Easy"},
		{85.0, "Easy"},
		{80.0, "Easy"},
		{75.0, "Fairly Easy"},
		{70.0, "Fairly Easy"},
		{65.0, "Standard"},
		{60.0, "Standard"},
		{55.0, "Fairly Difficult"},
		{50.0, "Fairly Difficult"},
		{45.0, "Difficult"},
		{30.0, "Difficult"},
		{25.0, "Very Difficult"},
		{0.0, "Very Difficult"},
		{-10.0, "Very Difficult"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := readabilityLabel(tt.score)
			if got != tt.want {
				t.Errorf("readabilityLabel(%.1f) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

// --- Markdown Output Tests ---

func TestMarkdown_Output(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "doc1.md",
			Status: "pass",
			Structural: analyzer.Structural{
				Lines: 50,
				Words: 200,
			},
			Readability: analyzer.Readability{
				FleschKincaidGrade: 8.5,
				FleschReadingEase:  65.0,
				ARI:                9.0,
			},
		},
		{
			File:   "doc2.md",
			Status: "fail",
			Structural: analyzer.Structural{
				Lines: 100,
				Words: 400,
			},
			Readability: analyzer.Readability{
				FleschKincaidGrade: 15.5,
				FleschReadingEase:  35.0,
				ARI:                16.0,
			},
			Diagnostics: []analyzer.Diagnostic{
				{Rule: "readability/grade-level"},
			},
		},
	}

	var buf bytes.Buffer
	Markdown(&buf, results)

	output := buf.String()

	// Check summary table
	if !strings.Contains(output, "| Files | 2 |") {
		t.Errorf("Expected '| Files | 2 |' in output")
	}
	// Check results table headers
	if !strings.Contains(output, "| Status | File | Lines |") {
		t.Errorf("Expected results table headers")
	}
	// Check pass emoji
	if !strings.Contains(output, "✅") {
		t.Errorf("Expected ✅ for passing file")
	}
	// Check fail emoji
	if !strings.Contains(output, "❌") {
		t.Errorf("Expected ❌ for failing file")
	}
}

func TestSummary_AllPass(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "doc1.md",
			Status: "pass",
			Structural: analyzer.Structural{Lines: 50, Words: 200},
			Readability: analyzer.Readability{FleschReadingEase: 75.0},
		},
	}

	var buf bytes.Buffer
	Summary(&buf, results)

	output := buf.String()
	if !strings.Contains(output, "All files pass") {
		t.Errorf("Expected 'All files pass' message")
	}
}

func TestSummary_WithFailures(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "bad.md",
			Status: "fail",
			Structural: analyzer.Structural{Lines: 50, Words: 200},
			Readability: analyzer.Readability{
				FleschKincaidGrade: 18.0,
				FleschReadingEase:  25.0,
				ARI:                18.0,
			},
		},
	}

	var buf bytes.Buffer
	Summary(&buf, results)

	output := buf.String()
	if !strings.Contains(output, "1 file(s) failed") {
		t.Errorf("Expected failure count")
	}
	if !strings.Contains(output, "Failed Files") {
		t.Errorf("Expected 'Failed Files' section")
	}
}

func TestReport_Output(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "doc.md",
			Status: "pass",
			Structural: analyzer.Structural{Lines: 50, Words: 200},
			Readability: analyzer.Readability{FleschReadingEase: 65.0},
		},
	}

	var buf bytes.Buffer
	Report(&buf, results)

	output := buf.String()
	if !strings.Contains(output, "## Documentation Readability") {
		t.Errorf("Expected header")
	}
	if !strings.Contains(output, "All files pass") {
		t.Errorf("Expected pass message")
	}
	if !strings.Contains(output, "<details>") {
		t.Errorf("Expected collapsible distribution")
	}
}

func TestReport_WithFailures(t *testing.T) {
	results := []*analyzer.Result{
		{
			File:   "bad.md",
			Status: "fail",
			Structural: analyzer.Structural{Lines: 500, Words: 200},
			Readability: analyzer.Readability{
				FleschKincaidGrade: 18.0,
				FleschReadingEase:  25.0,
				ARI:                18.0,
			},
		},
	}

	var buf bytes.Buffer
	Report(&buf, results)

	output := buf.String()
	if !strings.Contains(output, "1/1 failed") {
		t.Errorf("Expected failure count, got %q", output)
	}
	if !strings.Contains(output, "Files Requiring Attention") {
		t.Errorf("Expected attention section")
	}
}

func TestAggregateCounts(t *testing.T) {
	results := []*analyzer.Result{
		{Status: "pass", Structural: analyzer.Structural{Words: 100, Lines: 20}},
		{Status: "pass", Structural: analyzer.Structural{Words: 200, Lines: 40}},
		{Status: "fail", Structural: analyzer.Structural{Words: 150, Lines: 30}},
	}

	passed, failed, totalWords, totalLines := aggregateCounts(results)

	if passed != 2 {
		t.Errorf("Passed = %d, want 2", passed)
	}
	if failed != 1 {
		t.Errorf("Failed = %d, want 1", failed)
	}
	if totalWords != 450 {
		t.Errorf("TotalWords = %d, want 450", totalWords)
	}
	if totalLines != 90 {
		t.Errorf("TotalLines = %d, want 90", totalLines)
	}
}

func TestIdentifyIssue(t *testing.T) {
	tests := []struct {
		name   string
		result *analyzer.Result
		want   string
	}{
		{
			name: "high grade level",
			result: &analyzer.Result{
				Readability: analyzer.Readability{
					FleschKincaidGrade: 18.0,
					ARI:                10.0,
					FleschReadingEase:  50.0,
				},
				Structural: analyzer.Structural{Lines: 50},
			},
			want: "Grade level too high",
		},
		{
			name: "high ARI",
			result: &analyzer.Result{
				Readability: analyzer.Readability{
					FleschKincaidGrade: 10.0,
					ARI:                18.0,
					FleschReadingEase:  50.0,
				},
				Structural: analyzer.Structural{Lines: 50},
			},
			want: "ARI too high",
		},
		{
			name: "low reading ease",
			result: &analyzer.Result{
				Readability: analyzer.Readability{
					FleschKincaidGrade: 10.0,
					ARI:                10.0,
					FleschReadingEase:  25.0,
				},
				Structural: analyzer.Structural{Lines: 50},
			},
			want: "Reading ease too low",
		},
		{
			name: "too many lines",
			result: &analyzer.Result{
				Readability: analyzer.Readability{
					FleschKincaidGrade: 10.0,
					ARI:                10.0,
					FleschReadingEase:  50.0,
				},
				Structural: analyzer.Structural{Lines: 500},
			},
			want: "Too many lines",
		},
		{
			name: "no obvious issue",
			result: &analyzer.Result{
				Readability: analyzer.Readability{
					FleschKincaidGrade: 10.0,
					ARI:                10.0,
					FleschReadingEase:  50.0,
				},
				Structural: analyzer.Structural{Lines: 50},
			},
			want: "Threshold exceeded",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := identifyIssue(tt.result)
			if got != tt.want {
				t.Errorf("identifyIssue() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCalculateDistribution(t *testing.T) {
	results := []*analyzer.Result{
		{Readability: analyzer.Readability{FleschReadingEase: 95.0}},  // Very Easy
		{Readability: analyzer.Readability{FleschReadingEase: 85.0}},  // Easy
		{Readability: analyzer.Readability{FleschReadingEase: 75.0}},  // Fairly Easy
		{Readability: analyzer.Readability{FleschReadingEase: 65.0}},  // Standard
		{Readability: analyzer.Readability{FleschReadingEase: 55.0}},  // Fairly Difficult
		{Readability: analyzer.Readability{FleschReadingEase: 40.0}},  // Difficult
		{Readability: analyzer.Readability{FleschReadingEase: 20.0}},  // Very Difficult
	}

	dist := calculateDistribution(results)

	// Should have 7 categories with 1 each
	if len(dist) != 7 {
		t.Errorf("Expected 7 distribution entries, got %d", len(dist))
	}

	// Each should have count 1 and ~14.3% (1/7)
	for _, d := range dist {
		if d.Count != 1 {
			t.Errorf("Expected count 1 for %s, got %d", d.Label, d.Count)
		}
		if d.Percent < 14.0 || d.Percent > 15.0 {
			t.Errorf("Expected ~14%% for %s, got %.1f", d.Label, d.Percent)
		}
	}
}

func TestCalculateDistribution_Empty(t *testing.T) {
	dist := calculateDistribution([]*analyzer.Result{})

	if len(dist) != 0 {
		t.Errorf("Expected empty distribution for empty results, got %d", len(dist))
	}
}

func TestCleanPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"./docs/file.md", "docs/file.md"},
		{"../docs/file.md", "docs/file.md"},
		{"../../docs/file.md", "docs/file.md"},
		{"docs/file.md", "docs/file.md"},
		{"file.md", "file.md"},
		{"./file.md", "file.md"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := cleanPath(tt.path)
			if got != tt.want {
				t.Errorf("cleanPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}

func TestMw_Methods(t *testing.T) {
	var buf bytes.Buffer
	m := mw{&buf}

	m.println("hello")
	m.printf("world %d", 42)

	output := buf.String()
	if !strings.Contains(output, "hello") {
		t.Errorf("Expected 'hello' in output")
	}
	if !strings.Contains(output, "world 42") {
		t.Errorf("Expected 'world 42' in output")
	}
}
