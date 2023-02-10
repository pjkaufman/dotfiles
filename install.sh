#!/bin/bash

# helper functions

install_managed_package() {

  if ! command -v git &> /dev/null
  then
      echo "installing $1"
      sudo apt install $1
  else
    echo "$1 is already installed"
  fi
}

# actual setup
echo "starting environment setup"
echo ""

echo "installing managed packages:"
echo ""

install_managed_package "git"
install_managed_package "tmux"
install_managed_package "btop"
install_managed_package "python3"
install_managed_package "python3-pip"

echo ""
echo "environment setup complete"
