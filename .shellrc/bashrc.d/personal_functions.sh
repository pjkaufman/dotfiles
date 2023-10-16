#!/usr/bin/env bash

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
  ebook-lint epub compress-and-lint -i
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
# the first param is expected to be an epub file
# the second param is expected to be a Markdown file
epubreplaceallstrings() {
  ebook-lint epub replace-strings -f "$1" -e "$2"
}

# epubmanualfixes helps go through manually fixable epub issues
# the first param is expected to be an epub file
epubmanualfixes() {
  # TODO: see about swapping the logic to check the param count and based on the param
  # count either take in all params provided as is or just take in the epub value
  ebook-lint epub replace-strings -f "$1" -a
}