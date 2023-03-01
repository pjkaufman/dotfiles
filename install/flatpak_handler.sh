#!/bin/bash

# only try to add the flatpaks if on a personal computer
if [ ${COMPUTER_TYPE} = "work" ]
then
  echo "Skipping flatpak installation"
  return
fi

declare -A flatpak_packages_to_install=(
  ["Brave Browser"]="com.brave.Browser"
  ["Minecraft"]="com.mojang.Minecraft"
  ["GnuCash"]="com.gnucash.GnuCash"
  ["Sigil"]="com.sigil_ebook.Sigil"
  ["Calibre"]="com.calibre_ebook.calibre"
  ["Obsidian"]="md.obsidian.Obsidian"
  ["Only Office"]="org.onlyoffice.desktopeditors"
)

if ! command -v flatpak &> /dev/null
then
  echo "Flatpak not installed. Please install it."
else
  for pkg in "${!flatpak_packages_to_install[@]}"; do install_flatpak_package "$pkg" "${apt_packages_to_install[$pkg]}"; done
fi
