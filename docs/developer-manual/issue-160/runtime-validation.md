# Component 4: Runtime Validation

## Overview

Add JSON Schema validation to the Go CLI to catch configuration errors at runtime. This provides a safety net beyond IDE validation—ensuring invalid configs are detected even when edited without schema-aware tools.

## Technical Approach

### Integration Point

Validation should occur in the config loading pipeline:

```go
// pkg/config/config.go:52-64 (current code)
func Load(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    cfg := DefaultConfig()
    if err := yaml.Unmarshal(data, cfg); err != nil {
        return nil, err
    }

    // NEW: Add schema validation here
    if err := validateAgainstSchema(data, cfg); err != nil {
        return nil, fmt.Errorf("schema validation failed: %w", err)
    }

    return cfg, nil
}
```

### Validation Flow

1. **Load YAML** → Unmarshal to Go struct (existing behavior)
2. **Validate Schema** → Check against JSON Schema (new step)
3. **Return Config** or **Return Detailed Error**

This preserves backward compatibility—existing YAML parsing still works, schema validation adds extra safety.

## Go JSON Schema Libraries

### Library Comparison

| Library | Stars | Last Updated | Draft Support | Pros | Cons | Verdict |
|---------|-------|--------------|---------------|------|------|---------|
| [gojsonschema](https://github.com/xeipuuv/gojsonschema) | 5.3k | 2023-06 | 2020-12 ✅ | Mature, widely used | Maintenance concerns | ⚠️ |
| [jsonschema](https://github.com/santhosh-tekuri/jsonschema) | 1.1k | 2024-11 | 2020-12 ✅ | Actively maintained, fast | Smaller community | ✅ **Recommended** |
| [go-jsschema](https://github.com/lestrrat-go/go-jsschema) | 26 | 2021-05 | Draft-07 ❌ | - | Outdated, inactive | ❌ |
| [fastjsonschema](https://github.com/romapres/fastjsonschema) | 12 | 2024-03 | Draft-07 ❌ | Fast code generation | No 2020-12 support | ❌ |

### Recommended Library: santhosh-tekuri/jsonschema

**GitHub**: https://github.com/santhosh-tekuri/jsonschema
**Stars**: 1.1k
**License**: MIT
**Latest Release**: v6.2.0 (November 2024)

**Why This Library**:
- ✅ **Actively maintained** - Regular updates, responsive maintainer
- ✅ **Draft 2020-12 support** - Latest JSON Schema spec
- ✅ **Excellent error messages** - Detailed validation failures with JSON pointers
- ✅ **Embedded schema support** - Can bundle schema in binary
- ✅ **YAML support** - Direct YAML validation (no intermediate JSON conversion)
- ✅ **Custom validators** - Extensible for future needs
- ✅ **Zero dependencies** - No transitive dependency bloat

**Installation**:
```bash
go get github.com/santhosh-tekuri/jsonschema/v6
```

## Implementation Design

### Approach 1: Validate YAML Directly (Recommended)

```go
package config

import (
    _ "embed"
    "fmt"

    "github.com/santhosh-tekuri/jsonschema/v6"
    "gopkg.in/yaml.v3"
)

//go:embed schemas/readability-config.schema.json
var schemaJSON []byte

var compiledSchema *jsonschema.Schema

func init() {
    // Compile schema once at startup
    compiler := jsonschema.NewCompiler()
    if err := compiler.AddResource("readability-config.schema.json", bytes.NewReader(schemaJSON)); err != nil {
        panic(fmt.Sprintf("failed to load schema: %v", err))
    }

    var err error
    compiledSchema, err = compiler.Compile("readability-config.schema.json")
    if err != nil {
        panic(fmt.Sprintf("failed to compile schema: %v", err))
    }
}

func Load(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    // Parse YAML to generic interface{} for schema validation
    var yamlData interface{}
    if err := yaml.Unmarshal(data, &yamlData); err != nil {
        return nil, fmt.Errorf("invalid YAML: %w", err)
    }

    // Validate against schema
    if err := compiledSchema.Validate(yamlData); err != nil {
        return nil, formatSchemaError(err)
    }

    // Parse into typed Config struct
    cfg := DefaultConfig()
    if err := yaml.Unmarshal(data, cfg); err != nil {
        return nil, err
    }

    return cfg, nil
}

func formatSchemaError(err error) error {
    if validationErr, ok := err.(*jsonschema.ValidationError); ok {
        // Extract detailed error information
        var messages []string
        for _, err := range validationErr.DetailedErrors() {
            // err.InstanceLocation gives JSON pointer path (e.g., /thresholds/max_grade)
            // err.Message gives human-readable error
            messages = append(messages, fmt.Sprintf("%s: %s", err.InstanceLocation, err.Message))
        }
        return fmt.Errorf("configuration validation failed:\n  %s", strings.Join(messages, "\n  "))
    }
    return err
}
```

**Pros**:
- ✅ Direct YAML validation (no conversion to JSON)
- ✅ Schema embedded in binary (no runtime file loading)
- ✅ Compiled once (fast validation on subsequent calls)
- ✅ Detailed error messages with field paths

**Cons**:
- ⚠️ Adds ~500KB to binary size (schema + library)

**Verdict**: **Recommended** - Best balance of performance and usability

### Approach 2: Convert YAML to JSON First

```go
func Load(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    // Parse YAML
    var yamlData interface{}
    if err := yaml.Unmarshal(data, &yamlData); err != nil {
        return nil, err
    }

    // Convert YAML to JSON
    jsonData, err := json.Marshal(yamlData)
    if err != nil {
        return nil, err
    }

    // Validate JSON against schema
    if err := compiledSchema.Validate(bytes.NewReader(jsonData)); err != nil {
        return nil, formatSchemaError(err)
    }

    // Parse into Config
    cfg := DefaultConfig()
    if err := yaml.Unmarshal(data, cfg); err != nil {
        return nil, err
    }

    return cfg, nil
}
```

**Pros**:
- ✅ More libraries support JSON validation

**Cons**:
- ❌ Extra conversion step (YAML → interface{} → JSON)
- ❌ Potential data loss (YAML features not in JSON)
- ❌ Slower

**Verdict**: **Not Recommended** - Unnecessary complexity

### Approach 3: Optional Validation (Flag-Gated)

```go
var enableSchemaValidation = true // or use flag

func Load(path string) (*Config, error) {
    // ... load YAML ...

    if enableSchemaValidation {
        if err := compiledSchema.Validate(yamlData); err != nil {
            return nil, formatSchemaError(err)
        }
    }

    // ... parse to Config ...
}
```

**Pros**:
- ✅ Can disable for performance
- ✅ Opt-in for users who want it

**Cons**:
- ❌ Users might disable and miss errors
- ❌ Adds configuration complexity

**Verdict**: **Not Recommended** - Validation should always run (performance impact is negligible)

## Error Message Design

### Problem

JSON Schema validation errors are often cryptic:
```
validation failed: /thresholds/max_grade: expected number, got string
```

### Enhanced Error Formatting

```go
func formatSchemaError(err error) error {
    validationErr, ok := err.(*jsonschema.ValidationError)
    if !ok {
        return err
    }

    var buf strings.Builder
    buf.WriteString("Configuration validation failed:\n\n")

    for _, e := range validationErr.DetailedErrors() {
        // Convert JSON pointer to YAML path
        yamlPath := jsonPointerToYAMLPath(e.InstanceLocation)

        buf.WriteString(fmt.Sprintf("  • %s\n", yamlPath))
        buf.WriteString(fmt.Sprintf("    %s\n", e.Message))

        // Add suggestion if possible
        if suggestion := getSuggestion(e); suggestion != "" {
            buf.WriteString(fmt.Sprintf("    Suggestion: %s\n", suggestion))
        }
        buf.WriteString("\n")
    }

    buf.WriteString("See https://github.com/markcheret/readability/blob/main/docs/cli/config-file.md for configuration reference.\n")

    return fmt.Errorf(buf.String())
}

func jsonPointerToYAMLPath(pointer string) string {
    // Convert /thresholds/max_grade → thresholds.max_grade
    return strings.ReplaceAll(strings.TrimPrefix(pointer, "/"), "/", ".")
}

func getSuggestion(err *jsonschema.DetailedError) string {
    switch {
    case strings.Contains(err.Message, "expected number, got string"):
        return "Remove quotes around numeric values"
    case strings.Contains(err.Message, "additional properties"):
        return "Check for typos in field names"
    case strings.Contains(err.Message, "required"):
        return "This field is mandatory"
    default:
        return ""
    }
}
```

**Example Output**:
```
Configuration validation failed:

  • thresholds.max_grade
    expected number, got string
    Suggestion: Remove quotes around numeric values

  • thresholds.unknown_field
    property "unknown_field" not allowed
    Suggestion: Check for typos in field names

See https://github.com/markcheret/readability/blob/main/docs/cli/config-file.md for configuration reference.
```

## CLI Flag: --validate-config

Add a dedicated flag to test configuration without running analysis:

```go
// cmd/readability/root.go
var validateConfigFlag bool

func init() {
    rootCmd.Flags().BoolVar(&validateConfigFlag, "validate-config", false, "validate configuration and exit")
}

func run(cmd *cobra.Command, args []string) error {
    // Load config
    cfg, err := config.Load(configPath)
    if err != nil {
        return err
    }

    if validateConfigFlag {
        fmt.Println("✓ Configuration is valid")
        return nil
    }

    // Continue with normal analysis...
}
```

**Usage**:
```bash
$ readability --validate-config
✓ Configuration is valid

$ readability --validate-config
Configuration validation failed:

  • thresholds.max_grade
    expected number, got string
    Suggestion: Remove quotes around numeric values
```

**Use Cases**:
- CI pipeline: validate config before running analysis
- Pre-commit hooks: catch errors early
- Config editing: quick feedback loop

## Embedding Schema in Binary

### Why Embed?

**Benefits**:
- ✅ No external file dependency (schema bundled with CLI)
- ✅ Works in any environment (no need to download schema)
- ✅ Version-locked (schema matches CLI version)
- ✅ Faster startup (no file I/O)

**Trade-offs**:
- ⚠️ Increases binary size (~50KB for schema, ~500KB for library)
- ⚠️ Schema updates require recompiling CLI

**Verdict**: **Worth it** - Reliability > binary size

### Implementation

```go
package config

import _ "embed"

//go:embed schemas/readability-config.schema.json
var embeddedSchema []byte
```

**Build Requirements**:
- Go 1.16+ (for `//go:embed`)
- Schema file must be in same module

**File Structure**:
```
readability/
├── schemas/
│   └── readability-config.schema.json
├── pkg/
│   └── config/
│       ├── config.go
│       └── validate.go (new file)
└── go.mod
```

## Viability Analysis

### ✅ High Viability

**Evidence**:
1. **Mature Libraries**: `santhosh-tekuri/jsonschema` is production-ready
2. **Low Complexity**: Integration requires ~100 lines of code
3. **Backward Compatible**: Doesn't break existing config loading
4. **Performance**: Schema compilation is one-time cost, validation is fast (<1ms)
5. **Proven Pattern**: Used by major Go projects (Kubernetes, Helm, etc.)

### Performance Impact

**Benchmark** (estimated):

| Operation | Without Validation | With Validation | Overhead |
|-----------|-------------------|-----------------|----------|
| Load `.readability.yml` | 0.5ms | 1.0ms | +0.5ms |
| Schema compilation (startup) | 0ms | 5ms | +5ms (one-time) |
| Binary size | 8MB | 8.5MB | +500KB |

**Verdict**: Negligible impact for CLI tool (most time spent in markdown parsing/analysis)

## Alternatives Considered

### Alternative 1: No Runtime Validation

**Approach**: Rely solely on IDE validation, skip runtime checks

**Pros**:
- ✅ Zero code changes
- ✅ Smaller binary

**Cons**:
- ❌ Users editing without IDE support get no validation
- ❌ CI/CD pipelines can't validate configs
- ❌ Errors discovered late (during analysis)
- ❌ Poor user experience for non-IDE users (vim, nano, etc.)

**Verdict**: **Rejected** - Leaves too many users without safety net

### Alternative 2: External Validator Tool

**Approach**: Provide separate `readability-validate-config` binary

**Pros**:
- ✅ Keeps main CLI binary small
- ✅ Optional for users who want it

**Cons**:
- ❌ Extra tool to install
- ❌ Users won't discover it
- ❌ Maintenance burden (two binaries)
- ❌ Inconsistent experience

**Verdict**: **Rejected** - Fragmentation is worse than binary size

### Alternative 3: Online Validator Service

**Approach**: Web service that validates configs (like JSONLint)

**Pros**:
- ✅ No CLI changes needed
- ✅ Pretty web UI

**Cons**:
- ❌ Requires internet connection
- ❌ Users must copy/paste config
- ❌ Privacy concerns (uploading config files)
- ❌ Hosting costs
- ❌ Not integrated with workflow

**Verdict**: **Rejected** - Doesn't fit CLI-first tool

### Alternative 4: Lazy Schema Compilation

**Approach**: Only compile schema if validation is explicitly requested

```go
var compiledSchema *jsonschema.Schema
var schemaCompiled bool

func Load(path string) (*Config, error) {
    // ... load YAML ...

    if !schemaCompiled {
        compiledSchema = compileSchema()
        schemaCompiled = true
    }

    // ... validate ...
}
```

**Pros**:
- ✅ Slightly faster startup if validation never used

**Cons**:
- ❌ Validation should always run (it's cheap)
- ❌ Adds complexity
- ❌ First validation call is slower

**Verdict**: **Rejected** - Premature optimization

## Implementation Roadmap

### Phase 1: Basic Validation (Week 1)

1. Add `santhosh-tekuri/jsonschema` dependency
2. Embed schema in binary
3. Add validation to `config.Load()`
4. Basic error formatting

**Deliverable**: Invalid configs are rejected with error message

### Phase 2: Enhanced Errors (Week 2)

1. Implement detailed error formatting
2. Add field path translation (JSON pointer → YAML path)
3. Add helpful suggestions
4. Link to documentation

**Deliverable**: Clear, actionable error messages

### Phase 3: CLI Flag (Week 2)

1. Add `--validate-config` flag
2. Dedicated validation command
3. Exit code handling (0 = valid, 1 = invalid)

**Deliverable**: CI-friendly validation command

### Phase 4: Testing (Week 3)

1. Unit tests for validation logic
2. Integration tests with invalid configs
3. Error message snapshot tests
4. Performance benchmarks

**Deliverable**: Comprehensive test coverage

## Testing Requirements

### Unit Tests

```go
// pkg/config/validate_test.go

func TestSchemaValidation(t *testing.T) {
    tests := []struct {
        name    string
        yaml    string
        wantErr bool
        errMsg  string
    }{
        {
            name: "valid config",
            yaml: `
thresholds:
  max_grade: 12
`,
            wantErr: false,
        },
        {
            name: "invalid type",
            yaml: `
thresholds:
  max_grade: "twelve"
`,
            wantErr: true,
            errMsg:  "expected number, got string",
        },
        {
            name: "unknown field",
            yaml: `
thresholds:
  typo_field: 100
`,
            wantErr: true,
            errMsg:  "additional properties",
        },
        {
            name: "out of range",
            yaml: `
thresholds:
  max_grade: 1000
`,
            wantErr: true,
            errMsg:  "maximum",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateYAML([]byte(tt.yaml))
            if tt.wantErr {
                require.Error(t, err)
                assert.Contains(t, err.Error(), tt.errMsg)
            } else {
                require.NoError(t, err)
            }
        })
    }
}
```

### Integration Tests

```go
func TestLoadWithValidation(t *testing.T) {
    // Create temp config file
    tmpfile, _ := os.CreateTemp("", "config-*.yml")
    defer os.Remove(tmpfile.Name())

    // Write invalid config
    tmpfile.WriteString(`
thresholds:
  max_grade: "invalid"
`)
    tmpfile.Close()

    // Should fail to load
    _, err := Load(tmpfile.Name())
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "validation failed")
}
```

### Error Message Tests

```go
func TestErrorFormatting(t *testing.T) {
    yaml := `
thresholds:
  max_grade: "twelve"
  unknown: 100
`
    err := validateYAML([]byte(yaml))
    require.Error(t, err)

    errMsg := err.Error()

    // Should contain field paths
    assert.Contains(t, errMsg, "thresholds.max_grade")
    assert.Contains(t, errMsg, "thresholds.unknown")

    // Should contain helpful messages
    assert.Contains(t, errMsg, "expected number")
    assert.Contains(t, errMsg, "not allowed")

    // Should link to docs
    assert.Contains(t, errMsg, "docs/cli/config-file.md")
}
```

## Success Metrics

- ✅ Invalid configs are caught at load time
- ✅ Error messages include field paths and suggestions
- ✅ `--validate-config` flag works in CI pipelines
- ✅ Schema validation adds <2ms to config loading
- ✅ Binary size increase <1MB
- ✅ Test coverage >90% for validation code
- ✅ Zero false positives (valid configs don't error)

## Next Steps

After runtime validation:
1. Expand [Testing Strategy](testing-strategy.md) with validation test cases
2. Update documentation to mention runtime validation
3. Add validation to CI pipeline (GitHub Actions)
4. Monitor user feedback on error messages

## References

- [santhosh-tekuri/jsonschema](https://github.com/santhosh-tekuri/jsonschema)
- [JSON Schema Validation Spec](https://json-schema.org/draft/2020-12/json-schema-validation.html)
- [Go embed Directive](https://pkg.go.dev/embed)
- [Cobra CLI Framework](https://github.com/spf13/cobra)
