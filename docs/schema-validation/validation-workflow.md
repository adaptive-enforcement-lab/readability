# Validation Workflow and Debugging

This page covers the validation workflow, pre-commit integration, CI validation, and debugging strategies. Use this guide to set up automated validation and troubleshoot issues effectively.

!!! note "Common Errors"
    For detailed examples of type errors, range errors, property errors, and YAML syntax errors, see [Validation Guide](validation-guide.md#common-errors).

## Validation Workflow

Follow this four-step process to validate and fix your configuration file. The workflow helps you identify errors quickly and resolve them systematically.

### Step 1: Run Validation

Start by running the validation command. This checks your configuration against the JSON Schema and reports any issues.

```bash
readability --validate-config
```

### Step 2: Read Error Messages

The tool provides structured error messages to help you fix issues. Each error message includes three components. The location tells you which field has the problem. The problem describes what is wrong. The suggestion guides you to the fix.

**Example error**:
```
  • thresholds.max_grade
    got string, want number
    Suggestion: Remove quotes around numeric values
```

### Step 3: Fix Errors

Apply the suggested fix to your configuration file. Start by locating the field mentioned in the error. Then apply the specific suggestion provided. Finally, save your changes to the file.

### Step 4: Re-validate

Run the validation command again to confirm your fixes. Repeat this process until validation passes.

```bash
readability --validate-config
```

When successful, you will see this message:
```
✓ Configuration is valid
```

## Pre-commit Validation

Catch configuration errors before they reach your repository. Pre-commit hooks validate your files automatically when you commit changes.

Install and configure pre-commit with these commands:

```bash
# Install pre-commit
pip install pre-commit
pre-commit install

# Manually run pre-commit hooks
pre-commit run --all-files
```

The pre-commit hooks perform two validation checks. First, they verify that the schema file is valid JSON Schema. Second, they confirm your config matches the schema requirements.

## CI Validation

GitHub Actions runs automated validation on every pull request. This ensures that schema changes and configuration updates are always valid before merging.

The CI pipeline validates three aspects. It checks the schema file on every pull request. It validates the config file against the schema. It also runs runtime validation tests to catch integration issues.

When CI fails with validation errors, follow these steps. First, check the output from the job named "Validate JSON Schema" in GitHub Actions. Read the error message carefully. Finally, fix the issues locally and push your changes again.

## Debugging Tips

When validation fails, these debugging strategies can help you identify and fix the root cause.

### Check Schema URL

The schema URL must be exactly correct for IDE validation to work. Verify that your config file includes this exact line at the top:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
```

Common mistakes include missing the protocol, misspelling the domain, or using the wrong path. The URL must start with `https://` and end with `/config.json`.

### Verify YAML Syntax

YAML syntax errors can cause validation to fail before schema checks even run. Use a dedicated YAML linter to catch these issues early.

```bash
# Install yamllint
pip install yamllint

# Lint YAML file
yamllint .readability.yml
```

### Test with Minimal Config

When facing complex validation errors, start simple. Create a minimal configuration that you know is valid. Then add fields one at a time to identify which addition causes the failure.

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade: 12
```

This minimal config should always validate. Add your custom fields incrementally to isolate the problematic configuration.

### Check Field Names

IDE autocomplete helps you avoid typos in field names. Open your config file in your editor. Type `thresholds:` on a new line. Press Ctrl+Space on Windows or Linux, or Cmd+Space on macOS. Your IDE will show all available fields.

### Compare with Examples

Review working examples to see correct usage patterns. The documentation includes comprehensive examples in several pages. Check the Configuration File guide, Schema Reference, and Validation Guide for working code.

## Getting Help

When validation errors persist despite troubleshooting, these resources can help you find a solution.

First, consult the Schema Reference for complete field documentation. The reference explains every configuration option with examples and constraints.

For more detailed error output, run the validation tool in verbose mode. This provides additional context about what failed and why.

```bash
check-jsonschema --verbose --schemafile docs/schemas/config.json .readability.yml
```

Check your editor logs for YAML language server errors. These logs often contain additional details not shown in the validation output.

If you still cannot resolve the issue, file a bug report on GitHub. Include your complete config file, the full error message, and the output from running `readability --validate-config`. This information helps maintainers reproduce and fix the problem quickly.

## Next Steps

- [Validation Guide](validation-guide.md): Common error examples and fixes
- [Schema Reference](schema-reference.md): Complete field documentation
- [IDE Setup](ide-setup.md): Configure validation in your editor
- [Configuration File](../cli/config-file.md): Usage examples
