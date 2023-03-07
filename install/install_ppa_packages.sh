#!/bin/bash

add_ppa_and_install_package "neovim-ppa/unstable" "neovim"

add_ppa_and_install_package "git-core/ppa" "git"

# only try to install remaining PPA packages if on a personal computer
if [ ${COMPUTER_TYPE} = "work" ]
then
  return
fi

add_ppa_and_install_package "syncthing/stable" "syncthing"
