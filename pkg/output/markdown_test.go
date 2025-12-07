package output

import "testing"

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
