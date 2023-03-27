#!/bin/bash

# setup the bash and terminal related files 

tmux_dir="$HOME/.config/tmux"
declare -A bash_related_file_sylink_info=( 
  ["$HOME/dotfiles/.shellrc/bash_profile"]="$HOME/.bash_profile"
  ["$HOME/dotfiles/.shellrc/bashrc"]="$HOME/.bashrc" 
  ["$HOME/dotfiles/.shellrc/hushlogin"]="$HOME/.hushlogin" # make sure that certain logs are not shown on startup
  ["$HOME/dotfiles/tmux/tmux.conf"]="$tmux_dir/tmux.conf"
)

tmux_dir="$HOME/.config/tmux"
if [ ! -d "$HOME/.config/tmux" ]; then 
  mkdir -p "$tmux_dir"
fi

for file in "${!bash_related_file_sylink_info[@]}"; do ensure_file_symlink_is_in_place "$file" "${bash_related_file_sylink_info[$file]}"; done

declare -A apt_packages_to_install=( 
  ["tmux"]="tmux" 
  ["grep"]="grep" 
  ["curl"]="curl" 
  ["rg"]="ripgrep" # better version of grep
  ["btop"]="btop" # linux task manager
  ["python3"]="python3"
  ["pip3"]="python3-pip" # python 3 package installer
  ["rename"]="rename" # easier renaming of files
)

for pkg in "${!apt_packages_to_install[@]}"; do install_apt_package "$pkg" "${apt_packages_to_install[$pkg]}"; done

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
