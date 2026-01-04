#!/usr/bin/env python3
"""
Verify test coverage using Codecov API.

This script automates coverage verification by:
1. Querying Codecov API for current and parent commit coverage
2. Calculating actual coverage improvement
3. Comparing file-level coverage changes
4. Validating against coverage thresholds
5. Detecting discrepancies with local tools

Usage:
    # Verify current branch coverage
    python .claude/skills/codecov-api/verify_coverage.py

    # Verify specific branch
    python .claude/skills/codecov-api/verify_coverage.py --branch develop

    # Check against threshold
    python .claude/skills/codecov-api/verify_coverage.py --threshold 99.0

    # Compare specific commits
    python .claude/skills/codecov-api/verify_coverage.py --before abc123 --after def456

    # Output as JSON
    python .claude/skills/codecov-api/verify_coverage.py --format json
"""

import argparse
import json
import os
import sys
from typing import Dict, List, Optional, Tuple
import urllib.request
import urllib.error


class CodecovAPIError(Exception):
    """Codecov API request failed."""
    pass


class CoverageVerifier:
    """Verify coverage using Codecov API."""

    BASE_URL = "https://api.codecov.io/api/v2/github"
    OWNER = "adaptive-enforcement-lab"
    REPO = "readability"

    def __init__(self, token: Optional[str] = None):
        """Initialize verifier with Codecov token.

        Args:
            token: Codecov API token. If None, reads from CODECOV_TOKEN env var.

        Raises:
            ValueError: If token is not provided and not in environment.
        """
        self.token = token or os.getenv("CODECOV_TOKEN")
        if not self.token:
            raise ValueError(
                "CODECOV_TOKEN not found. Set environment variable or pass as argument.\n"
                "Run: source ~/.zshrc"
            )

    def _make_request(self, endpoint: str) -> Dict:
        """Make authenticated request to Codecov API.

        Args:
            endpoint: API endpoint path (e.g., /branches/main)

        Returns:
            Parsed JSON response

        Raises:
            CodecovAPIError: If request fails
        """
        url = f"{self.BASE_URL}/{self.OWNER}/repos/{self.REPO}{endpoint}"
        headers = {"Authorization": f"Bearer {self.token}"}

        try:
            req = urllib.request.Request(url, headers=headers)
            with urllib.request.urlopen(req) as response:
                data = json.loads(response.read().decode())

                # Check for API error in response
                if "detail" in data:
                    raise CodecovAPIError(f"API error: {data['detail']}")

                return data
        except urllib.error.HTTPError as e:
            if e.code == 401:
                raise CodecovAPIError("Invalid or expired token")
            elif e.code == 404:
                raise CodecovAPIError(f"Not found: {endpoint}")
            elif e.code == 429:
                raise CodecovAPIError("Rate limit exceeded (100 requests/minute)")
            else:
                raise CodecovAPIError(f"HTTP {e.code}: {e.reason}")
        except urllib.error.URLError as e:
            raise CodecovAPIError(f"Network error: {e.reason}")

    def get_branch_coverage(self, branch: str = "main") -> Dict:
        """Get coverage for a branch.

        Args:
            branch: Branch name (default: main)

        Returns:
            Branch coverage data with head_commit details
        """
        return self._make_request(f"/branches/{branch}")

    def get_commit_coverage(self, sha: str) -> Dict:
        """Get coverage for a specific commit.

        Args:
            sha: Commit SHA (full or short >= 7 chars)

        Returns:
            Commit coverage data
        """
        return self._make_request(f"/commits/{sha}")

    def verify_improvement(
        self,
        before_sha: Optional[str] = None,
        after_sha: Optional[str] = None,
        branch: str = "main",
    ) -> Dict:
        """Verify coverage improvement between two commits.

        Args:
            before_sha: Parent commit SHA. If None, uses parent of after_sha
            after_sha: Current commit SHA. If None, uses branch head
            branch: Branch to query if after_sha not specified

        Returns:
            Dictionary with before/after coverage and file-level changes
        """
        # Get after commit data
        if after_sha:
            after_data = self.get_commit_coverage(after_sha)
            after_coverage = after_data["totals"]["coverage"]
            files_after = {
                f["name"]: f["totals"]["coverage"]
                for f in after_data.get("report", {}).get("files", [])
            }
        else:
            branch_data = self.get_branch_coverage(branch)
            after_coverage = branch_data["head_commit"]["totals"]["coverage"]
            after_sha = branch_data["head_commit"]["commitid"]
            before_sha = before_sha or branch_data["head_commit"]["parent"]
            files_after = {
                f["name"]: f["totals"]["coverage"]
                for f in branch_data["head_commit"].get("report", {}).get("files", [])
            }

        # Get before commit data
        before_data = self.get_commit_coverage(before_sha)
        before_coverage = before_data["totals"]["coverage"]
        files_before = {
            f["name"]: f["totals"]["coverage"]
            for f in before_data.get("report", {}).get("files", [])
        }

        # Calculate improvement
        improvement = round(after_coverage - before_coverage, 2)

        # Find changed files
        changed_files = {}
        for filename, after_cov in files_after.items():
            before_cov = files_before.get(filename)
            if before_cov is None:
                changed_files[filename] = {
                    "before": None,
                    "after": after_cov,
                    "change": None,
                    "status": "new",
                }
            elif abs(after_cov - before_cov) >= 0.01:
                change = round(after_cov - before_cov, 2)
                changed_files[filename] = {
                    "before": before_cov,
                    "after": after_cov,
                    "change": change,
                    "status": "improved" if change > 0 else "regressed",
                }

        return {
            "before": {
                "sha": before_sha[:7],
                "coverage": before_coverage,
            },
            "after": {
                "sha": after_sha[:7],
                "coverage": after_coverage,
            },
            "improvement": improvement,
            "changed_files": changed_files,
        }

    def check_threshold(
        self, branch: str = "main", threshold: float = 99.0
    ) -> Tuple[bool, float, float]:
        """Check if branch coverage meets threshold.

        Args:
            branch: Branch name
            threshold: Minimum coverage percentage

        Returns:
            Tuple of (meets_threshold, current_coverage, deficit)
        """
        data = self.get_branch_coverage(branch)
        current = data["head_commit"]["totals"]["coverage"]
        meets = current >= threshold
        deficit = round(max(0, threshold - current), 2)

        return meets, current, deficit

    def get_files_below_threshold(
        self, branch: str = "main", threshold: float = 95.0
    ) -> List[Dict]:
        """Get files below coverage threshold.

        Args:
            branch: Branch name
            threshold: Coverage threshold

        Returns:
            List of files below threshold with coverage details
        """
        data = self.get_branch_coverage(branch)
        files = data["head_commit"].get("report", {}).get("files", [])

        below_threshold = []
        for file_data in files:
            coverage = file_data["totals"]["coverage"]
            if coverage < threshold:
                below_threshold.append(
                    {
                        "name": file_data["name"],
                        "coverage": coverage,
                        "lines": file_data["totals"]["lines"],
                        "hits": file_data["totals"]["hits"],
                        "misses": file_data["totals"]["misses"],
                    }
                )

        return sorted(below_threshold, key=lambda x: x["coverage"])


