# Component 3: YAML Integration

## Overview

Enable schema-driven validation and autocomplete in `.readability.yml` files by adding `$schema` references. This component bridges the JSON Schema (created in Component 1) with actual YAML configuration files used by end-users.

## Technical Approach

### Schema Reference Syntax

YAML files can reference JSON Schemas using special comments recognized by YAML language servers:

```yaml
# yaml-language-server: $schema=<schema-url>

thresholds:
  max_grade: 12
```

This comment instructs the YAML language server to validate the file against the specified schema.

### Supported Formats

#### Format 1: YAML Language Server Directive (Recommended)

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json

thresholds:
  max_grade: 12
```

**Pros**:
- ✅ Explicit and clear
- ✅ Works in all editors with yaml-language-server
- ✅ Overrides automatic detection
- ✅ Can specify exact schema version/URL

**Cons**:
- ⚠️ Requires users to add comment (unless SchemaStore provides automatic detection)

**Tooling Support**:
- VS Code ✅ (via Red Hat YAML extension)
- IntelliJ/WebStorm ✅ (native support)
- Vim/Neovim ✅ (via yaml-language-server LSP)
- Emacs ✅ (via lsp-mode + yaml-language-server)

#### Format 2: JSON Schema Property (Alternative)

```yaml
$schema: https://json.schemastore.org/readability.json

thresholds:
  max_grade: 12
```

**Pros**:
- ✅ Standard JSON Schema convention
- ✅ No special comment syntax

**Cons**:
- ❌ `$schema` becomes part of config data structure
- ❌ Go YAML parser will read it as a field
- ❌ Need to filter/ignore in config loading code
- ⚠️ Less common for YAML files

**Verdict**: **Not Recommended** - Pollutes config structure

#### Format 3: Automatic Detection (SchemaStore Only)

```yaml
# No schema reference needed - automatic!

thresholds:
  max_grade: 12
