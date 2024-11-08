#!/usr/bin/env bash

# only add these functions if on a personal computer
if is_work_computer; then
  return
fi

# personal computer functions

function hibernate() {
  sudo systemctl hibernate
}

# enable the use of brightness since it is locked by admin permissions by default and I need to modify it using user permission
function enablebright() {
  sudo chmod a+wr /sys/class/backlight/amdgpu_bl0/brightness
}

# ebook functions

# compressepub helps with compressing epubs so they take up less space
function compressepub() {
  ebook-lint epub compress-and-lint -i
}

# convertcbrtocbz helps with converting cbrs to cbzs
function convertcbrtocbz() {
  ebook-lint cbr to-cbz "$@"
}

# compresscbz helps with compressing cbzs so they take up less space
function compresscbz() {
  ebook-lint cbz compress "$@"
}

# epubreplaceallstrings helps with replacing a bunch of strings in an epub file
# the first param is expected to be an epub file
# the second param is expected to be a Markdown file
function epubreplaceallstrings() {
  ebook-lint epub replace-strings -f "$1" -e "$2"
}

# epubmanualfixes helps go through manually fixable epub issues
# the first param is expected to be an epub file
function epubmanualfixes() {
  # TODO: see about swapping the logic to check the param count and based on the param
  # count either take in all params provided as is or just take in the epub value
  ebook-lint epub fixable -f "$1" -a
}

# Obsidian functions for opening differnt vaults that I have

function notes() {
  xdg-open obsidian://open?vault=Obsidian
}

function testVault() {
  xdg-open obsidian://open?vault=TestVault
}
