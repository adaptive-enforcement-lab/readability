# Component 2: Schema Publishing (MkDocs Hosting)

## Overview

Publish the JSON Schema to your existing MkDocs Material documentation site, leveraging the infrastructure already in place at `readability.adaptive-enforcement-lab.com`. This provides immediate availability with zero new infrastructure setup.

## Why This Approach

**Existing Infrastructure**:
- ✅ MkDocs Material site already configured
- ✅ Custom domain with HTTPS
- ✅ GitHub Pages deployment pipeline
- ✅ Mike versioning support
- ✅ Automatic CDN and caching

**Benefits**:
- **Zero setup time** - reuse existing deployment
- **Professional URL** - branded domain already configured
- **Same workflow** - schema deploys with docs
- **Version control** - can use mike for schema versions
- **Immediate availability** - no waiting for external approvals

## Current Setup Analysis

**Site Configuration** (from `mkdocs.yml`):
- **Site URL**: `https://readability.adaptive-enforcement-lab.com`
- **Platform**: MkDocs Material
- **Versioning**: Mike (configured)
- **Deployment**: GitHub Pages with custom domain
- **HTTPS**: Enabled by default

**Schema URL**: `https://readability.adaptive-enforcement-lab.com/schemas/config.json`

## Implementation Steps

### Step 1: Create Schema Directory

```bash
# Create schemas directory in docs
mkdir -p docs/schemas
```

**Location**: `docs/schemas/` (alongside other documentation)

**Rationale**:
- MkDocs automatically serves files from `docs/`
- Keeps schema versioned with documentation
- Single source of truth for both docs and runtime validation

### Step 2: Add Schema File

```bash
# When schema is ready, add to docs
cp schemas/readability-config.schema.json docs/schemas/config.json
```

**File**: `docs/schemas/config.json`

**Important**: Update `$id` in schema to match hosted URL:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://readability.adaptive-enforcement-lab.com/schemas/config.json",
  "title": "Readability Configuration",
  "description": "Configuration schema for readability markdown analyzer",
  // ... rest of schema
}
```

### Step 3: Configure MkDocs (Optional)

MkDocs will automatically serve `docs/schemas/config.json`, but you can explicitly configure if needed:

**Option A**: No configuration needed
- MkDocs serves all files in `docs/` by default
- `docs/schemas/config.json` → `https://site.com/schemas/config.json`

