#!/bin/bash

# pull in the package installation helper functions
source ./install/package_install_functions.sh

# pull in the symlink helper functions
source ./install/symlink_functions.sh

setup_header_text() {
  echo ""
  echo "==============================="
  echo "$1"
  echo "==============================="
  echo ""
}

# actual setup

echo "starting environment setup"

# determine whether or not this env is work or personal

setup_header_text "get computer type:"
source ./install/computer_type_handler.sh

# apt packages

setup_header_text "apt packages:"
source ./install/install_apt_packages.sh

# cargo packages

setup_header_text "cargo packages:"
source ./install/install_cargo_packages.sh

# PPA additions

setup_header_text "PPA additions"
source ./install/install_ppa_packages.sh

# flatpak packages

setup_header_text "Flatpak packages:"
source ./install/flatpak_handler.sh

# special package managers like nvm and gvm

setup_header_text "gvm and nvm install:"
source ./install/install_gvm_and_nvm.sh

# TODO: add logic for wkhtml to pdf
# TODO: add logic around ssh agent for github

setup_header_text "ssh setup:"
source ./install/install_ssh.sh

# i3 setup

setup_header_text "i3 setup:"
source ./install/install_i3_packages.sh

# neovim setup

setup_header_text "Neovim setup:"
source ./install/install_neovim_packages.sh

# setup config symlinks

setup_header_text "Symlink setup:"
source ./install/setup_symlinks.sh

echo ""
echo "environment setup complete"
