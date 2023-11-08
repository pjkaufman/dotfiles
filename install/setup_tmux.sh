#!/usr/bin/env bash

# setup tmux 

tmux_dir="$XDG_CONFIG_HOME/tmux"
if [ ! -d "$tmux_dir" ]; then 
  mkdir -p "$tmux_dir"
fi

ensure_file_symlink_is_in_place "$DOTFILES/tmux/tmux.conf" "$tmux_dir/tmux.conf"

install_apt_package "tmux"

# install package manager for tmux
tmux_plugin_dir="$XDG_DATA_HOME/tmux/plugins/tpm"
if [ ! -d "$tmux_plugin_dir" ]; then 
  git clone https://github.com/tmux-plugins/tpm "$tmux_plugin_dir"
  # install plugins right after tpm is installed
  "$tmux_plugin_dir/scripts/install_plugins.sh"
fi
