#!/usr/bin/env bash
# tmux aliases

function killsess() {
  tmux kill-session -t "$@"
}

function tls() {
  tmux ls
}
