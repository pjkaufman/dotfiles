#!/bin/bash

check_if_command_exists_and_run_install_command_otherwise "gvm" "curl -sSL https://github.com/soulteary/gvm/raw/master/binscripts/gvm-installer | bash"

# gvm setup if possible
# [[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"

# gvmforceuse $GO_VERSION

go build -C "$HOME/dotfiles/go-tools/git-helper" -o "$HOME/.local/bin/git-helper"
