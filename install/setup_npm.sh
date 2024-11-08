#!/usr/bin/env bash

# nvm is special and loads its command via autocompletion and checking the created variable
# is more reliable than checking if the method exists
if [ ! -d "$NVM_DIR" ]; then
  echo "installing nvm"

  mkdir -p "$NVM_DIR"
  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash

  # source bashrc to make sure npm command is available
  # shellcheck source=./bash/bashrc
  source "$HOME/.bashrc"
else
  echo "nvm is already installed at $NVM_DIR"
fi
