#!/bin/bash

declare -A apt_packages_to_install=( 
  ["tmux"]="tmux" 
  ["grep"]="grep" 
  ["ripgrep"]="rg" # better version of grep
  ["curl"]="curl" 
  ["btop"]="btop" # linux task manager
  ["python3"]="python3"
  ["pip3"]="python3-pip" # python 3 package installer
  ["rename"]="rename" # easier renaming of files
  ["cargo"]="cargo" # package manager for rust
  ["ssh-keygen"]="openssh-client" # ssh client for github
  ["rkhunter"]="rkhunter" # rootkit checker
)

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