**Option B**: Exclude from navigation (if you don't want it in nav)

```yaml
# mkdocs.yml
plugins:
  - search
  - social
  # No need to add anything - JSON files served but not in nav by default
```

**Option C**: Add to nav (for discoverability)

```yaml
# mkdocs.yml
nav:
  # ... existing nav ...
  - Developer Resources:
      - JSON Schema: schemas/config.json
```

**Recommendation**: **Option A** - Let MkDocs serve automatically, no config needed.

### Step 4: Embed Schema in Go Binary

Go code should reference the schema from the docs directory:

```go
// pkg/config/validate.go
package config

import _ "embed"

//go:embed ../../docs/schemas/config.json
var embeddedSchema []byte
```

**Why embed from docs/**:
- Single source of truth
- Schema in docs is deployed to web
- Go binary embeds same file for runtime validation
- No build step to sync separate files

### Step 5: Deploy

Schema deploys automatically with next documentation deployment:

```bash
# Using mike (versioning)
mike deploy --push --update-aliases latest

# OR standard MkDocs deploy
mkdocs gh-deploy

# OR let GitHub Actions deploy automatically on merge to main
```

**Deployment triggers**:
- Manual: `mike deploy` or `mkdocs gh-deploy`
- Automatic: GitHub Actions on push to main
- Preview: Local `mkdocs serve` for testing

### Step 6: Verify Deployment

After deployment, verify schema is accessible:

```bash
# Check schema is accessible
curl -I https://readability.adaptive-enforcement-lab.com/schemas/config.json

# Expected response:
# HTTP/2 200
# content-type: application/json
# content-length: ...
```

Test schema in IDE:

```yaml
# .readability.yml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/schemas/config.json

thresholds:
  max_grade: 12
```

Open in VS Code - autocomplete should work.

## File Structure

**Recommended Layout**:

```
readability/
├── docs/
│   ├── index.md
│   ├── cli/
│   ├── metrics/
│   └── schemas/
│       └── config.json          # Canonical source, deployed to web
├── pkg/
│   └── config/
│       └── validate.go          # Embeds from docs/schemas/config.json
└── mkdocs.yml
```

**Why this structure**:
- ✅ Single source of truth (`docs/schemas/config.json`)
- ✅ Automatically deployed with docs
- ✅ Versioned with mike (if using versioning)
- ✅ Go embeds same file for runtime validation
- ✅ No build steps to keep files in sync

## Schema Versioning Strategy

### Current Approach: Unversioned

Use single, stable URL that always points to latest schema:

```
https://readability.adaptive-enforcement-lab.com/schemas/config.json
```

**Rationale**:
- Config files are forward-compatible (old tools ignore new fields)
- Config files are backward-compatible (new fields are optional)
- Users expect latest validation rules
- Simplifies documentation

### Future: Version with Mike (Optional)

If versioning becomes necessary:

```bash
# Deploy schema with specific version
mike deploy v1.0 latest
mike set-default latest

# Schema available at both:
# https://site.com/schemas/config.json (latest)
# https://site.com/v1.0/schemas/config.json (pinned)
```

**When to version**:
- Breaking changes to schema structure (unlikely)
- Major version bumps of readability CLI
- Different tools need different schemas

**Current recommendation**: Start unversioned, add versioning only if needed.

## Testing

### Pre-Deployment Tests

```bash
# 1. Validate schema file itself
npm install -g ajv-cli
ajv compile -s docs/schemas/config.json

# 2. Test schema locally
# Start local MkDocs server
mkdocs serve

# In another terminal, test schema endpoint
curl http://localhost:8000/schemas/config.json

# 3. Validate example configs against schema
readability --config .readability.yml --validate-config
```

### Post-Deployment Tests

```bash
# 1. Verify schema is accessible
curl -I https://readability.adaptive-enforcement-lab.com/schemas/config.json

# 2. Verify correct Content-Type
curl -s -o /dev/null -w "%{content_type}\n" \
  https://readability.adaptive-enforcement-lab.com/schemas/config.json
# Should output: application/json

# 3. Verify schema is valid JSON
curl -s https://readability.adaptive-enforcement-lab.com/schemas/config.json | jq .
```

### IDE Integration Tests

**Manual Testing**:

1. Create test `.readability.yml`:
   ```yaml
   # yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/schemas/config.json

   thresholds:
     max_grade:
   ```

2. Open in VS Code
3. Trigger autocomplete after `max_grade:` (Ctrl+Space)
4. Verify autocomplete suggests valid values
5. Add invalid value: `max_grade: "invalid"`
6. Verify red squiggle appears with error message

## Maintenance

### Updating the Schema

When schema needs updates:

1. **Edit**: Update `docs/schemas/config.json`
2. **Test locally**: `mkdocs serve` and verify changes
3. **Validate**: `ajv compile -s docs/schemas/config.json`
4. **Commit**: Commit schema changes
5. **Deploy**: Push to main, GitHub Actions deploys automatically
6. **Verify**: Check live URL after deployment

**Timeline**: Schema updates are live within minutes of merge to main.

### Schema Update Checklist

- [ ] Edit `docs/schemas/config.json`
- [ ] Update `$id` if URL changes
- [ ] Validate with ajv-cli
- [ ] Test locally with `mkdocs serve`
- [ ] Update Go code if schema structure changes
- [ ] Update documentation examples
- [ ] Commit and push
- [ ] Verify deployment

## Rollback

If schema deployment has issues:

```bash
# Revert to previous version (if using mike)
mike deploy v1.0 latest --force

# OR revert commit and redeploy
git revert <commit-hash>
git push
# GitHub Actions redeploys automatically
```

**Impact**: Minimal - schema rollback is same as docs rollback.

## Success Metrics

- ✅ Schema accessible at canonical URL within 5 minutes of deployment
- ✅ HTTPS enabled (no certificate warnings)
- ✅ Correct Content-Type: `application/json`
- ✅ Schema validates with ajv-cli
- ✅ IDE autocomplete works with hosted schema
- ✅ No 404s or 500s
- ✅ Schema updates deploy within minutes

## Timeline

| Milestone | Duration | Blocker |
|-----------|----------|---------|
| Create `docs/schemas/` directory | 1 min | None |
| Add schema file | 2 min | Schema finalized |
| Update schema `$id` | 1 min | None |
| Deploy to site | Immediate | Next docs deployment |
| Verify accessibility | 2 min | Deployment complete |
| **Total** | **~10 minutes** | - |

## Next Steps

After schema is published on your site:

1. Update [YAML Integration](yaml-integration.md) examples to use canonical URL
2. Document schema URL in user-facing documentation
3. Test IDE integration with live schema
4. (Later) Consider [SchemaStore Submission](schemastore-submission.md) for automatic discovery

## Advantages Over Alternatives

| Approach | Setup Time | URL Quality | Maintenance | Verdict |
|----------|-----------|-------------|-------------|---------|
| **Existing MkDocs** | 10 min | ⭐⭐⭐⭐⭐ | Same as docs | ✅ **OPTIMAL** |
| New GitHub Pages | 1 day | ⭐⭐⭐⭐ | Separate pipeline | ⚠️ Unnecessary |
| jsDelivr CDN | Instant | ⭐⭐⭐ | Manual uploads | ⚠️ Backup only |
| Raw GitHub | Instant | ⭐ | Auto (but not suitable) | ❌ Not recommended |

## References

- [MkDocs Material Documentation](https://squidfunk.github.io/mkdocs-material/)
- [Mike Versioning](https://github.com/jimporter/mike)
- [JSON Schema Specification](https://json-schema.org/)
- [Your Site](https://readability.adaptive-enforcement-lab.com)
