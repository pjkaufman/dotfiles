#!/usr/bin/env bash

# only try to install remaining PPA packages if on a personal computer
if is_work_computer; then
  echo "skipping syncthing install"
  return
fi

add_ppa_and_install_package "syncthing/stable" "syncthing"
