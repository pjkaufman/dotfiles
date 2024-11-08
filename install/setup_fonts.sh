#!/usr/bin/env bash

# only try to add the fonts if on a personal computer
if is_work_computer; then
  echo "skipping font installation and setup"
  return
fi

fonts_dir="$HOME/.local/share/fonts"
mkdir -p "$fonts_dir"

cd /tmp || exit
fonts=(
  "Hack"
)

installed_fonts="$(fc-list)"
install_made=0
for font in "${fonts[@]}"; do
  if ! echo "$installed_fonts" | grep -q "$fonts_dir/$font/"; then
    wget "https://github.com/ryanoasis/nerd-fonts/releases/download/v2.3.3/$font.zip"
    unzip "$font.zip" -d "$fonts_dir/$font/"
    rm "$font.zip"
    install_made=1
  else
    echo "$font is already installed."
  fi
done

if [ $install_made = 1 ]; then
  fc-cache
fi
