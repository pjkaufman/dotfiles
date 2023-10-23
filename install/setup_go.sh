#!/usr/bin/env bash

install-go "$GO_VERSION"

cd "$HOME/dotfiles/go-tools" && make install
