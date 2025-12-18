# IDE Setup Guide

This guide shows how to configure different editors for YAML schema validation with `.readability.yml` files.

## VS Code

### Installation

1. Install the [YAML extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml) by Red Hat:

    ```bash
    code --install-extension redhat.vscode-yaml
    ```

    Or install from the Extensions marketplace (Ctrl+Shift+X / Cmd+Shift+X).

2. Open your `.readability.yml` file

3. Ensure the first line contains the schema reference:

    ```yaml
    # yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
    ```

4. Start typing - autocomplete will appear automatically

### Verification

1. Type `thresholds:`
2. Press space, then Ctrl+Space (or Cmd+Space on macOS)
3. You should see autocomplete suggestions: `max_grade`, `max_ari`, etc.
4. Hover over a field to see its description

### Troubleshooting

**Autocomplete not working?**

- Restart VS Code after installing the extension
- Check that the file is named `.readability.yml` (not `.yaml`)
- Verify the schema URL is correct in the first line
- Check VS Code output: View → Output → Select "YAML Support" from dropdown

**Schema not loading?**

- Ensure you have internet connectivity (schema is fetched from URL)
- Check the schema URL is accessible: visit it in your browser
- Try a workspace reload: Ctrl+Shift+P → "Developer: Reload Window"

## JetBrains IDEs

JetBrains IDEs (IntelliJ IDEA, WebStorm, PyCharm, GoLand, etc.) have built-in YAML schema support.

### Setup

1. Open your `.readability.yml` file

2. Add the schema reference to the first line:

    ```yaml
    # yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
    ```

3. The IDE automatically detects and loads the schema

### Verification

1. Type `thresholds:`
2. Press Ctrl+Space (Windows/Linux) or Cmd+Space (macOS)
3. Autocomplete suggestions appear
4. Hover over fields to see documentation

### Troubleshooting

**Schema not detected?**

- File → Invalidate Caches → Invalidate and Restart
- Settings → Languages & Frameworks → Schemas and DTOs → JSON Schema Mappings
- Ensure "YAML" plugin is enabled: Settings → Plugins → search for "YAML"

## Neovim

### Prerequisites

- Neovim 0.8+ (for native LSP support)
- A package manager (lazy.nvim, packer, vim-plug, etc.)

### Installation with Mason

