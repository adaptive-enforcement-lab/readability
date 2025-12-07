package output

import (
	"testing"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
)

func TestReadingTime(t *testing.T) {
	tests := []struct {
		name  string
		words int
		want  string
	}{
		{"zero words", 0, "<1m"},
		{"negative words", -10, "<1m"},
		{"1 word", 1, "1m"},
		{"199 words", 199, "1m"},
		{"200 words exactly", 200, "1m"},
		{"201 words rounds up", 201, "2m"},
		{"399 words", 399, "2m"},
		{"400 words", 400, "2m"},
		{"401 words rounds up", 401, "3m"},
		{"500 words", 500, "3m"},
		{"1000 words", 1000, "5m"},
		{"1001 words rounds up", 1001, "6m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readingTime(tt.words)
			if got != tt.want {
				t.Errorf("readingTime(%d) = %q, want %q", tt.words, got, tt.want)
			}
		})
	}
}

func TestIdentifyIssues(t *testing.T) {
	tests := []struct {
		name        string
		diagnostics []analyzer.Diagnostic
		want        string
	}{
		{
			name:        "no diagnostics",
			diagnostics: nil,
			want:        "Threshold exceeded",
		},
		{
			name: "single issue - grade",
			diagnostics: []analyzer.Diagnostic{
				{Rule: "readability/grade-level"},
			},
			want: "Grade",
		},
		{
			name: "multiple readability issues",
			diagnostics: []analyzer.Diagnostic{
				{Rule: "readability/grade-level"},
				{Rule: "readability/ari"},
				{Rule: "readability/gunning-fog"},
				{Rule: "readability/flesch-ease"},
			},
			want: "Grade, ARI, Fog, Ease",
		},
		{
			name: "structure and content issues",
			diagnostics: []analyzer.Diagnostic{
				{Rule: "structure/max-lines"},
				{Rule: "content/admonitions"},
			},
			want: "Lines, Admonitions",
		},
		{
			name: "all issue types",
			diagnostics: []analyzer.Diagnostic{
				{Rule: "readability/grade-level"},
				{Rule: "readability/ari"},
				{Rule: "content/admonitions"},
			},
			want: "Grade, ARI, Admonitions",
		},
		{
			name: "unknown rule uses rule ID",
			diagnostics: []analyzer.Diagnostic{
				{Rule: "custom/unknown-rule"},
			},
			want: "custom/unknown-rule",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &analyzer.Result{Diagnostics: tt.diagnostics}
			got := identifyIssues(r)
			if got != tt.want {
				t.Errorf("identifyIssues() = %q, want %q", got, tt.want)
			}
		})
	}
}
