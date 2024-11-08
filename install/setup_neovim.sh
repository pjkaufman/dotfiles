#!/usr/bin/env bash

add_ppa_and_install_package "neovim-ppa/unstable" "neovim"

ensure_folder_symlink_is_in_place "$DOTFILES/nvim" "$XDG_CONFIG_HOME/nvim"

# null-ls lsp formatters and diagnostics
neovim_pip_packages_to_install=(
  "codespell"
  "black"
  "flake8"
  "beautysh"
)
for pkg in "${neovim_pip_packages_to_install[@]}"; do pip_install_package "$pkg"; done

if is_work_computer; then
  declare -A neovim_go_packages_to_install=(
    ["github.com/yoheimuta/protolint/cmd/protolint"]="protolint"
    ["golang.org/x/tools/cmd/goimports"]="golangci-lint"
  )

  for pkg in "${!neovim_go_packages_to_install[@]}"; do go_install_package "$pkg" "${neovim_go_packages_to_install[$pkg]}"; done
fi

declare -A neovim_common_go_packages_to_install=(
  ["golang.org/x/tools/cmd/goimports"]="goimports"
  ["github.com/go-delve/delve/cmd/dlv"]="dlv" # dlv install for go debugging
)
for pkg in "${!neovim_common_go_packages_to_install[@]}"; do go_install_package "$pkg" "${neovim_common_go_packages_to_install[$pkg]}"; done

npm_install_package "eslint"

# TODO: install via a different mechanism
# cargo_install_package "stylua"

# TODO: handle google_java_format install
