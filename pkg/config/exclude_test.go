package config

import (
	"testing"
)

func TestIsExcluded(t *testing.T) {
	cfg := &Config{
		Overrides: []PathOverride{
			{
				Path:    "docs/includes/",
				Exclude: true,
			},
			{
				Path:    "CHANGELOG.md",
				Exclude: true,
			},
			{
				Path: "docs/api/",
				Thresholds: Thresholds{
					MaxGrade: 20,
				},
			},
		},
	}

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "relative path matches exclude override",
			path: "docs/includes/abbreviations.md",
			want: true,
		},
		{
			name: "relative path with ./ prefix matches",
			path: "./docs/includes/abbreviations.md",
			want: true,
		},
		{
			name: "relative path with ../ prefix matches",
			path: "../docs/includes/abbreviations.md",
			want: true,
		},
		{
			name: "absolute path matches exclude override",
			path: "/home/runner/work/repo/docs/includes/abbreviations.md",
			want: true,
		},
		{
			name: "specific file matches exclude override",
			path: "CHANGELOG.md",
			want: true,
		},
		{
			name: "specific file with path prefix matches",
			path: "./CHANGELOG.md",
			want: true,
		},
		{
			name: "path outside exclude returns false",
			path: "docs/guide/quickstart.md",
			want: false,
		},
		{
			name: "path with thresholds but not exclude returns false",
			path: "docs/api/endpoints.md",
			want: false,
		},
		{
			name: "empty path returns false",
			path: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cfg.IsExcluded(tt.path); got != tt.want {
				t.Errorf("IsExcluded(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestIsExcluded_FirstMatchWins(t *testing.T) {
	cfg := &Config{
		Overrides: []PathOverride{
			{
				Path:    "docs/",
				Exclude: true,
			},
			{
				Path: "docs/api/",
				Thresholds: Thresholds{
					MaxGrade: 20,
				},
			},
		},
	}

	// First override matches and has exclude:true
	if got := cfg.IsExcluded("docs/api/endpoints.md"); !got {
		t.Error("IsExcluded(docs/api/endpoints.md) = false, want true (first match wins)")
	}
}

func TestValidatePathOverrides(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		wantErr bool
	}{
		{
			name: "exclude without thresholds is valid",
			cfg: &Config{
				Overrides: []PathOverride{
					{
						Path:    "docs/includes/",
						Exclude: true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "thresholds without exclude is valid",
			cfg: &Config{
				Overrides: []PathOverride{
					{
						Path: "docs/api/",
						Thresholds: Thresholds{
							MaxGrade: 20,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "exclude with zero thresholds is valid",
			cfg: &Config{
				Overrides: []PathOverride{
					{
						Path:       "docs/includes/",
						Exclude:    true,
						Thresholds: Thresholds{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "exclude with non-zero thresholds is invalid",
			cfg: &Config{
				Overrides: []PathOverride{
					{
						Path:    "docs/includes/",
						Exclude: true,
						Thresholds: Thresholds{
							MaxGrade: 20,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "multiple overrides with mixed exclude/thresholds is valid",
			cfg: &Config{
				Overrides: []PathOverride{
					{
						Path:    "docs/includes/",
						Exclude: true,
					},
					{
						Path: "docs/api/",
						Thresholds: Thresholds{
							MaxGrade: 20,
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePathOverrides(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePathOverrides() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
