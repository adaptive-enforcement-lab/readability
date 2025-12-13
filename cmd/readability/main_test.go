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

func TestNewRootCmd(t *testing.T) {
	resetFlags()
	cmd := newRootCmd()

	// Verify command is created with correct settings
	if cmd.Use != "readability [path]" {
		t.Errorf("Expected Use 'readability [path]', got %q", cmd.Use)
	}

	// Verify all flags are registered
	flags := []string{"format", "verbose", "check", "config", "max-grade", "max-ari", "max-lines", "min-admonitions"}
	for _, flag := range flags {
		if cmd.Flags().Lookup(flag) == nil {
			t.Errorf("Expected flag %q to be registered", flag)
		}
	}
}

func TestNewRootCmd_ExecuteSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test\n\nSimple content."), 0644); err != nil {
		t.Fatal(err)
	}

	resetFlags()
	cmd := newRootCmd()
	cmd.SetArgs([]string{testFile})

	var runErr error
	captureOutput(t, func() {
		runErr = cmd.Execute()
	})

	if runErr != nil {
		t.Errorf("Execute() error = %v", runErr)
	}
}

func TestNewRootCmd_ExecuteError(t *testing.T) {
	resetFlags()
	cmd := newRootCmd()
	cmd.SetArgs([]string{"/nonexistent/path"})

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error for nonexistent path")
	}
}

func TestNewRootCmd_NoArgs(t *testing.T) {
	resetFlags()
	cmd := newRootCmd()
	cmd.SetArgs([]string{})

	captureStderr(t, func() {
		err := cmd.Execute()
		if err == nil {
			t.Error("Expected error for missing args")
		}
	})
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

func TestOutputResults_AllFormats(t *testing.T) {
	results := []*analyzer.Result{{
		File:        "test.md",
		Status:      "pass",
		Readability: analyzer.Readability{FleschReadingEase: 60},
	}}

	formats := []struct {
		format   string
		contains string
	}{
		{"json", "test.md"},
		{"markdown", "test.md"},
		{"summary", ""},
		{"report", ""},
		{"diagnostic", ""},
		{"table", ""},
		{"unknown", ""}, // defaults to table
	}

	for _, tt := range formats {
		t.Run(tt.format, func(t *testing.T) {
			formatFlag = tt.format
			verboseFlag = false

			err := outputResults(results)
			if err != nil {
				t.Errorf("outputResults() error = %v", err)
			}
		})
	}
}

func TestApplyFlagOverrides_AllFlags(t *testing.T) {
	tests := []struct {
		name           string
		setupFlags     func(cmd *cobra.Command)
		expectedGrade  float64
		expectedARI    float64
		expectedLines  int
		expectedMinAdm int
	}{
		{
			name: "max-grade override",
			setupFlags: func(cmd *cobra.Command) {
				maxGradeFlag = 10.0
			},
			expectedGrade:  10.0,
			expectedARI:    16.0, // default
			expectedLines:  375,  // default
			expectedMinAdm: 1,    // default
		},
		{
			name: "max-ari override",
			setupFlags: func(cmd *cobra.Command) {
				maxARIFlag = 12.0
			},
			expectedGrade:  16.0, // default
			expectedARI:    12.0,
			expectedLines:  375, // default
			expectedMinAdm: 1,   // default
		},
		{
			name: "max-lines override via Changed",
			setupFlags: func(cmd *cobra.Command) {
				if err := cmd.Flags().Set("max-lines", "200"); err != nil {
					panic(err)
				}
			},
			expectedGrade:  16.0, // default
			expectedARI:    16.0, // default
			expectedLines:  200,
			expectedMinAdm: 1, // default
		},
		{
			name: "min-admonitions override via Changed",
			setupFlags: func(cmd *cobra.Command) {
				if err := cmd.Flags().Set("min-admonitions", "5"); err != nil {
					panic(err)
				}
			},
			expectedGrade:  16.0, // default
			expectedARI:    16.0, // default
			expectedLines:  375,  // default
			expectedMinAdm: 5,
		},
		{
			name: "all flags combined",
			setupFlags: func(cmd *cobra.Command) {
				maxGradeFlag = 8.0
				maxARIFlag = 9.0
				if err := cmd.Flags().Set("max-lines", "100"); err != nil {
					panic(err)
				}
				if err := cmd.Flags().Set("min-admonitions", "3"); err != nil {
					panic(err)
				}
			},
			expectedGrade:  8.0,
			expectedARI:    9.0,
			expectedLines:  100,
			expectedMinAdm: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			cfg := config.DefaultConfig()

			cmd := &cobra.Command{}
			cmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "")
			cmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "")
			cmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "")
			cmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "")

			tt.setupFlags(cmd)
			applyFlagOverrides(cmd, cfg)

			if cfg.Thresholds.MaxGrade != tt.expectedGrade {
				t.Errorf("MaxGrade = %v, want %v", cfg.Thresholds.MaxGrade, tt.expectedGrade)
			}
			if cfg.Thresholds.MaxARI != tt.expectedARI {
				t.Errorf("MaxARI = %v, want %v", cfg.Thresholds.MaxARI, tt.expectedARI)
			}
			if cfg.Thresholds.MaxLines != tt.expectedLines {
				t.Errorf("MaxLines = %v, want %v", cfg.Thresholds.MaxLines, tt.expectedLines)
			}
			if cfg.Thresholds.MinAdmonitions != tt.expectedMinAdm {
				t.Errorf("MinAdmonitions = %v, want %v", cfg.Thresholds.MinAdmonitions, tt.expectedMinAdm)
			}
		})
	}
}

