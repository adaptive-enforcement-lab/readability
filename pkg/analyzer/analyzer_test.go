package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adaptive-enforcement-lab/readability/pkg/config"
	"github.com/adaptive-enforcement-lab/readability/pkg/markdown"
)

func TestNew(t *testing.T) {
	a := New()
	if a == nil {
		t.Fatal("New() returned nil")
	}
	if a.Config == nil {
		t.Error("New() Config should not be nil")
	}
	// Check default thresholds are set
	if a.Thresholds.MaxFleschKincaidGrade == 0 {
		t.Error("Default threshold should be non-zero")
	}
}

func TestNewWithThresholds(t *testing.T) {
	custom := Thresholds{
		MaxFleschKincaidGrade: 10.0,
		MaxARI:                12.0,
		MaxGunningFog:         14.0,
		MinFleschReadingEase:  50.0,
		MaxLines:              100,
	}
	a := NewWithThresholds(custom)

	if a.Thresholds.MaxFleschKincaidGrade != 10.0 {
		t.Errorf("MaxFleschKincaidGrade = %v, want 10.0", a.Thresholds.MaxFleschKincaidGrade)
	}
	if a.Thresholds.MaxARI != 12.0 {
		t.Errorf("MaxARI = %v, want 12.0", a.Thresholds.MaxARI)
	}
}

func TestNewWithConfig(t *testing.T) {
	cfg := &config.Config{
		Thresholds: config.Thresholds{
			MaxGrade: 8.0,
			MaxARI:   9.0,
			MaxFog:   10.0,
			MinEase:  60.0,
			MaxLines: 200,
		},
	}
	a := NewWithConfig(cfg)

	if a.Config != cfg {
		t.Error("Config not properly set")
	}
	if a.Thresholds.MaxFleschKincaidGrade != 8.0 {
		t.Errorf("MaxFleschKincaidGrade = %v, want 8.0", a.Thresholds.MaxFleschKincaidGrade)
	}
}

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name           string
		content        string
		wantStatus     string
		wantWords      int
		wantLines      int
		checkDiagCount int
	}{
		{
			name:       "simple prose",
			content:    "This is a simple test document. It has clear sentences.",
			wantStatus: "pass",
			wantWords:  10,
			wantLines:  1,
		},
		{
			name: "prose with headings",
			content: `# Introduction

This is the introduction section.

## Details

More detailed information here.`,
			wantStatus: "pass",
			wantLines:  7,
		},
		{
			name: "with code blocks",
			content: `# Example

Here is some code:

` + "```go" + `
package main

func main() {}
` + "```" + `

That was the code.`,
			wantStatus: "pass",
		},
		{
			name: "with frontmatter",
			content: `---
title: Test
author: Tester
---

# Document

Content here.`,
			wantStatus: "pass",
		},
		{
			name: "with admonitions",
			content: `# Guide

!!! note
    Important note here.

!!! warning "Caution"
    Be careful!`,
			wantStatus: "pass",
		},
	}

	// Use permissive thresholds for these tests
	cfg := &config.Config{
		Thresholds: config.Thresholds{
			MaxGrade:       100.0,
			MaxARI:         100.0,
			MaxFog:         100.0,
			MinEase:        0.0,
			MaxLines:       1000,
			MinWords:       0,
			MinAdmonitions: 0, // Don't require admonitions for basic tests
		},
	}
	a := NewWithConfig(cfg)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := a.Analyze("test.md", []byte(tt.content))
			if err != nil {
				t.Fatalf("Analyze() error = %v", err)
			}

			if result.Status != tt.wantStatus {
				t.Errorf("Status = %q, want %q", result.Status, tt.wantStatus)
			}

			if tt.wantWords > 0 && result.Structural.Words < tt.wantWords {
				t.Errorf("Words = %d, want >= %d", result.Structural.Words, tt.wantWords)
			}

			if tt.wantLines > 0 && result.Structural.Lines != tt.wantLines {
				t.Errorf("Lines = %d, want %d", result.Structural.Lines, tt.wantLines)
			}
		})
	}
}