[Mason](https://github.com/williamboman/mason.nvim) is the recommended way to install LSP servers:

```lua
-- In your Neovim config (init.lua)
require('mason').setup()
require('mason-lspconfig').setup({
    ensure_installed = { 'yamlls' }
})

-- Configure yaml-language-server
require('lspconfig').yamlls.setup({
    settings = {
        yaml = {
            schemaStore = {
                enable = true,
                url = "https://www.schemastore.org/api/json/catalog.json",
            },
            schemas = {},
            validate = true,
        }
    }
})
```

### Manual Installation

1. Install yaml-language-server:

    ```bash
    npm install -g yaml-language-server
    ```

2. Configure in your Neovim config:

    ```lua
    require('lspconfig').yamlls.setup({
        settings = {
            yaml = {
                validate = true,
                schemaStore = { enable = true },
            }
        }
    })
    ```

3. Add the schema reference to your `.readability.yml`:

    ```yaml
    # yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
    ```

### Verification

1. Open `.readability.yml` in Neovim
2. Use `:LspInfo` to verify yaml-language-server is attached
3. Type `thresholds:` and trigger completion (usually Ctrl+X Ctrl+O)
4. Hover over fields with `K` to see documentation

### Troubleshooting

**LSP not attaching?**

- Check `:LspInfo` shows yaml-language-server
- Verify server is installed: `which yaml-language-server`
- Check Neovim logs: `:messages`
- Ensure filetype is set: `:set filetype?` should show `yaml`

**Autocomplete not working?**

- Verify completion plugin is installed (nvim-cmp, coq_nvim, etc.)
- Try manual completion: Ctrl+X Ctrl+O in insert mode
- Check LSP capabilities: `:lua =vim.lsp.get_active_clients()[1].server_capabilities`

## Vim

Vim requires a plugin to use LSP servers.

### Option 1: coc.nvim with coc-yaml

1. Install [coc.nvim](https://github.com/neoclide/coc.nvim):

    ```vim
    Plug 'neoclide/coc.nvim', {'branch': 'release'}
    ```

2. Install coc-yaml:

    ```vim
    :CocInstall coc-yaml
    ```

3. Add schema reference to `.readability.yml`:

    ```yaml
    # yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
    ```

4. Restart Vim

### Option 2: ALE with yaml-language-server

1. Install [ALE](https://github.com/dense-analysis/ale):

    ```vim
    Plug 'dense-analysis/ale'
    ```

2. Install yaml-language-server globally:

    ```bash
    npm install -g yaml-language-server
    ```

3. Configure ALE in `.vimrc`:

    ```vim
    let g:ale_linters = {'yaml': ['yamllint', 'yamlls']}
    let g:ale_completion_enabled = 1
    ```

4. Add schema reference to `.readability.yml`

### Verification

- Open `.readability.yml`
- Trigger completion (usually Ctrl+X Ctrl+O or Tab with coc.nvim)
- Check ALE status: `:ALEInfo`

## Emacs

### Installation with lsp-mode

1. Install lsp-mode and yaml-language-server:

    ```elisp
    ;; In your Emacs config
    (use-package lsp-mode
      :ensure t
      :hook ((yaml-mode . lsp-deferred))
      :commands (lsp lsp-deferred))

    (use-package yaml-mode
      :ensure t
      :mode "\\.ya?ml\\'")
    ```

2. Install yaml-language-server globally:

    ```bash
    npm install -g yaml-language-server
    ```

3. Add schema reference to `.readability.yml`:

    ```yaml
    # yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
    ```

4. Restart Emacs

### With eglot

Alternatively, use eglot (built-in to Emacs 29+):

```elisp
(use-package eglot
  :ensure t
  :hook (yaml-mode . eglot-ensure))
```

### Verification

- Open `.readability.yml`
- Check LSP status: `M-x lsp-describe-session`
- Trigger completion: `M-x completion-at-point` or configured keybinding

## Schema Reference Location

All editors load the schema from:

```
https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
```

The `/latest/` path always points to the most recent schema version. For version-specific schemas, use:

```
https://readability.adaptive-enforcement-lab.com/v1.11.0/schemas/config.json
```

## Offline Usage

For offline development:

1. Download the schema file:

    ```bash
    curl -O https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
    ```

2. Reference the local file:

    ```yaml
    # yaml-language-server: $schema=file:///path/to/config.json
    ```

!!! warning "Absolute Paths"
    Local schema references must use absolute paths. Relative paths are not supported by most YAML language servers.

## Testing Your Setup

Create a test `.readability.yml` file:

```yaml
# yaml-language-server: $schema=https://readability.adaptive-enforcement-lab.com/latest/schemas/config.json
---
thresholds:
  max_grade:  # Press Ctrl+Space here - should show type hint "number"
```

Try these tests:

1. **Autocomplete**: Type `max_` and trigger completion - should suggest `max_grade`, `max_ari`, `max_fog`, etc.
2. **Validation**: Type `max_grade: "twelve"` - should show error (expects number, got string)
3. **Documentation**: Hover over `max_grade` - should show description
4. **Range check**: Type `max_grade: 200` - should warn (exceeds maximum of 100)

All tests passing? Your setup is working correctly!

## Next Steps

- [Schema Reference](schema-reference.md) - Complete schema documentation
- [Validation Guide](validation-guide.md) - Common error examples and fixes
- [Validation Workflow](validation-workflow.md) - Step-by-step validation process
- [Configuration File](../index.md) - Learn all config options