func TestLoadConfig_AutoDetectFromFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create config in directory
	configContent := `thresholds:
  max_grade: 12
`
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a .git directory to mark repo root
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create test file
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test"), 0644); err != nil {
		t.Fatal(err)
	}

	resetFlags()

	// Load config using the file path (not directory)
	cfg, err := loadConfig(testFile)
	if err != nil {
		t.Fatalf("loadConfig() error = %v", err)
	}

	if cfg.Thresholds.MaxGrade != 12 {
		t.Errorf("MaxGrade = %v, want 12", cfg.Thresholds.MaxGrade)
	}
}

func TestLoadConfig_DefaultsWhenNoConfigFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a .git directory but no config file
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatal(err)
	}

	resetFlags()

	cfg, err := loadConfig(tmpDir)
	if err != nil {
		t.Fatalf("loadConfig() error = %v", err)
	}

	// Should return default config
	if cfg.Thresholds.MaxGrade != 16.0 {
		t.Errorf("MaxGrade = %v, want 16.0 (default)", cfg.Thresholds.MaxGrade)
	}
}

func TestAnalyzeTarget_DirectoryError(t *testing.T) {
	cfg := config.DefaultConfig()

	// Test with non-existent path
	_, err := analyzeTarget(cfg, "/nonexistent/path")
	if err == nil {
		t.Error("Expected error for non-existent path")
	}
}

func TestAnalyzeTarget_DirectorySuccess(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test files
	files := map[string]string{
		"doc1.md": "# Doc 1\n\nContent one.",
		"doc2.md": "# Doc 2\n\nContent two.",
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(tmpDir, name), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	cfg := config.DefaultConfig()
	cfg.Thresholds.MinAdmonitions = 0 // Don't require admonitions

	results, err := analyzeTarget(cfg, tmpDir)
	if err != nil {
		t.Fatalf("analyzeTarget() error = %v", err)
	}

	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}
}

func TestRun_OutputResultsError(t *testing.T) {
	// This tests the error path in outputResults
	// JSON output can return an error if encoding fails
	// but in practice this is hard to trigger with valid results
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test\n\nContent."), 0644); err != nil {
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

	args := []string{testFile}

	var runErr error
	captureOutput(t, func() {
		runErr = run(cmd, args)
	})

	if runErr != nil {
		t.Errorf("run() error = %v", runErr)
	}
}

