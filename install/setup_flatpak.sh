#!/usr/bin/env bash

# only try to add the flatpaks if on a personal computer
if is_work_computer; then
	echo "skipping flatpak installation and setup"
	return
fi

# Only Office Settings
function setup_only_office_settings() {
	only_office_dark_theme_setting="UITheme2=theme-dark"
	only_office_settings_file="$HOME/.var/app/org.onlyoffice.desktopeditors/config/onlyoffice/DesktopEditors.conf"
	if ! grep -q "$only_office_dark_theme_setting" "$only_office_settings_file"; then
		echo "Adding dark theme setting for Only Office"
		sed -i "s/\[General\]/\[General\]\n$only_office_dark_theme_setting/g" "$only_office_settings_file"
	fi
}

# GnuCash Settings
function setup_gnucash_settings() {
	gnucashConfigDir="$HOME/.var/app/org.gnucash.GnuCash/config"
	ensure_folder_symlink_is_in_place "$DOTFILES/gnucash" "$gnucashConfigDir/gnucash"
}

# Sigil Settings
function setup_sigil_settings() {
	ensure_file_symlink_is_in_place "$DOTFILES/sigil/qt_styles.qss" "$HOME/.var/app/com.sigil_ebook.Sigil/data/sigil-ebook/sigil/qt_styles.qss"
	ensure_folder_symlink_is_in_place "$DOTFILES/sigil/user-dictionaries" "$HOME/.var/app/com.sigil_ebook.Sigil/data/sigil-ebook/sigil/user_dictionaries"
}

# Obsidian Settings
function setup_obsidian_settings() {
	# based on https://forum.obsidian.md/t/meta-post-linux-tips-tricks-solutions-to-common-problems/6291/17
	desktop_folder="$HOME/.local/share/applications"
	desktop_file=obsidian.desktop
	obsidian_desktop="$desktop_folder/$desktop_file"

  if [ ! -s "$desktop_folder/obsidian.desktop" ]; then
    ensure_file_symlink_is_in_place "$HOME/dotfiles/obsidian/$desktop_file" "$obsidian_desktop"
  fi

  if [ ! "$(xdg-mime query default x-scheme-handler/obsidian)" == "$desktop_file" ]; then
    xdg-mime default "$desktop_file" x-scheme-handler/obsidian
    update-desktop-database
  fi
}

if ! command -v flatpak &>/dev/null; then
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
setup_obsidian_settings

sudo flatpak override --filesystem="$HOME/.themes"
sudo flatpak override --filesystem="$HOME/.config/gtk-3.0"
sudo flatpak override --env=GTK_THEME="$GTK_THEME"
# sudo flatpak override --env=XDG_CURRENT_DESKTOP="$XDG_CURRENT_DESKTOP"
# sudo flatpak override --env=QT_QPA_PLATFORM=wayland # this currently does not seem to work
