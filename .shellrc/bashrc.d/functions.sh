#!/bin/bash

# Add to path
prepend-path() {
  [ -d $1 ] && PATH="$1:$PATH"
}

reload() {
  [ -f "$HOME/.bash_profile" ] && source "$HOME/.bash_profile"
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

# tmux aliases
function starttmux() {
  source "$HOME/dotfiles/bin/starttmux"
}

function killsess() {
  tmux kill-session -t "$@"
}

function tls() {
  tmux ls
}
