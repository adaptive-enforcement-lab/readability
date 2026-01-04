package analyzer

import (
	"testing"

	"github.com/adaptive-enforcement-lab/readability/pkg/config"
)

func TestDetermineStatus_Excluded(t *testing.T) {
	cfg := &config.Config{
		Thresholds: config.Thresholds{
			MaxGrade: 16,
			MaxARI:   16,
		},
		Overrides: []config.PathOverride{
			{
				Path:    "docs/includes/",
				Exclude: true,
			},
		},
	}

	a := NewWithConfig(cfg)

	tests := []struct {
		name       string
		filePath   string
		wantStatus string
	}{
		{
			name:       "excluded file gets excluded status",
			filePath:   "docs/includes/abbreviations.md",
			wantStatus: "excluded",
		},
		{
			name:       "non-excluded file gets normal status",
			filePath:   "docs/guide/quickstart.md",
			wantStatus: "pass", // Assuming no diagnostics
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := &Result{
				File:        tt.filePath,
				Diagnostics: []Diagnostic{}, // No diagnostics
			}

			status := a.determineStatus(result)
			if status != tt.wantStatus {
				t.Errorf("determineStatus() = %q, want %q", status, tt.wantStatus)
			}
		})
	}
}

func TestDetermineStatus_ExcludedWithDiagnostics(t *testing.T) {
	cfg := &config.Config{
		Thresholds: config.Thresholds{
			MaxGrade: 16,
		},
		Overrides: []config.PathOverride{
			{
				Path:    "docs/includes/",
				Exclude: true,
			},
		},
	}

	a := NewWithConfig(cfg)

	// File is excluded even with error diagnostics
	result := &Result{
		File: "docs/includes/abbreviations.md",
		Diagnostics: []Diagnostic{
			{
				Severity: SeverityError,
				Rule:     "readability/grade-level",
				Message:  "Grade level too high",
			},
		},
	}

	status := a.determineStatus(result)
	if status != "excluded" {
		t.Errorf("determineStatus() = %q, want %q (excluded takes priority over errors)", status, "excluded")
	}
}

func TestAnalyze_ExcludedFileStillComputesMetrics(t *testing.T) {
	cfg := &config.Config{
		Thresholds: config.Thresholds{
			MaxGrade: 16,
		},
		Overrides: []config.PathOverride{
			{
				Path:    "test.md",
				Exclude: true,
			},
		},
	}

	a := NewWithConfig(cfg)

	content := []byte("This is a test document with some words.")

	result, err := a.Analyze("test.md", content)
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	// Verify status is excluded
	if result.Status != "excluded" {
		t.Errorf("Status = %q, want %q", result.Status, "excluded")
	}

	// Verify metrics are still computed
	if result.Structural.Words == 0 {
		t.Error("Structural.Words = 0, want > 0 (metrics should still be computed)")
	}

	if result.Readability.FleschKincaidGrade == 0 {
		t.Error("FleschKincaidGrade = 0, want > 0 (readability should still be computed)")
	}
}