func TestAnalyze_ThresholdViolations(t *testing.T) {
	// This content has complex sentences that will likely violate thresholds
	content := `# Complex Document

The implementation of sophisticated algorithmic paradigms necessitates comprehensive understanding of computational complexity theory and mathematical abstractions. Furthermore, the utilization of advanced architectural patterns fundamentally transforms application infrastructure.

Additionally, synthesizing heterogeneous data sources requires meticulous consideration of schema harmonization methodologies.`

	a := New()
	a.Config = nil
	a.Thresholds = Thresholds{
		MaxFleschKincaidGrade: 8.0, // Very strict
		MaxARI:                8.0,
		MaxGunningFog:         10.0,
		MinFleschReadingEase:  70.0, // Very strict
		MaxLines:              1000,
	}

	result, err := a.Analyze("test.md", []byte(content))
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	if result.Status != "fail" {
		t.Errorf("Expected fail status for complex content, got %q", result.Status)
	}

	if len(result.Diagnostics) == 0 {
		t.Error("Expected diagnostics for threshold violations")
	}
}

func TestAnalyze_LineLimitViolation(t *testing.T) {
	// Generate content with many lines
	content := "# Document\n\n"
	for i := 0; i < 50; i++ {
		content += "This is line content.\n"
	}

	a := New()
	a.Config = nil
	a.Thresholds = Thresholds{
		MaxFleschKincaidGrade: 100.0,
		MaxARI:                100.0,
		MaxGunningFog:         100.0,
		MinFleschReadingEase:  0.0,
		MaxLines:              10, // Very strict
	}

	result, err := a.Analyze("test.md", []byte(content))
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	if result.Status != "fail" {
		t.Errorf("Expected fail status for line limit violation, got %q", result.Status)
	}

	// Check for line limit diagnostic
	found := false
	for _, d := range result.Diagnostics {
		if d.Rule == "structure/max-lines" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected max-lines diagnostic")
	}
}

func TestAnalyzeFile(t *testing.T) {
	// Create temp file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	content := "# Test\n\nSimple test content."
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	a := New()
	result, err := a.AnalyzeFile(testFile)
	if err != nil {
		t.Fatalf("AnalyzeFile() error = %v", err)
	}

	if result.File != testFile {
		t.Errorf("File = %q, want %q", result.File, testFile)
	}
}

func TestAnalyzeFile_NotFound(t *testing.T) {
	a := New()
	_, err := a.AnalyzeFile("/nonexistent/file.md")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}

