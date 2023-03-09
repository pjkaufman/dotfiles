#!/bin/bash

declare -A apt_packages_to_install=( ["tmux"]="tmux" ["grep"]="grep" ["ripgrep"]="rg" ["curl"]="curl" ["btop"]="btop" ["python3"]="python3" ["pip3"]="python3-pip" ["rename"]="rename" ["cargo"]="cargo")

for pkg in "${!apt_packages_to_install[@]}"; do install_apt_package "$pkg" "${apt_packages_to_install[$pkg]}"; done

# only try to add the remaining packages if on a personal computer
if [ ${COMPUTER_TYPE} = "work" ]
then
  return
fi

personal_apt_packages_to_install=(
  "imgp" # image compression
  "pandoc" # document conversion
  "flameshot" # screenshots
  "kitty" # terminal
  "evince" # pdf editor and viewer
)
for pkg in "${personal_apt_packages_to_install[@]}"; do install_apt_package "$pkg"; done
