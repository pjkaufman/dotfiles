#!/usr/bin/env bash

# setup the bash and terminal related files 

declare -A bash_related_file_sylink_info=( 
  ["$HOME/dotfiles/bash/bashrc"]="$HOME/.bashrc" 
  ["$HOME/dotfiles/bash/hushlogin"]="$HOME/.hushlogin" # make sure that certain logs are not shown on startup
)

for file in "${!bash_related_file_sylink_info[@]}"; do ensure_file_symlink_is_in_place "$file" "${bash_related_file_sylink_info[$file]}"; done

declare -A apt_packages_to_install=( 
  ["grep"]="grep" 
  ["curl"]="curl" 
  ["rg"]="ripgrep" # better version of grep
  ["btop"]="btop" # linux task manager
  ["python3"]="python3"
  ["pip3"]="python3-pip" # python 3 package installer
  ["rename"]="rename" # easier renaming of files
)

for pkg in "${!apt_packages_to_install[@]}"; do install_apt_package "$pkg" "${apt_packages_to_install[$pkg]}"; done

# source exports so later on we have all of our env variables ready to go
source "$HOME/.bashrc"

ensure_file_symlink_is_in_place "$HOME/bash/inputrc" "$INPUTRC" 

# only try to add the remaining packages if on a personal computer
if is_work_computer; then
  return
fi

personal_apt_packages_to_install=(
  "imgp" # image compression
  "flameshot" # screenshots
  "evince" # pdf editor and viewer
)

for pkg in "${personal_apt_packages_to_install[@]}"; do install_apt_package "$pkg"; done

# gtk-3.0
gtk_config_dir="$HOME/.config/gtk-3.0"
if [ ! -d "$gtk_config_dir" ]; then
  mkdir -p "$gtk_config_dir"
fi

ensure_file_symlink_is_in_place "$HOME/dotfiles/gtk-3.0/settings.ini" "$gtk_config_dir/settings.ini"
