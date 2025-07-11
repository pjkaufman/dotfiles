#!/usr/bin/env bash

# flatpak functions

# only add these functions if on a personal computer
if is_work_computer; then
  return
fi

# allows for easy running of obsidian via terminal
function obsidian() {
  flatpak run md.obsidian.Obsidian --ozone-platform-hint=auto --enable-features=WaylandWindowDecorations &
}

# allows for easy running of brave via terminal
function brave() {
  flatpak run com.brave.Browser "$@" --ozone-platform-hint=auto --enable-features=WaylandWindowDecorations &
}

# allows for easy running of GnuCash via terminal
function gnucash() {
  flatpak run org.gnucash.GnuCash &
}

# allows for easy running of Minecraft via terminal
function minecraft() {
  flatpak run com.mojang.Minecraft &
}

# allows for easy running of Calibre via terminal
function calibre() {
  flatpak run com.calibre_ebook.calibre &
}

# allows for easy running of Calibre editor view via terminal
function editepub() {
  flatpak run --command="ebook-edit" com.calibre_ebook.calibre --detach "$@"
}

# allows for easy running of Calibre reader view via terminal
function vieweebook() {
  flatpak run --command="ebook-viewer" com.calibre_ebook.calibre --detach "$@"
}

function convertebook() {
  if [ "$#" -le 1 ]; then
    echo "No arguments supplied"
    echo "Usage convertebook [epub-file] [new-epub-file]"
  fi

  flatpak run --command="ebook-convert" com.calibre_ebook.calibre "$1" "$2"
}
