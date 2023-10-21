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
  if [[ -z $4 ]]; then
    rename "s/$1/$2/" $3
    return;
  fi

  rename -n "s/$1/$2/" $3
}

# keyboard setup

# set keyboard layout to English international to allow alt+character to have accented or Spanish characters
function eson(){
  setxkbmap -layout us -variant intl
}

# set keyboard layout to regular English to make it easier to program and do other things from the cli
function esoff(){
  setxkbmap -layout us
}

# scan the computer for viruses and other issues
function scan() {
  sudo rkhunter --check --sk --rwo
}
