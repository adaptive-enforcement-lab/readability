package output

import (
	"encoding/json"
	"io"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
)

// JSON writes results as JSON array.
func JSON(w io.Writer, results []*analyzer.Result) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	// Always return an array for consistent parsing
	return encoder.Encode(results)
}
