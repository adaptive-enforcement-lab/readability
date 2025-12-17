# Component 6: Documentation Updates

## Overview

Comprehensive documentation updates to guide users through JSON Schema features, IDE setup, and configuration best practices. Documentation must be clear, accurate, and actionable.

## Documentation Strategy

### Audience Segmentation

| Audience | Needs | Documentation Focus |
|----------|-------|-------------------|
| **New Users** | Quick start, basic config | Getting started guide, simple examples |
| **Power Users** | Advanced config, overrides | Complete reference, path matching rules |
| **IDE Users** | Setup autocomplete | Editor-specific guides |
| **CI/CD Users** | Automated validation | `--validate-config` flag, exit codes |
| **Contributors** | Schema maintenance | Schema structure, update process |

## Files to Update

### 1. README.md

**Section**: Configuration

**Current State**: Basic YAML example without schema reference

**Proposed Update**:

```markdown
### Configuration

Create `.readability.yml` in your repository root:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json

thresholds:
  max_grade: 12      # Maximum Flesch-Kincaid grade level
  max_ari: 12        # Maximum Automated Readability Index
  min_ease: 40       # Minimum Flesch Reading Ease (0-100 scale)
  max_lines: 300     # Maximum lines of prose per file
```

**IDE Support**: Your editor will provide autocomplete, validation, and inline documentation as you type. See [Configuration Guide](docs/cli/config-file.md#ide-support) for setup instructions.

**Validation**: Test your configuration before running analysis:
```bash
readability --validate-config
```
```

**Rationale**:
- Introduce schema immediately (first exposure)
- Highlight IDE benefits
- Show validation flag for CI use cases

---

### 2. docs/cli/config-file.md

**Section**: New section "IDE Support"

**Location**: After "Configuration File Format" section

**Content**:

```markdown
## IDE Support

The `.readability.yml` configuration file has full JSON Schema support, providing autocomplete, real-time validation, and inline documentation in your editor.

### Automatic Schema Detection

If you use **VS Code**, **IntelliJ IDEA**, or **WebStorm**, schema validation works automatically—no setup needed. The YAML language server detects `.readability.yml` files and loads the schema from [SchemaStore](https://schemastore.org/).

Open `.readability.yml` and start typing to see autocomplete suggestions.

### Manual Schema Reference

To explicitly specify the schema (for version pinning or custom URLs):

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json

thresholds:
  max_grade: 12
```

Add this comment as the first line of your configuration file.

### Editor Setup

#### VS Code

**No setup required** if you have the [YAML extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml) installed (usually pre-installed).

**Manual installation**:
1. Open VS Code Extensions (Cmd/Ctrl+Shift+X)
2. Search for "YAML" by Red Hat
3. Click Install
4. Reload VS Code

Open `.readability.yml` and autocomplete will work automatically.

#### IntelliJ IDEA / WebStorm

**No setup required**. JSON Schema support is built-in.

Open `.readability.yml` and start typing to see autocomplete suggestions.

#### Vim / Neovim

1. Install [yaml-language-server](https://github.com/redhat-developer/yaml-language-server):
   ```bash
   npm install -g yaml-language-server
   ```

2. Configure your LSP client (e.g., `nvim-lspconfig`):
   ```lua
   require('lspconfig').yamlls.setup{}
   ```

3. Open `.readability.yml`—autocomplete and validation will work

#### Emacs

1. Install `lsp-mode` and `yaml-language-server`
2. Enable `lsp-mode` for YAML files:
   ```elisp
   (add-hook 'yaml-mode-hook #'lsp)
   ```
3. Open `.readability.yml`

### Features

With schema support, your editor provides:

- **Autocomplete**: Press Ctrl+Space to see available fields
- **Validation**: Real-time error checking as you type
- **Documentation**: Hover over fields to see descriptions and examples
- **Type Checking**: Catch typos and invalid values before running analysis

### Troubleshooting

**Autocomplete doesn't work**:
- Verify YAML language server is installed
- Check for YAML syntax errors (invalid indentation, missing colons)
- Restart your editor/language server
- Add explicit `$schema` reference (see Manual Schema Reference above)

**Wrong schema loaded**:
- Add explicit `$schema` reference to override automatic detection
- Clear language server cache (restart editor)

**Validation errors on valid config**:
- Ensure you're using the latest schema version
- Check if your config uses experimental/deprecated fields
- Report issue at https://github.com/markcheret/readability/issues

### Runtime Validation

In addition to IDE validation, the CLI validates configuration at runtime:

```bash
# Validate configuration and exit
readability --validate-config

# Validation happens automatically during analysis
readability docs/
```

Invalid configurations will fail with detailed error messages:

```
Configuration validation failed:

  • thresholds.max_grade
    expected number, got string
    Suggestion: Remove quotes around numeric values

See https://github.com/markcheret/readability/blob/main/docs/cli/config-file.md for reference.
```
```

**Rationale**:
- Covers all major editors
- Step-by-step setup for each
- Troubleshooting for common issues
- Links to runtime validation

---

### 3. CHANGELOG.md

**Section**: Unreleased

**Entry**:

```markdown
### Added
- JSON Schema support for `.readability.yml` configuration files
  - IDE autocomplete and validation in VS Code, IntelliJ, Vim, and other schema-aware editors
  - Schema published to [SchemaStore](https://schemastore.org/) for automatic discovery
  - Runtime validation with detailed error messages and suggestions
  - `--validate-config` flag to test configuration without running analysis
  - Schema available at:
    - https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json (canonical)
    - https://json.schemastore.org/readability.json (after SchemaStore approval)

### Changed
- Configuration loading now validates against JSON Schema, providing earlier error detection
- Error messages for invalid configurations are more detailed and actionable
```

---

### 4. Contributing Guide (docs/contributing.md or CONTRIBUTING.md)

**Section**: New section "Updating Configuration Schema"

**Content**:

```markdown
## Updating Configuration Schema

When adding new configuration fields to `pkg/config/config.go`, update the JSON Schema:

### 1. Update Schema File

Edit `schemas/readability-config.schema.json`:

```json
{
  "properties": {
    "thresholds": {
      "properties": {
        "new_field": {
          "type": "number",
          "description": "Description of what this field does",
          "minimum": 0,
          "default": 10,
          "examples": [5, 10, 15]
        }
      }
    }
  }
}
```

### 2. Validate Schema

```bash
npm install -g ajv-cli
ajv compile -s schemas/readability-config.schema.json
```

### 3. Update Tests

Add test case to `pkg/config/schema_test.go`:

```go
func TestSchemaCompleteness(t *testing.T) {
    // ... existing tests ...

    // Add new field to expected list
    expectedFields := []string{
        // ... existing fields ...
        "new_field",
    }
}
```

### 4. Update Documentation

- Add field to examples in `docs/cli/config-file.md`
- Document behavior in appropriate sections
- Add migration notes if changing existing fields

### 5. Version Schema (if breaking change)

For breaking changes (rare):
- Increment schema `$id` version
- Document migration path
- Consider backward compatibility period

### Schema Maintenance Checklist

- [ ] Go struct updated
- [ ] JSON Schema updated
- [ ] Schema validates with ajv-cli
- [ ] Tests updated
- [ ] Documentation updated
- [ ] Examples updated
- [ ] CHANGELOG entry added
```

**Rationale**:
- Clear process for contributors
- Prevents schema drift
- Ensures consistency

---

### 5. GitHub Issue Templates

**Template**: `.github/ISSUE_TEMPLATE/config-error.md`

**New Template**:

```markdown
---
name: Configuration Error
about: Report an issue with .readability.yml configuration
title: '[Config] '
labels: 'configuration'
---

## Configuration Error

**Config File**:
```yaml
# Paste your .readability.yml here
```

**Error Message**:
```
# Paste error output here
```

**Expected Behavior**:
<!-- Describe what you expected to happen -->

**Schema Validation**:
Did you validate your config?
```bash
readability --validate-config
```
Output:
```
# Paste validation output
```

**IDE Used**:
- [ ] VS Code
- [ ] IntelliJ IDEA
- [ ] Vim/Neovim
- [ ] Other: _____

**Schema Loaded in IDE**:
- [ ] Yes (autocomplete worked)
- [ ] No
- [ ] Unknown

**Additional Context**:
<!-- Any other information that might help -->
```

**Rationale**:
- Guides users to validate config first
- Collects relevant diagnostics
- Reduces back-and-forth in issue triage

---

## Documentation Testing

### Validation Checklist

Before publishing documentation:

- [ ] All YAML examples validate against schema
- [ ] All code blocks have correct syntax highlighting
- [ ] All links are accessible (no 404s)
- [ ] Schema URLs return 200 OK
- [ ] Editor setup steps tested in actual editors
- [ ] Troubleshooting steps are accurate
- [ ] Examples match current schema version

### Automated Checks

```yaml
# .github/workflows/docs-test.yml
name: Documentation Tests
on: [push, pull_request]

jobs:
  validate-examples:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
      - run: go build ./cmd/readability

      # Extract YAML examples from markdown docs
      - run: |
          # Find all YAML code blocks in docs/
          # Save to temp files
          # Validate each with readability --validate-config
          ./scripts/validate-doc-examples.sh

  check-links:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: lycheeverse/lychee-action@v2
        with:
          args: --verbose --no-progress 'docs/**/*.md' 'README.md'
```

**Script**: `scripts/validate-doc-examples.sh`

```bash
#!/bin/bash
set -e

# Extract YAML code blocks from markdown files
# and validate them

DOCS_DIR="docs"
TEMP_DIR=$(mktemp -d)

echo "Extracting YAML examples from documentation..."

# Find all .md files and extract YAML blocks
find "$DOCS_DIR" -name "*.md" -exec grep -A 50 '^```yaml' {} \; | \
    grep -B 1 -A 50 '^```yaml' | \
    sed -n '/^```yaml$/,/^```$/p' | \
    sed '/^```/d' > "$TEMP_DIR/examples.yml"

echo "Validating examples..."
./readability --config "$TEMP_DIR/examples.yml" --validate-config

echo "✓ All documentation examples are valid"
```

---

## Documentation Style Guide

### YAML Examples

**Always include**:
- Schema reference comment
- Inline comments explaining values
- Realistic, production-ready values

**Good Example**:
```yaml
# yaml-language-server: $schema=https://json.schemastore.org/readability.json

thresholds:
  max_grade: 12      # High school senior reading level
  max_ari: 12        # Automated Readability Index
  min_ease: 40       # Flesch Reading Ease (higher = easier)
  max_lines: 300     # Maximum prose lines per file
```

**Bad Example**:
```yaml
thresholds:
  max_grade: 12
  max_ari: 12
```

### Field Descriptions

**Format**: `<field>: <value> # <description>`

**Good**:
```yaml
max_grade: 12      # Maximum Flesch-Kincaid grade level (12 = high school)
```

**Bad**:
```yaml
max_grade: 12      # grade
```

### Error Messages in Docs

When documenting error messages, use actual output:

**Good**:
```
Configuration validation failed:

  • thresholds.max_grade
    expected number, got string
    Suggestion: Remove quotes around numeric values
```

**Bad**:
"You'll get an error if the type is wrong."

---

## Documentation Maintenance

### Update Triggers

Update documentation when:
1. **New config field added** → Update schema, examples, reference
2. **Field behavior changes** → Update descriptions, migration notes
3. **Error message changes** → Update troubleshooting section
4. **New editor support** → Add to IDE setup guide
5. **Schema URL changes** → Update all references

### Review Cycle

- **Weekly**: Check for broken links
- **Per release**: Verify all examples validate
- **Per major version**: Review entire docs for accuracy

### User Feedback Loop

Monitor for documentation issues:
- GitHub Issues tagged "documentation"
- Questions in Discussions
- Recurring support requests

Integrate feedback:
- Add FAQ entries
- Improve unclear sections
- Add more examples

---

## Success Metrics

Documentation is successful when:

- ✅ Users can set up IDE support without assistance
- ✅ Common config errors are self-service (troubleshooting guide works)
- ✅ GitHub Issues about config decrease by >50%
- ✅ Positive feedback on IDE integration (Discussions, social media)
- ✅ All examples validate against current schema
- ✅ Zero broken links in documentation

---

## Next Steps

After documentation updates:
1. Review [Viability Summary](08-viability-summary.md) for overall assessment
2. Begin implementation following phased approach
3. Update docs incrementally as features ship
4. Collect user feedback and iterate

---

## References

- [Writing Great Documentation](https://documentation.divio.com/)
- [Markdown Style Guide](https://google.github.io/styleguide/docguide/style.html)
- [SchemaStore Catalog](https://www.schemastore.org/json/)
- [YAML Language Server](https://github.com/redhat-developer/yaml-language-server)
