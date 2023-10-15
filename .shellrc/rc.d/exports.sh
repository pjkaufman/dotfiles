#!/bin/sh
# Based on https://github.com/mathiasbynens/dotfiles/blob/master/.exports

# XDG vars
export XDG_DATA_HOME=$HOME/.local/share
export XDG_CONFIG_HOME=$HOME/.config
export XDG_STATE_HOME=$HOME/.local/state
export XDG_CACHE_HOME=$HOME/.cache

# Set the default editor to neovim
export EDITOR="nvim"

# Increase Bash history size. Allow 32³ entries; the default is 500.
export HISTSIZE='32768';
export HISTFILESIZE="${HISTSIZE}";

# make history file respect xdg dirs
export HISTFILE="${XDG_STATE_HOME}"/bash/history

# Omit duplicates and commands that begin with a space from history.
export HISTCONTROL='ignoreboth';

# Prefer US English and use UTF-8.
export LANG='en_US.UTF-8';
export LC_ALL='en_US.UTF-8';

# Highlight section titles in manual pages.
export LESS_TERMCAP_md="${yellow}";

# nvm
export NVM_VERSION=18.0
export NVM_DIR="$HOME/.nvm"

# Golang
export GO_BINARY_BASE_URL=https://go.dev/dl
export GO_VERSION=1.21
export GOROOT=$HOME/.gvm/gos/go$GO_VERSION/

# set SSH_AUTH_SOCK env var to a fixed value
export SSH_AUTH_SOCK=$HOME/.ssh/ssh-agent.sock

export TERM=xterm-256color # needed to prevent an error on load

# wine will get placed in the xdg config
export WINEPREFIX="$XDG_CONFIG_HOME"/wine/

# make sure bash completion gets put in xdg folders
export BASH_COMPLETION_USER_DIR="$XDG_CONFIG_HOME"/bash-completion
export BASH_COMPLETION_USER_FILE="$BASH_COMPLETION_USER_DIR"/bash_completion

# get cargo to install under xdg folders
export CARGO_HOME="$XDG_DATA_HOME"/cargo

# make sure that less history is under xdg folders
export LESSHISTFILE="$XDG_STATE_HOME"/less/history
