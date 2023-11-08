#!/usr/bin/env bash

# only try to install sway packages if on a personal computer
if is_work_computer ; then
  echo "skipping sway setup"
  return
fi

sway_packages_to_install=(
  "sway"
  "scrot" # screenshots
  "light" # brightness changes
  "swaylock" # lockscreen setup
  "waybar" # status bar
  "wofi" # app selector
)
for pkg in "${sway_packages_to_install[@]}"; do install_apt_package "$pkg"; done

install_apt_package "pactl" "pulseaudio-utils" # sound changes

# clipboard manager
go_install_package "cliphist" "go.senan.xyz/cliphist"

ensure_folder_symlink_is_in_place "$DOTFILES/sway" "$HOME/.config/sway"
ensure_folder_symlink_is_in_place "$DOTFILES/waybar" "$HOME/.config/waybar"

# echo "Copying images"

# # make sure to overwrite the existing image if there is one since I cannot use symlinks
# cp -f "$HOME/dotfiles/i3/Laminin.png" "$HOME/.config/i3/Laminin.png"
# cp -f "$HOME/dotfiles/i3/CharlesOutside.jpg" "$HOME/.config/i3/CharlesOutside.jpg"
