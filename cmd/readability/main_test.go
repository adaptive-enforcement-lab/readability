package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

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

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := run(cmd, args)

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = oldStdout

	if err != nil {
		t.Errorf("run() error = %v", err)
	}

	output := buf.String()
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

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := run(cmd, args)

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = oldStdout

	if err != nil {
		t.Errorf("run() error = %v", err)
	}

	output := buf.String()
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

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := run(cmd, args)

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = oldStdout

	if err != nil {
		t.Errorf("run() error = %v", err)
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

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := run(cmd, args)

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = oldStdout

	// Should pass since content is simple
	if err != nil {
		t.Logf("Output: %s", buf.String())
		t.Errorf("Expected no error for passing check, got %v", err)
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

			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := run(cmd, args)

			w.Close()
			var buf bytes.Buffer
			buf.ReadFrom(r)
			os.Stdout = oldStdout

			if err != nil {
				t.Errorf("Format %q: run() error = %v", format, err)
			}

			if buf.Len() == 0 {
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
	cmd.Flags().Set("max-lines", "100")
	cmd.Flags().Set("min-admonitions", "0")

	args := []string{testFile}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := run(cmd, args)

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = oldStdout

	if err != nil {
		t.Errorf("run() error = %v", err)
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

	// Capture stderr for "No markdown files found" message
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	err := run(cmd, args)

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stderr = oldStderr

	// Should not error, just print message
	if err != nil {
		t.Errorf("run() error = %v for empty dir", err)
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

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := run(cmd, args)

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = oldStdout

	if err != nil {
		t.Errorf("run() error = %v", err)
	}

	output := buf.String()
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
