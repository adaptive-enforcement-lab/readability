# Component 5: Testing Strategy

## Overview

Comprehensive test coverage for JSON Schema implementation across all components: schema creation, publishing, YAML integration, and runtime validation. Tests must verify functionality, catch regressions, and ensure excellent user experience.

## Testing Pyramid

```
           ╱╲
          ╱  ╲
         ╱ E2E ╲          ← Integration tests (IDE behavior, CLI workflow)
        ╱--------╲
       ╱  Unit    ╲       ← Schema validation, error formatting
      ╱____________╲
     ╱   Schema     ╲     ← Schema validity, example validation
    ╱________________╲
```

## Layer 1: Schema Validation Tests

### Purpose

Verify the JSON Schema itself is valid, complete, and matches Go struct definitions.

### Test Categories

#### 1.1 Schema Validity

**Test**: Schema conforms to JSON Schema Draft 2020-12 specification

```bash
# Using ajv-cli
npm install -g ajv-cli
ajv compile -s schemas/readability-config.schema.json
```

**Expected**: Schema compiles without errors

**Automation**:
```yaml
# .github/workflows/test-schema.yml
name: Schema Tests
on: [push, pull_request]

jobs:
  validate-schema:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
      - run: npm install -g ajv-cli
      - run: ajv compile -s schemas/readability-config.schema.json
```

#### 1.2 Schema Completeness

**Test**: All Go struct fields have corresponding schema properties

```go
// pkg/config/schema_test.go

func TestSchemaCompleteness(t *testing.T) {
    // Load schema
    schemaData, err := os.ReadFile("../../schemas/readability-config.schema.json")
    require.NoError(t, err)

    var schema map[string]interface{}
    require.NoError(t, json.Unmarshal(schemaData, &schema))

    // Check Thresholds fields
    thresholdProps := getNestedMap(schema, "properties", "thresholds", "properties")

    expectedFields := []string{
        "max_grade", "max_ari", "max_fog", "min_ease",
        "max_lines", "min_words", "min_admonitions", "max_dash_density",
    }

    for _, field := range expectedFields {
        assert.Contains(t, thresholdProps, field,
            "Schema missing field: %s", field)
    }

    // Check PathOverride structure
    overrideItemProps := getNestedMap(schema, "properties", "overrides", "items", "properties")
    assert.Contains(t, overrideItemProps, "path")
    assert.Contains(t, overrideItemProps, "thresholds")
}
```

**Expected**: No missing fields

#### 1.3 Example Validation

**Test**: All documented YAML examples validate against schema

```go
func TestDocumentedExamplesValidate(t *testing.T) {
    examples := []struct {
        name     string
        yamlPath string
    }{
        {"Repository config", "../../.readability.yml"},
        {"README example", testdata/example-readme.yml"},
        {"Docs example", "testdata/example-docs.yml"},
        {"Override example", "testdata/example-overrides.yml"},
    }

    schema := loadSchema(t)

    for _, ex := range examples {
        t.Run(ex.name, func(t *testing.T) {
            data, err := os.ReadFile(ex.yamlPath)
            require.NoError(t, err)

            var yamlData interface{}
            require.NoError(t, yaml.Unmarshal(data, &yamlData))

            err = schema.Validate(yamlData)
            assert.NoError(t, err, "Example should validate: %s", ex.name)
        })
    }
}
```

**Expected**: All examples pass validation

#### 1.4 Type Constraints

**Test**: Schema enforces correct types for each field

