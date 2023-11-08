#!/usr/bin/env bash

install-go "$GO_VERSION"

cd "$DOTFILES/go-tools" && make install
