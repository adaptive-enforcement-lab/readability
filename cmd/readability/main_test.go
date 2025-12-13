package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

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