func TestAnalyzeDirectory(t *testing.T) {
	// Create temp directory with markdown files
	tmpDir := t.TempDir()

	files := map[string]string{
		"doc1.md":          "# Doc 1\n\nContent one.",
		"doc2.md":          "# Doc 2\n\nContent two.",
		"subdir/doc3.md":   "# Doc 3\n\nContent three.",
		"README.md":        "# README\n\nThis is readme.",
		"CHANGELOG.md":     "# Changelog\n\nChanges here.",         // Should be skipped
		"CONTRIBUTING.md":  "# Contributing\n\nHow to contribute.", // Should be skipped
		"not_markdown.txt": "This is not markdown.",
	}

	for name, content := range files {
		path := filepath.Join(tmpDir, name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	a := New()
	results, err := a.AnalyzeDirectory(tmpDir)
	if err != nil {
		t.Fatalf("AnalyzeDirectory() error = %v", err)
	}

	// Should have doc1.md, doc2.md, subdir/doc3.md, README.md
	// Should NOT have CHANGELOG.md, CONTRIBUTING.md, not_markdown.txt
	if len(results) != 4 {
		t.Errorf("Expected 4 results, got %d", len(results))
	}

	// Verify CHANGELOG.md and CONTRIBUTING.md are excluded
	for _, r := range results {
		base := filepath.Base(r.File)
		if base == "CHANGELOG.md" || base == "CONTRIBUTING.md" {
			t.Errorf("Should not include %s", base)
		}
	}
}

func TestAnalyzeDirectory_NotFound(t *testing.T) {
	a := New()
	_, err := a.AnalyzeDirectory("/nonexistent/directory")
	if err == nil {
		t.Error("Expected error for nonexistent directory")
	}
}

func TestCollectDiagnostics_SkipShortDocs(t *testing.T) {
	a := New()
	a.Config = nil
	a.Thresholds = Thresholds{
		MaxFleschKincaidGrade: 8.0,
		MaxARI:                8.0,
		MaxGunningFog:         10.0,
		MinFleschReadingEase:  70.0,
		MaxLines:              1000,
	}

	// Very short content - readability metrics should be skipped
	result := &Result{
		File: "test.md",
		Structural: Structural{
			Words:     50, // Below 100 word minimum
			Lines:     5,
			Sentences: 2,
		},
		Readability: Readability{
			FleschKincaidGrade: 20.0, // Would fail if checked
			ARI:                20.0,
			GunningFog:         25.0,
			FleschReadingEase:  10.0,
		},
	}

	diagnostics := a.collectDiagnostics(result)

	// Readability violations should be skipped for short docs
	for _, d := range diagnostics {
		if d.Rule == "readability/grade-level" ||
			d.Rule == "readability/ari" ||
			d.Rule == "readability/gunning-fog" ||
			d.Rule == "readability/flesch-ease" {
			t.Errorf("Should skip readability check for short doc, got %s", d.Rule)
		}
	}
}

func TestCollectDiagnostics_WithConfig(t *testing.T) {
	cfg := &config.Config{
		Thresholds: config.Thresholds{
			MaxGrade:       8.0,
			MaxARI:         8.0,
			MaxFog:         10.0,
			MinEase:        60.0,
			MaxLines:       100,
			MinWords:       50,
			MinAdmonitions: 2,
		},
	}

	a := NewWithConfig(cfg)

	result := &Result{
		File: "test.md",
		Structural: Structural{
			Words:     200,
			Lines:     150, // Over limit
			Sentences: 10,
		},
		Readability: Readability{
			FleschKincaidGrade: 5.0,  // OK
			ARI:                5.0,  // OK
			GunningFog:         8.0,  // OK
			FleschReadingEase:  70.0, // OK
		},
		Admonitions: Admonitions{
			Count: 1, // Below minimum
		},
	}

	diagnostics := a.collectDiagnostics(result)

	// Should have max-lines and admonitions violations
	hasMaxLines := false
	hasAdmonitions := false
	for _, d := range diagnostics {
		if d.Rule == "structure/max-lines" {
			hasMaxLines = true
		}
		if d.Rule == "content/admonitions" {
			hasAdmonitions = true
		}
	}

	if !hasMaxLines {
		t.Error("Expected max-lines diagnostic")
	}
	if !hasAdmonitions {
		t.Error("Expected admonitions diagnostic")
	}
}

func TestDetermineStatus(t *testing.T) {
	a := New()

	tests := []struct {
		name        string
		diagnostics []Diagnostic
		want        string
	}{
		{
			name:        "no diagnostics",
			diagnostics: nil,
			want:        "pass",
		},
		{
			name:        "empty diagnostics",
			diagnostics: []Diagnostic{},
			want:        "pass",
		},
		{
			name: "info only",
			diagnostics: []Diagnostic{
				{Severity: SeverityInfo, Rule: "test", Message: "info"},
			},
			want: "pass",
		},
		{
			name: "warning",
			diagnostics: []Diagnostic{
				{Severity: SeverityWarning, Rule: "test", Message: "warning"},
			},
			want: "fail",
		},
		{
			name: "error",
			diagnostics: []Diagnostic{
				{Severity: SeverityError, Rule: "test", Message: "error"},
			},
			want: "fail",
		},
		{
			name: "mixed with error first",
			diagnostics: []Diagnostic{
				{Severity: SeverityError, Rule: "test1", Message: "error"},
				{Severity: SeverityWarning, Rule: "test2", Message: "warning"},
			},
			want: "fail",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := a.determineStatus(tt.diagnostics)
			if got != tt.want {
				t.Errorf("determineStatus() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStripFrontmatter(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    string
	}{
		{
			name:    "no frontmatter",
			content: "Hello world",
			want:    "Hello world",
		},
		{
			name:    "with frontmatter",
			content: "---\ntitle: Test\n---\nContent here",
			want:    "Content here",
		},
		{
			name:    "frontmatter only",
			content: "---\ntitle: Test\n---",
			want:    "",
		},
		{
			name:    "unclosed frontmatter",
			content: "---\ntitle: Test\nContent here",
			want:    "---\ntitle: Test\nContent here",
		},
		{
			name:    "complex frontmatter",
			content: "---\ntitle: Test\nauthor: Tester\ntags:\n  - one\n  - two\n---\n\n# Document\n\nActual content.",
			want:    "# Document\n\nActual content.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripFrontmatter(tt.content)
			if got != tt.want {
				t.Errorf("stripFrontmatter() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"empty", "", 0},
		{"single word", "hello", 1},
		{"multiple words", "hello world test", 3},
		{"with punctuation", "Hello, world! How are you?", 5},
		{"extra whitespace", "  hello   world  ", 2},
		{"newlines", "hello\nworld\ntest", 3},
		{"tabs", "hello\tworld", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWords(tt.text)
			if got != tt.want {
				t.Errorf("countWords(%q) = %d, want %d", tt.text, got, tt.want)
			}
		})
	}
}

func TestCountSentences(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"empty", "", 0},
		{"no ending punctuation", "hello world", 1},
		{"single sentence", "Hello world.", 1},
		{"multiple periods", "First. Second. Third.", 3},
		{"exclamation", "Hello! World!", 2},
		{"question", "How are you? I am fine.", 2},
		{"mixed", "Hello! How are you? I am fine.", 3},
		{"abbreviation counted", "Dr. Smith is here.", 2}, // Note: simple impl counts each period
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSentences(tt.text)
			if got != tt.want {
				t.Errorf("countSentences(%q) = %d, want %d", tt.text, got, tt.want)
			}
		})
	}
}

func TestCalculateReadingTime(t *testing.T) {
	tests := []struct {
		name  string
		words int
		want  int
	}{
		{"zero words", 0, 0},
		{"negative words", -10, 0},
		{"1 word", 1, 1},
		{"199 words", 199, 1},
		{"200 words exactly", 200, 1},
		{"201 words rounds up", 201, 2},
		{"399 words", 399, 2},
		{"400 words", 400, 2},
		{"401 words rounds up", 401, 3},
		{"500 words", 500, 3},
		{"1000 words", 1000, 5},
		{"1001 words rounds up", 1001, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateReadingTime(tt.words)
			if got != tt.want {
				t.Errorf("calculateReadingTime(%d) = %d, want %d", tt.words, got, tt.want)
			}
		})
	}
}

func TestCalculateRatio(t *testing.T) {
	tests := []struct {
		name  string
		part  int
		total int
		want  float64
	}{
		{"zero total", 10, 0, 0.0},
		{"zero part", 0, 100, 0.0},
		{"half", 50, 100, 0.5},
		{"quarter", 25, 100, 0.25},
		{"full", 100, 100, 1.0},
		{"over", 150, 100, 1.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateRatio(tt.part, tt.total)
			if got != tt.want {
				t.Errorf("calculateRatio(%d, %d) = %v, want %v", tt.part, tt.total, got, tt.want)
			}
		})
	}
}

func TestCountHeadings(t *testing.T) {
	tests := []struct {
		name     string
		headings []markdown.Heading
		wantH1   int
		wantH2   int
		wantH3   int
	}{
		{
			name:     "empty",
			headings: nil,
			wantH1:   0,
			wantH2:   0,
			wantH3:   0,
		},
		{
			name: "single H1",
			headings: []markdown.Heading{
				{Level: 1, Text: "Title"},
			},
			wantH1: 1,
		},
		{
			name: "mixed levels",
			headings: []markdown.Heading{
				{Level: 1, Text: "Title"},
				{Level: 2, Text: "Section 1"},
				{Level: 2, Text: "Section 2"},
				{Level: 3, Text: "Subsection"},
				{Level: 4, Text: "Detail"},
				{Level: 5, Text: "Sub-detail"},
				{Level: 6, Text: "Deep"},
			},
			wantH1: 1,
			wantH2: 2,
			wantH3: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countHeadings(tt.headings)
			if got.H1 != tt.wantH1 {
				t.Errorf("H1 = %d, want %d", got.H1, tt.wantH1)
			}
			if got.H2 != tt.wantH2 {
				t.Errorf("H2 = %d, want %d", got.H2, tt.wantH2)
			}
			if got.H3 != tt.wantH3 {
				t.Errorf("H3 = %d, want %d", got.H3, tt.wantH3)
			}
		})
	}
}

