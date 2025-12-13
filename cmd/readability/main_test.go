package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
	"github.com/adaptive-enforcement-lab/readability/pkg/config"
	"github.com/spf13/cobra"
)

// captureOutput captures stdout during function execution and returns the output
func captureOutput(t *testing.T, fn func()) string {
	t.Helper()
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("Failed to close writer: %v", err)
	}
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}
	os.Stdout = oldStdout

	return buf.String()
}

// captureStderr captures stderr during function execution and returns the output
func captureStderr(t *testing.T, fn func()) string {
	t.Helper()
	oldStderr := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stderr = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("Failed to close writer: %v", err)
	}
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}
	os.Stderr = oldStderr

	return buf.String()
}

func TestRun_SingleFile(t *testing.T) {
	// Create temp file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	content := "# Test Document\n\nThis is a simple test document with clear sentences."
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// Reset flags
	resetFlags()
	formatFlag = "json"

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	args := []string{testFile}

	var runErr error
	output := captureOutput(t, func() {
		runErr = run(cmd, args)
	})

	if runErr != nil {
		t.Errorf("run() error = %v", runErr)
	}

	if !strings.Contains(output, "test.md") {
		t.Errorf("Expected output to contain file name, got %q", output)
	}
}

func TestRun_Directory(t *testing.T) {
	// Create temp directory with files
	tmpDir := t.TempDir()
	files := map[string]string{
		"doc1.md": "# Doc 1\n\nContent one is here.",
		"doc2.md": "# Doc 2\n\nContent two is here.",
	}
	for name, content := range files {
		path := filepath.Join(tmpDir, name)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Reset flags
	resetFlags()
	formatFlag = "json"

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	args := []string{tmpDir}

	var runErr error
	output := captureOutput(t, func() {
		runErr = run(cmd, args)
	})

	if runErr != nil {
		t.Errorf("run() error = %v", runErr)
	}

	if !strings.Contains(output, "doc1.md") {
		t.Errorf("Expected output to contain 'doc1.md'")
	}
	if !strings.Contains(output, "doc2.md") {
		t.Errorf("Expected output to contain 'doc2.md'")
	}
}

func TestRun_NotFound(t *testing.T) {
	resetFlags()

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	args := []string{"/nonexistent/path/file.md"}

	err := run(cmd, args)

	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
	if !strings.Contains(err.Error(), "cannot access") {
		t.Errorf("Expected 'cannot access' error, got %v", err)
	}
}

func TestRun_WithConfig(t *testing.T) {
	// Create temp directory with config and md file
	tmpDir := t.TempDir()

	configContent := `thresholds:
  max_grade: 20
  max_ari: 20
  max_fog: 20
  min_ease: 0
  max_lines: 1000
  min_admonitions: 0
`
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test\n\nSimple content."), 0644); err != nil {
		t.Fatal(err)
	}

	resetFlags()
	configFlag = configPath
	formatFlag = "json"

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	args := []string{testFile}

	var runErr error
	captureOutput(t, func() {
		runErr = run(cmd, args)
	})

	if runErr != nil {
		t.Errorf("run() error = %v", runErr)
	}
}

func TestRun_CheckMode_Pass(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	content := "# Test\n\nThis is a simple test. It is easy to read."
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	resetFlags()
	formatFlag = "json"
	checkFlag = true

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	args := []string{testFile}

	var runErr error
	output := captureOutput(t, func() {
		runErr = run(cmd, args)
	})

	// Should pass since content is simple
	if runErr != nil {
		t.Logf("Output: %s", output)
		t.Errorf("Expected no error for passing check, got %v", runErr)
	}
}

func TestRun_Formats(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	content := "# Test\n\nSimple content."
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	formats := []string{"table", "json", "markdown", "summary", "report", "diagnostic"}

	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			resetFlags()
			formatFlag = format

			cmd := &cobra.Command{}
			cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
			cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
			cmd.Flags().BoolVar(&checkFlag, "check", false, "")
			cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
			cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
			cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
			cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
			cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

			args := []string{testFile}

			var runErr error
			output := captureOutput(t, func() {
				runErr = run(cmd, args)
			})

			if runErr != nil {
				t.Errorf("Format %q: run() error = %v", format, runErr)
			}

			if len(output) == 0 {
				t.Errorf("Format %q: expected non-empty output", format)
			}
		})
	}
}