```go
func TestSchemaTypeConstraints(t *testing.T) {
    schema := loadSchema(t)

    tests := []struct {
        name    string
        yaml    string
        wantErr bool
    }{
        {
            name: "max_grade as string",
            yaml: `thresholds: {max_grade: "twelve"}`,
            wantErr: true,
        },
        {
            name: "max_grade as number",
            yaml: `thresholds: {max_grade: 12}`,
            wantErr: false,
        },
        {
            name: "max_lines as float",
            yaml: `thresholds: {max_lines: 12.5}`,
            wantErr: true,
        },
        {
            name: "max_lines as integer",
            yaml: `thresholds: {max_lines: 375}`,
            wantErr: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var data interface{}
            yaml.Unmarshal([]byte(tt.yaml), &data)

            err := schema.Validate(data)
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

#### 1.5 Range Validation

**Test**: Schema enforces min/max bounds

```go
func TestSchemaRangeValidation(t *testing.T) {
    schema := loadSchema(t)

    tests := []struct {
        field   string
        value   interface{}
        wantErr bool
    }{
        {"max_grade", 10, false},      // Valid
        {"max_grade", 0, false},       // Min edge
        {"max_grade", 30, false},      // Max edge
        {"max_grade", -5, true},       // Below min
        {"max_grade", 100, true},      // Above max
        {"min_ease", -100, false},     // Negative allowed (disable check)
        {"min_ease", 101, true},       // Above max
        {"min_admonitions", -1, false}, // Negative allowed
        {"min_admonitions", -2, true},  // Below min
    }

    for _, tt := range tests {
        t.Run(fmt.Sprintf("%s=%v", tt.field, tt.value), func(t *testing.T) {
            yaml := fmt.Sprintf(`thresholds: {%s: %v}`, tt.field, tt.value)
            var data interface{}
            yaml.Unmarshal([]byte(yaml), &data)

            err := schema.Validate(data)
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

## Layer 2: Runtime Validation Tests

### Purpose

Test the Go implementation of schema validation in `pkg/config`.

### Test Categories

#### 2.1 Valid Configuration Loading

```go
func TestLoadValidConfig(t *testing.T) {
    tests := []struct {
        name   string
        yaml   string
        expect *Config
    }{
        {
            name: "minimal config",
            yaml: `thresholds: {max_grade: 12}`,
            expect: &Config{
                Thresholds: Thresholds{
                    MaxGrade: 12,
                    MaxARI:   16, // Defaults
                    // ...
                },
            },
        },
        {
            name: "complete config",
            yaml: `
thresholds:
  max_grade: 14
  max_ari: 14
  max_fog: 16
  min_ease: 30
  max_lines: 400
  min_words: 150
  min_admonitions: 2
  max_dash_density: 5
`,
            expect: &Config{
                Thresholds: Thresholds{
                    MaxGrade:       14,
                    MaxARI:         14,
                    MaxFog:         16,
                    MinEase:        30,
                    MaxLines:       400,
                    MinWords:       150,
                    MinAdmonitions: 2,
                    MaxDashDensity: 5,
                },
            },
        },
        {
            name: "config with overrides",
            yaml: `
thresholds:
  max_grade: 16
overrides:
  - path: docs/guide/
    thresholds:
      max_grade: 12
`,
            expect: &Config{
                Thresholds: Thresholds{MaxGrade: 16, /* ... */},
                Overrides: []PathOverride{
                    {Path: "docs/guide/", Thresholds: Thresholds{MaxGrade: 12}},
                },
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cfg, err := loadYAMLString(tt.yaml)
            require.NoError(t, err)
            assert.Equal(t, tt.expect, cfg)
        })
    }
}
```

#### 2.2 Invalid Configuration Rejection

```go
func TestLoadInvalidConfig(t *testing.T) {
    tests := []struct {
        name      string
        yaml      string
        errSubstr string // Substring expected in error message
    }{
        {
            name:      "wrong type",
            yaml:      `thresholds: {max_grade: "twelve"}`,
            errSubstr: "expected number",
        },
        {
            name:      "unknown field",
            yaml:      `thresholds: {typo_field: 100}`,
            errSubstr: "not allowed",
        },
        {
            name:      "out of range",
            yaml:      `thresholds: {max_grade: 1000}`,
            errSubstr: "maximum",
        },
        {
            name:      "negative where not allowed",
            yaml:      `thresholds: {max_grade: -5}`,
            errSubstr: "minimum",
        },
        {
            name:      "missing required path in override",
            yaml:      `overrides: [{thresholds: {max_grade: 10}}]`,
            errSubstr: "required",
        },
        {
            name:      "invalid YAML syntax",
            yaml:      `thresholds: {max_grade: 12`,
            errSubstr: "invalid YAML",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := loadYAMLString(tt.yaml)
            require.Error(t, err)
            assert.Contains(t, err.Error(), tt.errSubstr)
        })
    }
}
```

#### 2.3 Error Message Quality

```go
func TestErrorMessageFormatting(t *testing.T) {
    yaml := `
thresholds:
  max_grade: "twelve"
  unknown_field: 100
`
    _, err := loadYAMLString(yaml)
    require.Error(t, err)

    errMsg := err.Error()

    // Should be multi-line with clear structure
    assert.Contains(t, errMsg, "Configuration validation failed")

    // Should list all errors
    assert.Contains(t, errMsg, "thresholds.max_grade")
    assert.Contains(t, errMsg, "thresholds.unknown_field")

    // Should provide context
    assert.Contains(t, errMsg, "expected number")
    assert.Contains(t, errMsg, "not allowed")

    // Should link to docs
    assert.Contains(t, errMsg, "docs/cli/config-file.md")
}
```

#### 2.4 CLI Flag: --validate-config

```go
func TestValidateConfigFlag(t *testing.T) {
    tests := []struct {
        name       string
        configYAML string
        wantExit   int
        wantOutput string
    }{
        {
            name: "valid config",
            configYAML: `thresholds: {max_grade: 12}`,
            wantExit:   0,
            wantOutput: "✓ Configuration is valid",
        },
        {
            name: "invalid config",
            configYAML: `thresholds: {max_grade: "invalid"}`,
            wantExit:   1,
            wantOutput: "Configuration validation failed",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create temp config
            tmpfile := createTempConfig(t, tt.configYAML)
            defer os.Remove(tmpfile)

            // Run CLI with --validate-config
            cmd := exec.Command("readability", "--config", tmpfile, "--validate-config")
            output, err := cmd.CombinedOutput()

            if tt.wantExit == 0 {
                assert.NoError(t, err)
            } else {
                assert.Error(t, err)
            }

            assert.Contains(t, string(output), tt.wantOutput)
        })
    }
}
```

## Layer 3: IDE Integration Tests

### Purpose

Verify schema works correctly in real IDEs (not mocked).

### Test Categories

#### 3.1 VS Code Integration

**Manual Test Procedure**:

1. Install VS Code with YAML extension
2. Create test `.readability.yml` file:
   ```yaml
   # yaml-language-server: $schema=https://json.schemastore.org/readability.json

   thresholds:
     max_
   ```
3. Verify autocomplete appears after typing `max_`
4. Add invalid value: `max_grade: "twelve"`
5. Verify red squiggle appears
6. Hover over `max_grade`
7. Verify tooltip shows description from schema

**Expected Results**:
- ✅ Autocomplete works
- ✅ Validation errors shown
- ✅ Tooltips display descriptions
- ✅ No false positives

**Automation** (Headless):
```yaml
# .github/workflows/test-ide.yml
name: IDE Integration
on: [push]

jobs:
  vscode-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
      - run: npm install -g yaml-language-server
      - run: |
          # Start yaml-language-server
          # Send LSP requests (didOpen, completion, hover)
          # Verify responses contain expected data
          ./scripts/test-lsp.sh
```

#### 3.2 IntelliJ Integration

**Manual Test**: Same as VS Code

**Expected**: Native JSON Schema support works without plugins

#### 3.3 Vim/Neovim Integration

**Manual Test**:
1. Install yaml-language-server
2. Configure LSP client (e.g., nvim-lspconfig)
3. Open `.readability.yml`
4. Test completion, diagnostics, hover

**Expected**: Works if yaml-language-server is configured

## Layer 4: End-to-End Tests

### Purpose

Test complete workflows from config creation to analysis.

### Test Scenarios

#### 4.1 Happy Path

```bash
# User creates config
cat > .readability.yml << 'EOF'
# yaml-language-server: $schema=https://json.schemastore.org/readability.json

thresholds:
  max_grade: 12
  min_ease: 40
EOF

# Config validates
readability --validate-config
# Output: ✓ Configuration is valid

# Analysis runs
readability docs/
# Output: Analysis results...
```

**Expected**: No errors, analysis completes

#### 4.2 Invalid Config Caught Early

```bash
# User creates invalid config
cat > .readability.yml << 'EOF'
thresholds:
  max_grade: "twelve"
EOF

# Validation fails
readability --validate-config
# Output: Configuration validation failed...

# Analysis also fails
readability docs/
# Output: Configuration validation failed... (same error)
```

**Expected**: Clear error message, non-zero exit code

#### 4.3 Schema Update Workflow

```bash
# Developer updates schema (adds new field)
# Schema published to SchemaStore
# User updates config (IDE autocompletes new field)
# Old CLI version ignores new field (forward compatibility)
# New CLI version validates new field
```

**Expected**: Graceful handling of version mismatches

## Test Coverage Goals

| Component | Coverage Target | Rationale |
|-----------|----------------|-----------|
| Schema validation logic | >90% | Critical path |
| Error formatting | >85% | User-facing |
| Config loading | >80% | Well-tested existing code |
| CLI flags | >70% | Integration tests |
| Overall | >80% | Production quality |

## Pre-Commit Hooks

### Schema Validation Hook

Add to `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: local
    hooks:
      # ... existing hooks ...

      - id: validate-json-schema
        name: Validate JSON Schema
        entry: >
          bash -c '
          command -v check-jsonschema >/dev/null || pipx install check-jsonschema;
          check-jsonschema --check-metaschema docs/schemas/config.json
          '
        language: system
        files: '^docs/schemas/config\.json$'
        pass_filenames: false

      - id: validate-readability-config
        name: Validate .readability.yml against schema
        entry: >
          bash -c '
          command -v check-jsonschema >/dev/null || pipx install check-jsonschema;
          check-jsonschema --schemafile docs/schemas/config.json .readability.yml
          '
        language: system
        files: '^\.readability\.yml$'
        pass_filenames: false
```

**What This Does**:
- **validate-json-schema**: Validates `docs/schemas/config.json` against JSON Schema Draft 2020-12 meta-schema
- **validate-readability-config**: Validates `.readability.yml` against the schema

**Triggers**:
- Runs automatically on `git commit` when either file changes
- Prevents commits with invalid schema or config
- Catches errors before they reach CI

**Installation**:
```bash
# Install pre-commit
pip install pre-commit

# Install hooks
pre-commit install

# Test hooks manually
pre-commit run validate-json-schema --all-files
pre-commit run validate-readability-config --all-files
```

## Continuous Integration

### GitHub Actions Workflow

```yaml
# .github/workflows/test.yml
name: Tests
on: [push, pull_request]

jobs:
  test-schema:
    name: Schema Validity
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: '3.11'
      - run: pipx install check-jsonschema
      - name: Validate schema meta-schema compliance
        run: check-jsonschema --check-metaschema docs/schemas/config.json
      - name: Validate .readability.yml against schema
        run: check-jsonschema --schemafile docs/schemas/config.json .readability.yml
      - name: Validate test configs against schema
        run: |
          check-jsonschema --schemafile docs/schemas/config.json \
            pkg/config/testdata/valid-*.yml

  test-go:
    name: Go Tests
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      - run: go test -v -race -coverprofile=coverage.out ./...
      - run: go tool cover -func=coverage.out

  test-examples:
    name: Example Validation
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
      - run: go build ./cmd/readability
      - run: ./readability --validate-config
      - run: |
          for f in pkg/config/testdata/*.yml; do
            ./readability --config "$f" --validate-config || true
          done
```

**Key Changes from Baseline**:
1. **test-schema job**: Added schema meta-schema validation and config file validation using `check-jsonschema`
2. **Python setup**: Added to install `check-jsonschema` (better Draft 2020-12 support than ajv)
3. **Multiple validation targets**: Validates production config and test fixtures

**Why check-jsonschema**:
- Native Draft 2020-12 support (no --spec flag needed)
- Better error messages for YAML files
- Handles both meta-schema and instance validation
- Available via pipx (no npm required)

## Test Maintenance

### When to Update Tests

1. **Schema changes**: Update type/range tests
2. **New config fields**: Add to completeness tests
3. **Error message changes**: Update snapshot tests
4. **New Go version**: Test with latest Go release
5. **New IDE versions**: Verify integration still works

### Test Documentation

Each test file should include:
```go
// pkg/config/schema_test.go

// Package config tests schema validation functionality.
//
// Test Categories:
// - Schema Validity: Tests that verify the JSON Schema itself is correct
// - Runtime Validation: Tests that verify Go code validates configs correctly
// - Error Messages: Tests that verify helpful error messages
//
// Running Tests:
//   go test ./pkg/config -v -run TestSchema
//
// Adding Tests:
//   When adding new config fields, update TestSchemaCompleteness.
//   When changing validation rules, update TestSchemaTypeConstraints.
```

## Performance Testing

### Benchmark Goals

```go
func BenchmarkSchemaValidation(b *testing.B) {
    cfg := `
thresholds:
  max_grade: 12
  max_ari: 12
overrides:
  - path: docs/
    thresholds:
      max_grade: 10
`
    data := []byte(cfg)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Load(data)
    }
}
```

**Target**: <1ms per validation (on modern hardware)

### Performance Regression Tests

```yaml
# .github/workflows/benchmark.yml
name: Benchmarks
on: [push]

jobs:
  benchmark:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
      - run: go test -bench=. -benchmem ./pkg/config > new.txt
      - uses: actions/cache@v4
        with:
          path: old.txt
          key: benchmark-${{ github.base_ref }}
      - run: benchstat old.txt new.txt || true
```

## Success Metrics

- ✅ Schema validity tests pass
- ✅ All documented examples validate
- ✅ Invalid configs trigger clear errors
- ✅ Test coverage >80%
- ✅ No flaky tests
- ✅ IDE integration verified in 2+ editors
- ✅ CI pipeline runs tests on every PR
- ✅ Benchmarks show <2ms validation overhead

## Next Steps

After implementing testing strategy:
1. Review [Component Viability Summary](08-viability-summary.md)
2. Begin implementation starting with [Schema Creation](01-schema-creation.md)
3. Run tests continuously during development
4. Update tests as feedback is received

## References

- [Go Testing Documentation](https://pkg.go.dev/testing)
- [Table-Driven Tests in Go](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- [ajv-cli Documentation](https://github.com/ajv-validator/ajv-cli)
- [YAML Language Server](https://github.com/redhat-developer/yaml-language-server)
