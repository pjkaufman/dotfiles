#!/bin/bash

# only add these functions if on a personal computer
if [ ${COMPUTER_TYPE} = "personal" ]
then
  return
fi

# personal computer aliases
hibernate() {
  sudo systemctl hibernate
}

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
  flatpak run org.onlyoffice.desktopeditors &
}

# allows for easy running of Sigil via terminal
sigil() {
  flatpak run com.sigil_ebook.Sigil &
}

# allows for easy running of Calibre via terminal
calibre() {
  flatpak run com.calibre_ebook.calibre &
}

# enable the use of brightness since it is locked by admin permissions by default and I need to modify it using user permission
enablebright() {
  sudo chmod a+wr /sys/class/backlight/amdgpu_bl0/brightness
}

compressepub() {
  source "$HOME/dotfiles/bin/compressepub"
}
