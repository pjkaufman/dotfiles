#!/bin/bash

ensure_folder_symlink_is_in_place "$HOME/dotfiles/nvim" "$HOME/.config/nvim"

declare -A file_symlink_info=( 
  ["$HOME/dotfiles/git/.gitconfig"]="$HOME/.gitconfig" 
  ["$HOME/dotfiles/.shellrc/bash_profile"]="$HOME/.bash_profile" 
  ["$HOME/dotfiles/.shellrc/bashrc"]="$HOME/.bashrc" 
  ["$HOME/dotfiles/.shellrc/hushlogin"]="$HOME/.hushlogin"
  ["$HOME/dotfiles/tmux/.tmux.conf"]="$HOME/.tmux.conf"
  ["$HOME/dotfiles/topgrade/topgrade.toml"]="$HOME/.config/topgrade.toml"
)

for file in "${!file_symlink_info[@]}"; do ensure_file_symlink_is_in_place "$file" "${file_symlink_info[$file]}"; done

# only setup the remaining symlinks if on a personal computer
if [ ${COMPUTER_TYPE} = "work" ]
then
  return
fi

declare -A personal_file_symlink_info=( 
  ["$HOME/dotfiles/kitty/kitty.conf"]="$HOME/.config/kitty/kitty.conf"
  ["$HOME/dotfiles/i3/config"]="$HOME/.config/i3/config"
)

for file in "${!personal_file_symlink_info[@]}"; do ensure_file_symlink_is_in_place "$file" "${personal_file_symlink_info[$file]}"; done