def format_output(data: Dict, format: str = "text") -> str:
    """Format verification results.

    Args:
        data: Verification results
        format: Output format (text, json, markdown)

    Returns:
        Formatted string
    """
    if format == "json":
        return json.dumps(data, indent=2)

    elif format == "markdown":
        output = ["## Coverage Verification\n"]

        if "before" in data and "after" in data:
            output.append(f"- **Before**: {data['before']['coverage']}% ({data['before']['sha']})")
            output.append(f"- **After**: {data['after']['coverage']}% ({data['after']['sha']})")

            change = data["improvement"]
            sign = "+" if change >= 0 else ""
            output.append(f"- **Change**: {sign}{change}%\n")

            if data.get("changed_files"):
                output.append("### Changed Files\n")
                for filename, changes in sorted(data["changed_files"].items()):
                    if changes["status"] == "new":
                        output.append(f"- `{filename}`: NEW ({changes['after']}%)")
                    else:
                        sign = "+" if changes["change"] >= 0 else ""
                        output.append(
                            f"- `{filename}`: {changes['before']}% → {changes['after']}% ({sign}{changes['change']}%)"
                        )

        return "\n".join(output)

    else:  # text
        output = []

        if "before" in data and "after" in data:
            output.append("Coverage Verification")
            output.append("=" * 50)
            output.append(f"Before: {data['before']['coverage']}% ({data['before']['sha']})")
            output.append(f"After:  {data['after']['coverage']}% ({data['after']['sha']})")

            change = data["improvement"]
            sign = "+" if change >= 0 else ""
            output.append(f"Change: {sign}{change}%")

            if data.get("changed_files"):
                output.append("\nChanged Files:")
                output.append("-" * 50)
                for filename, changes in sorted(data["changed_files"].items()):
                    if changes["status"] == "new":
                        output.append(f"  {filename}")
                        output.append(f"    NEW: {changes['after']}%")
                    else:
                        sign = "+" if changes["change"] >= 0 else ""
                        output.append(f"  {filename}")
                        output.append(
                            f"    {changes['before']}% → {changes['after']}% ({sign}{changes['change']}%)"
                        )

        return "\n".join(output)


