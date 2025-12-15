# Component 7: SchemaStore Submission (Future)

## Overview

Submit the JSON Schema to [SchemaStore.org](https://schemastore.org/) for automatic discovery in IDEs. This is a **future enhancement** to be done after the schema is stable, tested, and has been used in production.

## Why Wait?

**Do this AFTER**:
- ✅ Schema is published on your own domain
- ✅ Schema has been tested with real users
- ✅ Schema structure is stable (no frequent changes)
- ✅ Common configuration questions have been answered
- ✅ Edge cases have been discovered and documented

**Timeline**: Recommend waiting **2-4 weeks** after initial schema deployment before submitting to SchemaStore.

**Rationale**:
- Avoid submitting a schema that needs frequent updates
- Ensure schema validation rules match real-world usage
- Gather feedback on schema descriptions and examples
- Reduce back-and-forth in SchemaStore PR review

## What is SchemaStore?

[SchemaStore.org](https://schemastore.org/) is a centralized registry of JSON schemas maintained by the community. It powers automatic schema detection in:
- VS Code (built-in)
- Visual Studio
- JetBrains IDEs (IntelliJ, WebStorm, PyCharm, etc.)
- Rider
- Eclipse

### How It Works

When you open `.readability.yml` in VS Code, the YAML language server:
1. Extracts the filename pattern `.readability.yml`
2. Queries SchemaStore's catalog
3. Automatically loads the associated schema
4. Provides validation and autocomplete

**No user configuration required** - it's completely automatic.

## Benefits of SchemaStore

**With SchemaStore** (automatic):
```yaml
# User just opens .readability.yml - no schema reference needed
thresholds:
  max_grade: 12  # Autocomplete works automatically!
```

**Without SchemaStore** (manual):
```yaml
# User must add schema reference manually
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/schemas/config.json

thresholds:
  max_grade: 12  # Autocomplete only works with explicit reference
```

**Impact**:
- ✅ Better first-run experience for new users
- ✅ No need to document schema setup
- ✅ Works "out of the box" in VS Code
- ✅ Reduced support burden

## Prerequisites

Before submitting to SchemaStore:

- [ ] Schema has been published at `readability.adaptive-enforcement-lab.com` for 2+ weeks
- [ ] Schema `$id` references your canonical URL
- [ ] Schema has been tested in multiple IDEs (VS Code, IntelliJ, Vim)
- [ ] Schema includes comprehensive descriptions for all fields
- [ ] Schema includes examples for common values
- [ ] Schema validation rules have been tested with real configs
- [ ] No major schema changes expected in near future
- [ ] User feedback on schema has been positive

## Submission Process

### Step 1: Fork SchemaStore Repository

```bash
git clone https://github.com/SchemaStore/schemastore.git
cd schemastore
git checkout -b add-readability-schema
```

### Step 2: Add Schema to Catalog

Edit `src/api/json/catalog.json` and add your entry:

```json
{
  "name": "Readability Configuration",
  "description": "Configuration file for readability markdown analyzer",
  "fileMatch": [
    ".readability.yml",
    ".readability.yaml"
  ],
  "url": "https://json.schemastore.org/readability.json"
}
```

**Key Fields**:
- `name`: Display name in IDE schema pickers
- `description`: Shows in IDE documentation
- `fileMatch`: Filename patterns that trigger automatic schema loading
- `url`: Where SchemaStore will host the schema (their CDN)

**Important**: Insert alphabetically in the catalog.

### Step 3: Add Schema File

Copy your schema to SchemaStore:

```bash
# Copy from your docs directory
cp /path/to/readability/docs/schemas/config.json \
   src/schemas/json/readability.json
```

**Schema Requirements**:
- Must be valid JSON (not YAML)
- Must validate against JSON Schema meta-schema
- `$id` should reference your canonical URL: `https://readability.adaptive-enforcement-lab.com/schemas/config.json`
- SchemaStore will serve a **copy** at `https://json.schemastore.org/readability.json`

**Important**: Your domain remains the authoritative source. SchemaStore is a mirror for discovery.

### Step 4: Test Locally

SchemaStore has automated tests:

```bash
# Install dependencies
npm install

# Run tests for your schema
npm test -- --schema readability

# Run all tests (recommended)
npm test
```

Tests verify:
- ✅ Schema is valid JSON
- ✅ Schema conforms to JSON Schema spec
- ✅ All `$ref` references resolve
- ✅ No syntax errors
- ✅ Catalog entry is valid

Fix any errors before submitting.

### Step 5: Submit Pull Request

```bash
git add src/api/json/catalog.json src/schemas/json/readability.json
git commit -m "Add schema for readability configuration (.readability.yml)"
git push origin add-readability-schema
```

Create PR at: https://github.com/SchemaStore/schemastore/pulls

**PR Template**:

```markdown
## Description
Adds JSON Schema for [readability](https://github.com/adaptive-enforcement-lab/readability), a markdown documentation analyzer and GitHub Action.

## Details
- **File Match**: `.readability.yml`, `.readability.yaml`
- **Project**: https://github.com/adaptive-enforcement-lab/readability
- **Documentation**: https://readability.adaptive-enforcement-lab.com
- **Schema Spec**: Draft 2020-12
- **Canonical URL**: https://readability.adaptive-enforcement-lab.com/schemas/config.json

## Schema Features
- Comprehensive field descriptions
- Validation for all config properties
- Examples for common values
- Type checking (number, integer, string, etc.)
- Range validation (min/max values)

## Testing
- [x] Schema validates against meta-schema
- [x] All example values are valid
- [x] Tests pass locally (`npm test -- --schema readability`)
- [x] Schema tested in VS Code, IntelliJ, and Vim
- [x] Schema has been used in production for 2+ weeks

## Checklist
- [x] Schema file is valid JSON
- [x] Catalog entry added alphabetically
- [x] `$id` references canonical URL
- [x] All fields have descriptions
- [x] Tests pass
- [x] PR follows contribution guidelines
```

### Step 6: Respond to Review

**Review Timeline**: 1-7 days

**Common Review Feedback**:
- "Add more examples" → Update schema with additional `examples` arrays
- "Improve descriptions" → Make field descriptions more detailed
- "Fix validation rules" → Adjust min/max values based on feedback
- "Update catalog entry" → Refine name/description text

**Response Time**: Respond to feedback within 24-48 hours to keep PR active.

### Step 7: Schema Goes Live

After PR is merged:
- Schema is deployed to SchemaStore CDN
- Available at `https://json.schemastore.org/readability.json`
- Automatic discovery works in all supported IDEs
- Users no longer need explicit `$schema` references

**Propagation**: Changes may take 1-24 hours to propagate to all IDEs.

## Maintaining Schema in SchemaStore

### When to Update

Update schema in SchemaStore when:
- New config fields are added
- Validation rules change
- Field descriptions improve
- Examples are added/updated

**Frequency**: Batch updates quarterly or after major releases.

### Update Process

1. Update schema on your domain first
2. Test updated schema for 1-2 weeks
3. Submit PR to SchemaStore with changes:
   ```bash
   cd schemastore
   git checkout -b update-readability-schema
   cp /path/to/updated/config.json src/schemas/json/readability.json
   git commit -m "Update readability schema: add new fields X, Y"
   git push origin update-readability-schema
   # Create PR
   ```

### Breaking Changes

If making breaking changes to schema:
- Version the schema URL (e.g., `v2/config.json`)
- Submit both v1 and v2 to SchemaStore
- Update `fileMatch` to distinguish versions
- Document migration path

**Avoid breaking changes when possible** - use optional fields instead.

## Viability Assessment

### ✅ HIGH VIABILITY

**Evidence**:
- 1000+ schemas successfully published to SchemaStore
- Active maintainers and responsive community
- Clear contribution guidelines
- Automated testing infrastructure
- No gatekeeping—legitimate projects accepted
- Fast review cycle (1-7 days)

**Success Rate**: 95%+ for well-documented schemas

**Risk Level**: 🟢 LOW

### Factors for Success

**High Success Probability**:
- ✅ Schema is well-documented (comprehensive descriptions)
- ✅ Schema has been tested in production
- ✅ Schema includes examples
- ✅ Tests pass locally
- ✅ PR follows guidelines

**Common Rejection Reasons** (avoid these):
- ❌ Schema is incomplete or poorly documented
- ❌ Schema hasn't been tested
- ❌ Project is inactive/abandoned
- ❌ Schema is for internal/private tool
- ❌ Tests fail

## Rollback Plan

If SchemaStore submission is rejected or delayed:

**Impact**: Minimal
- Users can still use explicit `$schema` reference
- Schema works perfectly from your domain
- IDE support is identical (just requires one extra line)

**Fallback**:
```yaml
# Users add this line (one-time setup)
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/schemas/config.json

thresholds:
  max_grade: 12
```

**No functionality is lost** - SchemaStore just makes it more convenient.

## Timeline

| Milestone | Duration | Blocker |
|-----------|----------|---------|
| Wait for schema stability | 2-4 weeks | Schema must be tested |
| Fork SchemaStore repo | 5 min | None |
| Add catalog entry | 10 min | None |
| Copy schema file | 2 min | Schema finalized |
| Run tests locally | 5 min | None |
| Submit PR | 10 min | Tests pass |
| Community review | 1-7 days | Responsiveness |
| Schema goes live | Immediate | PR merged |
| **Total** | **3-5 weeks** | Schema stability |

## Success Metrics

- ✅ PR submitted within 1 month of schema stabilization
- ✅ PR reviewed within 1 week
- ✅ PR merged within 2 weeks (including any revisions)
- ✅ Schema live on SchemaStore CDN
- ✅ Automatic discovery works in VS Code
- ✅ Users report "it just works" experience
- ✅ No schema bugs reported after SchemaStore deployment

## Documentation Updates

After SchemaStore approval, update docs to mention automatic discovery:

**README.md**:
```markdown
### Configuration

Create `.readability.yml` in your repository root:

```yaml
thresholds:
  max_grade: 12
  min_ease: 40
```

**IDE Support**: Autocomplete and validation work automatically in VS Code, IntelliJ, and other editors. No setup required!
```

**docs/cli/config-file.md**:
```markdown
## IDE Support

### Automatic Schema Detection

The YAML language server automatically detects `.readability.yml` files and loads the schema from [SchemaStore](https://schemastore.org/). No configuration needed!

Open `.readability.yml` and start typing - autocomplete will suggest available fields.
```

## Next Steps

**Now** (immediate):
1. Publish schema on your domain (see [Schema Publishing](schema-publishing.md))
2. Test with users
3. Gather feedback
4. Iterate on schema

**Later** (2-4 weeks from now):
1. Review this document
2. Verify prerequisites are met
3. Submit to SchemaStore following steps above
4. Update documentation after approval

## References

- [SchemaStore.org](https://www.schemastore.org/)
- [SchemaStore GitHub Repository](https://github.com/SchemaStore/schemastore)
- [Contribution Guidelines](https://github.com/SchemaStore/schemastore/blob/master/CONTRIBUTING.md)
- [YAML Language Server](https://github.com/redhat-developer/yaml-language-server)
- [Example Schemas](https://www.schemastore.org/json/) (browse for inspiration)
