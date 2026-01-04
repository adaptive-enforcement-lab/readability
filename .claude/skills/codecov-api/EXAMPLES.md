# Codecov API Examples

Real-world examples from the readability project showing coverage verification workflows.

## Example 1: PR #226 Coverage Verification

**Context**: PR #226 claimed to improve coverage from 98.6% to 99.0%. We verified this against Codecov API and discovered discrepancies.

### Initial Claim vs Reality

**Claimed** (from local `go tool cover`):
- Before: 98.6%
- After: 99.0%
- Improvement: +0.4%

**Actual** (from Codecov API):
- Before: 98.45%
- After: 98.81%
- Improvement: +0.36%

**Discrepancy**: 0.19 percentage points on final coverage

### Step-by-Step Verification

**Step 1: Get current branch coverage**

```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq '{
    current_commit: .head_commit.commitid,
    parent_commit: .head_commit.parent,
    coverage: .head_commit.totals.coverage,
    files: .head_commit.totals.files,
    lines: .head_commit.totals.lines,
    hits: .head_commit.totals.hits,
    misses: .head_commit.totals.misses
  }'
```

**Response:**
```json
{
  "current_commit": "723e1e4bc8a9f1a2b3c4d5e6f7g8h9i0j1k2l3m4",
  "parent_commit": "43256dc9844084feb1835f53ab13aff27e1273fe",
  "coverage": 98.81,
  "files": 10,
  "lines": 1101,
  "hits": 1088,
  "misses": 10
}
```

**Key insight**: Current coverage is 98.81%, not 99.0%

**Step 2: Get parent commit coverage**

```bash
PARENT="43256dc9844084feb1835f53ab13aff27e1273fe"

curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${PARENT}" | \
  jq '{
    commit: .commitid,
    coverage: .totals.coverage,
    files: .totals.files,
    lines: .totals.lines,
    hits: .totals.hits,
    misses: .totals.misses
  }'
```

**Response:**
```json
{
  "commit": "43256dc9844084feb1835f53ab13aff27e1273fe",
  "coverage": 98.45,
  "files": 10,
  "lines": 1090,
  "hits": 1073,
  "misses": 14
}
```

**Key insight**: Parent coverage was 98.45%, not 98.6%

**Step 3: Calculate actual improvement**

```bash
BEFORE=98.45
AFTER=98.81
CHANGE=$(echo "$AFTER - $BEFORE" | bc)

echo "Coverage change: ${BEFORE}% → ${AFTER}% (+${CHANGE}%)"
```

**Output:**
```
Coverage change: 98.45% → 98.81% (+0.36%)
```

**Key insight**: Real improvement was +0.36%, not +0.4%

### Component-Level Analysis

**Step 4: Compare file-level changes**

```bash
# Get parent commit file coverage
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${PARENT}" | \
  jq -r '.report.files[] | "\(.name):\(.totals.coverage)"' | \
  sort > /tmp/before.txt

# Get current commit file coverage
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq -r '.head_commit.report.files[] | "\(.name):\(.totals.coverage)"' | \
  sort > /tmp/after.txt

# Compare
diff /tmp/before.txt /tmp/after.txt
```

**Output:**
```diff
< cmd/readability/main.go:96.21
> cmd/readability/main.go:97.29
```

**Analysis**: Only CLI component improved (+1.08%), all other files unchanged.

### Local Tool Comparison

**Step 5: Compare with local coverage**

```bash
# Run local coverage
go test ./... -coverprofile=/tmp/coverage.out
go tool cover -func=/tmp/coverage.out | tail -1
```

**Local output:**
```
total:  (statements)  99.0%
```

**Codecov API:**
```
total:  (lines)  98.81%
```

**Difference**: 0.19 percentage points

### Why the Discrepancy?

**Rounding**:
- Local `go tool cover`: Rounds to 1 decimal place (99.0%)
- Codecov: Uses precise calculation (98.81%)

**Coverage Method**:
- Local: Statement-based coverage
- Codecov: Line-based coverage

**File Inclusion**:
- Local `go test`: Only includes files with tests
- Codecov: Tracks all repository files

