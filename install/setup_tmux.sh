#!/bin/bash

# setup tmux 

tmux_dir="$HOME/.config/tmux"
if [ ! -d "$tmux_dir" ]; then 
  mkdir -p "$tmux_dir"
fi

ensure_file_symlink_is_in_place "$HOME/dotfiles/tmux/tmux.conf" "$tmux_dir/tmux.conf"

install_apt_package "tmux"

# install package manager for tmux
tmux_plugin_dir="$HOME/.tmux/plugins/tpm"
if [ ! -d "$tmux_plugin_dir" ]; then 
  git clone https://github.com/tmux-plugins/tpm $HOME/.tmux/plugins/tpm
  # install plugins right after tpm is installed
  $HOME/.config/tmux/plugins/tpm/scripts/install_plugins.sh
fi