func TestRun_CLIOverrides(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	content := "# Test\n\nSimple content."
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	resetFlags()
	formatFlag = "json"
	maxGradeFlag = 5.0
	maxARIFlag = 5.0

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	// Mark flags as changed to trigger override behavior
	if err := cmd.Flags().Set("max-lines", "100"); err != nil {
		t.Fatalf("Failed to set max-lines flag: %v", err)
	}
	if err := cmd.Flags().Set("min-admonitions", "0"); err != nil {
		t.Fatalf("Failed to set min-admonitions flag: %v", err)
	}

	args := []string{testFile}

	var runErr error
	captureOutput(t, func() {
		runErr = run(cmd, args)
	})

	if runErr != nil {
		t.Errorf("run() error = %v", runErr)
	}
}

func TestRun_EmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	// Create a subdirectory with no markdown files
	subDir := filepath.Join(tmpDir, "empty")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatal(err)
	}

	resetFlags()
	formatFlag = "json"

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	args := []string{subDir}

	var runErr error
	captureStderr(t, func() {
		runErr = run(cmd, args)
	})

	// Should not error, just print message
	if runErr != nil {
		t.Errorf("run() error = %v for empty dir", runErr)
	}
}

func TestRun_InvalidConfig(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test\n\nContent."), 0644); err != nil {
		t.Fatal(err)
	}

	resetFlags()

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	// Set the config flag to an invalid path AFTER binding
	configFlag = "/nonexistent/config.yml"

	args := []string{testFile}

	err := run(cmd, args)

	if err == nil {
		t.Fatal("Expected error for invalid config path")
	}
	if !strings.Contains(err.Error(), "cannot load config") {
		t.Errorf("Expected 'cannot load config' error, got %v", err)
	}
}

func TestRun_VerboseOutput(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	content := "# Test\n\nSimple content."
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	resetFlags()

	cmd := &cobra.Command{}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "")
	cmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
	cmd.Flags().BoolVar(&checkFlag, "check", false, "")
	cmd.Flags().StringVarP(&configFlag, "config", "c", "", "")
	cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
	cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
	cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
	cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

	// Set flags AFTER binding
	formatFlag = "table"
	verboseFlag = true

	args := []string{testFile}

	var runErr error
	output := captureOutput(t, func() {
		runErr = run(cmd, args)
	})

	if runErr != nil {
		t.Errorf("run() error = %v", runErr)
	}

	// Verbose table output should include additional metrics
	if !strings.Contains(output, "Additional metrics") {
		t.Errorf("Verbose output should contain 'Additional metrics', got: %s", output)
	}
}

// resetFlags resets all global flags to their defaults
func resetFlags() {
	formatFlag = "table"
	verboseFlag = false
	checkFlag = false
	configFlag = ""
	maxGradeFlag = 0
	maxARIFlag = 0
	maxLinesFlag = 0
	minAdmonitionsFlag = -1
}

