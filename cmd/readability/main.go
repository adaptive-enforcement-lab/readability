package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adaptive-enforcement-lab/readability/pkg/analyzer"
	"github.com/adaptive-enforcement-lab/readability/pkg/config"
	"github.com/adaptive-enforcement-lab/readability/pkg/output"
	"github.com/spf13/cobra"
)

// version is set via ldflags at build time
var version = "dev"

var (
	formatFlag         string
	verboseFlag        bool
	checkFlag          bool
	validateConfigFlag bool
	configFlag         string
	maxGradeFlag       float64
	maxARIFlag         float64
	maxLinesFlag       int
	minAdmonitionsFlag int
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

// newRootCmd creates and returns the root command with all flags configured.
func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "readability [path]",
		Short:   "Analyze markdown documentation for readability and structure",
		Version: version,
		Long: `A tool for analyzing documentation quality, readability, and structure.

Computes readability metrics (Flesch-Kincaid, ARI, Coleman-Liau, etc.),
structural analysis (headings, line counts), and content composition.

Configuration:
  Reads .readability.yml from the target directory or git root.
  CLI flags override config file values.

Examples:
  readability docs/quickstart.md
  readability docs/
  readability docs/ --format json
  readability docs/ --format markdown
  readability docs/ --check
  readability docs/ --config .readability.yml`,
		Args: cobra.ExactArgs(1),
		RunE: run,
	}

	rootCmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "Output format: table, json, markdown, summary, report, diagnostic")
	rootCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "Show all metrics")
	rootCmd.Flags().BoolVar(&checkFlag, "check", false, "Check against thresholds (exit 1 on failure)")
	rootCmd.Flags().BoolVar(&validateConfigFlag, "validate-config", false, "Validate configuration and exit (no analysis)")
	rootCmd.Flags().StringVarP(&configFlag, "config", "c", "", "Path to config file (default: auto-detect .readability.yml)")
	rootCmd.Flags().Float64Var(&maxGradeFlag, "max-grade", 0, "Maximum Flesch-Kincaid grade level (overrides config)")
	rootCmd.Flags().Float64Var(&maxARIFlag, "max-ari", 0, "Maximum ARI score (overrides config)")
	rootCmd.Flags().IntVar(&maxLinesFlag, "max-lines", 0, "Maximum lines per file (overrides config, 0 to disable)")
	rootCmd.Flags().IntVar(&minAdmonitionsFlag, "min-admonitions", -1, "Minimum MkDocs-style admonitions (overrides config, 0 to disable)")

	return rootCmd
}

func run(cmd *cobra.Command, args []string) error {
	// Determine path for config loading
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		// If no path provided, use current directory for config search
		path = "."
	}

	cfg, err := loadConfig(path)
	if err != nil {
		return err
	}

	// If --validate-config flag is set, just validate and exit
	if validateConfigFlag {
		fmt.Println("✓ Configuration is valid")
		return nil
	}

	// For normal operation, require a path argument
	if len(args) == 0 {
		return fmt.Errorf("requires a path argument")
	}

	applyFlagOverrides(cmd, cfg)

	results, err := analyzeTarget(cfg, path)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		fmt.Fprintln(os.Stderr, "No markdown files found")
		return nil
	}

	if err := outputResults(results); err != nil {
		return err
	}

	if checkFlag {
		return checkResults(results, cfg)
	}

	return nil
}

// loadConfig loads configuration from file or returns defaults.
func loadConfig(path string) (*config.Config, error) {
	if configFlag != "" {
		cfg, err := config.Load(configFlag)
		if err != nil {
			return nil, fmt.Errorf("cannot load config %s: %w", configFlag, err)
		}
		return cfg, nil
	}

	// Auto-detect config file
	startDir := path
	if info, err := os.Stat(path); err == nil && !info.IsDir() {
		startDir = filepath.Dir(path)
	}

	configPath := config.FindConfigFile(startDir)
	if configPath != "" {
		if cfg, err := config.Load(configPath); err == nil {
			return cfg, nil
		}
	}

	return config.DefaultConfig(), nil
}

// applyFlagOverrides applies CLI flag values to the config.
func applyFlagOverrides(cmd *cobra.Command, cfg *config.Config) {
	if maxGradeFlag > 0 {
		cfg.Thresholds.MaxGrade = maxGradeFlag
	}
	if maxARIFlag > 0 {
		cfg.Thresholds.MaxARI = maxARIFlag
	}
	if cmd.Flags().Changed("max-lines") {
		cfg.Thresholds.MaxLines = maxLinesFlag
	}
	if cmd.Flags().Changed("min-admonitions") {
		cfg.Thresholds.MinAdmonitions = minAdmonitionsFlag
	}
}

