# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 1.x     | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security issue, please report it responsibly.

### Private Disclosure (Preferred)

For security vulnerabilities, please use **GitHub's private vulnerability reporting**:

1. Go to the [Security Advisories](https://github.com/adaptive-enforcement-lab/readability/security/advisories) page
2. Click "Report a vulnerability"
3. Provide details about the vulnerability

This ensures the issue is handled privately until a fix is available.

### What to Include

When reporting a vulnerability, please include:

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Any suggested fixes (optional)

### Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 7 days
- **Resolution Target**: Within 90 days (depending on severity)

### After Reporting

1. We will acknowledge receipt of your report
2. We will investigate and validate the issue
3. We will work on a fix and coordinate disclosure
4. We will credit you in the security advisory (unless you prefer anonymity)

## Security Measures

This project employs several security measures:

- **Trivy scanning**: Automated vulnerability scanning on every CI build
- **Dependabot**: Automated dependency updates for security patches
- **SBOM generation**: Software Bill of Materials for transparency
- **Code review**: All changes require pull request review

## Scope

This security policy applies to:

- The `readability` CLI tool
- The GitHub Action
- Official container images

Third-party forks and modifications are not covered by this policy.
