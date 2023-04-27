#!/bin/bash

# setup tmux 

tmux_dir="$HOME/.config/tmux"
if [ ! -d "$HOME/.config/tmux" ]; then 
  mkdir -p "$tmux_dir"
fi

ensure_file_symlink_is_in_place "$HOME/dotfiles/tmux/tmux.conf" "$tmux_dir/tmux.conf"

install_apt_package "tmux"

# install package manager for tmux
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm
