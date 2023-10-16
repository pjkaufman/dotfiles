#!/usr/bin/env bash
# flatpak aliases

# only add these functions if on a personal computer
if [ ${COMPUTER_TYPE} = "work" ]
then
  return
fi

# allows for easy running of obsidian via terminal
obsidian() {
  flatpak run md.obsidian.Obsidian &
}
# allows for easy running of brave via terminal
brave() {
  flatpak run com.brave.Browser &
}

# allows for easy running of GnuCash via terminal
gnucash() {
  flatpak run org.gnucash.GnuCash &
}

# allows for easy running of Minecraft via terminal
minecraft() {
  flatpak run com.mojang.Minecraft &
}

# allows for easy running of Only Office via terminal
office() {
  eson
  flatpak run org.onlyoffice.desktopeditors &
  wait; esoff # freezes the current terminal until the previous command finishes and it will then turn Spanish characters off for the same of terminal
  # typing
}

# allows for easy running of Sigil via terminal
sigil() {
  flatpak run com.sigil_ebook.Sigil &
}

# allows for easy running of Calibre via terminal
calibre() {
  flatpak run com.calibre_ebook.calibre &
}
