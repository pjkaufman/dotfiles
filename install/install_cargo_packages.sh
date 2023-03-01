#!/bin/bash

cargo_packages_to_install=("topgrade" "cargo-update")
for pkg in "${cargo_packages_to_install[@]}"; do cargo_install_package "$pkg"; done
