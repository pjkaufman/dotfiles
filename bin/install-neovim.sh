#!/usr/bin/env bash

set -e

version=latest
if [ "$#" -ge 1 ]; then
    version=$1
fi

curl -LO "https://github.com/neovim/neovim/releases/$version/download/nvim.appimage"
mv nvim.appimage nvim
chmod +x nvim
sudo mv nvim /usr/local/bin
