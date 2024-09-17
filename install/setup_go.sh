#!/usr/bin/env bash

install-go "$GO_VERSION"

git submodule init

cd "$DOTFILES/go-go-gadgets" && git checkout master && git pull

# reload source so we can have go available in the path
source "$HOME/.bashrc"

make install
