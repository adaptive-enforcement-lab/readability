# Component 2: Schema Publishing

## Overview

Distribute the JSON Schema file to enable automatic discovery by IDEs and tools. Users should not need to manually configure schema paths—their editors should "just work" when they open `.readability.yml`.

## Technical Approach

Schema publishing involves two strategies:

1. **SchemaStore.org** - Global schema registry for automatic discovery
2. **Direct Web Hosting** - Stable URL for explicit `$schema` references

Both approaches are complementary and should be implemented.

## Strategy 1: SchemaStore.org Submission

### What is SchemaStore?

[SchemaStore.org](https://schemastore.org/) is a centralized registry of JSON schemas maintained by the community. It powers automatic schema detection in:
- VS Code (built-in)
- Visual Studio
- JetBrains IDEs (IntelliJ, WebStorm, etc.)
- Rider
- Eclipse

### How It Works

When you open `.readability.yml` in VS Code, the YAML language server:
1. Extracts the filename pattern `.readability.yml`
2. Queries SchemaStore's catalog
3. Automatically loads the associated schema
4. Provides validation and autocomplete

**No user configuration required** - it's automatic.

### Submission Process

#### Step 1: Fork SchemaStore Repository

```bash
git clone https://github.com/SchemaStore/schemastore.git
cd schemastore
git checkout -b add-readability-schema
```

#### Step 2: Add Schema to Catalog

Edit `src/api/json/catalog.json` and add:

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
- `url`: Where the schema will be hosted (SchemaStore's CDN)

#### Step 3: Add Schema File

Copy your schema to `src/schemas/json/readability.json`:

```bash
cp /path/to/readability/schemas/readability-config.schema.json \
   src/schemas/json/readability.json
```

**Important**: SchemaStore expects:
- Schema must be valid JSON (not YAML)
- Schema must validate against JSON Schema meta-schema
- `$id` should match the CDN URL: `https://json.schemastore.org/readability.json`

#### Step 4: Test Locally

SchemaStore has a test suite to validate schemas:

```bash
npm install
npm test -- --schema readability
```

This verifies:
- Schema is valid JSON
- Schema conforms to JSON Schema spec
- All referenced `$ref` paths exist
- No syntax errors

#### Step 5: Submit Pull Request

```bash
git add src/api/json/catalog.json src/schemas/json/readability.json
git commit -m "Add schema for readability configuration (.readability.yml)"
git push origin add-readability-schema
```

Create PR at https://github.com/SchemaStore/schemastore/pulls

**PR Template**:
```markdown
## Description
Adds JSON Schema for [readability](https://github.com/markcheret/readability), a markdown documentation analyzer.

## Details
- **File Match**: `.readability.yml`, `.readability.yaml`
- **Project**: https://github.com/markcheret/readability
- **Schema Spec**: Draft 2020-12

## Checklist
- [x] Schema validates against meta-schema
- [x] All example values are valid
- [x] Tests pass locally
- [x] Schema includes descriptions for all fields
```

### Review Timeline

**Expected**: 1-7 days for community review
**Factors**:
- Schema quality (good descriptions, examples)
- Comprehensive validation rules
- Test coverage
- Responsiveness to review feedback

### Viability Assessment: ✅ HIGH

**Evidence**:
- 1000+ schemas in SchemaStore (proven process)
- Active maintainers and community
- Clear contribution guidelines
- Automated testing infrastructure
- No gatekeeping—legitimate projects are accepted

**Risk**: Low
- Well-documented process
- Fast review cycle
- Can use fallback (direct hosting) while waiting

## Strategy 2: Direct Web Hosting

### Purpose

Provide a stable URL for:
1. **Explicit `$schema` references** in YAML files
2. **Fallback during SchemaStore review** (can use immediately)
3. **Documentation links** (users can browse schema)
4. **Version pinning** (future: support schema versioning)

### Hosting Options

#### Option A: GitHub Pages (Recommended)

**Approach**: Host schema at `https://markcheret.github.io/readability/schemas/config.json`

**Pros**:
- ✅ Free, reliable hosting
- ✅ Automatic HTTPS
- ✅ Easy to update (git push)
- ✅ Version with Git tags
- ✅ No infrastructure maintenance

**Cons**:
- ⚠️ URL tied to GitHub username (not ideal for long-term projects)

**Setup**:

```bash
# Enable GitHub Pages in repo settings
# Settings → Pages → Source: gh-pages branch

# Create gh-pages branch
git checkout --orphan gh-pages
git rm -rf .
mkdir -p schemas
cp /path/to/readability-config.schema.json schemas/config.json

# Update $id in schema
sed -i 's|readability.dev/schemas|markcheret.github.io/readability/schemas|g' schemas/config.json

git add schemas/
git commit -m "Publish JSON Schema to GitHub Pages"
git push origin gh-pages
```

**Automation** (publish on release):

```yaml
# .github/workflows/publish-schema.yml
name: Publish Schema
on:
  release:
    types: [published]

jobs:
  publish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Deploy schema to GitHub Pages
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./schemas
          destination_dir: schemas
```

#### Option B: Custom Domain (readability.dev)

**Approach**: Host schema at `https://readability.dev/schemas/config.json`

**Pros**:
- ✅ Professional, stable URL
- ✅ Project branding
- ✅ Long-term stability (URL won't change if repo moves)

**Cons**:
- ⚠️ Requires domain registration (~$12/year)
- ⚠️ Needs DNS configuration
- ⚠️ Hosting setup (can still use GitHub Pages + custom domain)

**Viability**: ⚠️ **MEDIUM** (cost/effort trade-off)

**Recommendation**: Start with GitHub Pages (Option A), upgrade to custom domain if project gains traction.

#### Option C: CDN (jsDelivr, unpkg)

**Approach**: Use free CDN to serve schema from GitHub releases

**URL**: `https://cdn.jsdelivr.net/gh/markcheret/readability@latest/schemas/readability-config.schema.json`

**Pros**:
- ✅ Free, fast global CDN
- ✅ No setup required
- ✅ Version pinning via Git tags

**Cons**:
- ❌ Long, ugly URLs
- ❌ Less professional
- ❌ Third-party dependency

**Viability**: ✅ **HIGH** (good fallback option)

**Use Case**: Temporary hosting while waiting for SchemaStore approval

#### Option D: Raw GitHub

**Approach**: Link directly to `raw.githubusercontent.com`

**URL**: `https://raw.githubusercontent.com/markcheret/readability/main/schemas/readability-config.schema.json`

**Pros**:
- ✅ Zero setup
- ✅ Always up-to-date with main branch

**Cons**:
- ❌ Not recommended for production use (caching issues)
- ❌ Content-Type headers may not be optimal
- ❌ Rate limiting

**Viability**: ⚠️ **LOW** (development/testing only)

**Verdict**: **Avoid for production**

### Recommended Hosting Strategy

**Phase 1** (Immediate):
```
https://cdn.jsdelivr.net/gh/markcheret/readability@latest/schemas/config.json
```
Use jsDelivr for instant availability while preparing SchemaStore submission.

**Phase 2** (Post-Release):
```
https://markcheret.github.io/readability/schemas/config.json
```
Set up GitHub Pages for cleaner, official-looking URL.

**Phase 3** (After SchemaStore Approval):
```
https://json.schemastore.org/readability.json
```
Users get automatic schema loading—no explicit URL needed.

**Future** (If project grows):
```
https://readability.dev/schemas/config.json
```
Custom domain for long-term stability.

## Schema Versioning Strategy

### Problem

Should we version the schema? What happens when we add new config fields?

### Approach: Unversioned URLs

**Recommendation**: Use a single, unversioned URL that always points to the latest schema.

**Rationale**:
- Config files are **forward-compatible** (old tools ignore new fields)
- Config files are **backward-compatible** (new fields are optional)
- Users expect latest validation rules
- Simplifies documentation (one URL to remember)

**Example**:
```yaml
# yaml-language-server: $schema=https://json.schemastore.org/readability.json
```

This URL always serves the latest schema, even as new fields are added.

### When to Version

Consider versioned schemas only if:
- **Breaking changes** to config structure (unlikely)
- **Different tools** need different schemas (e.g., v1.x vs v2.x)
- **Pinned validation** required for reproducible builds

For readability, versioning is **not needed** in Phase 1-3.

## Implementation Guidance

### Timeline

| Milestone | Duration | Blocker |
|-----------|----------|---------|
| Create schema file | 1 day | None |
| jsDelivr hosting | Immediate | Merge to main + release |
| GitHub Pages setup | 1 day | Repo settings |
| SchemaStore PR submission | 1 day | Schema finalized |
| SchemaStore review | 1-7 days | Community review |
| Total | ~2 weeks | - |

### Step-by-Step

1. **Week 1, Day 1-2**: Create and test schema (see [Schema Creation](schema-creation.md))
2. **Week 1, Day 3**: Submit SchemaStore PR
3. **Week 1, Day 4**: Set up GitHub Pages hosting
4. **Week 1, Day 5**: Update documentation with schema URLs
5. **Week 2**: Respond to SchemaStore review feedback
6. **Week 2+**: Schema goes live on SchemaStore CDN

### Rollback Plan

If SchemaStore submission is rejected or delayed:
1. Continue using GitHub Pages URL
2. Users add explicit `$schema` reference
3. Validation still works, just not automatic

**Impact**: Minor inconvenience, not a blocker

## Testing Requirements

### Pre-Submission Tests

1. **Schema Validation**: Schema validates against JSON Schema meta-schema
2. **Example Configs**: All documented examples pass schema validation
3. **IDE Integration**: Schema loads correctly in VS Code, IntelliJ
4. **CORS Headers**: Hosted schema allows cross-origin requests (for web-based tools)

### Post-Deployment Tests

1. **URL Accessibility**: Schema URL returns 200 OK, correct Content-Type
2. **Caching**: Schema updates propagate within reasonable time
3. **HTTPS**: Schema served over HTTPS (required by some tools)
4. **Stability**: URL doesn't change, no 404s

### SchemaStore-Specific Tests

SchemaStore runs automated tests on all submissions:
```bash
# Run locally before submitting
npm install
npm test -- --schema readability
```

Tests verify:
- JSON syntax
- JSON Schema spec compliance
- No broken `$ref` references
- Valid `fileMatch` patterns

## Success Metrics

- ✅ Schema accessible at stable URL within 1 week
- ✅ SchemaStore PR submitted within 2 weeks of schema creation
- ✅ Schema appears in SchemaStore catalog within 1 month
- ✅ Users report automatic validation in VS Code (no manual setup)
- ✅ Schema URL documented in README and configuration guide

## Alternatives Comparison

| Approach | Setup Time | URL Quality | Auto-Discovery | Cost | Verdict |
|----------|-----------|-------------|----------------|------|---------|
| SchemaStore | 1 day | ⭐⭐⭐⭐⭐ | ✅ Yes | Free | ✅ Primary |
| GitHub Pages | 1 day | ⭐⭐⭐⭐ | ❌ No | Free | ✅ Fallback |
| jsDelivr CDN | Instant | ⭐⭐⭐ | ❌ No | Free | ✅ Temporary |
| Custom Domain | 2 days | ⭐⭐⭐⭐⭐ | ❌ No | $12/year | ⚠️ Future |
| Raw GitHub | Instant | ⭐ | ❌ No | Free | ❌ Dev only |

## Next Steps

After schema publishing:
1. Update [YAML Integration](yaml-integration.md) with published schema URL
2. Add schema URL to documentation
3. Test automatic discovery in multiple IDEs
4. Monitor SchemaStore PR for review feedback

## References

- [SchemaStore.org](https://www.schemastore.org/)
- [SchemaStore Contribution Guide](https://github.com/SchemaStore/schemastore/blob/master/CONTRIBUTING.md)
- [GitHub Pages Documentation](https://docs.github.com/en/pages)
- [jsDelivr CDN](https://www.jsdelivr.com/)
