#!/bin/bash

# only try to install i3 packages if on a personal computer
if [ ${COMPUTER_TYPE} = "work" ]
then
  echo "skipping i3 setup"
  return
fi

i3_packages_to_install=(
  "i3"
  "scrot" # screenshots
  "light" # brightness changes
  "feh" # background image
  "i3lock" # lockscreen setup
  "i3status" # status info
  "pactl" # sound changes
  "dmenu" # app selector
)
for pkg in "${i3_packages_to_install[@]}"; do install_apt_package "$pkg"; done

pip_install_package "bumblebee-status" # status bar
