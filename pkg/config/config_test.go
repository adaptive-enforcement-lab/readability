package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestThresholdsForPath_RelativePaths(t *testing.T) {
	cfg := &Config{
		Thresholds: Thresholds{
			MaxGrade:       16,
			MinAdmonitions: 1,
		},
		Overrides: []PathOverride{
			{
				Path: "docs/developer-guide/",
				Thresholds: Thresholds{
					MaxGrade: 50,
					// Note: MinAdmonitions not set - inherits from base (1)
					// To disable, user must set min_admonitions: -1
				},
			},
			{
				Path: "docs/api/",
				Thresholds: Thresholds{
					MaxGrade:       30,
					MinAdmonitions: -1, // Explicitly disabled
				},
			},
		},
	}

	tests := []struct {
		name               string
		path               string
		wantMaxGrade       float64
		wantMinAdmonitions int
	}{
		{
			name:               "relative path matches override, inherits base MinAdmonitions",
			path:               "docs/developer-guide/test.md",
			wantMaxGrade:       50,
			wantMinAdmonitions: 1, // Inherited from base since not specified in override
		},
		{
			name:               "relative path with ./ prefix matches override",
			path:               "./docs/developer-guide/test.md",
			wantMaxGrade:       50,
			wantMinAdmonitions: 1, // Inherited from base
		},
		{
			name:               "relative path with ../ prefix matches override",
			path:               "../docs/developer-guide/test.md",
			wantMaxGrade:       50,
			wantMinAdmonitions: 1, // Inherited from base
		},
		{
			name:               "path outside override uses defaults",
			path:               "docs/user-guide/test.md",
			wantMaxGrade:       16,
			wantMinAdmonitions: 1,
		},
		{
			name:               "path with explicit MinAdmonitions: -1 override",
			path:               "docs/api/endpoints.md",
			wantMaxGrade:       30,
			wantMinAdmonitions: -1, // Explicitly set to -1 in override
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thresholds := cfg.ThresholdsForPath(tt.path)
			if thresholds.MaxGrade != tt.wantMaxGrade {
				t.Errorf("MaxGrade = %v, want %v", thresholds.MaxGrade, tt.wantMaxGrade)
			}
			if thresholds.MinAdmonitions != tt.wantMinAdmonitions {
				t.Errorf("MinAdmonitions = %v, want %v", thresholds.MinAdmonitions, tt.wantMinAdmonitions)
			}
		})
	}
}

func TestThresholdsForPath_AbsolutePaths(t *testing.T) {
	cfg := &Config{
		Thresholds: Thresholds{
			MaxGrade:       16,
			MinAdmonitions: 1,
		},
		Overrides: []PathOverride{
			{
				Path: "docs/developer-guide/",
				Thresholds: Thresholds{
					MaxGrade:       50,
					MinAdmonitions: -1, // Explicitly disabled
				},
			},
		},
	}

	tests := []struct {
		name               string
		path               string
		wantMaxGrade       float64
		wantMinAdmonitions int
	}{
		{
			name:               "absolute path matches override (Linux style)",
			path:               "/home/runner/work/repo/docs/developer-guide/test.md",
			wantMaxGrade:       50,
			wantMinAdmonitions: -1,
		},
		{
			name:               "absolute path matches override (short)",
			path:               "/tmp/docs/developer-guide/test.md",
			wantMaxGrade:       50,
			wantMinAdmonitions: -1,
		},
		{
			name:               "absolute path outside override uses defaults",
			path:               "/home/runner/work/repo/docs/user-guide/test.md",
			wantMaxGrade:       16,
			wantMinAdmonitions: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thresholds := cfg.ThresholdsForPath(tt.path)
			if thresholds.MaxGrade != tt.wantMaxGrade {
				t.Errorf("MaxGrade = %v, want %v", thresholds.MaxGrade, tt.wantMaxGrade)
			}
			if thresholds.MinAdmonitions != tt.wantMinAdmonitions {
				t.Errorf("MinAdmonitions = %v, want %v", thresholds.MinAdmonitions, tt.wantMinAdmonitions)
			}
		})
	}
}

