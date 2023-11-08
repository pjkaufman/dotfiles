#!/usr/bin/env bash

install_apt_package "rkhunter"

ensure_file_symlink_is_in_place_as_sudo "$DOTFILES/rkhunter/rkhunter.conf" "/etc/rkhunter.conf"