def main():
    """Main entry point."""
    parser = argparse.ArgumentParser(
        description="Verify test coverage using Codecov API",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Verify current branch
  python .claude/skills/codecov-api/verify_coverage.py

  # Check against threshold
  python .claude/skills/codecov-api/verify_coverage.py --threshold 99.0

  # Compare specific commits
  python .claude/skills/codecov-api/verify_coverage.py --before abc123 --after def456

  # Find files below threshold
  python .claude/skills/codecov-api/verify_coverage.py --files-below 95.0

  # Output as JSON
  python .claude/skills/codecov-api/verify_coverage.py --format json
        """,
    )

    parser.add_argument(
        "--branch",
        default="main",
        help="Branch to verify (default: main)",
    )
    parser.add_argument(
        "--before",
        help="Parent commit SHA (defaults to branch parent)",
    )
    parser.add_argument(
        "--after",
        help="Current commit SHA (defaults to branch head)",
    )
    parser.add_argument(
        "--threshold",
        type=float,
        help="Coverage threshold to check against",
    )
    parser.add_argument(
        "--files-below",
        type=float,
        metavar="THRESHOLD",
        help="Show files below coverage threshold",
    )
    parser.add_argument(
        "--format",
        choices=["text", "json", "markdown"],
        default="text",
        help="Output format (default: text)",
    )
    parser.add_argument(
        "--token",
        help="Codecov API token (defaults to CODECOV_TOKEN env var)",
    )

    args = parser.parse_args()

    try:
        verifier = CoverageVerifier(token=args.token)

        if args.files_below:
            # Show files below threshold
            files = verifier.get_files_below_threshold(
                branch=args.branch, threshold=args.files_below
            )

            if args.format == "json":
                print(json.dumps(files, indent=2))
            else:
                print(f"Files below {args.files_below}% coverage:")
                print("=" * 50)
                for file_data in files:
                    print(f"\n{file_data['name']}")
                    print(f"  Coverage: {file_data['coverage']}%")
                    print(f"  Lines: {file_data['lines']}")
                    print(f"  Hits: {file_data['hits']}")
                    print(f"  Misses: {file_data['misses']}")

        elif args.threshold:
            # Check threshold
            meets, current, deficit = verifier.check_threshold(
                branch=args.branch, threshold=args.threshold
            )

            if args.format == "json":
                result = {
                    "threshold": args.threshold,
                    "current": current,
                    "meets_threshold": meets,
                    "deficit": deficit,
                }
                print(json.dumps(result, indent=2))
            else:
                print(f"Coverage threshold: {args.threshold}%")
                print(f"Current coverage: {current}%")
                if meets:
                    print(f"✓ Coverage meets threshold")
                    sys.exit(0)
                else:
                    print(f"✗ Coverage below threshold")
                    print(f"  Need to improve by {deficit}%")
                    sys.exit(1)

        else:
            # Verify improvement
            result = verifier.verify_improvement(
                before_sha=args.before,
                after_sha=args.after,
                branch=args.branch,
            )

            print(format_output(result, format=args.format))

            # Exit with error if coverage regressed
            if result["improvement"] < 0:
                sys.exit(1)

    except CodecovAPIError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
    except KeyboardInterrupt:
        print("\nInterrupted", file=sys.stderr)
        sys.exit(130)


if __name__ == "__main__":
    main()