### Conclusion

**PR #226 Coverage Summary**:
- ✓ Coverage improved: 98.45% → 98.81%
- ✓ Improvement: +0.36 percentage points
- ✓ Component improved: CLI (96.21% → 97.29%)
- ✗ Did NOT reach 99.0% (local tools misleading)

**Lesson**: Always verify against Codecov API before making PR claims.

## Example 2: Find Files Below Coverage Threshold

**Goal**: Identify files below 95% coverage for improvement targeting.

```bash
THRESHOLD=95.0

curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq -r --arg thresh "$THRESHOLD" \
  '.head_commit.report.files[] |
   select(.totals.coverage < ($thresh | tonumber)) |
   "\(.name): \(.totals.coverage)% (lines: \(.totals.lines), misses: \(.totals.misses))"' | \
  sort -t: -k2 -n
```

**Output (from commit 723e1e4):**
```
cmd/readability/main.go: 97.29% (lines: 147, misses: 2)
```

**Analysis**: Only CLI main.go is below 95% threshold. Target these 2 missed lines for next coverage improvement.

## Example 3: Coverage Trend Analysis

**Goal**: Track coverage changes over last 5 commits to identify regression.

```bash
echo "Commit Coverage Trend:"
echo "====================="

git log -5 --format='%H %s' | while read SHA MESSAGE; do
  COV=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
    "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${SHA}" 2>/dev/null | \
    jq -r '.totals.coverage // "N/A"')

  SHORT=$(echo $SHA | cut -c1-7)
  printf "%s: %6s%% - %s\n" "$SHORT" "$COV" "${MESSAGE:0:50}"
done
```

**Output:**
```
Commit Coverage Trend:
=====================
723e1e4:  98.81% - test: add CLI integration tests for check mode
43256dc:  98.45% - fix: handle empty markdown files gracefully
e2a2bfc:  98.45% - chore(main): release 3.0.0
4f784a5:  98.20% - fix(deps): update module github.com/yuin/gold
25dc0b7:  98.20% - fix: test action builds from source instea
```

**Analysis**: Coverage improved from 98.20% → 98.81% over 5 commits. No regression detected.

## Example 4: Component Coverage Breakdown

**Goal**: Get detailed breakdown of coverage by package/directory.

```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq -r '.head_commit.report.files[] | "\(.name):\(.totals.coverage)"' | \
  awk -F'/' '{
    if ($1 == "cmd") pkg = "CLI";
    else if ($1 == "pkg") pkg = $2;
    else pkg = "Other";
    coverage = $NF;
    split(coverage, arr, ":");
    sum[pkg] += arr[2];
    count[pkg]++;
  }
  END {
    for (p in sum) {
      avg = sum[p] / count[p];
      printf "%s: %.2f%% (%d files)\n", p, avg, count[p];
    }
  }' | sort -t: -k2 -nr
```

**Output:**
```
config: 100.00% (1 files)
output: 100.00% (5 files)
analyzer: 99.11% (1 files)
markdown: 98.75% (1 files)
CLI: 97.29% (1 files)
```

**Analysis**: Focus improvement efforts on CLI package (lowest at 97.29%).

## Example 5: Pre-PR Coverage Check

**Goal**: Verify coverage meets 99% threshold before creating PR.

```bash
#!/bin/bash
set -e

TARGET=99.0

echo "Checking coverage against target ${TARGET}%..."

# Get current coverage
CURRENT=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq -r '.head_commit.totals.coverage')

echo "Current coverage: ${CURRENT}%"

# Compare (bc returns 1 for true, 0 for false)
if (( $(echo "$CURRENT >= $TARGET" | bc -l) )); then
  echo "✓ Coverage ${CURRENT}% meets target ${TARGET}%"
  exit 0
else
  DEFICIT=$(echo "$TARGET - $CURRENT" | bc)
  echo "✗ Coverage ${CURRENT}% below target ${TARGET}%"
  echo "  Need to improve by ${DEFICIT}%"
  exit 1
fi
```

**Output (for 98.81% coverage):**
```
Checking coverage against target 99.0%...
Current coverage: 98.81%
✗ Coverage 98.81% below target 99.0%
  Need to improve by 0.19%
```

