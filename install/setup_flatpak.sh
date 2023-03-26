#!/bin/bash

# only try to add the flatpaks if on a personal computer
if is_work_computer; then
  echo "skipping flatpak installation and setup"
  return
fi

declare -A flatpak_packages_to_install=(
  ["Brave Browser"]="com.brave.Browser"
  ["Minecraft"]="com.mojang.Minecraft"
  ["GnuCash"]="org.gnucash.GnuCash"
  ["Sigil"]="com.sigil_ebook.Sigil"
  ["Calibre"]="com.calibre_ebook.calibre"
  ["Obsidian"]="md.obsidian.Obsidian"
  ["Only Office"]="org.onlyoffice.desktopeditors"
)

if ! command -v flatpak &> /dev/null; then
  echo "Flatpak not installed. Please install it."
else
  install_flatpak_package "Brave Browser" "com.brave.Browser"
  install_flatpak_package "Minecraft" "com.mojang.Minecraft"
  install_flatpak_package "GnuCash" "org.gnucash.GnuCash"
  install_flatpak_package "Sigil" "com.sigil_ebook.Sigil"
  install_flatpak_package "Calibre" "com.calibre_ebook.calibre"
  install_flatpak_package "Obsidian" "md.obsidian.Obsidian"
  install_flatpak_package "Only Office" "org.onlyoffice.desktopeditors"

  # for pkg in "${!flatpak_packages_to_install[@]}"; do install_flatpak_package "$pkg" "${apt_packages_to_install[$pkg]}"; done
fi
