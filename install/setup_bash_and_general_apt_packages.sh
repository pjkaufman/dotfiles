#!/usr/bin/env bash

# create xdg file paths

[ -d "$XDG_DATA_HOME" ] || mkdir -p "$XDG_DATA_HOME"
[ -d "$XDG_CONFIG_HOME" ] || mkdir -p "$XDG_CONFIG_HOME"
[ -d "$XDG_STATE_HOME" ] || mkdir -p "$XDG_STATE_HOME"
[ -d "$XDG_CACHE_HOME" ] || mkdir -p "$XDG_CACHE_HOME"

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
  ["fzf"]="fzf" # fuzzy finder for the terminal
  ["btop"]="btop" # linux task manager
  ["python3"]="python3"
  ["pip3"]="python3-pip" # python 3 package installer
  ["pipx"]="pipx" # python 3 package installer
  ["rename"]="rename" # easier renaming of files
)

for pkg in "${!apt_packages_to_install[@]}"; do install_apt_package "$pkg" "${apt_packages_to_install[$pkg]}"; done

# source exports so later on we have all of our env variables ready to go
source "$HOME/.bashrc"

ensure_file_symlink_is_in_place "$DOTFILES/bash/inputrc" "$INPUTRC" 

ensure_file_symlink_is_in_place "$DOTFILES/btop/btop.conf" "$XDG_CONFIG_HOME/btop/btop.conf"
ensure_folder_symlink_is_in_place "$DOTFILES/btop/themes/" "$XDG_CONFIG_HOME/btop/themes"

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
gtk_config_dir="$XDG_CONFIG_HOME/gtk-3.0"
if [ ! -d "$gtk_config_dir" ]; then
  mkdir -p "$gtk_config_dir"
fi

ensure_file_symlink_is_in_place "$DOTFILES/gtk-3.0/settings.ini" "$gtk_config_dir/settings.ini"
ensure_folder_symlink_is_in_place "$DOTFILES/gtk-3.0/themes" "$HOME/.themes"
