#!/usr/bin/env bash

function killsess() {
  tmux kill-session -t "$@"
}

function tls() {
  tmux ls
}
