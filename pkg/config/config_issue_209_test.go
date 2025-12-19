package config

import (
	"os"
	"path/filepath"
	"testing"
)

// TestIssue209_PathOverrideMatching reproduces issue #209
// Path overrides not matching in v1.14.1
func TestIssue209_PathOverrideMatching(t *testing.T) {
	// Create config matching the reported issue
	content := `# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade: 16
  max_ari: 16
  max_fog: 18
  min_ease: 25

overrides:
  - path: docs/patterns/
    thresholds:
      max_grade: 50
      max_ari: 60
      max_fog: 50
      min_ease: -100

  - path: docs/build/
    thresholds:
      max_grade: 50
      max_ari: 60
      max_fog: 50
      min_ease: -100
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	// Test files from the issue
	testCases := []struct {
		filePath         string
		expectedMaxGrade float64
		expectedMaxARI   float64
		expectedMinEase  float64
	}{
		{
			filePath:         "docs/patterns/architecture/separation-of-concerns/index.md",
			expectedMaxGrade: 50,
			expectedMaxARI:   60,
			expectedMinEase:  -100,
		},
		{
			filePath:         "docs/build/go-cli-architecture/packaging/github-actions.md",
			expectedMaxGrade: 50,
			expectedMaxARI:   60,
			expectedMinEase:  -100,
		},
		{
			filePath:         "docs/other/file.md", // No override
			expectedMaxGrade: 16,
			expectedMaxARI:   16,
			expectedMinEase:  25,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.filePath, func(t *testing.T) {
			thresholds := cfg.ThresholdsForPath(tc.filePath)

			if thresholds.MaxGrade != tc.expectedMaxGrade {
				t.Errorf("MaxGrade = %.1f, want %.1f", thresholds.MaxGrade, tc.expectedMaxGrade)
			}
			if thresholds.MaxARI != tc.expectedMaxARI {
				t.Errorf("MaxARI = %.1f, want %.1f", thresholds.MaxARI, tc.expectedMaxARI)
			}
			if thresholds.MinEase != tc.expectedMinEase {
				t.Errorf("MinEase = %.1f, want %.1f", thresholds.MinEase, tc.expectedMinEase)
			}
		})
	}
}