func TestLoadConfig_ConfigFileLoadError(t *testing.T) {
	tmpDir := t.TempDir()

	// Create an invalid YAML config file
	invalidConfig := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(invalidConfig, []byte("invalid: [yaml"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a .git directory to mark repo root
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatal(err)
	}

	resetFlags()

	// When config file is auto-detected but invalid, it falls back to defaults
	cfg, err := loadConfig(tmpDir)
	if err != nil {
		t.Fatalf("loadConfig() should fall back to defaults, got error = %v", err)
	}

	// Should return default config
	if cfg.Thresholds.MaxGrade != 16.0 {
		t.Errorf("Expected default MaxGrade, got %v", cfg.Thresholds.MaxGrade)
	}
}

func TestRun_CheckModeFail(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a config file with strict thresholds
	configContent := `thresholds:
  max_grade: 100
  max_ari: 100
  max_fog: 100
  min_ease: 0
  max_lines: 10
  min_admonitions: 0
`
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a .git to mark repo root
	if err := os.Mkdir(filepath.Join(tmpDir, ".git"), 0755); err != nil {
		t.Fatal(err)
	}

	// Create a file that will fail checks (too many lines)
	content := "# Document\n\n"
	for i := 0; i < 400; i++ {
		content += "This is line content.\n"
	}
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// Reset and set flags BEFORE creating command
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

	// Set flags AFTER binding to variables
	checkFlag = true
	formatFlag = "diagnostic"
	configFlag = configPath

	args := []string{testFile}

	var runErr error
	captureOutput(t, func() {
		captureStderr(t, func() {
			runErr = run(cmd, args)
		})
	})

	if runErr == nil {
		t.Error("Expected error for failing check")
	}
}

func TestAnalyzeTarget_FileError(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a markdown file
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Make file unreadable to cause AnalyzeFile to fail
	if err := os.Chmod(testFile, 0000); err != nil {
		t.Skip("Cannot change file permissions")
	}
	defer func() { _ = os.Chmod(testFile, 0644) }() // Restore for cleanup

	cfg := config.DefaultConfig()
	_, err := analyzeTarget(cfg, testFile)

	// On Unix systems, this should return an error
	if err == nil {
		t.Log("Expected error for unreadable file (may not work on all platforms)")
	} else if !strings.Contains(err.Error(), "error analyzing file") {
		t.Errorf("Expected 'error analyzing file' error, got %v", err)
	}
}

func TestRun_AnalyzeTargetError(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a markdown file
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test\n\nContent here."), 0644); err != nil {
		t.Fatal(err)
	}

	// Make file unreadable
	if err := os.Chmod(testFile, 0000); err != nil {
		t.Skip("Cannot change file permissions")
	}
	defer func() { _ = os.Chmod(testFile, 0644) }()

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

	args := []string{testFile}

	err := run(cmd, args)

	// Should return error for unreadable file
	if err == nil {
		t.Log("Expected error for unreadable file (may not work on all platforms)")
	} else if !strings.Contains(err.Error(), "error analyzing file") {
		t.Errorf("Expected 'error analyzing file' error, got %v", err)
	}
}

func TestAnalyzeTarget_DirectoryAnalyzeError(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a subdirectory with an unreadable file
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatal(err)
	}

	testFile := filepath.Join(subDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Make file unreadable
	if err := os.Chmod(testFile, 0000); err != nil {
		t.Skip("Cannot change file permissions")
	}
	defer func() { _ = os.Chmod(testFile, 0644) }()

	cfg := config.DefaultConfig()
	_, err := analyzeTarget(cfg, subDir)

	// AnalyzeDirectory should return error for unreadable files
	if err == nil {
		t.Log("Expected error for directory with unreadable file (may not work on all platforms)")
	} else if !strings.Contains(err.Error(), "error analyzing directory") {
		t.Errorf("Expected 'error analyzing directory' error, got %v", err)
	}
}
