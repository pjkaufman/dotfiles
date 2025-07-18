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
  if [ "$#" -eq 0 ]; then
    epub-lint compress-and-lint -i
  else
    epub-lint compress-and-lint -i -l "$1"
  fi
}

# epubreplaceallstrings helps with replacing a bunch of strings in an epub file
# the first param is expected to be an epub file
# the second param is expected to be a Markdown file
function epubreplaceallstrings() {
  epub-lint replace-strings -f "$1" -e "$2"
}

# epubmanualfixes helps go through manually fixable epub issues
# the first param is expected to be an epub file
function epubmanualfixes() {
  # TODO: see about swapping the logic to check the param count and based on the param
  # count either take in all params provided as is or just take in the epub value
  if [ "$#" -eq 0 ]; then
    echo "No arguments supplied"
    echo "Usage epubmanualfixes [epub-file] [any value to indicate that you want to use the TUI] [optional TUI log file]"
  elif [ "$#" -eq 1 ]; then
    epub-lint fixable -f "$1" -a
  elif [ "$#" -eq 2 ]; then
    epub-lint fixable -f "$1" -a --use-tui
  else
    epub-lint fixable -f "$1" -a --use-tui --log-file "$3"
  fi
}

function validateepub() {
  if [ "$#" -eq 0 ]; then
    echo "No arguments supplied"
    echo "Usage validateepub [epub-file] [optional-json-output-file]"
  elif [ "$#" -eq 1 ]; then
    epub-lint validate -f "$1"
  else
    epub-lint validate -f "$1" --json-file "$2"
  fi
}

function fixepub() {
  if [ ! "$#" -eq 2 ]; then
    echo "Incorrect number of arguments supplied"
    echo "Usage fixepub  [epub-file] [json-file]"
  else
    epub-lint fix-validation -f "$1" --issue-file "$2" --cleanup-jnovels
  fi
}

# Obsidian functions for opening different vaults that I have

function notes() {
  xdg-open obsidian://open?vault=Obsidian
}

function testVault() {
  xdg-open obsidian://open?vault=TestVault
}