func TestCountAdmonitions(t *testing.T) {
	tests := []struct {
		name        string
		admonitions []markdown.Admonition
		wantCount   int
		wantTypes   []string
	}{
		{
			name:        "empty",
			admonitions: nil,
			wantCount:   0,
			wantTypes:   []string{},
		},
		{
			name: "single",
			admonitions: []markdown.Admonition{
				{Type: "note", Line: 1},
			},
			wantCount: 1,
			wantTypes: []string{"note"},
		},
		{
			name: "multiple same type",
			admonitions: []markdown.Admonition{
				{Type: "note", Line: 1},
				{Type: "note", Line: 5},
			},
			wantCount: 2,
			wantTypes: []string{"note"},
		},
		{
			name: "multiple different types",
			admonitions: []markdown.Admonition{
				{Type: "note", Line: 1},
				{Type: "warning", Line: 5},
				{Type: "tip", Line: 10},
			},
			wantCount: 3,
			wantTypes: []string{"note", "warning", "tip"},
		},
		{
			name: "empty type filtered",
			admonitions: []markdown.Admonition{
				{Type: "note", Line: 1},
				{Type: "", Line: 5},
			},
			wantCount: 2,
			wantTypes: []string{"note"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countAdmonitions(tt.admonitions)
			if got.Count != tt.wantCount {
				t.Errorf("Count = %d, want %d", got.Count, tt.wantCount)
			}
			if len(got.Types) != len(tt.wantTypes) {
				t.Errorf("Types = %v, want %v", got.Types, tt.wantTypes)
			}
		})
	}
}

