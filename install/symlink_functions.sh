#!/usr/bin/env bash

ensure_file_symlink_is_in_place() {
  if [ -L "$2" ]; then
    if [ -e "$2" ]; then
      echo "'$2' is already symlinked"
    else
      echo "'$2' is a broken symlink"
    fi
  elif [ -e "$2" ]; then
    echo "'$2' exists, but is not symlinked"
    mv "$2" "$2.bak"
  else
    echo "'$2' does not exist"
  fi

  dir="$(dirname "$2")"
  [[ ! -d "$dir" ]] && mkdir -p "$dir"

  ln -sf "$1" "$2"
}

ensure_file_symlink_is_in_place_as_sudo() {
  if [ -L "$2" ]; then
    if [ -e "$2" ]; then
      echo "'$2' is already symlinked"
    else
      echo "'$2' is a broken symlink"
    fi
  elif [ -e "$2" ]; then
    echo "'$2' exists, but is not symlinked"
    sudo mv "$2" "$2.bak"
  else
    echo "'$2' does not exist"
  fi

  dir="$(dirname "$2")"
  [[ ! -d "$dir" ]] && mkdir -p "$dir"

  sudo ln -sf "$1" "$2"
}

ensure_folder_symlink_is_in_place() {
  if [ -L "$2" ]; then
    if [ -d "$2" ]; then
      echo "'$2' is already symlinked, please check that it is the correct symlink"
      return
    else
      echo "'$2' is a broken symlink"
      ln -s "$1" "$2"
    fi
  elif [ -d "$2" ]; then
    echo "'$2' exists, but is not symlinked (implementation needed)"
    return
  else
    echo "'$2' does not exist"
    ln -s "$1" "$2"
  fi
}
