#!/usr/bin/env bash

if is_work_computer; then
  echo "skipping kitty setup"
  return
fi

install_apt_package "kitty"

ensure_file_symlink_is_in_place "$DOTFILES/kitty/kitty.conf" "$XDG_CONFIG_HOME/kitty/kitty.conf"
