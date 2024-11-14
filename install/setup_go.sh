#!/usr/bin/env bash

install-go "$GO_VERSION"

go_install_package "mvdan.cc/sh/v3/cmd/shfmt" "shfmt"

git submodule init

cd "$DOTFILES/go-go-gadgets" && git checkout master && git pull

# reload source so we can have go available in the path
# shellcheck source=./bash/bashrc
source "$HOME/.bashrc"

make install
