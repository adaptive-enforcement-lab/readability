# Contributing to Readability

Thank you for your interest in contributing to Readability! This document provides guidelines and information for contributors.

## How to Contribute

### Reporting Bugs

If you find a bug, please [open an issue](https://github.com/adaptive-enforcement-lab/readability/issues/new) with:

- A clear, descriptive title
- Steps to reproduce the issue
- Expected vs actual behavior
- Your environment (OS, Go version, etc.)
- Sample input files if applicable

### Suggesting Enhancements

Feature requests are welcome! Please [open an issue](https://github.com/adaptive-enforcement-lab/readability/issues/new) with:

- A clear description of the proposed feature
- The problem it solves or use case it enables
- Any relevant examples or references

### Submitting Pull Requests

1. **Fork the repository** and create your branch from `main`
2. **Make your changes** following the code style guidelines below
3. **Add tests** for any new functionality
4. **Ensure all tests pass** with `go test -race ./...`
5. **Run the linter** with `golangci-lint run`
6. **Submit a pull request** with a clear description of your changes

## Development Setup

### Prerequisites

- Go 1.23 or later
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [gotestsum](https://github.com/gotestyourself/gotestsum) (optional, for better test output)

### Building

```bash
git clone https://github.com/adaptive-enforcement-lab/readability.git
cd readability
go build ./cmd/readability
```

### Running Tests

```bash
# Run all tests
go test -race ./...

# With coverage
go test -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Using gotestsum (recommended)
gotestsum -- -race -coverprofile=coverage.out ./...
```

### Linting

```bash
golangci-lint run
```

## Code Style

- Follow standard Go conventions and [Effective Go](https://go.dev/doc/effective_go)
- Use `gofmt` for formatting (enforced by CI)
- Write clear, descriptive commit messages
- Add comments for exported functions and complex logic
- Maintain test coverage above 95%

## Testing Requirements

- All new features must include tests
- All bug fixes should include regression tests
- Tests must pass on all supported platforms
- Coverage must remain above 95% for each package

## Code of Conduct

Be respectful and constructive in all interactions. We are committed to providing a welcoming and inclusive environment for all contributors.

## Questions?

If you have questions about contributing, feel free to [open an issue](https://github.com/adaptive-enforcement-lab/readability/issues/new) with the "question" label.

## License

By contributing to Readability, you agree that your contributions will be licensed under the [MIT License](LICENSE).
