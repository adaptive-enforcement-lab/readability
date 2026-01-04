# Codecov API Reference

Comprehensive reference for Codecov API v2 endpoints used for coverage verification.

## Base URL

```
https://api.codecov.io/api/v2/github/{owner}/repos/{repository}
```

For this repository:
```
https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability
```

## Authentication

All requests require Bearer token authentication:

```bash
curl -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/...'
```

**Security**: Never hardcode tokens in scripts or commit them to repositories. Use environment variables.

## Rate Limiting

- **Rate limit**: 100 requests per minute per token
- **Headers**: Check `X-RateLimit-Remaining` and `X-RateLimit-Reset` in responses
- **Best practice**: Cache responses when possible, especially for historical commits

## Endpoints

### GET /branches/{branch}

Get coverage information for a specific branch.

**Parameters:**
- `branch` (path, required): Branch name (e.g., `main`, `develop`)

**Example Request:**
```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main'
```

**Response Schema:**
```json
{
  "name": "main",
  "head_commit": {
    "commitid": "723e1e4bc8a9f...",
    "parent": "43256dc9844084...",
    "totals": {
      "coverage": 98.81,
      "files": 10,
      "lines": 1101,
      "hits": 1088,
      "misses": 10,
      "partials": 3,
      "branches": 0,
      "methods": 0,
      "messages": 0,
      "sessions": 1,
      "complexity": 0,
      "complexity_total": 0,
      "diff": 0
    },
    "report": {
      "files": [
        {
          "name": "cmd/readability/main.go",
          "totals": {
            "coverage": 97.29,
            "lines": 147,
            "hits": 143,
            "misses": 2,
            "partials": 2
          }
        }
      ]
    }
  }
}
```

**Key Fields:**
- `head_commit.totals.coverage`: Overall branch coverage percentage
- `head_commit.parent`: SHA of parent commit (useful for before/after comparison)
- `head_commit.report.files`: Array of file-level coverage details

### GET /commits/{sha}

Get coverage information for a specific commit.

**Parameters:**
- `sha` (path, required): Full commit SHA (40 chars) or short SHA (7+ chars)

**Example Request:**
```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/43256dc9844084feb1835f53ab13aff27e1273fe'
```

**Response Schema:**
```json
{
  "commitid": "43256dc9844084...",
  "totals": {
    "coverage": 98.45,
    "files": 10,
    "lines": 1090,
    "hits": 1073,
    "misses": 14,
    "partials": 3
  },
  "report": {
    "files": [
      {
        "name": "pkg/analyzer/analyzer.go",
        "totals": {
          "coverage": 99.11,
          "lines": 112,
          "hits": 111,
          "misses": 1
        }
      }
    ]
  }
}
```

**Key Fields:**
- `totals.coverage`: Overall commit coverage percentage
- `report.files`: Array of file-level coverage details

**SHA Requirements:**
- Full SHA: 40 hexadecimal characters
- Short SHA: Minimum 7 characters (must be unique in repository)
- Invalid: Fewer than 7 characters will return 404

### GET /components/{component_id}

Get coverage for a specific component (custom Codecov grouping).

**Note**: This endpoint requires component configuration in Codecov settings. Most repositories won't use this.

**Parameters:**
- `component_id` (path, required): Component identifier

**Example Request:**
```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/components/cli'
```

## Response Fields Reference

### Coverage Totals

| Field | Type | Description |
|-------|------|-------------|
| `coverage` | float | Overall coverage percentage (0-100) |
| `files` | integer | Number of files tracked |
| `lines` | integer | Total lines of code |
| `hits` | integer | Lines covered by tests |
| `misses` | integer | Lines not covered by tests |
| `partials` | integer | Lines partially covered (e.g., multi-condition statements) |
| `branches` | integer | Branch coverage count |
| `methods` | integer | Method coverage count |
| `sessions` | integer | Number of test sessions (CI runs) |
| `complexity` | float | Cyclomatic complexity score |
| `diff` | float | Coverage change from parent commit |

### File Totals

Each file in `report.files[]` contains:

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | File path relative to repository root |
| `totals.coverage` | float | File coverage percentage |
| `totals.lines` | integer | Total lines in file |
| `totals.hits` | integer | Lines covered in file |
| `totals.misses` | integer | Lines not covered in file |
| `totals.partials` | integer | Partially covered lines in file |

## Error Codes

### 401 Unauthorized

**Cause**: Invalid or missing authentication token

**Response:**
```json
{
  "detail": "Invalid token"
}
```

**Fix:**
```bash
# Verify token is set
echo $CODECOV_TOKEN

# Source from shell config if empty
source ~/.zshrc
```

### 404 Not Found

