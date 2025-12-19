package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:generate go run ../../internal/genschema/main.go

// Config represents the content analyzer configuration.
type Config struct {
	Thresholds Thresholds     `yaml:"thresholds" json:"thresholds" jsonschema:"description=Base readability thresholds applied to all files"`
	Overrides  []PathOverride `yaml:"overrides,omitempty" json:"overrides,omitempty" jsonschema:"description=Path-specific threshold overrides (first match wins)"`
}

// Thresholds defines limits for pass/fail checks.
type Thresholds struct {
	MaxGrade       float64 `yaml:"max_grade" json:"max_grade" jsonschema:"minimum=0,maximum=100,default=16,examples=12;14;16,description=Maximum Flesch-Kincaid grade level (12 = high school senior\\, 16 = college senior)"`
	MaxARI         float64 `yaml:"max_ari" json:"max_ari" jsonschema:"minimum=0,maximum=100,default=16,examples=12;14;16,description=Maximum Automated Readability Index (similar to grade level)"`
	MaxFog         float64 `yaml:"max_fog" json:"max_fog" jsonschema:"minimum=0,maximum=100,default=18,examples=14;16;18,description=Maximum Gunning Fog index (years of formal education needed)"`
	MinEase        float64 `yaml:"min_ease" json:"min_ease" jsonschema:"minimum=-100,maximum=100,default=25,examples=30;40;50;-100,description=Minimum Flesch Reading Ease (0-100 scale\\, higher = easier). Use negative value to disable."`
	MaxLines       int     `yaml:"max_lines" json:"max_lines" jsonschema:"minimum=1,maximum=10000,default=375,examples=250;375;500,description=Maximum lines of prose per file"`
	MinWords       int     `yaml:"min_words" json:"min_words" jsonschema:"minimum=0,maximum=10000,default=100,examples=50;100;150,description=Minimum words before applying readability formulas (sparse docs are unreliable)"`
	MinAdmonitions int     `yaml:"min_admonitions" json:"min_admonitions" jsonschema:"minimum=-1,maximum=100,default=1,examples=0;1;2;-1,description=Minimum MkDocs-style admonitions required (!!! note\\, !!! warning). Use -1 to disable."`
	MaxDashDensity float64 `yaml:"max_dash_density" json:"max_dash_density" jsonschema:"minimum=-1,maximum=500,default=0,examples=0;2;5;-1,description=Maximum mid-sentence dash pairs per 100 sentences (detects AI-generated slop). Use -1 to disable. 0 = no dashes allowed."`
}

// PathOverride allows different thresholds for specific paths.
type PathOverride struct {
	Path       string     `yaml:"path" json:"path" jsonschema:"minLength=1,examples=docs/developer-guide/;docs/user-guide/;api/;README.md,description=Path prefix to match (e.g.\\, 'docs/developer-guide/' or 'api/')"`
	Thresholds Thresholds `yaml:"thresholds" json:"thresholds" jsonschema:"description=Threshold overrides for this path (inherits unspecified values from base)"`
}

// DefaultConfig returns sensible defaults for technical documentation.
func DefaultConfig() *Config {
	return &Config{
		Thresholds: Thresholds{
			MaxGrade:       16.0, // College senior
			MaxARI:         16.0,
			MaxFog:         18.0,
			MinEase:        25.0,
			MaxLines:       375,
			MinWords:       100, // Skip readability for very short/code-heavy docs
			MinAdmonitions: 1,   // Require at least one MkDocs-style admonition
			MaxDashDensity: 0,   // No mid-sentence dashes allowed (prevents AI slop)
		},
	}
}

// Load reads configuration from a YAML file and validates it against the embedded schema.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Parse YAML to generic interface{} for schema validation
	var yamlData interface{}
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return nil, fmt.Errorf("invalid YAML syntax: %w", err)
	}

	// Validate against embedded JSON Schema
	if err := ValidateAgainstSchema(yamlData); err != nil {
		return nil, err
	}

	// Parse into typed config struct
	cfg := DefaultConfig()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// ValidateConfig validates a config file against the JSON schema.
// This is provided for the --validate-config flag but Load() also validates automatically.
func ValidateConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Parse YAML to generic interface{} for schema validation
	var yamlData interface{}
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return fmt.Errorf("invalid YAML syntax: %w", err)
	}

	// Validate against JSON Schema
	return ValidateAgainstSchema(yamlData)
}

// LoadOrDefault tries to load config from path, returns default if not found.
func LoadOrDefault(path string) *Config {
	cfg, err := Load(path)
	if err != nil {
		return DefaultConfig()
	}
	return cfg
}

// ConfigFileName is the default configuration file name.
const ConfigFileName = ".readability.yml"

// FindConfigFile searches for .readability.yml in the given directory
// and parent directories up to the git root.
func FindConfigFile(startDir string) string {
	dir := startDir
	for {
		configPath := filepath.Join(dir, ConfigFileName)
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}

		// Check for git root (stop searching)
		gitDir := filepath.Join(dir, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			// We're at git root, already checked above
			return ""
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root
			return ""
		}
		dir = parent
	}
}

// ThresholdsForPath returns the appropriate thresholds for a given file path.
func (c *Config) ThresholdsForPath(filePath string) Thresholds {
	// Normalize path separators
	normalizedPath := filepath.ToSlash(filePath)

	// Strip leading ../ sequences (for when running from subdirectory)
	for strings.HasPrefix(normalizedPath, "../") {
		normalizedPath = normalizedPath[3:]
	}
	// Strip leading ./
	normalizedPath = strings.TrimPrefix(normalizedPath, "./")

	// Check overrides in order (first match wins)
	for _, override := range c.Overrides {
		overridePath := filepath.ToSlash(override.Path)
		// Check if override path appears anywhere in the file path
		// This handles both relative paths (docs/guide.md) and
		// absolute paths (/home/runner/work/repo/docs/guide.md)
		if strings.HasPrefix(normalizedPath, overridePath) || strings.Contains(normalizedPath, "/"+overridePath) {
			// Merge with defaults - override only specified values
			return mergeThresholds(c.Thresholds, override.Thresholds)
		}
	}

	return c.Thresholds
}

// mergeThresholds returns base thresholds with non-zero override values applied.
// Zero values in the override are treated as "not specified" and inherit from base.
// To explicitly disable a check via override, use a negative value:
//   - MinEase: use any negative value (e.g., -100) to allow very low readability
//   - MinAdmonitions: use -1 to disable the admonition requirement
//   - MaxDashDensity: use -1 to disable dash density check
func mergeThresholds(base, override Thresholds) Thresholds {
	result := base
	if override.MaxGrade > 0 {
		result.MaxGrade = override.MaxGrade
	}
	if override.MaxARI > 0 {
		result.MaxARI = override.MaxARI
	}
	if override.MaxFog > 0 {
		result.MaxFog = override.MaxFog
	}
	if override.MinEase != 0 {
		result.MinEase = override.MinEase
	}
	if override.MaxLines > 0 {
		result.MaxLines = override.MaxLines
	}
	if override.MinWords > 0 {
		result.MinWords = override.MinWords
	}
	if override.MinAdmonitions != 0 {
		result.MinAdmonitions = override.MinAdmonitions
	}
	if override.MaxDashDensity >= 0 {
		result.MaxDashDensity = override.MaxDashDensity
	}
	return result
}
