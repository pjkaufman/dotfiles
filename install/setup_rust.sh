#!/usr/bin/env bash

install_apt_package "cargo"

cargo_packages_to_install=(
  "topgrade" # package that helps manage updating all package managers
  "cargo-update" # package for updating rust
)
for pkg in "${cargo_packages_to_install[@]}"; do cargo_install_package "$pkg"; done

ensure_file_symlink_is_in_place "$DOTFILES/topgrade/topgrade.toml" "$XDG_CONFIG_HOME/topgrade.toml"
