#!/bin/sh
# Based on https://github.com/mathiasbynens/dotfiles/blob/master/.exports

# Set the default editor to neovim
export EDITOR="nvim"

# Increase Bash history size. Allow 32³ entries; the default is 500.
export HISTSIZE='32768';
export HISTFILESIZE="${HISTSIZE}";
# Omit duplicates and commands that begin with a space from history.
export HISTCONTROL='ignoreboth';

# Prefer US English and use UTF-8.
export LANG='en_US.UTF-8';
export LC_ALL='en_US.UTF-8';

# Highlight section titles in manual pages.
export LESS_TERMCAP_md="${yellow}";

# nvm
export NVM_VERSION=16.0

# Golang
export GO_BINARY_BASE_URL=https://go.dev/dl
export GO_VERSION=1.19
export GOROOT=/home/peterk/.gvm/gos/go$GO_VERSION/

# nvm 
export NVM_DIR="$HOME/.nvm"
