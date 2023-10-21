#!/usr/bin/env bash

# make sure bash fails if any error happens
set -e

setup_header_text() {
  echo ""
  echo "==============================="
  echo "$1"
  echo "==============================="
  echo ""
}

include_file() {
  source "$HOME/dotfiles/$1"
}

# pull in the package installation helper functions
include_file "install/package_install_functions.sh"

# pull in the symlink helper functions
include_file "install/symlink_functions.sh"

# include computer type functions
source "$HOME/dotfiles/bash/functions/computer_type_functions.sh" 

# actual setup

echo "starting environment setup"

install_script_section_text=(
  "setup computer type" 
  "setup bash and common packages"
  "setup git"
  "setup tmux"
  "setup rust"
  "setup go"
  "setup npm"
  "setup kitty"
  "setup rkhunter"
  "setup syncthing"
  "setup flatpaks"
  "setup i3"
  "setup neovim"
  "setup wkhtmltopdf"
  "setup wine"
)

 declare -A install_script_sections_files=( 
  ["setup computer type"]="install/setup_computer_type.sh" 
  ["setup bash and common packages"]="install/setup_bash_and_general_apt_packages.sh"
  ["setup git"]="install/setup_git.sh"
  ["setup tmux"]="install/setup_tmux.sh"
  ["setup rust"]="install/setup_rust.sh"
  ["setup syncthing"]="install/setup_syncthing.sh"
  ["setup flatpaks"]="install/setup_flatpak.sh"
  ["setup go"]="install/setup_go.sh"
  ["setup npm"]="install/setup_npm.sh"
  ["setup kitty"]="install/setup_kitty.sh"
  ["setup rkhunter"]="install/setup_rkhunter.sh"
  ["setup i3"]="install/setup_i3.sh"
  ["setup neovim"]="install/setup_neovim.sh"
  ["setup wkhtmltopdf"]="install/setup_wkhtmltopdf.sh"
  ["setup wine"]="install/setup_wine.sh"
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