func TestCountFailures(t *testing.T) {
	tests := []struct {
		name     string
		results  []*analyzer.Result
		minAdm   int
		expected failureStats
	}{
		{
			name:     "no failures",
			results:  []*analyzer.Result{{Status: "pass"}},
			minAdm:   0,
			expected: failureStats{failed: 0},
		},
		{
			name: "too long",
			results: []*analyzer.Result{{
				Status:      "fail",
				Structural:  analyzer.Structural{Lines: 400},
				Readability: analyzer.Readability{FleschReadingEase: 60}, // Valid readability to avoid triggering lowReadability
			}},
			minAdm:   0,
			expected: failureStats{failed: 1, tooLong: 1},
		},
		{
			name: "low readability - high grade",
			results: []*analyzer.Result{{
				Status:      "fail",
				Readability: analyzer.Readability{FleschKincaidGrade: 16},
			}},
			minAdm:   0,
			expected: failureStats{failed: 1, lowReadability: 1},
		},
		{
			name: "low readability - high ARI",
			results: []*analyzer.Result{{
				Status:      "fail",
				Readability: analyzer.Readability{ARI: 16},
			}},
			minAdm:   0,
			expected: failureStats{failed: 1, lowReadability: 1},
		},
		{
			name: "low readability - low ease",
			results: []*analyzer.Result{{
				Status:      "fail",
				Readability: analyzer.Readability{FleschReadingEase: 20},
			}},
			minAdm:   0,
			expected: failureStats{failed: 1, lowReadability: 1},
		},
		{
			name: "missing admonitions",
			results: []*analyzer.Result{{
				Status:      "fail",
				Readability: analyzer.Readability{FleschReadingEase: 60}, // Valid readability to avoid triggering lowReadability
				Admonitions: analyzer.Admonitions{Count: 0},
			}},
			minAdm:   1,
			expected: failureStats{failed: 1, missingAdmonitions: 1},
		},
		{
			name: "multiple issues",
			results: []*analyzer.Result{{
				Status:      "fail",
				Structural:  analyzer.Structural{Lines: 500},
				Readability: analyzer.Readability{FleschKincaidGrade: 18},
				Admonitions: analyzer.Admonitions{Count: 0},
			}},
			minAdm:   1,
			expected: failureStats{failed: 1, tooLong: 1, lowReadability: 1, missingAdmonitions: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.DefaultConfig()
			cfg.Thresholds.MinAdmonitions = tt.minAdm

			stats := countFailures(tt.results, cfg)

			if stats.failed != tt.expected.failed {
				t.Errorf("failed: got %d, want %d", stats.failed, tt.expected.failed)
			}
			if stats.tooLong != tt.expected.tooLong {
				t.Errorf("tooLong: got %d, want %d", stats.tooLong, tt.expected.tooLong)
			}
			if stats.lowReadability != tt.expected.lowReadability {
				t.Errorf("lowReadability: got %d, want %d", stats.lowReadability, tt.expected.lowReadability)
			}
			if stats.missingAdmonitions != tt.expected.missingAdmonitions {
				t.Errorf("missingAdmonitions: got %d, want %d", stats.missingAdmonitions, tt.expected.missingAdmonitions)
			}
		})
	}
}

func TestCheckResults_Fail(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.Thresholds.MinAdmonitions = 0

	results := []*analyzer.Result{{
		Status:     "fail",
		Structural: analyzer.Structural{Lines: 500},
	}}

	err := checkResults(results, cfg)

	if err == nil {
		t.Error("Expected error for failing results")
	}
	if !strings.Contains(err.Error(), "1 file(s) failed") {
		t.Errorf("Expected failure message, got %v", err)
	}
}

func TestCheckResults_Pass(t *testing.T) {
	cfg := config.DefaultConfig()

	results := []*analyzer.Result{{Status: "pass"}}

	err := checkResults(results, cfg)

	if err != nil {
		t.Errorf("Expected no error for passing results, got %v", err)
	}
}

func TestPrintFailureGuidance(t *testing.T) {
	tests := []struct {
		name     string
		stats    failureStats
		contains []string
	}{
		{
			name:     "too long",
			stats:    failureStats{failed: 1, tooLong: 1},
			contains: []string{"IMPORTANT:", "SPLIT"},
		},
		{
			name:     "low readability",
			stats:    failureStats{failed: 1, lowReadability: 1},
			contains: []string{"READABILITY:", "Break long sentences"},
		},
		{
			name:     "missing admonitions",
			stats:    failureStats{failed: 1, missingAdmonitions: 1},
			contains: []string{"ADMONITIONS:", "!!! note"},
		},
		{
			name:     "all issues",
			stats:    failureStats{failed: 1, tooLong: 1, lowReadability: 1, missingAdmonitions: 1},
			contains: []string{"IMPORTANT:", "READABILITY:", "ADMONITIONS:"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureStderr(t, func() {
				printFailureGuidance(tt.stats)
			})

			for _, want := range tt.contains {
				if !strings.Contains(output, want) {
					t.Errorf("Expected output to contain %q, got %q", want, output)
				}
			}
		})
	}
}

func TestPrintLengthGuidance(t *testing.T) {
	output := captureStderr(t, func() {
		printLengthGuidance()
	})

	if !strings.Contains(output, "SPLIT") {
		t.Errorf("Expected length guidance to mention SPLIT, got %q", output)
	}
}

func TestPrintReadabilityGuidance(t *testing.T) {
	output := captureStderr(t, func() {
		printReadabilityGuidance()
	})

	if !strings.Contains(output, "Break long sentences") {
		t.Errorf("Expected readability guidance, got %q", output)
	}
}

func TestPrintAdmonitionGuidance(t *testing.T) {
	output := captureStderr(t, func() {
		printAdmonitionGuidance()
	})

	if !strings.Contains(output, "!!! note") {
		t.Errorf("Expected admonition guidance, got %q", output)
	}
}
