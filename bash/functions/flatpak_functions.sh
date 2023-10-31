#!/usr/bin/env bash

# flatpak functions

# only add these functions if on a personal computer
if is_work_computer; then
	return
fi

# allows for easy running of obsidian via terminal
function obsidian() {
	flatpak run md.obsidian.Obsidian &
}

# allows for easy running of brave via terminal
function brave() {
	flatpak run com.brave.Browser "$@" &
}

# allows for easy running of GnuCash via terminal
function gnucash() {
	flatpak run org.gnucash.GnuCash &
}

# allows for easy running of Minecraft via terminal
function minecraft() {
	flatpak run com.mojang.Minecraft &
}

# allows for easy running of Only Office via terminal
function office() {
	eson
	flatpak run org.onlyoffice.desktopeditors &
	wait
	esoff # freezes the current terminal until the previous command finishes and it will then turn Spanish characters off for the same of terminal
	# typing
}

# allows for easy running of Sigil via terminal
function sigil() {
	flatpak run com.sigil_ebook.Sigil "$@" &
}

# allows for easy running of Calibre via terminal
function calibre() {
	flatpak run com.calibre_ebook.calibre &
}
