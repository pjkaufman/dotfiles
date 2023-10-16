#!/usr/bin/env bash

# nvm is special and loads its command via autocompletion and checking the created variable
# is more reliable than checking if the method exists
if [ -z ${NVM_DIR} ]
then
  echo "installing nvm"
  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
else
  echo "nvm is already installed"
fi
