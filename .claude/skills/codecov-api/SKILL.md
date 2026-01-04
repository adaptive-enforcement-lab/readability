---
name: codecov-api-verification
description: Verifies test coverage using Codecov API, compares coverage between commits/branches, validates coverage claims against thresholds. Use when verifying PR coverage improvements, investigating local tool discrepancies, or validating CI/CD coverage reports.
allowed-tools: Read, Bash(python:*), Bash(curl:*), Bash(jq:*), Bash(bc:*), Bash(source:*)
model: haiku
---

# Codecov API Verification

Verify coverage claims against Codecov API instead of trusting local tools like `go tool cover`.

## Prerequisites

Ensure `CODECOV_TOKEN` is set:

```bash
# Check if set
echo $CODECOV_TOKEN

# If not set, source from shell config
source ~/.zshrc
```

## Quick Start

### Verify Current Coverage

```bash
python .claude/skills/codecov-api/verify_coverage.py
```

Returns before/after coverage and file-level changes.

### Check Against Threshold

```bash
python .claude/skills/codecov-api/verify_coverage.py --threshold 99.0
```

Exits with code 1 if below threshold.

### Find Files Below Target

```bash
python .claude/skills/codecov-api/verify_coverage.py --files-below 95.0
```

Lists files needing coverage improvement.

## Common Tasks

### Task 1: Verify PR Coverage Claim

Before making coverage claims in a PR:

```bash
# Run automated verification
python .claude/skills/codecov-api/verify_coverage.py --format markdown > coverage-report.md

# Check the report shows actual Codecov numbers
cat coverage-report.md
```

**Critical**: Always use Codecov API numbers, not `go tool cover` output.

### Task 2: Compare Commits

```bash
python .claude/skills/codecov-api/verify_coverage.py \
  --before abc123 \
  --after def456 \
  --format json
```

Returns JSON with improvement details and changed files.

### Task 3: Manual API Query

```bash
# Get current branch coverage
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq '.head_commit.totals.coverage'
```

### Task 4: Get Component Breakdown

```bash
curl -s -H 'Authorization: Bearer $CODECOV_TOKEN' \
  'https://api.codecov.io/api/v2/github/adaptive-enforcement-lab/repos/readability/branches/main' | \
  jq -r '.head_commit.report.files[] | "\(.name): \(.totals.coverage)%"'
```

## Coverage Verification Workflow

Use this checklist when verifying coverage:

- [ ] Get parent commit coverage (before changes)
- [ ] Get current commit coverage (after changes)
- [ ] Calculate actual coverage improvement
- [ ] Get file-level breakdown
- [ ] Compare against local tool output
- [ ] Document discrepancies if any
- [ ] Use Codecov numbers for PR claims

**Automated workflow**:

```bash
# Single command verification
python .claude/skills/codecov-api/verify_coverage.py
```

## Understanding Coverage Discrepancies

Local tools (`go tool cover`, etc.) often show different numbers than Codecov:

**Why this happens**:
1. **Rounding**: Different algorithms (local: 99.0%, Codecov: 98.81%)
2. **File inclusion**: Codecov tracks all repo files; `go test` only files with tests
3. **Calculation**: Codecov uses line coverage; Go uses statement coverage

**Example from PR #226**:
- Local `go tool cover`: 99.0%
- Codecov API: 98.81%
- Difference: 0.19 percentage points

**Rule**: Always use Codecov API data for PR coverage claims.

## Advanced Reference

- **API endpoints and methods**: [API-REFERENCE.md](API-REFERENCE.md)
- **Real-world examples**: [EXAMPLES.md](EXAMPLES.md)
- **Automated verification**: Run `python verify_coverage.py --help`

## Verification Script Options

```bash
# Script location
.claude/skills/codecov-api/verify_coverage.py

# Available options
--branch BRANCH          # Branch to verify (default: main)
--before SHA             # Parent commit SHA
--after SHA              # Current commit SHA
--threshold PERCENT      # Coverage threshold to check
--files-below PERCENT    # Show files below threshold
--format FORMAT          # Output format (text, json, markdown)
--token TOKEN            # Codecov API token (defaults to env var)
```

## Error Handling

### Token Not Found

```bash
# Verify token is set
echo $CODECOV_TOKEN

# If empty, source shell config
source ~/.zshrc
```

### Invalid Commit SHA

Commit SHAs must be full (40 chars) or at least 7 chars:

```bash
# Good
.../commits/43256dc9844084feb1835f53ab13aff27e1273fe
.../commits/43256dc

# Bad (too short)
.../commits/432
```

### Branch Not Found

Branch names are case-sensitive:

```bash
# Good
/branches/main

# Bad
/branches/Main
/branches/master
```

## Response Structure

### Branch Response

```json
{
  "head_commit": {
    "commitid": "abc123...",
    "parent": "def456...",
    "totals": {
      "coverage": 98.81,
      "files": 10,
      "lines": 1101,
      "hits": 1088,
      "misses": 10
    },
    "report": {
      "files": [
        {
          "name": "cmd/readability/main.go",
          "totals": {
            "coverage": 97.29
          }
        }
      ]
    }
  }
}
```

**Key fields**:
- `head_commit.totals.coverage`: Overall coverage percentage
- `head_commit.parent`: Parent commit SHA for before/after comparison
- `head_commit.report.files`: File-level coverage details

### Commit Response

```json
{
  "totals": {
    "coverage": 98.45
  },
  "report": {
    "files": [
      {
        "name": "pkg/analyzer/analyzer.go",
        "totals": {
          "coverage": 99.11
        }
      }
    ]
  }
}
```

**Key fields**:
- `totals.coverage`: Commit coverage percentage
- `report.files`: Array of file coverage objects

## Summary

This skill provides:
- ✅ Automated coverage verification against Codecov API
- ✅ Before/after comparison workflows
- ✅ Threshold validation (with exit codes for CI/CD)
- ✅ File-level coverage analysis
- ✅ Local tool vs Codecov discrepancy explanation
- ✅ Zero-dependency Python verification script

**Critical rule**: Always use Codecov API data for PR coverage claims, not local tool output.
