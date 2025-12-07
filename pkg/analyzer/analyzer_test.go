package analyzer

import "testing"

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