```

**How It Works**:
1. User opens `.readability.yml` in VS Code
2. YAML language server checks filename against SchemaStore catalog
3. Matches `.readability.yml` → loads schema automatically
4. Validation and autocomplete "just work"

**Pros**:
- ✅ Zero user configuration
- ✅ Best user experience
- ✅ No comments needed

**Cons**:
- ⚠️ Only works after SchemaStore approval
- ⚠️ Users must have internet connection (for schema download)
- ⚠️ Can't pin schema versions

**Availability**: Available after [Schema Publishing](02-schema-publishing.md) Phase 3 (SchemaStore approval)

## Implementation Strategy

### Phase 1: Explicit References (Immediate)

**Status**: ✅ **COMPLETE** (PR #185)

**Action**: Add `$schema` comments to all example configurations

**Files Updated**:

1. **Repository's own config** (`.readability.yml`):
   - ✅ Already updated in PR #180 (Component 1)
   - Schema reference: `./docs/schemas/config.json` (relative path)
   - YAML document marker (`---`) added

2. **Examples in documentation** (`docs/cli/config-file.md`):
   - ✅ Updated all 5 YAML examples with schema references:
     - Quick Start example
     - All Options example
     - Different Rules for Different Folders example
     - Path ordering example
     - Disabling Checks example
   - Schema reference: `https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json`
   - YAML document markers (`---`) added to all examples

3. **Example in README.md**:
   - ✅ Updated main configuration example
   - Schema reference: `https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json`
   - YAML document marker (`---`) added
   - Enhanced inline comments explaining each field purpose
   - Added link to IDE setup documentation (forward reference to Component 6)

4. **Test fixtures** (`pkg/config/testdata/*.yml`):
   - ⚠️ **N/A** - No YAML test fixtures exist in this codebase
   - Tests use in-memory Go structs in `pkg/config/config_test.go`

**Verification**:
- ✅ All examples validated against schema using `check-jsonschema`
- ✅ Schema URL accessible and returns valid JSON
- ✅ HTTP 200, Content-Type: `application/json; charset=utf-8`

### Phase 2: Documentation (Post-Schema Publishing)

**Action**: Document how to use schema references in user guides

**Add to `docs/cli/config-file.md`**:

```markdown
## IDE Support

The `.readability.yml` configuration file has JSON Schema support for autocomplete and validation in your editor.

### Automatic Detection (Recommended)

If you use VS Code or IntelliJ, schema validation is automatic—no setup needed. The YAML language server will detect `.readability.yml` files and load the schema from [SchemaStore](https://schemastore.org/).

### Manual Schema Reference

To explicitly specify the schema (for version pinning or custom URLs):

```yaml
# yaml-language-server: $schema=https://json.schemastore.org/readability.json

thresholds:
  max_grade: 12
```

### Editor Setup

#### VS Code

1. Install [YAML extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml) (usually pre-installed)
2. Open `.readability.yml` - autocomplete and validation work automatically

#### IntelliJ IDEA / WebStorm

JSON Schema support is built-in. Open `.readability.yml` and start typing.

#### Vim / Neovim

1. Install [yaml-language-server](https://github.com/redhat-developer/yaml-language-server)
2. Configure your LSP client (e.g., `nvim-lspconfig`)
3. Open `.readability.yml`

#### Emacs

1. Install `lsp-mode` and `yaml-language-server`
2. Enable `lsp-mode` for YAML files
3. Open `.readability.yml`

### Troubleshooting

If autocomplete doesn't work:
- Verify YAML language server is installed
- Check for syntax errors in your YAML
- Restart your editor/language server
- Add explicit `$schema` reference (see above)
```

### Phase 3: Automated Schema Injection (Optional)

**Idea**: CLI automatically adds `$schema` comment to generated configs

**Example** (if readability has a `--init` command):

```go
// cmd/readability/init.go (hypothetical)
func generateDefaultConfig() string {
    return `# yaml-language-server: $schema=https://json.schemastore.org/readability.json

thresholds:
  max_grade: 16
  max_ari: 16
  # ...
`
}
```

**Viability**: ⚠️ **LOW PRIORITY** (readability doesn't have init command)

**Verdict**: Skip unless init/scaffold command is added

## Viability Analysis

### ✅ High Viability

**Evidence**:
1. **Proven Pattern**: Widely used by major projects:
   - ESLint: `.eslintrc.json` has `$schema` references
   - GitHub Actions: Workflows include schema directives
   - MkDocs Material: `mkdocs.yml` supports schema
2. **Mature Tooling**: YAML language servers are stable and well-maintained
3. **Zero Breaking Changes**: Adding comments doesn't affect YAML parsing
4. **Immediate Value**: Users get IDE support as soon as schema is published

### Schema URL Selection

**For Documentation/Examples**: Use canonical domain URL (available immediately)
```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
```

**For Repository's Own Config**: Use hosted URL from day one
```yaml
# Production (immediate)
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json

# Alternative: After SchemaStore approval (automatic discovery, explicit reference optional)
# yaml-language-server: $schema=https://json.schemastore.org/readability.json
```

**For Testing/CI**: Use relative path to ensure tests don't depend on external services
```yaml
# Test fixture: pkg/config/testdata/valid.yml
# yaml-language-server: $schema=../../../schemas/readability-config.schema.json
```

## Alternatives Considered

### Alternative 1: Modeline-Style Comments

**Approach**: Use vim-style modelines

```yaml
# vim: set schema=https://json.schemastore.org/readability.json :

thresholds:
  max_grade: 12
```

**Pros**:
- Familiar to vim users

**Cons**:
- ❌ Not recognized by YAML language servers
- ❌ Non-standard
- ❌ Only works in vim (with custom plugin)

**Verdict**: **Rejected** - Doesn't solve IDE support problem

### Alternative 2: Separate Schema Mapping File

**Approach**: Users create `.vscode/settings.json` to map files to schemas

```json
{
  "yaml.schemas": {
    "https://json.schemastore.org/readability.json": ".readability.yml"
  }
}
```

**Pros**:
- ✅ Works in VS Code
- ✅ No comments in YAML files

**Cons**:
- ❌ Requires per-repository setup
- ❌ Doesn't work in other editors (IntelliJ, vim)
- ❌ Manual configuration burden
- ❌ Not discoverable by new users

**Verdict**: **Rejected as primary method** - Can document as fallback option

### Alternative 3: Embedded Schema in YAML

**Approach**: Embed schema directly in YAML files

```yaml
$schema:
  type: object
  properties:
    thresholds:
      type: object
      # ... entire schema inline

thresholds:
  max_grade: 12
```

**Pros**:
- Self-contained

**Cons**:
- ❌ Massive duplication (every file has full schema)
- ❌ Pollutes config files (hundreds of lines)
- ❌ Hard to maintain
- ❌ Not how JSON Schema is meant to work

**Verdict**: **Rejected** - Completely impractical

### Alternative 4: No Schema Reference (Rely on SchemaStore)

**Approach**: Don't add `$schema` comments, rely solely on SchemaStore automatic detection

**Pros**:
- ✅ Cleanest YAML files (no comments)
- ✅ Zero user effort

**Cons**:
- ⚠️ Only works after SchemaStore approval
- ⚠️ No version pinning
- ⚠️ Doesn't help users who want to test schema before publication

**Verdict**: **Use for final state** - Add explicit `$schema` for Phase 1-2, remove later (optional)

## Implementation Guidance

### Step 1: Update Repository Config

```bash
# Edit .readability.yml
sed -i '1i# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json\n' .readability.yml
```

Verify in VS Code that autocomplete works.

### Step 2: Update Documentation Examples

Update all YAML examples in:
- `README.md`
- `docs/cli/config-file.md`
- Any other documentation with config examples

Add schema reference to the top of each example.

### Step 3: Update Test Fixtures

For local testing, use relative paths:

```yaml
# pkg/config/testdata/valid-minimal.yml
# yaml-language-server: $schema=../../../schemas/readability-config.schema.json

thresholds:
  max_grade: 10
```

### Step 4: Document IDE Setup

Add "IDE Support" section to `docs/cli/config-file.md` (see Phase 2 example above).

### Step 5: Create Quick Start Example

Add to README.md:

```markdown
### Configuration with IDE Support

Create `.readability.yml` in your repository root:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json

thresholds:
  max_grade: 12      # High school senior reading level
  max_ari: 12        # Automated Readability Index
  min_ease: 40       # Flesch Reading Ease (0-100)
  max_lines: 300     # Maximum prose lines per file
```

Your editor will provide autocomplete and validation as you type.
```

## Testing Requirements

### Manual IDE Tests

Test schema integration in multiple editors:

| Editor | Test Steps | Expected Behavior |
|--------|-----------|-------------------|
| VS Code | Open `.readability.yml`, type `thresholds:` | Autocomplete suggests fields |
| VS Code | Set `max_grade: "twelve"` | Red squiggle, error message |
| VS Code | Hover over `max_grade` | Tooltip shows description |
| IntelliJ | Same as VS Code | Same behavior |
| Vim (with LSP) | Same as VS Code | Same behavior (if yaml-language-server configured) |

### Automated Tests

1. **YAML Validity**: All example YAML files parse correctly
2. **Schema Reference Syntax**: `$schema` comments are well-formed
3. **Schema URL Accessibility**: Schema URL returns 200 OK
4. **Relative Path Resolution**: Test fixtures can load local schema

### Documentation Tests

1. **Example Validation**: All documented examples validate against schema
2. **Code Block Syntax**: YAML code blocks have correct triple-backtick formatting
3. **Link Verification**: Schema URLs in docs are accessible

## Success Metrics

- ✅ All repository YAML files include `$schema` reference
- ✅ Documentation includes IDE setup instructions
- ✅ Example configurations validate in VS Code
- ✅ Autocomplete works in at least 2 IDEs (VS Code, IntelliJ)
- ✅ Users report successful schema-driven editing (GitHub Discussions/Issues)

## Rollout Plan

### Week 1: Preparation
- Add `$schema` to repo's `.readability.yml`
- Test IDE integration locally
- Update one documentation example as proof-of-concept

### Week 2: Documentation
- Update all documentation examples
- Add IDE setup guide
- Update README with schema reference example

### Week 3: Schema Publishing
- Publish schema (see [Schema Publishing](02-schema-publishing.md))
- Update `$schema` URLs to published location
- Test remote schema loading

### Week 4: User Communication
- Announce in README/CHANGELOG
- Create GitHub Discussion about IDE support
- Update issue templates to reference schema validation

## Next Steps

After YAML integration:
1. Implement [Runtime Validation](04-runtime-validation.md) for CLI-based schema checking
2. Expand [Testing Strategy](05-testing-strategy.md) to cover schema scenarios
3. Monitor user feedback and iterate on schema/documentation

## References

- [YAML Language Server Directives](https://github.com/redhat-developer/yaml-language-server#language-server-settings)
- [VS Code YAML Extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml)
- [JSON Schema in YAML](https://json-schema.org/understanding-json-schema/reference/schema.html)
