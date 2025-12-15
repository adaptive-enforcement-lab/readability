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
cp /path/to/readability/docs/schemas/config.json \
   src/schemas/json/readability.json
```

**Important**: SchemaStore expects:
- Schema must be valid JSON (not YAML)
- Schema must validate against JSON Schema meta-schema
- `$id` should reference your canonical URL: `https://readability.adaptive-enforcement-lab.com/schemas/config.json`
- SchemaStore will serve a copy at `https://json.schemastore.org/readability.json`

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

### Existing Infrastructure: MkDocs Material + GitHub Pages

**Current Setup**:
- **Site URL**: `https://readability.adaptive-enforcement-lab.com`
- **Platform**: MkDocs Material hosted on GitHub Pages
- **Versioning**: Mike (already configured in `mkdocs.yml`)
- **Domain**: Custom domain with HTTPS

**Schema URL**: `https://readability.adaptive-enforcement-lab.com/schemas/config.json`

### Hosting Strategy: Leverage Existing MkDocs Deployment

**Approach**: Add schema file to existing MkDocs site structure

**Pros**:
- ✅ **Zero new infrastructure** - reuses existing GitHub Pages deployment
- ✅ **Professional domain** - already configured and stable
- ✅ **Same deployment pipeline** - schema updates with docs
- ✅ **Versioning support** - can use mike for schema versions if needed
- ✅ **HTTPS by default** - already configured
- ✅ **No additional cost** - uses existing setup

**Cons**:
- None (ideal solution)

**Verdict**: ✅ **OPTIMAL** - Best of all options

### Implementation Steps

#### Step 1: Add Schema to MkDocs Site

```bash
# Create schemas directory in docs
mkdir -p docs/schemas

# Copy schema file
cp schemas/readability-config.schema.json docs/schemas/config.json

# Update $id in schema to match hosted URL
# Edit docs/schemas/config.json:
# "$id": "https://readability.adaptive-enforcement-lab.com/schemas/config.json"
```

#### Step 2: Configure MkDocs to Serve Schema

Update `mkdocs.yml`:

```yaml
# Add to existing mkdocs.yml
nav:
  # ... existing nav items ...
  - Schemas:
      - schemas/config.json  # Makes schema available at /schemas/config.json

# OR use extra_files plugin to serve without navigation entry
plugins:
  - search
  - social
  # Add this if you want schema served but not in navigation:
  - exclude:
      glob:
        - schemas/*  # Exclude from navigation but still serve
```

**Alternative** (if MkDocs plugins don't work well for JSON):

Add schema to `docs/` and configure as static file:

```yaml
# mkdocs.yml
extra_files:
  - schemas/config.json
```

Or simply place in `docs/schemas/config.json` and MkDocs will serve it automatically.

#### Step 3: Update Schema $id

Edit `docs/schemas/config.json`:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://readability.adaptive-enforcement-lab.com/schemas/config.json",
  "title": "Readability Configuration",
  // ... rest of schema
}
```

#### Step 4: Deploy with Existing Pipeline

Schema deploys automatically with next docs deployment:

```bash
# If using mike for versioning
mike deploy --push --update-aliases latest

# Or standard MkDocs deploy
mkdocs gh-deploy
```

#### Step 5: Verify Schema Accessibility

```bash
# After deployment, verify schema is accessible
curl -I https://readability.adaptive-enforcement-lab.com/schemas/config.json

# Should return:
# HTTP/2 200
# content-type: application/json
```

### Automation: Keep Schema in Sync

**Option A**: Single Source (Recommended)

Keep schema in `docs/schemas/config.json` as the canonical source:

```bash
# Project structure
readability/
├── docs/
│   └── schemas/
│       └── config.json  # Canonical source, deployed to site
├── pkg/config/
│   └── validate.go      # Embeds schema from docs/
└── mkdocs.yml
```

**Option B**: Copy During Build

Keep schema in `schemas/` directory, copy to `docs/` during build:

```yaml
# .github/workflows/docs.yml
- name: Copy schema to docs
  run: cp schemas/readability-config.schema.json docs/schemas/config.json

- name: Deploy docs
  run: mkdocs gh-deploy --force
```

### Recommended Approach

Use **Option A** (single source in `docs/schemas/`):
- ✅ Single source of truth
- ✅ No build step needed
- ✅ Simple to maintain
- ✅ Automatically versioned with docs (via mike)

Then **embed schema in Go binary** from docs directory:

```go
// pkg/config/validate.go
import _ "embed"

//go:embed ../../docs/schemas/config.json
var embeddedSchema []byte
```

### Recommended Hosting Strategy

**Phase 1** (Immediate):
```
https://readability.adaptive-enforcement-lab.com/schemas/config.json
```
Add schema to existing MkDocs site - **already have professional domain and infrastructure**.

**Phase 2** (Parallel with Phase 1):
```
https://json.schemastore.org/readability.json
```
Submit to SchemaStore for automatic discovery in IDEs.

**Phase 3** (After SchemaStore Approval):
```
https://json.schemastore.org/readability.json (primary)
https://readability.adaptive-enforcement-lab.com/schemas/config.json (canonical)
```
Users get automatic schema loading from SchemaStore, with your domain as the authoritative source.

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
| Add schema to MkDocs | 10 min | Schema finalized |
| Deploy to existing site | Immediate | Next docs deployment |
| SchemaStore PR submission | 1 day | Schema finalized |
| SchemaStore review | 1-7 days | Community review |
| Total | ~1 week | - |

### Step-by-Step

1. **Day 1-2**: Create and test schema (see [Schema Creation](schema-creation.md))
2. **Day 2**: Add schema to `docs/schemas/config.json` and update `$id`
3. **Day 2**: Deploy docs (schema goes live immediately)
4. **Day 3**: Submit SchemaStore PR
5. **Day 4-5**: Update documentation with schema URLs
6. **Week 2**: Respond to SchemaStore review feedback
7. **Week 2+**: Schema goes live on SchemaStore CDN

### Rollback Plan

If SchemaStore submission is rejected or delayed:
1. Continue using `readability.adaptive-enforcement-lab.com` URL (already live)
2. Users add explicit `$schema` reference
3. Validation still works, just not automatic

**Impact**: Minimal - users have a stable, professional URL to reference

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
| Existing MkDocs Site | 10 min | ⭐⭐⭐⭐⭐ | ❌ No | Free | ✅ **OPTIMAL** |
| jsDelivr CDN | Instant | ⭐⭐⭐ | ❌ No | Free | ⚠️ Backup only |
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
