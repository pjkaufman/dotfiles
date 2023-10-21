#!/usr/bin/env bash

# sshstart starts ssh if it has not already been started
function sshstart() {
  # test whether $SSH_AUTH_SOCK is valid
  ssh-add -l 2>/dev/null >/dev/null

  # if not valid, then start ssh-agent using $SSH_AUTH_SOCK
  [ $? -ge 2 ] && ssh-agent -a "$SSH_AUTH_SOCK" >/dev/null
}
