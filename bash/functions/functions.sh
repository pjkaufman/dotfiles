#!/usr/bin/env bash

function reload() {
  # shellcheck source=./bash/bashrc
  [ -f "$HOME/.bashrc" ] && source "$HOME/.bashrc"
}

function c() {
  clear && clear
}

function gh() {
  history | grep "$@"
}

function python() {
  command python3 "$@"
}

function update() {
  sudo apt update -y && sudo apt upgrade
  pipx upgrade-all

  is_personal_computer && fwupdmgr update
  [ -n "$(which flatpak)" ] && flatpak update -y
}

function rn() {
  if [[ -z "$4" ]]; then
    # shellcheck disable=SC2086
    rename "s/$1/$2/" $3
    return
  fi

  # shellcheck disable=SC2086
  rename -n "s/$1/$2/" $3
}

# scan the computer for viruses and other issues
function scan() {
  sudo rkhunter --check --sk --rwo
}
