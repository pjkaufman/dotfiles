#!/usr/bin/env bash

check_if_command_exists_and_run_install_command_otherwise "gvm" "curl -sSL https://github.com/soulteary/gvm/raw/master/binscripts/gvm-installer | bash"

cd "$HOME/dotfiles/go-tools" && make install