// analyzeTarget analyzes a file or directory and returns results.
func analyzeTarget(cfg *config.Config, path string) ([]*analyzer.Result, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("cannot access %s: %w", path, err)
	}

	a := analyzer.NewWithConfig(cfg)

	if info.IsDir() {
		results, err := a.AnalyzeDirectory(path)
		if err != nil {
			return nil, fmt.Errorf("error analyzing directory: %w", err)
		}
		return results, nil
	}

	result, err := a.AnalyzeFile(path)
	if err != nil {
		return nil, fmt.Errorf("error analyzing file: %w", err)
	}
	return []*analyzer.Result{result}, nil
}

// outputResults writes results in the specified format.
func outputResults(results []*analyzer.Result) error {
	switch formatFlag {
	case "json":
		return output.JSON(os.Stdout, results)
	case "markdown":
		output.Markdown(os.Stdout, results)
	case "summary":
		output.Summary(os.Stdout, results)
	case "report":
		output.Report(os.Stdout, results)
	case "diagnostic":
		output.Diagnostic(os.Stdout, results)
		output.DiagnosticSummary(os.Stdout, results)
	default:
		output.Table(os.Stdout, results, verboseFlag)
	}
	return nil
}

// checkResults validates results against thresholds and returns an error if any fail.
func checkResults(results []*analyzer.Result, cfg *config.Config) error {
	stats := countFailures(results, cfg)

	if stats.failed == 0 {
		return nil
	}

	printFailureGuidance(stats)
	return fmt.Errorf("%d file(s) failed readability checks", stats.failed)
}

// failureStats holds counts of different failure types.
type failureStats struct {
	failed             int
	tooLong            int
	lowReadability     int
	missingAdmonitions int
}

// countFailures counts failures by category.
func countFailures(results []*analyzer.Result, cfg *config.Config) failureStats {
	stats := failureStats{}
	minAdm := cfg.Thresholds.MinAdmonitions

	for _, r := range results {
		if r.Status != "fail" {
			continue
		}
		stats.failed++
		if r.Structural.Lines > 375 {
			stats.tooLong++
		}
		if r.Readability.FleschKincaidGrade > 14 || r.Readability.ARI > 14 || r.Readability.FleschReadingEase < 30 {
			stats.lowReadability++
		}
		if minAdm > 0 && r.Admonitions.Count < minAdm {
			stats.missingAdmonitions++
		}
	}
	return stats
}

// printFailureGuidance prints helpful guidance for each failure type.
func printFailureGuidance(stats failureStats) {
	if stats.tooLong > 0 {
		printLengthGuidance()
	}
	if stats.lowReadability > 0 {
		printReadabilityGuidance()
	}
	if stats.missingAdmonitions > 0 {
		printAdmonitionGuidance()
	}
}

func printLengthGuidance() {
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "IMPORTANT: Files exceeding line limits should be SPLIT into smaller documents.")
	fmt.Fprintln(os.Stderr, "Do NOT remove content to meet thresholds. Split logically by topic or section.")
	fmt.Fprintln(os.Stderr, "")
}

func printReadabilityGuidance() {
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "READABILITY: High grade level indicates complex sentence structure or dense vocabulary.")
	fmt.Fprintln(os.Stderr, "- Break long sentences into shorter ones (aim for 15-20 words per sentence)")
	fmt.Fprintln(os.Stderr, "- Replace jargon with plain language where possible")
	fmt.Fprintln(os.Stderr, "- Add brief introductory sentences before bullet lists, code blocks, or admonitions")
	fmt.Fprintln(os.Stderr, "- Use transitional phrases to connect dense technical sections")
	fmt.Fprintln(os.Stderr, "Do NOT remove technical content. Rewrite for clarity while preserving accuracy.")
	fmt.Fprintln(os.Stderr, "")
}

func printAdmonitionGuidance() {
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "ADMONITIONS: Files are missing MkDocs-style admonitions (note, warning, tip, etc.).")
	fmt.Fprintln(os.Stderr, "Admonitions improve documentation by highlighting important information:")
	fmt.Fprintln(os.Stderr, "- Use !!! note for supplementary information")
	fmt.Fprintln(os.Stderr, "- Use !!! warning for potential pitfalls or breaking changes")
	fmt.Fprintln(os.Stderr, "- Use !!! tip for best practices or shortcuts")
	fmt.Fprintln(os.Stderr, "- Use !!! example for code samples or use cases")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Example syntax:")
	fmt.Fprintln(os.Stderr, "  !!! note \"Optional Title\"")
	fmt.Fprintln(os.Stderr, "      Content indented by 4 spaces.")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Do NOT add empty or meaningless admonitions. Add value with relevant context.")
	fmt.Fprintln(os.Stderr, "")
}
