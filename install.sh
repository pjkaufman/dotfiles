#!/usr/bin/env bash

# make sure bash fails if any error happens
set -e

SCRIPTPATH="$(
  cd -- "$(dirname "$0")" > /dev/null 2>&1
  pwd -P
)"

setup_header_text() {
  echo ""
  echo "==============================="
  echo "$1"
  echo "==============================="
  echo ""
}

include_file() {
  # shellcheck disable=SC1090
  source "$SCRIPTPATH/$1"
}

# pull in the package installation helper functions
include_file "install/package_install_functions"

# pull in the symlink helper functions
include_file "install/symlink_functions"

# include computer type functions
# shellcheck source=./bash/functions/computer_type_functions.sh
source "$SCRIPTPATH/bash/functions/computer_type_functions.sh"

# actual setup

echo "starting environment setup"

install_script_section_text=(
  "setup computer type"
  "setup bash and common packages"
  "setup fonts"
  "setup git"
  "setup tmux"
  "setup go"
  "setup npm"
  "setup kitty"
  "setup rkhunter"
  "setup syncthing"
  "setup flatpaks"
  "setup i3"
  "setup vscode"
  "setup neovim"
  "setup doc converts"
)

declare -A install_script_sections_files=(
  ["setup computer type"]="install/setup_computer_type"
  ["setup bash and common packages"]="install/setup_bash_and_general_apt_packages"
  ["setup fonts"]="install/setup_fonts"
  ["setup git"]="install/setup_git"
  ["setup tmux"]="install/setup_tmux"
  ["setup syncthing"]="install/setup_syncthing"
  ["setup flatpaks"]="install/setup_flatpak"
  ["setup go"]="install/setup_go"
  ["setup npm"]="install/setup_npm"
  ["setup kitty"]="install/setup_kitty"
  ["setup rkhunter"]="install/setup_rkhunter"
  ["setup i3"]="install/setup_i3"
  ["setup neovim"]="install/setup_neovim"
  ["setup vscode"]="install/setup_vscode"
  ["setup doc converts"]="install/setup_doc_converters"
)

for i in "${!install_script_section_text[@]}"; do
  header="${install_script_section_text[$i]}"
  setup_header_text "${header}:"
  include_file "${install_script_sections_files[$header]}"
done

unset header

# remove any no longer needed packages
sudo apt autoremove -y

echo ""
echo "environment setup complete"
