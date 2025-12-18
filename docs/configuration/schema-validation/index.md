# Schema Validation

The `.readability.yml` configuration file is powered by a comprehensive JSON Schema that provides IDE autocomplete, real-time validation, and inline documentation.

## Why Schema Validation?

Schema validation catches configuration errors **before** you commit, saving time and reducing CI failures.

### Benefits

- ✅ **Autocomplete** - IntelliSense suggests all available options
- ✅ **Real-time Validation** - Red squiggles show errors as you type
- ✅ **Inline Documentation** - Hover tooltips show descriptions, defaults, and examples
- ✅ **Type Safety** - Prevents invalid values and typos
- ✅ **CI Integration** - Automatic validation in pre-commit hooks and GitHub Actions

## Quick Start

Add the schema reference to the first line of your `.readability.yml`:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade: 12
  max_ari: 12
```

That's it! Your IDE will now provide autocomplete and validation.

!!! tip "Zero Configuration"
    Most modern editors automatically detect the schema reference and enable IDE features. No manual configuration needed.

## IDE Support

The following editors support YAML schema validation:

| Editor | Supported | Setup Required |
|--------|-----------|----------------|
| VS Code | ✅ | [Install YAML extension](ide-setup.md#vs-code) |
| JetBrains IDEs | ✅ | Built-in (IntelliJ, WebStorm, PyCharm) |
| Neovim | ✅ | [Install yaml-language-server](ide-setup.md#neovim) |
| Vim | ✅ | [Use coc-yaml or ALE](ide-setup.md#vim) |
| Emacs | ✅ | [Use lsp-mode](ide-setup.md#emacs) |

See [IDE Setup Guide](ide-setup.md) for detailed instructions.

## Validation Methods

### IDE Validation (Real-time)

Your editor validates the config file as you type. Errors appear as red squiggles.

### CLI Validation (Pre-commit)

Validate your config before committing:

```bash
readability --validate-config
```

### Pre-commit Hooks (Automatic)

Install pre-commit hooks to validate automatically:

```bash
pip install pre-commit
pre-commit install
```

The hooks validate:
- Schema file against JSON Schema Draft 2020-12 metaschema
- `.readability.yml` against the schema

### CI Validation (GitHub Actions)

The CI workflow automatically validates the schema and config files on every PR.

See [Validation Guide](validation-guide.md) for common error examples and [Validation Workflow](validation-workflow.md) for the complete validation process.

## Schema Reference

The schema defines all available configuration options with:

- **Type constraints** - Number, string, integer, object, array
- **Range validation** - Minimum and maximum values
- **Required fields** - Must-have properties
- **Descriptions** - Inline documentation for each field
- **Examples** - Sample values

See [Schema Reference](schema-reference.md) for field definitions and [Schema Overrides and Validation](schema-overrides.md) for path-specific overrides and validation rules.

## Maintaining the Schema

For contributors: the schema must stay synchronized with Go structs.

- **Location**: `docs/schemas/config.json`
- **Tests**: `pkg/config/schema_test.go` verifies sync
- **CI**: Validates schema on every PR

See [Maintaining the Schema](maintaining-schema.md) for developer guidelines.

## Troubleshooting

### Schema Not Loading

If autocomplete doesn't work:

1. Check the schema URL in the first line
2. Verify your editor has YAML language server installed
3. Restart your editor
4. Check editor-specific setup in [IDE Setup Guide](ide-setup.md)

### Validation Errors

If you see validation errors:

1. Run `readability --validate-config` for detailed error messages
2. Check the [Schema Reference](schema-reference.md) for valid values
3. See common errors in [Validation Guide](validation-guide.md)
4. Follow the workflow in [Validation Workflow](validation-workflow.md)

### Schema Updates

The schema is versioned with the documentation:

- **Latest**: `https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json`
- **Specific version**: `https://readability.adaptive-enforcement-lab.com/v1.11.0/schemas/config.json`

Always use `/latest/` for the most recent schema.
