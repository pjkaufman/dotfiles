#!/usr/bin/env bash
# tmux aliases

starttmux() {
  source "$HOME/dotfiles/bin/starttmux"
}

killsess() {
  tmux kill-session -t "$@"
}

tls() {
  tmux ls
}
