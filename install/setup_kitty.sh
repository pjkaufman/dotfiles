#!/bin/bash

if is_work_computer; then
  echo "skipping kitty setup"
  return
fi

install_apt_package "kitty"

ensure_file_symlink_is_in_place "$HOME/dotfiles/kitty/kitty.conf" "$HOME/.config/kitty/kitty.conf"