func TestMergeThresholds(t *testing.T) {
	base := Thresholds{
		MaxGrade:       16,
		MaxARI:         16,
		MaxFog:         18,
		MinEase:        25,
		MaxLines:       375,
		MinWords:       100,
		MinAdmonitions: 1,
	}

	tests := []struct {
		name     string
		override Thresholds
		want     Thresholds
	}{
		{
			name:     "zero override keeps base values",
			override: Thresholds{},
			want:     base,
		},
		{
			name: "partial override merges correctly",
			override: Thresholds{
				MaxGrade: 50,
				MinEase:  -100,
			},
			want: Thresholds{
				MaxGrade:       50,
				MaxARI:         16,
				MaxFog:         18,
				MinEase:        -100,
				MaxLines:       375,
				MinWords:       100,
				MinAdmonitions: 1,
			},
		},
		{
			name: "explicit zero MinAdmonitions overrides",
			override: Thresholds{
				MinAdmonitions: 0, // This won't override due to != 0 check
			},
			want: Thresholds{
				MaxGrade:       16,
				MaxARI:         16,
				MaxFog:         18,
				MinEase:        25,
				MaxLines:       375,
				MinWords:       100,
				MinAdmonitions: 1, // Stays at base value
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeThresholds(base, tt.override)
			if got != tt.want {
				t.Errorf("mergeThresholds() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	// Create a temp config file
	content := `thresholds:
  max_grade: 14
  max_ari: 14
  min_words: 200

overrides:
  - path: docs/api/
    thresholds:
      max_grade: 20
`
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".readability.yml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Check base thresholds
	if cfg.Thresholds.MaxGrade != 14 {
		t.Errorf("MaxGrade = %v, want 14", cfg.Thresholds.MaxGrade)
	}
	if cfg.Thresholds.MinWords != 200 {
		t.Errorf("MinWords = %v, want 200", cfg.Thresholds.MinWords)
	}
	// Check defaults are preserved for unspecified values
	if cfg.Thresholds.MinAdmonitions != 1 {
		t.Errorf("MinAdmonitions = %v, want 1 (default)", cfg.Thresholds.MinAdmonitions)
	}

	// Check override
	if len(cfg.Overrides) != 1 {
		t.Fatalf("Expected 1 override, got %d", len(cfg.Overrides))
	}
	if cfg.Overrides[0].Path != "docs/api/" {
		t.Errorf("Override path = %v, want docs/api/", cfg.Overrides[0].Path)
	}
}

func TestFindConfigFile(t *testing.T) {
	// Create a temp directory structure with a config file
	tmpDir := t.TempDir()

	// Create nested directories
	subDir := filepath.Join(tmpDir, "sub", "nested")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirs: %v", err)
	}

	// Create .git directory at root (to stop search)
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("Failed to create .git: %v", err)
	}

	// Create config file at root
	configPath := filepath.Join(tmpDir, ConfigFileName)
	if err := os.WriteFile(configPath, []byte("thresholds:\n  max_grade: 16\n"), 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	// Test finding config from nested directory
	found := FindConfigFile(subDir)
	if found != configPath {
		t.Errorf("FindConfigFile() = %v, want %v", found, configPath)
	}

	// Test from root
	found = FindConfigFile(tmpDir)
	if found != configPath {
		t.Errorf("FindConfigFile() from root = %v, want %v", found, configPath)
	}
}

func TestFindConfigFile_NoConfig(t *testing.T) {
	// Create a temp directory with .git but no config
	tmpDir := t.TempDir()
	gitDir := filepath.Join(tmpDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("Failed to create .git: %v", err)
	}

	found := FindConfigFile(tmpDir)
	if found != "" {
		t.Errorf("FindConfigFile() = %v, want empty string", found)
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg == nil {
		t.Fatal("DefaultConfig() returned nil")
	}

	// Check default values
	if cfg.Thresholds.MaxGrade != 16.0 {
		t.Errorf("MaxGrade = %v, want 16.0", cfg.Thresholds.MaxGrade)
	}
	if cfg.Thresholds.MaxARI != 16.0 {
		t.Errorf("MaxARI = %v, want 16.0", cfg.Thresholds.MaxARI)
	}
	if cfg.Thresholds.MaxFog != 18.0 {
		t.Errorf("MaxFog = %v, want 18.0", cfg.Thresholds.MaxFog)
	}
	if cfg.Thresholds.MinEase != 25.0 {
		t.Errorf("MinEase = %v, want 25.0", cfg.Thresholds.MinEase)
	}
	if cfg.Thresholds.MaxLines != 375 {
		t.Errorf("MaxLines = %v, want 375", cfg.Thresholds.MaxLines)
	}
	if cfg.Thresholds.MinWords != 100 {
		t.Errorf("MinWords = %v, want 100", cfg.Thresholds.MinWords)
	}
	if cfg.Thresholds.MinAdmonitions != 1 {
		t.Errorf("MinAdmonitions = %v, want 1", cfg.Thresholds.MinAdmonitions)
	}
}

func TestLoadOrDefault_Success(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yml")
	content := "thresholds:\n  max_grade: 10\n"
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg := LoadOrDefault(configPath)

	if cfg.Thresholds.MaxGrade != 10 {
		t.Errorf("MaxGrade = %v, want 10", cfg.Thresholds.MaxGrade)
	}
}

func TestLoadOrDefault_Fallback(t *testing.T) {
	cfg := LoadOrDefault("/nonexistent/config.yml")

	// Should return default config
	if cfg.Thresholds.MaxGrade != 16.0 {
		t.Errorf("Expected default MaxGrade 16.0, got %v", cfg.Thresholds.MaxGrade)
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid.yml")
	content := "thresholds:\n  max_grade: [invalid\n" // Invalid YAML
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(configPath)
	if err == nil {
		t.Error("Expected error for invalid YAML")
	}
}

func TestLoad_NotFound(t *testing.T) {
	_, err := Load("/nonexistent/config.yml")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}