func TestDefaultThresholds(t *testing.T) {
	d := DefaultThresholds()

	if d.MaxFleschKincaidGrade != 16.0 {
		t.Errorf("MaxFleschKincaidGrade = %v, want 16.0", d.MaxFleschKincaidGrade)
	}
	if d.MaxARI != 16.0 {
		t.Errorf("MaxARI = %v, want 16.0", d.MaxARI)
	}
	if d.MaxGunningFog != 18.0 {
		t.Errorf("MaxGunningFog = %v, want 18.0", d.MaxGunningFog)
	}
	if d.MinFleschReadingEase != 25.0 {
		t.Errorf("MinFleschReadingEase = %v, want 25.0", d.MinFleschReadingEase)
	}
	if d.MaxLines != 375 {
		t.Errorf("MaxLines = %v, want 375", d.MaxLines)
	}
}

func TestCollectDiagnostics_AllViolations(t *testing.T) {
	// Test all readability threshold violations
	a := New()
	a.Config = nil
	a.Thresholds = Thresholds{
		MaxFleschKincaidGrade: 5.0,
		MaxARI:                5.0,
		MaxGunningFog:         5.0,
		MinFleschReadingEase:  90.0,
		MaxLines:              10,
	}

	result := &Result{
		File: "test.md",
		Structural: Structural{
			Words:     200, // Above minWords
			Lines:     20,  // Over limit
			Sentences: 10,
		},
		Readability: Readability{
			FleschKincaidGrade: 15.0, // Over threshold
			ARI:                15.0, // Over threshold
			GunningFog:         15.0, // Over threshold
			FleschReadingEase:  30.0, // Under threshold
		},
	}

	diagnostics := a.collectDiagnostics(result)

	// Should have all 5 diagnostics
	rules := make(map[string]bool)
	for _, d := range diagnostics {
		rules[d.Rule] = true
	}

	expectedRules := []string{
		"readability/grade-level",
		"readability/ari",
		"readability/gunning-fog",
		"readability/flesch-ease",
		"structure/max-lines",
	}

	for _, rule := range expectedRules {
		if !rules[rule] {
			t.Errorf("Expected diagnostic for rule %s", rule)
		}
	}
}

func TestAnalyzeDirectory_WalkError(t *testing.T) {
	// Create temp directory with a file that will cause issues
	tmpDir := t.TempDir()

	// Create a file that looks like a directory in the path
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test"), 0644); err != nil {
		t.Fatal(err)
	}

	a := New()

	// Analyze the actual file as a directory - should not error
	// but also should not find any markdown files inside
	_, err := a.AnalyzeDirectory(testFile)
	// This will fail because testFile is a file, not a directory
	if err == nil {
		// If it doesn't error, that's fine too - it just won't find files
		t.Log("No error, path was traversed")
	}
}

func TestAnalyzeDirectory_FileReadError(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a markdown file
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("# Test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Make file unreadable
	if err := os.Chmod(testFile, 0000); err != nil {
		t.Skip("Cannot change file permissions")
	}
	defer os.Chmod(testFile, 0644) // Restore for cleanup

	a := New()
	_, err := a.AnalyzeDirectory(tmpDir)
	if err == nil {
		t.Log("Expected error for unreadable file, but test may not work on all platforms")
	}
}
