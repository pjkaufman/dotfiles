#!/usr/bin/env bash

function reload() {
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
  command topgrade
}

function rn() {
  if [[ -z "$4" ]]; then
    rename "s/$1/$2/" $3
    return;
  fi

  rename -n "s/$1/$2/" $3
}

# scan the computer for viruses and other issues
function scan() {
  sudo rkhunter --check --sk --rwo
}
