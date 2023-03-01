#!/bin/bash

# null-ls lsp formatters and diagnostics 
neovim_pip_packages_to_install=(
  "codespell"
  "black"
  "flake8"
  "beautysh"
)
for pkg in "${neovim_pip_packages_to_install[@]}"; do pip_install_package "$pkg"; done

if [ ${COMPUTER_TYPE} = "work" ]
then
  declare -A neovim_go_packages_to_install=( 
    ["github.com/yoheimuta/protolint/cmd/protolint"]="protolint"
    ["golang.org/x/tools/cmd/goimports"]="golangci-lint"
  )

  for pkg in "${!neovim_go_packages_to_install[@]}"; do go_install_package "$pkg" "${neovim_go_packages_to_install[$pkg]}"; done
  # go_install_package "github.com/yoheimuta/protolint/cmd/protolint" "protolint"
  # go_install_package "golang.org/x/tools/cmd/goimports" "golangci-lint"
fi

go_install_package "golang.org/x/tools/cmd/goimports" "goimports"

npm_install_package "eslint"

cargo_install_package "stylua"

# TODO: handle google_java_format install
