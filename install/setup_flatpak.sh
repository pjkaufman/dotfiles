#!/usr/bin/env bash

# only try to add the flatpaks if on a personal computer
if is_work_computer; then
  echo "skipping flatpak installation and setup"
  return
fi

# Only Office Settings
setup_only_office_settings() {
  only_office_dark_theme_setting="UITheme2=theme-dark"
  only_office_settings_file="$HOME/.var/app/org.onlyoffice.desktopeditors/config/onlyoffice/DesktopEditors.conf"
  
  if [ -z $(grep "$only_office_dark_theme_setting" "$only_office_settings_file") ]; then
    echo "Adding dark theme setting for Only Office"
    sed -i "s/\[General\]/\[General\]\n$only_office_dark_theme_setting/g" $only_office_settings_file
  fi
}

# GnuCash Settings 
setup_gnucash_settings() {
  ensure_file_symlink_is_in_place "$HOME/dotfiles/gtk-3.0/settings.ini" "$HOME/.var/app/org.gnucash.GnuCash/config/gtk-3.0/settings.ini"
}

# Sigil Settings
setup_sigil_settings() {
  ensure_file_symlink_is_in_place "$HOME/dotfiles/sigil/qt_styles.qss"  "$HOME/.var/app/com.sigil_ebook.Sigil/data/sigil-ebook/sigil/qt_styles.qss"
  ensure_folder_symlink_is_in_place "$HOME/dotfiles/sigil/user-dictionaries" "$HOME/.var/app/com.sigil_ebook.Sigil/data/sigil-ebook/sigil/user_dictionaries"
}

if ! command -v flatpak &> /dev/null; then
  echo "Flatpak not installed. Please install it."
  return
else
  install_flatpak_package "Brave Browser" "com.brave.Browser"
  install_flatpak_package "Minecraft" "com.mojang.Minecraft"
  install_flatpak_package "GnuCash" "org.gnucash.GnuCash"
  install_flatpak_package "Sigil" "com.sigil_ebook.Sigil"
  install_flatpak_package "Calibre" "com.calibre_ebook.calibre"
  install_flatpak_package "Obsidian" "md.obsidian.Obsidian"
  install_flatpak_package "Only Office" "org.onlyoffice.desktopeditors"
fi

setup_only_office_settings
setup_gnucash_settings
setup_sigil_settings
