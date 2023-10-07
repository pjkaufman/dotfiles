#!/bin/bash

# only add these functions if on a personal computer
if [ ${COMPUTER_TYPE} = "work" ]
then
  return
fi

# personal computer aliases
hibernate() {
  sudo systemctl hibernate
}

# enable the use of brightness since it is locked by admin permissions by default and I need to modify it using user permission
enablebright() {
  sudo chmod a+wr /sys/class/backlight/amdgpu_bl0/brightness
}

# compressepub helps with compressing epubs so they take up less space
compressepub() {
  source "$HOME/dotfiles/bin/compressepub"
}

# convertcbrtocbz helps with converting cbrs to cbzs
convertcbrtocbz() {
  source "$HOME/dotfiles/bin/convertcbrtocbz"
}

# compresscbz helps with compressing cbzs so they take up less space
compresscbz() {
  source "$HOME/dotfiles/bin/compresscbz"
}

# epubreplaceallstrings helps with replacing a bunch of strings in an epub file
epubreplaceallstrings() {
  source "$HOME/dotfiles/bin/epubreplaceallstrings"
}

# epubmanualfixes helps go through manually fixable epub issues
epubmanualfixes() {
  source "$HOME/dotfiles/bin/epubmanualfixes"
}