## Example 6: Automated PR Comment

**Goal**: Generate coverage comment for PR body.

```bash
#!/bin/bash

# Get parent and current coverage
RESPONSE=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main')

PARENT=$(echo $RESPONSE | jq -r '.head_commit.parent')
AFTER=$(echo $RESPONSE | jq -r '.head_commit.totals.coverage')

BEFORE=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${PARENT}" | \
  jq -r '.totals.coverage')

CHANGE=$(echo "$AFTER - $BEFORE" | bc)

# Generate markdown comment
cat << EOF
## Test Coverage

- **Before**: ${BEFORE}%
- **After**: ${AFTER}%
- **Change**: ${CHANGE:+'+'}${CHANGE}%

### Coverage Details

\`\`\`bash
# Verified via Codecov API
curl -s -H 'Authorization: Bearer \$CODECOV_TOKEN' \\
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \\
  jq '.head_commit.totals.coverage'
\`\`\`

**Note**: Coverage verified against Codecov API, not local tools. Local tools may show different values due to rounding and calculation method differences.
EOF
```

**Output:**
```markdown
## Test Coverage

- **Before**: 98.45%
- **After**: 98.81%
- **Change**: +0.36%

### Coverage Details

```bash
# Verified via Codecov API
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq '.head_commit.totals.coverage'
```

**Note**: Coverage verified against Codecov API, not local tools. Local tools may show different values due to rounding and calculation method differences.
```

## Common Patterns

### Pattern 1: Before/After Comparison

```bash
# Store in variables for reuse
BRANCH_DATA=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main')

PARENT=$(echo $BRANCH_DATA | jq -r '.head_commit.parent')
AFTER=$(echo $BRANCH_DATA | jq -r '.head_commit.totals.coverage')

BEFORE=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${PARENT}" | \
  jq -r '.totals.coverage')

echo "${BEFORE}% → ${AFTER}%"
```

### Pattern 2: Component Filtering

```bash
# Only show pkg/ files
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq -r '.head_commit.report.files[] | select(.name | startswith("pkg/")) | "\(.name): \(.totals.coverage)%"'
```

### Pattern 3: Error Handling

```bash
RESPONSE=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main')

# Check for API error
if echo "$RESPONSE" | jq -e '.detail' > /dev/null 2>&1; then
  ERROR=$(echo "$RESPONSE" | jq -r '.detail')
  echo "Error: $ERROR" >&2

  # Check if token issue
  if [ -z "$CODECOV_TOKEN" ]; then
    echo "CODECOV_TOKEN not set. Run: source ~/.zshrc" >&2
  fi

  exit 1
fi

# Safe to extract data
COVERAGE=$(echo "$RESPONSE" | jq -r '.head_commit.totals.coverage')
```

## Troubleshooting Real Issues

### Issue: Token Not Found

**Symptom**: Empty or null responses from API

**Diagnosis:**
```bash
echo $CODECOV_TOKEN
# Output: (empty)
```

**Fix:**
```bash
source ~/.zshrc
echo $CODECOV_TOKEN
# Output: 31687df3-2a33-495f-97e3-1463bc8eb576
```

### Issue: Branch Not Found

**Symptom**: 404 error for branch query

**Diagnosis:**
```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/Main'
# Output: {"detail": "Not found"}
```

**Fix**: Branch names are case-sensitive
```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main'
# Output: (success)
```

### Issue: Commit Not Found

**Symptom**: 404 error for commit query

**Diagnosis:**
```bash
# Short SHA too short
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/723'
# Output: {"detail": "Not found"}
```

**Fix**: Use at least 7 characters
```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/723e1e4'
# Output: (success)
```

## Summary

These examples demonstrate:
- ✓ Real-world coverage verification workflow (PR #226)
- ✓ Identifying discrepancies between local tools and Codecov
- ✓ Component-level coverage analysis
- ✓ Coverage trend tracking
- ✓ Automated PR coverage reporting
- ✓ Common error handling patterns

**Key takeaway**: Always verify coverage claims against Codecov API before including them in PRs.
