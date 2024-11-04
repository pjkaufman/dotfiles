#!/usr/bin/env bash

# based loosely on https://github.com/miguelgfierro/scripts/blob/main/git_configure.sh

add_ppa_and_install_package "git-core/ppa" "git"

# ssh client for github
install_apt_package "ssh-keygen" "openssh-server"

mkdir -p "$XDG_CONFIG_HOME/git"

ensure_file_symlink_is_in_place "$DOTFILES/git/config" "$XDG_CONFIG_HOME/git/config" 
ensure_file_symlink_is_in_place "$DOTFILES/git/ignore" "$XDG_CONFIG_HOME/git/ignore"

# setup the ssh values for github

ssh_folder="$HOME/.ssh"

# create ssh folder if missing
if [ ! -d "$ssh_folder" ]; then
  mkdir "$ssh_folder"
fi

ssh_key_file="$ssh_folder/id_ed25519.pub"
if [ ! -f "$ssh_key_file" ]; then
  email=$(git config --global user.email)
  ssh-keygen -t ed25519 -C "$email"
  echo "This is your public key. To activate it in github, go to settings, SHH and GPG keys, New SSH key, and enter the following key:"
  cat "$ssh_key_file"
  echo -e "\nTo work with the ssh key, you have to clone all your repos with ssh instead of https."
else
  echo "You have already ssh-key. To activate it in github, got to settings, SHH and GPG keys, New SSH key, and enter the following key:"
  cat "$ssh_key_file"
fi

ensure_file_symlink_is_in_place "$DOTFILES/ssh/config" "$ssh_folder/config"
