#!/bin/bash

# only try to add the flatpaks if on a personal computer
if is_work_computer; then
  echo "skipping flatpak installation and setup"
  return
fi

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
fi

# Only Office Settings

only_office_dark_theme_setting="UITheme2=theme-dark"
only_office_settings_file="$HOME/.var/app/org.onlyoffice.desktopeditors/config/onlyoffice/DesktopEditors.conf"
echo "$only_office_settings_file"
if [ -z $(grep "$only_office_dark_theme_setting" "$only_office_settings_file") ]; then
  echo "Adding dark theme setting for Only Office"
  sed -i "s/\[General\]/\[General\]\n$only_office_dark_theme_setting/g" $only_office_settings_file
fi