**Causes:**
1. Invalid branch name (case-sensitive)
2. Invalid commit SHA (too short or doesn't exist)
3. Repository not found or no access

**Response:**
```json
{
  "detail": "Not found"
}
```

**Debug:**
```bash
# List available branches
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches'

# Verify commit exists
git log --oneline | grep <sha>
```

### 429 Too Many Requests

**Cause**: Rate limit exceeded (>100 requests/minute)

**Response:**
```json
{
  "detail": "Rate limit exceeded"
}
```

**Fix:**
```bash
# Check rate limit headers
curl -I -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main'

# Wait for reset time (X-RateLimit-Reset header)
```

## Query Patterns

### Get Coverage Improvement

Compare parent commit to current commit:

```bash
# Get current branch info with parent
RESPONSE=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main')

# Extract parent SHA
PARENT=$(echo $RESPONSE | jq -r '.head_commit.parent')

# Extract current coverage
AFTER=$(echo $RESPONSE | jq -r '.head_commit.totals.coverage')

# Get parent coverage
BEFORE=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${PARENT}" | \
  jq -r '.totals.coverage')

# Calculate change
CHANGE=$(echo "$AFTER - $BEFORE" | bc)
echo "Coverage: ${BEFORE}% → ${AFTER}% (${CHANGE:+'+'}${CHANGE}%)"
```

### Get Component Breakdown

Extract file-level coverage sorted by coverage percentage:

```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq -r '.head_commit.report.files[] | "\(.totals.coverage)%\t\(.name)"' | \
  sort -n
```

### Find Files Below Threshold

Filter files below specific coverage threshold:

```bash
THRESHOLD=95.0

curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq -r --arg thresh "$THRESHOLD" \
  '.head_commit.report.files[] |
   select(.totals.coverage < ($thresh | tonumber)) |
   "\(.name): \(.totals.coverage)%"'
```

### Track Coverage Over Time

Get coverage for multiple commits:

```bash
# Get last 5 commits
COMMITS=$(git log -5 --format='%H')

for SHA in $COMMITS; do
  COV=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
    "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${SHA}" | \
    jq -r '.totals.coverage')

  SHORT=$(echo $SHA | cut -c1-7)
  echo "${SHORT}: ${COV}%"
done
```

## Best Practices

### Caching Responses

Commit coverage doesn't change, so cache historical queries:

```bash
# Cache parent commit coverage
CACHE_FILE="/tmp/codecov_${PARENT}.json"

if [ -f "$CACHE_FILE" ]; then
  BEFORE=$(cat "$CACHE_FILE" | jq -r '.totals.coverage')
else
  RESPONSE=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
    "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${PARENT}")
  echo "$RESPONSE" > "$CACHE_FILE"
  BEFORE=$(echo "$RESPONSE" | jq -r '.totals.coverage')
fi
```

### Error Handling

Always check for API errors:

```bash
RESPONSE=$(curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main')

# Check for error response
if echo "$RESPONSE" | jq -e '.detail' > /dev/null 2>&1; then
  ERROR=$(echo "$RESPONSE" | jq -r '.detail')
  echo "Error: $ERROR" >&2
  exit 1
fi

# Proceed with data extraction
COVERAGE=$(echo "$RESPONSE" | jq -r '.head_commit.totals.coverage')
```

### Parallel Queries

Use background jobs for independent queries:

```bash
# Start both queries in parallel
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  "https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/commits/${PARENT}" \
  > /tmp/before.json &

curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' \
  > /tmp/after.json &

# Wait for both to complete
wait

# Extract results
BEFORE=$(cat /tmp/before.json | jq -r '.totals.coverage')
AFTER=$(cat /tmp/after.json | jq -r '.head_commit.totals.coverage')
```

## Limitations

### Coverage Calculation Differences

Codecov's coverage percentages may differ from local tools:

**Rounding**:
- Codecov: Uses specific rounding algorithm
- Local tools: May use different precision

**File Inclusion**:
- Codecov: Tracks all repository files
- `go test`: Only includes files with tests

**Calculation Method**:
- Codecov: Line-based coverage
- Go: Statement-based coverage

**Example from PR #226**:
- Local `go tool cover`: 99.0%
- Codecov API: 98.81%
- Difference: 0.19% (0.19 percentage points)

**Rule**: Always use Codecov API numbers for PR coverage claims.

### Historical Data Retention

- **Commits**: Available for 90 days after upload
- **Branches**: Available while branch exists
- **Deleted branches**: Coverage data removed after 30 days

### Component Limitations

Components require configuration in Codecov settings:
1. Define component paths (glob patterns)
2. Upload coverage with component flags
3. Query component endpoint

Most projects use file-level coverage instead.

## Further Reading

- [Codecov API Documentation](https://docs.codecov.com/reference)
- [Coverage Upload Guide](https://docs.codecov.com/docs/codecov-uploader)
- [GitHub Integration](https://docs.codecov.com/docs/github-integration)
